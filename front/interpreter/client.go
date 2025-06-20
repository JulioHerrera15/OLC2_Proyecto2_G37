package interpreter

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "strings"
    "sync"
    "time"
)

// Estructuras para parsear la respuesta JSON del backend
type SuccessResponse struct {
    Success bool           `json:"success"`
    Output  string         `json:"output"`
    Errors  []ErrorDetail  `json:"errors"`
    Symbols []SymbolDetail `json:"symbols"`
    Stats   ExecutionStats `json:"stats"`
}

type ErrorResponse struct {
    Success bool           `json:"success"`
    Output  string         `json:"output"`
    Errors  []ErrorDetail  `json:"errors"`
    Symbols []SymbolDetail `json:"symbols"` 
    Stats   ExecutionStats `json:"stats"`
}

type ErrorDetail struct {
    Line    int    `json:"line"`
    Column  int    `json:"column"`
    Type    string `json:"type"`
    Message string `json:"message"`
}

type SymbolDetail struct {
    ID         string `json:"id"`
    SymbolType string `json:"symbolType"`
    DataType   string `json:"dataType"`
    Scope      string `json:"scope"`
    Line       int    `json:"line"`
    Column     int    `json:"column"`
}

// AGREGAR: Estructura para estadísticas del backend
type ExecutionStats struct {
    ExecutionTime int64 `json:"executionTime"` // en millisegundos
    CodeSize      int   `json:"codeSize"`      // tamaño del código
    ErrorCount    int   `json:"errorCount"`    // número de errores
    SymbolCount   int   `json:"symbolCount"`   // número de símbolos
}

// Resultado estructurado OPTIMIZADO
type ExecutionResult struct {
    Success bool          
    Output  string
    Errors  []ErrorDetail
    Symbols []SymbolDetail
    Stats   ExecutionStats
}

// Client OPTIMIZADO maneja la comunicación con el backend del compilador
type Client struct {
    CompilerPath string
    buffer       *bytes.Buffer  
    tempDir      string        
    mu           sync.Mutex    
}

// OPTIMIZAR: Constructor con inicialización de recursos
func NewClient(compilerPath string) *Client {
    // Crear directorio temporal reutilizable
    tempDir, err := os.MkdirTemp("", "v_compiler_*")
    if err != nil {
        fmt.Printf("Warning: no se pudo crear directorio temporal: %v\n", err)
        tempDir = ""
    }
    
    return &Client{
        CompilerPath: compilerPath,
        buffer:       &bytes.Buffer{},
        tempDir:      tempDir,
        mu:           sync.Mutex{},
    }
}

// Método de limpieza
func (c *Client) Cleanup() error {
    if c.tempDir != "" {
        return os.RemoveAll(c.tempDir)
    }
    return nil
}

// OPTIMIZAR: Método Execute con mejor manejo de memoria y rendimiento
func (c *Client) Execute(code string) (*ExecutionResult, error) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    start := time.Now() // Medir tiempo de frontend
    
    // OPTIMIZAR: Verificación temprana de tamaño
    if len(code) == 0 {
        return &ExecutionResult{
            Success: true,
            Output:  "⚠️ No hay código para ejecutar",
            Errors:  []ErrorDetail{},
            Symbols: []SymbolDetail{},
            Stats: ExecutionStats{
                ExecutionTime: time.Since(start).Milliseconds(),
                CodeSize:      0,
                ErrorCount:    0,
                SymbolCount:   0,
            },
        }, nil
    }
    
    // OPTIMIZAR: Límite de tamaño para prevenir problemas de memoria
    const maxCodeSize = 1024 * 1024 // 1MB
    if len(code) > maxCodeSize {
        return &ExecutionResult{
            Success: false,
            Output:  "",
            Errors: []ErrorDetail{
                {
                    Line:    1,
                    Column:  1,
                    Type:    "MEMORY_ERROR",
                    Message: fmt.Sprintf("Código muy grande (%d bytes). Máximo permitido: %d bytes", len(code), maxCodeSize),
                },
            },
            Symbols: []SymbolDetail{},
            Stats: ExecutionStats{
                ExecutionTime: time.Since(start).Milliseconds(),
                CodeSize:      len(code),
                ErrorCount:    1,
                SymbolCount:   0,
            },
        }, nil
    }

    // OPTIMIZAR: Reutilizar buffer existente
    c.buffer.Reset()
    c.buffer.Grow(len(code) + 1024) // Pre-alocar memoria estimada

    // OPTIMIZAR: Configurar comando con límites de recursos y timeout
    ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
    defer cancel()
    
    cmd := exec.CommandContext(ctx, c.CompilerPath)
    cmd.Stdin = strings.NewReader(code)
    cmd.Stdout = c.buffer
    cmd.Stderr = c.buffer
    
    // OPTIMIZAR: Variables de entorno para mejor rendimiento
    cmd.Env = append(os.Environ(),
        "GOGC=50",              // GC más agresivo para liberar memoria rápido
        "GOMAXPROCS=2",         // Limitar a 2 CPUs para no saturar el sistema
        "GODEBUG=gctrace=0",    // Deshabilitar trazas de GC
    )

    // Ejecutar comando
    err := cmd.Run()
    frontendTime := time.Since(start).Milliseconds()
    
    if err != nil {
        if ctx.Err() == context.DeadlineExceeded {
            return &ExecutionResult{
                Success: false,
                Output:  "",
                Errors: []ErrorDetail{
                    {
                        Line:    1,
                        Column:  1,
                        Type:    "TIMEOUT_ERROR",
                        Message: "El compilador tardó más de 60 segundos en responder",
                    },
                },
                Symbols: []SymbolDetail{},
                Stats: ExecutionStats{
                    ExecutionTime: frontendTime,
                    CodeSize:      len(code),
                    ErrorCount:    1,
                    SymbolCount:   0,
                },
            }, nil
        }
        
        return &ExecutionResult{
            Success: false,
            Output:  c.buffer.String(),
            Errors: []ErrorDetail{
                {
                    Line:    1,
                    Column:  1,
                    Type:    "EXECUTION_ERROR",
                    Message: fmt.Sprintf("Error ejecutando compilador: %v", err),
                },
            },
            Symbols: []SymbolDetail{},
            Stats: ExecutionStats{
                ExecutionTime: frontendTime,
                CodeSize:      len(code),
                ErrorCount:    1,
                SymbolCount:   0,
            },
        }, nil
    }

    // OPTIMIZAR: Procesar salida eficientemente
    output := c.buffer.String()
    return c.parseOutputOptimized(output, frontendTime, len(code))
}

