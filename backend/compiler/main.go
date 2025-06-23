package main

import (
	c "backend/compiler/arm"
	"backend/compiler/visitor"
	parser "backend/parser"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/antlr4-go/antlr/v4"
)

// Estructura para la respuesta del compilador
type CompilerResponse struct {
    Success    bool   `json:"success"`
    Assembly   string `json:"assembly"`
    Error      string `json:"error,omitempty"`
    Stats      Stats  `json:"stats"`
}

type Stats struct {
    CompileTime string `json:"compileTime"`
    Lines       int    `json:"lines"`
    Size        string `json:"size"`
}

func main() {
    // Leer código desde stdin
    input, err := readStdin()
    if err != nil {
        sendError("Error reading input: " + err.Error())
        return
    }

    // Compilar
    startTime := time.Now()
    assembly, err := compile(input)
    compileTime := time.Since(startTime)

    if err != nil {
        sendError("Compilation error: " + err.Error())
        return
    }

    // Enviar respuesta exitosa
    response := CompilerResponse{
        Success:  true,
        Assembly: assembly,
        Stats: Stats{
            CompileTime: compileTime.String(),
            Lines:       strings.Count(input, "\n") + 1,
            Size:        fmt.Sprintf("%d bytes", len(assembly)),
        },
    }

    json.NewEncoder(os.Stdout).Encode(response)
}

func compile(code string) (string, error) {
    // Crear input stream
    input := antlr.NewInputStream(code)
    
    // Crear lexer
    lexer := parser.NewLanguageLexer(input)
    
    // Crear token stream
    stream := antlr.NewCommonTokenStream(lexer, 0)
    
    // Crear parser
    p := parser.NewLanguageParser(stream)
    
    // Parsear desde la regla 'program'
    tree := p.Program()
    
    // Crear visitor del compilador
    compilerVisitor := visitor.NewVisitor()
    
    // Recorrer el árbol
	compilerVisitor.Visit(tree)

	c.EndProgram()

    // Generar código ensamblador
    assembly := c.ToString()
    
    return assembly, nil
}

func readStdin() (string, error) {
    var result strings.Builder
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        result.WriteString(scanner.Text())
        result.WriteString("\n")
    }
    
    if err := scanner.Err(); err != nil && err != io.EOF {
        return "", err
    }
    
    return result.String(), nil
}

func sendError(message string) {
    response := CompilerResponse{
        Success: false,
        Error:   message,
    }
    json.NewEncoder(os.Stdout).Encode(response)
}