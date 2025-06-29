package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "backend/analyzer/cst"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "bytes"
    "fmt"
    "bufio"
    "time"
)

// Estructuras de request/response para el API
type ExecuteRequest struct {
    Code string `json:"code"`
}

type ExecuteResponse struct {
    Success   bool            `json:"success"`
    Output    string          `json:"output"`
    Errors    []ErrorDetail   `json:"errors"`
    Symbols   []SymbolDetail  `json:"symbols"`
    Stats     ExecutionStats  `json:"stats"`
    Optimized bool            `json:"optimized"`
}

type ErrorDetail struct {
    Type     string `json:"type"`
    Message  string `json:"message"`
    Line     int    `json:"line"`
    Column   int    `json:"column"`
    Severity string `json:"severity"`
}

type SymbolDetail struct {
    Name     string `json:"name"`
    Type     string `json:"type"`
    Scope    string `json:"scope"`
    Line     int    `json:"line"`
    Column   int    `json:"column"`
}

type ExecutionStats struct {
    CompileTime string `json:"compileTime"`
    Lines       int    `json:"lines"`
    Size        string `json:"size"`
}

// Estructuras del compilador (basadas en client.go)
type CompilerResponse struct {
    Success bool                    `json:"success"`
    Output  string                  `json:"output"`
    Errors  []CompilerErrorDetail   `json:"errors"`
    Symbols []CompilerSymbolDetail  `json:"symbols"`
    Stats   CompilerStats           `json:"stats"`
}

type CompilerErrorDetail struct {
    Line    int    `json:"line"`
    Column  int    `json:"column"`
    Type    string `json:"type"`
    Message string `json:"message"`
}

type CompilerSymbolDetail struct {
    ID         string `json:"id"`
    SymbolType string `json:"symbolType"`
    DataType   string `json:"dataType"`
    Scope      string `json:"scope"`
    Line       int    `json:"line"`
    Column     int    `json:"column"`
}

type CompilerStats struct {
    ExecutionTime int64 `json:"executionTime"`
    CodeSize      int   `json:"codeSize"`
    ErrorCount    int   `json:"errorCount"`
    SymbolCount   int   `json:"symbolCount"`
}

type CSTRequest struct {
    Code      string `json:"code"`
    Grammar   string `json:"grammar"`
    StartRule string `json:"startRule"`
}

type CSTResponse struct {
    Success     bool   `json:"success"`
    HTMLContent string `json:"htmlContent"`
    Error       string `json:"error,omitempty"`
}

type APIServer struct {
    analyzerPath string
    compilerPath string

}

type CompileAndRunRequest struct {
    Code string `json:"code"`
}

type CompileAndRunResponse struct {
    Success         bool   `json:"success"`
    CompilationLog  string `json:"compilationLog"`
    ExecutionOutput string `json:"executionOutput"`
    ExecutionError  string `json:"executionError,omitempty"`
    Assembly        string `json:"assembly,omitempty"`
    Error           string `json:"error,omitempty"`
}

func NewAPIServer() *APIServer {
    server := &APIServer{}
    server.findExecutables()
    return server
}

// Buscar el ejecutable del compilador
func (s *APIServer) findExecutables() {
    // Buscar analizador
    analyzerPaths := []string{
        "../analyzer/analyzer",
        "./analyzer/analyzer",
        "../backend/analyzer/analyzer",
    }
    
    for _, path := range analyzerPaths {
        if _, err := os.Stat(path); err == nil {
            absPath, _ := filepath.Abs(path)
            s.analyzerPath = absPath
            log.Printf("✅ Analizador encontrado: %s", absPath)
            break
        }
    }
    
    // Buscar compilador
    compilerPaths := []string{
        "../compiler/compiler",
        "./compiler/compiler",
        "../backend/compiler/compiler",
    }
    
    for _, path := range compilerPaths {
        if _, err := os.Stat(path); err == nil {
            absPath, _ := filepath.Abs(path)
            s.compilerPath = absPath
            log.Printf("✅ Compilador encontrado: %s", absPath)
            break
        }
    }
    
    if s.analyzerPath == "" {
        log.Fatal("❌ No se encontró el analizador")
    }
    if s.compilerPath == "" {
        log.Fatal("❌ No se encontró el compilador")
    }
}

