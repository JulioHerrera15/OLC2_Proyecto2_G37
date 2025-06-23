package main

import (
    "bufio"
    "backend/analyzer/errors"
    parser "backend/parser"
    "backend/analyzer/repl"
    "encoding/json"
    "fmt"
    "io" 
    "os"
    "strings"
    "time"

    "github.com/antlr4-go/antlr/v4"
)

// Estructura para devolver errores en formato JSON
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

// Estructura para símbolos
type SymbolDetail struct {
    ID         string `json:"id"`
    SymbolType string `json:"symbolType"`
    DataType   string `json:"dataType"`
    Scope      string `json:"scope"`
    Line       int    `json:"line"`
    Column     int    `json:"column"`
}

// Estructura para estadísticas
type ExecutionStats struct {
    ExecutionTime int64 `json:"executionTime"` 
    CodeSize      int   `json:"codeSize"`      
    ErrorCount    int   `json:"errorCount"`    
    SymbolCount   int   `json:"symbolCount"`   
}

// OPTIMIZAR: Función de lectura más eficiente
func readStdin() (string, error) {
    const maxInputSize = 10 * 1024 * 1024 // 10MB límite
    
    // Buffer más eficiente
    reader := bufio.NewReaderSize(os.Stdin, 64*1024) // 64KB buffer
    
    // Leer con límite
    limited := io.LimitReader(reader, maxInputSize)
    inputBytes, err := io.ReadAll(limited)
    if err != nil {
        return "", err
    }
    
    return string(inputBytes), nil
}

func main() {
    start := time.Now() // AGREGAR: Medir tiempo
    
    // Leer código fuente desde stdin (OPTIMIZADO)
    inputCode, err := readStdin()
    if err != nil {
        fmt.Println("Error leyendo entrada:", err)
        return
    }

    // OPTIMIZAR: Verificación de código vacío con respuesta
    trimmedCode := strings.TrimSpace(inputCode)
    if trimmedCode == "" {
        response := ErrorResponse{
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
        }
        jsonResponse, _ := json.Marshal(response)
        fmt.Print("SUCCESS:" + string(jsonResponse))
        return
    }

    // Verificar tamaño del código
    if len(inputCode) > 100000 { // 100KB límite
        response := ErrorResponse{
            Success: false,
            Output:  "",
            Errors: []ErrorDetail{
                {
                    Line:    1,
                    Column:  1,
                    Type:    "SIZE_ERROR",
                    Message: fmt.Sprintf("Código muy grande (%d caracteres). Límite: 100KB", len(inputCode)),
                },
            },
            Symbols: []SymbolDetail{},
            Stats: ExecutionStats{
                ExecutionTime: time.Since(start).Milliseconds(),
                CodeSize:      len(inputCode),
                ErrorCount:    1,
                SymbolCount:   0,
            },
        }
        jsonResponse, _ := json.Marshal(response)
        fmt.Print("SUCCESS:" + string(jsonResponse))
        return
    }

    // 1. Análisis Léxico
    lexicalErrorListener := errors.NewLexicalErrorListener()
    lexer := parser.NewLanguageLexer(antlr.NewInputStream(inputCode))
    lexer.RemoveErrorListeners()
    lexer.AddErrorListener(lexicalErrorListener)

    // 2. Tokens
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

    // 3. Parser + errores sintácticos
    parser := parser.NewLanguageParser(stream)
    parser.BuildParseTrees = true

    syntaxErrorListener := errors.NewSyntaxErrorListener(lexicalErrorListener.ErrorTable)
    parser.RemoveErrorListeners()
    parser.SetErrorHandler(errors.NewCustomErrorStrategy())
    parser.AddErrorListener(syntaxErrorListener)

    // 4. Generar árbol sintáctico
    tree := parser.Program()

    // 5. Ejecutar el código usando el visitor
    console := repl.NewConsole()
    visitor := repl.NewReplVisitor()
    visitor.Console = console
    visitor.ErrorTable = lexicalErrorListener.ErrorTable

    // SIEMPRE intentar ejecutar, incluso con errores
    visitor.Visit(tree)

    // 6. Capturar símbolos del visitor
    var symbolDetails []SymbolDetail
    for _, symbol := range visitor.SymbolsTable.Symbols {
        symbolDetails = append(symbolDetails, SymbolDetail{
            ID:         symbol.ID,
            SymbolType: symbol.SymbolType,
            DataType:   symbol.DataType,
            Scope:      symbol.Scope,
            Line:       symbol.Line,
            Column:     symbol.Column,
        })
    }

    // AGREGAR: Calcular estadísticas
    executionTime := time.Since(start).Milliseconds()
    stats := ExecutionStats{
        ExecutionTime: executionTime,
        CodeSize:      len(inputCode),
        ErrorCount:    len(lexicalErrorListener.ErrorTable.Errors),
        SymbolCount:   len(symbolDetails),
    }

    // 7. Devolver resultado
    output := console.GetOutput()
    
    if len(lexicalErrorListener.ErrorTable.Errors) > 0 {
        // HAY ERRORES: devolver errores + salida + símbolos + STATS
        var errorDetails []ErrorDetail
        for _, err := range lexicalErrorListener.ErrorTable.Errors {
            errorDetails = append(errorDetails, ErrorDetail{
                Line:    err.Line,
                Column:  err.Column,
                Type:    err.Type,
                Message: err.Msg,
            })
        }

        response := ErrorResponse{
            Success: false,
            Output:  output,
            Errors:  errorDetails,
            Symbols: symbolDetails,
            Stats:   stats,
        }

        jsonResponse, _ := json.Marshal(response)
        fmt.Print("ERROR_REPORT:" + string(jsonResponse))
    } else {
        // SIN ERRORES: comportamiento normal con símbolos + STATS
        response := ErrorResponse{
            Success: true,
            Output:  output,
            Errors:  []ErrorDetail{},
            Symbols: symbolDetails,
            Stats:   stats,
        }

        jsonResponse, _ := json.Marshal(response)
        fmt.Print("SUCCESS:" + string(jsonResponse))
    }
}