// OPTIMIZAR: Parser de salida más eficiente con mejor manejo de errores
func (c *Client) parseOutputOptimized(output string, frontendTime int64, codeSize int) (*ExecutionResult, error) {
    // OPTIMIZAR: Pre-alocar slices con capacidad estimada
    const estimatedErrors = 10
    const estimatedSymbols = 50
    
    result := &ExecutionResult{
        Success: false, // Por defecto false, se cambia según el resultado
        Errors:  make([]ErrorDetail, 0, estimatedErrors),
        Symbols: make([]SymbolDetail, 0, estimatedSymbols),
        Stats: ExecutionStats{
            ExecutionTime: frontendTime, // Tiempo del frontend por defecto
            CodeSize:      codeSize,
            ErrorCount:    0,
            SymbolCount:   0,
        },
    }
    
    // Buscar prefijos de manera eficiente
    switch {
    case strings.HasPrefix(output, "SUCCESS:"):
        result.Success = true
        jsonStr := strings.TrimPrefix(output, "SUCCESS:")
        return c.parseJSONOptimized(jsonStr, result)
        
    case strings.HasPrefix(output, "ERROR_REPORT:"):
        result.Success = false
        jsonStr := strings.TrimPrefix(output, "ERROR_REPORT:")
        return c.parseJSONOptimized(jsonStr, result)
        
    default:
        // Salida no reconocida - tratar como error pero mostrar output
        result.Success = false
        result.Output = output
        result.Errors = append(result.Errors, ErrorDetail{
            Line:    1,
            Column:  1,
            Type:    "PARSE_ERROR",
            Message: "Formato de salida no reconocido del compilador",
        })
        result.Stats.ErrorCount = 1
        return result, nil
    }
}

// OPTIMIZAR: Parser JSON más robusto con mejor manejo de errores
func (c *Client) parseJSONOptimized(jsonStr string, result *ExecutionResult) (*ExecutionResult, error) {
    // OPTIMIZAR: Decodificar directo a estructura conocida
    var response struct {
        Success bool `json:"success"`
        Output  string `json:"output"`
        Errors  []struct {
            Line    int    `json:"line"`
            Column  int    `json:"column"`
            Type    string `json:"type"`
            Message string `json:"message"`
        } `json:"errors"`
        Symbols []struct {
            ID         string `json:"id"`
            SymbolType string `json:"symbolType"`
            DataType   string `json:"dataType"`
            Scope      string `json:"scope"`
            Line       int    `json:"line"`
            Column     int    `json:"column"`
        } `json:"symbols"`
        Stats struct {
            ExecutionTime int64 `json:"executionTime"`
            CodeSize      int   `json:"codeSize"`
            ErrorCount    int   `json:"errorCount"`
            SymbolCount   int   `json:"symbolCount"`
        } `json:"stats"`
    }
    
    if err := json.Unmarshal([]byte(jsonStr), &response); err != nil {
        // Si el JSON está malformado, mostrar output crudo
        result.Success = false
        result.Output = jsonStr
        result.Errors = append(result.Errors, ErrorDetail{
            Line:    1,
            Column:  1,
            Type:    "JSON_PARSE_ERROR",
            Message: fmt.Sprintf("Error parseando JSON del backend: %v", err),
        })
        result.Stats.ErrorCount = 1
        return result, nil
    }
    
    // OPTIMIZAR: Copiar datos eficientemente
    result.Success = response.Success
    result.Output = response.Output
    
    // OPTIMIZAR: Pre-alocar con tamaño exacto
    result.Errors = make([]ErrorDetail, len(response.Errors))
    for i, err := range response.Errors {
        result.Errors[i] = ErrorDetail{
            Line:    err.Line,
            Column:  err.Column,
            Type:    err.Type,
            Message: err.Message,
        }
    }
    
    result.Symbols = make([]SymbolDetail, len(response.Symbols))
    for i, sym := range response.Symbols {
        result.Symbols[i] = SymbolDetail{
            ID:         sym.ID,
            SymbolType: sym.SymbolType,
            DataType:   sym.DataType,
            Scope:      sym.Scope,
            Line:       sym.Line,
            Column:     sym.Column,
        }
    }
    
    // OPTIMIZAR: Usar estadísticas del backend si están disponibles
    if response.Stats.ExecutionTime > 0 {
        result.Stats = ExecutionStats{
            ExecutionTime: response.Stats.ExecutionTime,
            CodeSize:      response.Stats.CodeSize,
            ErrorCount:    response.Stats.ErrorCount,
            SymbolCount:   response.Stats.SymbolCount,
        }
    } else {
        // Fallback a estadísticas calculadas localmente
        result.Stats.ErrorCount = len(result.Errors)
        result.Stats.SymbolCount = len(result.Symbols)
    }
    
    return result, nil
}