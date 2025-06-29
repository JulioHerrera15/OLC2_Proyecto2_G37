package cst

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "runtime"
)

type CSTResponse struct {
    SVGTree string `json:"svgtree"`
}

func ReadFile(filename string) string {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Printf("❌ Error leyendo archivo %s: %v\n", filename, err)
        return ""
    }
    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
        fmt.Printf("❌ Error leyendo contenido: %v\n", err)
        return ""
    }
    return string(content)
}

func CstReport(input string) string {
    // Obtener la ruta del archivo de gramática de manera más robusta
    _, filename, _, _ := runtime.Caller(0)
    
    // Buscar hacia arriba hasta encontrar la carpeta backend
    path := filepath.Dir(filename)
    for filepath.Base(path) != "backend" && path != "/" {
        path = filepath.Dir(path)
    }
    
    if filepath.Base(path) != "backend" {
        fmt.Printf("❌ No se pudo encontrar la carpeta backend desde: %s\n", filename)
        return ""
    }
    
    grammarPath := filepath.Join(path, "parser", "Language.g4")
    
    // Verificar que el archivo existe
    if _, err := os.Stat(grammarPath); os.IsNotExist(err) {
        fmt.Printf("❌ Archivo de gramática no encontrado: %s\n", grammarPath)
        return ""
    }
    
    fmt.Fprintf(os.Stderr, "📁 Usando gramática: %s\n", grammarPath)
    
    // Resto del código igual...
    
    // Leer la gramática
    grammarContent := ReadFile(grammarPath)
    if grammarContent == "" {
        fmt.Println("❌ No se pudo leer el archivo de gramática")
        return ""
    }
    
    // Convertir a JSON
    grammarJSON, err := json.Marshal(grammarContent)
    if err != nil {
        fmt.Printf("❌ Error convirtiendo gramática a JSON: %v\n", err)
        return ""
    }
    
    inputJSON, err := json.Marshal(input)
    if err != nil {
        fmt.Printf("❌ Error convirtiendo input a JSON: %v\n", err)
        return ""
    }
    
    // Crear payload para la API
    payload := []byte(fmt.Sprintf(`{
        "grammar": %s,
        "input": %s,
        "lexgrammar": %s,
        "start": "%s"
    }`, grammarJSON, inputJSON, grammarJSON, "program"))
    
    // Hacer petición a la API
    fmt.Fprintf(os.Stderr, "🌐 Enviando petición a ANTLR Lab...\n")
    
    req, err := http.NewRequest("POST", "http://lab.antlr.org/parse/", bytes.NewBuffer(payload))
    if err != nil {
        fmt.Printf("❌ Error creando petición: %v\n", err)
        return ""
    }
    
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "VLang-Compiler/1.0")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("❌ Error enviando petición: %v\n", err)
        return ""
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        fmt.Printf("❌ Error HTTP: %d %s\n", resp.StatusCode, resp.Status)
        return ""
    }
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("❌ Error leyendo respuesta: %v\n", err)
        return ""
    }
    
    // Parsear respuesta JSON
    var data map[string]interface{}
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Printf("❌ Error parseando JSON: %v\n", err)
        return ""
    }
    
    // Verificar que la respuesta tiene la estructura esperada
    result, ok := data["result"].(map[string]interface{})
    if !ok {
        fmt.Println("❌ Respuesta de API no tiene formato esperado")
        return ""
    }
    
    svgTree, ok := result["svgtree"].(string)
    if !ok {
        fmt.Println("❌ No se encontró el árbol SVG en la respuesta")
        return ""
    }
    
    fmt.Fprintf(os.Stderr, "✅ Árbol CST generado exitosamente\n")
    return svgTree
}