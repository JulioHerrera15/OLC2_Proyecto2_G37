package main

import (
    "bufio"
    "compiler/cst"
    "encoding/json"
    "fmt"
    "os"
    "strings"
)

type CSTResponse struct {
    Success bool   `json:"success"`
    CST     string `json:"cst"`
    Error   string `json:"error,omitempty"`
}

func readStdin() (string, error) {
    scanner := bufio.NewScanner(os.Stdin)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return "", err
    }
    return strings.Join(lines, "\n"), nil
}

func main() {
    // Leer código fuente desde stdin
    inputCode, err := readStdin()
    if err != nil {
        fmt.Printf("ERROR:Error leyendo entrada: %v\n", err)
        return
    }

    // Si no hay código, devolver error
    if strings.TrimSpace(inputCode) == "" {
        response := CSTResponse{
            Success: false,
            CST:     "",
            Error:   "No hay código para procesar",
        }
        jsonResponse, _ := json.Marshal(response)
        fmt.Printf("ERROR:%s\n", jsonResponse)
        return
    }

    fmt.Fprintf(os.Stderr, "🌐 Generando árbol CST con ANTLR Lab...\n")
    
    // Generar CST usando la función existente
    cstSVG := cst.CstReport(inputCode)
    
    var response CSTResponse
    if cstSVG == "" {
        fmt.Fprintf(os.Stderr, "❌ Error generando CST\n")
        response = CSTResponse{
            Success: false,
            CST:     "",
            Error:   "No se pudo generar el árbol CST",
        }
    } else {
        fmt.Fprintf(os.Stderr, "✅ Árbol CST generado exitosamente\n")
        response = CSTResponse{
            Success: true,
            CST:     cstSVG,
            Error:   "",
        }
    }

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        fmt.Printf("ERROR:Error al crear JSON: %v\n", err)
        return
    }

    fmt.Printf("SUCCESS:%s\n", jsonResponse)
}