package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
    "backend/analyzer/cst"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "bytes"
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
    
    // Parsear resultado del analyzer
    var analyzerResult struct {
        Success bool          `json:"success"`
        Output  string        `json:"output"`
        Errors  []ErrorDetail `json:"errors"`
        Symbols []SymbolDetail `json:"symbols"`
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
    
    // 2. ✅ SI HAY ERRORES: No compilar, solo devolver errores (MEJORADO)
    if !analyzerResult.Success || len(analyzerResult.Errors) > 0 {
        log.Printf("⚠️ Se encontraron %d errores, no se compilará", len(analyzerResult.Errors))
        
        result := ExecuteResponse{
            Success: false,
            Output:  "Se encontraron errores en el código. Revisa la tabla de errores.",
            Errors:  analyzerResult.Errors,   // ← Importante: Los errores se pasan
            Symbols: analyzerResult.Symbols,  // ← Los símbolos también
            Stats: ExecutionStats{
                CompileTime: "0ms (no compilado)",
                Lines:       strings.Count(req.Code, "\n") + 1,
                Size:        "0 bytes",
            },
            Optimized: false,
        }
        
        // ✅ Devolver 200 OK (no 500) con los errores
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(result)
        return
    }

    // 3. SI NO HAY ERRORES: Ejecutar compilador (resto del código igual)
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
            Symbols: analyzerResult.Symbols,
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
    
    // Parsear resultado del compilador (igual que antes)
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
        Output:  compilerResult.Assembly, // Código ARM64 generado
        Errors:  analyzerResult.Errors,
        Symbols: analyzerResult.Symbols,
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

// Configurar rutas (simplificado)
func (s *APIServer) setupRoutes() *mux.Router {
    r := mux.NewRouter()
    
    // Endpoints principales
    r.HandleFunc("/execute", s.executeCode).Methods("POST")
    r.HandleFunc("/cst", s.generateCST).Methods("POST")
    
    // Health check
    r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{
            "status":   "ok",
            "analyzer": s.analyzerPath,
            "time":     time.Now().Format(time.RFC3339),
        })
    }).Methods("GET")
    
    return r
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