// Función helper para mapear símbolos
func mapSymbols(analyzerSymbols []struct {
    ID         string `json:"id"`
    SymbolType string `json:"symbolType"`
    DataType   string `json:"dataType"`
    Scope      string `json:"scope"`
    Line       int    `json:"line"`
    Column     int    `json:"column"`
}) []SymbolDetail {
    var mapped []SymbolDetail
    for _, symbol := range analyzerSymbols {
        // Combinar SymbolType y DataType para mostrar información completa
        typeInfo := symbol.SymbolType
        if symbol.DataType != "" && symbol.DataType != symbol.SymbolType {
            typeInfo = symbol.SymbolType + " (" + symbol.DataType + ")"
        }
        
        mapped = append(mapped, SymbolDetail{
            Name:   symbol.ID,
            Type:   typeInfo,
            Scope:  symbol.Scope,
            Line:   symbol.Line,
            Column: symbol.Column,
        })
    }
    return mapped
}

// Endpoint para ejecutar código (principal)
func (s *APIServer) executeCode(w http.ResponseWriter, r *http.Request) {
    var req ExecuteRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    log.Printf("🔄 Analizando código (%d caracteres)", len(req.Code))

    // 1. PRIMERO: Ejecutar analyzer para detectar errores
    analyzerCmd := exec.Command(s.analyzerPath)
    analyzerCmd.Stdin = strings.NewReader(req.Code)
    
    analyzerOutput, analyzerErr := analyzerCmd.CombinedOutput()
    
    // Verificar si hay error en la ejecución del analyzer
    if analyzerErr != nil {
        log.Printf("❌ Error ejecutando analyzer: %v", analyzerErr)
        http.Error(w, "Error running analyzer: "+analyzerErr.Error(), http.StatusInternalServerError)
        return
    }
    
    // ✅ Limpiar la salida del analyzer (remover AMBOS prefijos)
    analyzerOutputStr := string(analyzerOutput)
    if after, ok := strings.CutPrefix(analyzerOutputStr, "SUCCESS:"); ok {
        analyzerOutputStr = after
    } else if after, ok := strings.CutPrefix(analyzerOutputStr, "ERROR_REPORT:"); ok {
        analyzerOutputStr = after
    }
    
    // ✅ Parsear resultado del analyzer con la estructura correcta del analyzer
    var analyzerResult struct {
        Success bool          `json:"success"`
        Output  string        `json:"output"`
        Errors  []ErrorDetail `json:"errors"`
        Symbols []struct {
            ID         string `json:"id"`         // ← Campos como los devuelve el analyzer
            SymbolType string `json:"symbolType"`
            DataType   string `json:"dataType"`
            Scope      string `json:"scope"`
            Line       int    `json:"line"`
            Column     int    `json:"column"`
        } `json:"symbols"`
        Stats   struct {
            ExecutionTime int64 `json:"executionTime"`
            CodeSize      int   `json:"codeSize"`
            ErrorCount    int   `json:"errorCount"`
            SymbolCount   int   `json:"symbolCount"`
        } `json:"stats"`
    }
    
    if err := json.Unmarshal([]byte(analyzerOutputStr), &analyzerResult); err != nil {
        log.Printf("❌ Error parseando salida del analyzer: %v", err)
        log.Printf("📄 Salida del analyzer (limpia): %s", analyzerOutputStr)
        http.Error(w, "Error parsing analyzer output: "+err.Error(), http.StatusInternalServerError)
        return
    }
    
    // ✅ Mapear símbolos del formato analyzer al formato frontend
    mappedSymbols := mapSymbols(analyzerResult.Symbols)
    
    // 2. ✅ SI HAY ERRORES: No compilar, solo devolver errores
    if !analyzerResult.Success || len(analyzerResult.Errors) > 0 {
        log.Printf("⚠️ Se encontraron %d errores, no se compilará", len(analyzerResult.Errors))
        
        result := ExecuteResponse{
            Success: false,
            Output:  "Se encontraron errores en el código. Revisa la tabla de errores.",
            Errors:  analyzerResult.Errors,
            Symbols: mappedSymbols, // ← Usar símbolos mapeados
            Stats: ExecutionStats{
                CompileTime: "0ms (no compilado)",
                Lines:       strings.Count(req.Code, "\n") + 1,
                Size:        "0 bytes",
            },
            Optimized: false,
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(result)
        return
    }

    // 3. SI NO HAY ERRORES: Ejecutar compilador
    log.Printf("✅ Análisis exitoso, compilando a ARM64...")
    
    compilerCmd := exec.Command(s.compilerPath)
    compilerCmd.Stdin = strings.NewReader(req.Code)
    
    // Separar stdout y stderr
    var stdout, stderr bytes.Buffer
    compilerCmd.Stdout = &stdout
    compilerCmd.Stderr = &stderr

    compilerErr := compilerCmd.Run()

    // Mostrar debug en logs del servidor
    if stderr.Len() > 0 {
        log.Printf("🔍 Debug del compilador:\n%s", stderr.String())
    }

    // Usar solo stdout para el JSON
    compilerOutput := stdout.Bytes()
    
    // Verificar si hay error en la ejecución del compilador
    if compilerErr != nil {
        log.Printf("❌ Error ejecutando compilador: %v", compilerErr)
        result := ExecuteResponse{
            Success: false,
            Output:  "Compiler execution failed: " + compilerErr.Error(),
            Errors:  analyzerResult.Errors,
            Symbols: mappedSymbols, // ← Usar símbolos mapeados
            Stats: ExecutionStats{
                CompileTime: "0ms",
                Lines:       strings.Count(req.Code, "\n") + 1,
                Size:        "0 bytes",
            },
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(result)
        return
    }
    
    // Parsear resultado del compilador
    var compilerResult struct {
        Success  bool   `json:"success"`
        Assembly string `json:"assembly"`
        Error    string `json:"error"`
        Stats    struct {
            CompileTime string `json:"compileTime"`
            Lines       int    `json:"lines"`
            Size        string `json:"size"`
        } `json:"stats"`
    }
    
    if err := json.Unmarshal(compilerOutput, &compilerResult); err != nil {
        log.Printf("❌ Error parseando salida del compilador: %v", err)
        log.Printf("📄 Salida del compilador: %s", string(compilerOutput))
        http.Error(w, "Error parsing compiler output: "+err.Error(), http.StatusInternalServerError)
        return
    }
    
    // 4. Devolver resultado final
    result := ExecuteResponse{
        Success: compilerResult.Success,
        Output:  compilerResult.Assembly,
        Errors:  analyzerResult.Errors,
        Symbols: mappedSymbols, // ← Usar símbolos mapeados
        Stats: ExecutionStats{
            CompileTime: compilerResult.Stats.CompileTime,
            Lines:       compilerResult.Stats.Lines,
            Size:        compilerResult.Stats.Size,
        },
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

// Endpoint para generar CST (solo SVG)
func (s *APIServer) generateCST(w http.ResponseWriter, r *http.Request) {
    var req CSTRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Code == "" {
        http.Error(w, "Code is required", http.StatusBadRequest)
        return
    }

    log.Printf("🌳 Generando CST para código de %d caracteres", len(req.Code))

    // Generar SVG usando el paquete CST
    svgTree := cst.CstReport(req.Code)
    
    if svgTree == "" {
        response := CSTResponse{
            Success: false,
            Error:   "No se pudo generar el árbol CST. El servicio ANTLR Lab puede estar temporalmente no disponible.",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
        return
    }

    // Devolver SVG puro
    response := CSTResponse{
        Success:     true,
        HTMLContent: svgTree, // SVG content
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// Agregar la ruta en setupRoutes()
func (s *APIServer) setupRoutes() *mux.Router {
    r := mux.NewRouter()
    
    // Endpoints principales
    r.HandleFunc("/execute", s.executeCode).Methods("POST")
    r.HandleFunc("/cst", s.generateCST).Methods("POST")
    r.HandleFunc("/compile-and-run", s.compileAndRun).Methods("POST") // ← Nueva ruta
    
    // Health check
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }).Methods("GET")
    
    return r
}

func (s *APIServer) compileAndRun(w http.ResponseWriter, r *http.Request) {
    var req CompileAndRunRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if req.Code == "" {
        http.Error(w, "Code is required", http.StatusBadRequest)
        return
    }

    log.Printf("🚀 Compilando y ejecutando código (%d caracteres)", len(req.Code))

    // Configurar headers para Server-Sent Events (streaming)
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // Función helper para enviar mensajes de streaming
    sendMessage := func(msgType, content string) {
        fmt.Fprintf(w, "data: {\"type\":\"%s\",\"content\":%q}\n\n", msgType, content)
        if flusher, ok := w.(http.Flusher); ok {
            flusher.Flush()
        }
    }

    sendMessage("info", "-- Iniciando compilación y ejecución...")
    time.Sleep(500 * time.Millisecond) // Pausa para efecto visual

    // 1. Compilar código Vlang a ARM64
    sendMessage("step", "- Paso 1: Compilando código Vlang a ARM64...")
    time.Sleep(300 * time.Millisecond)

    compilerCmd := exec.Command(s.compilerPath)
    compilerCmd.Stdin = strings.NewReader(req.Code)
    
    var stdout, stderr bytes.Buffer
    compilerCmd.Stdout = &stdout
    compilerCmd.Stderr = &stderr

    compilerErr := compilerCmd.Run()

    if compilerErr != nil {
        sendMessage("error", fmt.Sprintf("❌ Error compilando código: %v\n%s", compilerErr, stderr.String()))
        return
    }

    sendMessage("success", "Código Vlang compilado exitosamente")
    time.Sleep(200 * time.Millisecond)

    // 2. Parsear resultado
    var compilerResult struct {
        Success  bool   `json:"success"`
        Assembly string `json:"assembly"`
        Error    string `json:"error"`
    }

    if err := json.Unmarshal(stdout.Bytes(), &compilerResult); err != nil {
        sendMessage("error", fmt.Sprintf("❌ Error parseando salida del compilador: %v", err))
        return
    }

    if !compilerResult.Success {
        sendMessage("error", fmt.Sprintf("❌ Error en la compilación: %s", compilerResult.Error))
        return
    }

    // 3. Crear directorio temporal
    sendMessage("step", "- Paso 2: Creando directorio temporal...")
    tempDir, err := os.MkdirTemp("", "vlang_compile_*")
    if err != nil {
        sendMessage("error", fmt.Sprintf("❌ Error creando directorio temporal: %v", err))
        return
    }
    defer os.RemoveAll(tempDir)

    sendMessage("info", fmt.Sprintf("Directorio temporal: %s", tempDir))
    time.Sleep(200 * time.Millisecond)

    // 4. Guardar archivo .s
    sendMessage("step", "- Paso 3: Guardando código ARM64 en program.s...")
    asmFile := filepath.Join(tempDir, "program.s")
    if err := os.WriteFile(asmFile, []byte(compilerResult.Assembly), 0644); err != nil {
        sendMessage("error", fmt.Sprintf("❌ Error guardando archivo .s: %v", err))
        return
    }

    sendMessage("success", "Archivo program.s creado")
    time.Sleep(200 * time.Millisecond)

    // 5. Buscar script build.sh
    sendMessage("step", "- Paso 4: Buscando script build.sh...")
    buildScriptPaths := []string{
        "../assembler/build.sh",
        "./assembler/build.sh", 
        "../backend/assembler/build.sh",
        "./backend/assembler/build.sh",
    }
    
    var buildScriptPath string
    for _, path := range buildScriptPaths {
        if _, err := os.Stat(path); err == nil {
            absPath, _ := filepath.Abs(path)
            buildScriptPath = absPath
            break
        }
    }
    
    if buildScriptPath == "" {
        sendMessage("error", "Script build.sh no encontrado")
        return
    }

    sendMessage("success", fmt.Sprintf("Script encontrado: %s", buildScriptPath))
    time.Sleep(300 * time.Millisecond)

    // 6. Ejecutar ensamblado (simulando comandos paso a paso)
    sendMessage("step", "- Paso 5: Ejecutando ensamblado...")
    sendMessage("command", fmt.Sprintf("$ bash %s %s", filepath.Base(buildScriptPath), "program.s"))
    time.Sleep(500 * time.Millisecond)

    // Ejecutar el comando de ensamblado
    buildCmd := exec.Command("bash", buildScriptPath, asmFile)
    buildCmd.Dir = tempDir

    // Capturar salida línea por línea para streaming
    buildStdout, _ := buildCmd.StdoutPipe()
    buildStderr, _ := buildCmd.StderrPipe()

    if err := buildCmd.Start(); err != nil {
        sendMessage("error", fmt.Sprintf("Error iniciando build.sh: %v", err))
        return
    }

    // Leer stdout en tiempo real
    go func() {
        scanner := bufio.NewScanner(buildStdout)
        for scanner.Scan() {
            line := scanner.Text()
            if strings.Contains(line, "aarch64-linux-gnu-as") {
                sendMessage("command", "$ aarch64-linux-gnu-as -mcpu=cortex-a57 program.s -o program.o")
                time.Sleep(800 * time.Millisecond)
                sendMessage("success", "Ensamblado completado")
            } else if strings.Contains(line, "aarch64-linux-gnu-ld") {
                sendMessage("command", "$ aarch64-linux-gnu-ld program.o -o program")
                time.Sleep(600 * time.Millisecond)
                sendMessage("success", "Enlazado completado")
            } else if line != "" {
                sendMessage("output", line)
            }
        }
    }()

    // Leer stderr en tiempo real
    go func() {
        scanner := bufio.NewScanner(buildStderr)
        for scanner.Scan() {
            line := scanner.Text()
            if line != "" {
                sendMessage("warning", line)
            }
        }
    }()

    // Esperar a que termine el comando
    if err := buildCmd.Wait(); err != nil {
        sendMessage("error", fmt.Sprintf("❌ Error ejecutando build.sh: %v", err))
        return
    }

    time.Sleep(300 * time.Millisecond)

    // 7. Verificar ejecutable
    sendMessage("step", "- Paso 6: Verificando ejecutable generado...")
    exeFile := filepath.Join(tempDir, "program")
    if _, err := os.Stat(exeFile); err != nil {
        sendMessage("error", "❌ El ejecutable no fue generado correctamente")
        return
    }

    sendMessage("success", "Ejecutable 'program' generado exitosamente")
    time.Sleep(400 * time.Millisecond)

    // 8. Ejecutar con QEMU
    sendMessage("step", "- Paso 7: Ejecutando con QEMU...")
    sendMessage("command", "$ qemu-aarch64 ./program")
    time.Sleep(700 * time.Millisecond)

    qemuCmd := exec.Command("qemu-aarch64", exeFile)
    qemuCmd.Dir = tempDir
    
    // Capturar salida de QEMU en tiempo real
    qemuStdout, _ := qemuCmd.StdoutPipe()
    qemuStderr, _ := qemuCmd.StderrPipe()

    if err := qemuCmd.Start(); err != nil {
        sendMessage("error", fmt.Sprintf("❌ Error iniciando QEMU: %v", err))
        return
    }

    // Leer salida del programa
    go func() {
        scanner := bufio.NewScanner(qemuStdout)
        for scanner.Scan() {
            line := scanner.Text()
            sendMessage("program_output", line)
        }
    }()

    go func() {
        scanner := bufio.NewScanner(qemuStderr)
        for scanner.Scan() {
            line := scanner.Text()
            sendMessage("program_error", line)
        }
    }()

    qemuErr := qemuCmd.Wait()

    time.Sleep(200 * time.Millisecond)

    if qemuErr != nil {
        sendMessage("warning", fmt.Sprintf("Programa terminó con código: %v", qemuErr))
    } else {
        sendMessage("success", "Programa ejecutado exitosamente")
    }

    time.Sleep(300 * time.Millisecond)
    sendMessage("complete", "Compilación y ejecución completada")

    log.Printf("Ejecución completada para %d caracteres de código", len(req.Code))
}

func main() {
    server := NewAPIServer()
    
    router := server.setupRoutes()
    
    // Configurar CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
        AllowCredentials: true,
    })
    
    handler := c.Handler(router)
    
    port := "8080"
    log.Printf("🚀 Servidor API iniciado en puerto %s", port)
    log.Printf("🔗 Health check: http://localhost:%s/health", port)
    log.Printf("📡 CORS habilitado para: http://localhost:3000")
    
    if err := http.ListenAndServe(":"+port, handler); err != nil {
        log.Fatal("❌ Error iniciando servidor:", err)
    }
}