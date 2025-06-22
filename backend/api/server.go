package main

import (
    "encoding/json"
    "fmt"
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
    compilerPath string
}

func NewAPIServer() *APIServer {
    server := &APIServer{}
    server.findCompilerExecutable()
    return server
}

// Buscar el ejecutable del compilador
func (s *APIServer) findCompilerExecutable() {
    possiblePaths := []string{
        "../analyzer/analyzer",
        "./analyzer/analyzer",
        "../backend/analyzer/analyzer",
        "./backend/analyzer/analyzer",
        "../../backend/analyzer/analyzer",
        "../analyzer",
        "./analyzer",
    }

    for _, path := range possiblePaths {
        if _, err := os.Stat(path); err == nil {
            absPath, _ := filepath.Abs(path)
            s.compilerPath = absPath
            log.Printf("✅ Analizador encontrado en: %s", absPath)
            return
        }
    }

    log.Fatal("❌ No se pudo encontrar el ejecutable del analizador")
}

// Endpoint para ejecutar código (principal)
func (s *APIServer) executeCode(w http.ResponseWriter, r *http.Request) {
    var req ExecuteRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    log.Printf("🔄 Ejecutando código (%d caracteres)", len(req.Code))

    // Ejecutar compilador
    startTime := time.Now()
    cmd := exec.Command(s.compilerPath)
    cmd.Stdin = strings.NewReader(req.Code)
    
    output, err := cmd.CombinedOutput()
    compileTime := time.Since(startTime)

    // Parsear resultado
    result := s.parseCompilerOutput(string(output), err, req.Code, compileTime)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

// Parsear salida del compilador
func (s *APIServer) parseCompilerOutput(output string, execErr error, code string, compileTime time.Duration) ExecuteResponse {
    result := ExecuteResponse{
        Success:   false,
        Output:    "",
        Errors:    []ErrorDetail{},
        Symbols:   []SymbolDetail{},
        Optimized: true,
        Stats: ExecutionStats{
            CompileTime: fmt.Sprintf("%dms", compileTime.Milliseconds()),
            Lines:       len(strings.Split(code, "\n")),
            Size:        fmt.Sprintf("%.1f KB", float64(len(code))/1024),
        },
    }

    // Si hay error de ejecución del proceso
    if execErr != nil {
        result.Success = false
        result.Output = output
        result.Errors = append(result.Errors, ErrorDetail{
            Type:     "EXECUTION_ERROR",
            Message:  fmt.Sprintf("Error ejecutando compilador: %v", execErr),
            Line:     1,
            Column:   1,
            Severity: "high",
        })
        return result
    }

    // Parsear según el formato del compilador
    switch {
    case strings.HasPrefix(output, "SUCCESS:"):
        jsonStr := strings.TrimPrefix(output, "SUCCESS:")
        return s.parseCompilerJSON(jsonStr, result, true)
        
    case strings.HasPrefix(output, "ERROR_REPORT:"):
        jsonStr := strings.TrimPrefix(output, "ERROR_REPORT:")
        return s.parseCompilerJSON(jsonStr, result, false)
        
    default:
        // Salida no reconocida - mostrar output crudo
        result.Success = false
        result.Output = output
        result.Errors = append(result.Errors, ErrorDetail{
            Type:     "PARSE_ERROR",
            Message:  "Formato de salida no reconocido del compilador",
            Line:     1,
            Column:   1,
            Severity: "medium",
        })
        return result
    }
}

// Parsear JSON del compilador
func (s *APIServer) parseCompilerJSON(jsonStr string, result ExecuteResponse, _ bool) ExecuteResponse {
    var compilerResp CompilerResponse
    
    if err := json.Unmarshal([]byte(jsonStr), &compilerResp); err != nil {
        result.Success = false
        result.Output = jsonStr
        result.Errors = append(result.Errors, ErrorDetail{
            Type:     "JSON_PARSE_ERROR",
            Message:  fmt.Sprintf("Error parseando JSON: %v", err),
            Line:     1,
            Column:   1,
            Severity: "high",
        })
        return result
    }

    // Mapear respuesta del compilador a formato del API
    result.Success = compilerResp.Success
    result.Output = compilerResp.Output

    // Mapear errores
    for _, err := range compilerResp.Errors {
        result.Errors = append(result.Errors, ErrorDetail{
            Type:     err.Type,
            Message:  err.Message,
            Line:     err.Line,
            Column:   err.Column,
            Severity: "high",
        })
    }

    // Mapear símbolos
    for _, sym := range compilerResp.Symbols {
        result.Symbols = append(result.Symbols, SymbolDetail{
            Name:     sym.ID,
            Type:     sym.DataType,
            Scope:    sym.Scope,
            Line:     sym.Line,
            Column:   sym.Column,
        })
    }

    return result
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
            "compiler": s.compilerPath,
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