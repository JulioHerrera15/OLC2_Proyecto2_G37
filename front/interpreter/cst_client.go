package interpreter

import (
    "bytes"
    "encoding/json"
    "io"
    "os/exec"
    "strings"
)

type CSTResponse struct {
    Success bool   `json:"success"`
    CST     string `json:"cst"`
    Error   string `json:"error,omitempty"`
}

type CSTResult struct {
    CST   string
    Error string
}

type CSTClient struct {
    GeneratorPath string
}

func NewCSTClient(generatorPath string) *CSTClient {
    return &CSTClient{
        GeneratorPath: generatorPath,
    }
}

func (c *CSTClient) GenerateCST(code string) (*CSTResult, error) {
    cmd := exec.Command(c.GeneratorPath)

    stdin, err := cmd.StdinPipe()
    if err != nil {
        return nil, err
    }

    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    if err := cmd.Start(); err != nil {
        return nil, err
    }

    io.WriteString(stdin, code)
    stdin.Close()

    if err := cmd.Wait(); err != nil {
        return &CSTResult{
            CST:   "",
            Error: stderr.String(),
        }, nil
    }

    result := stdout.String()
    if result == "" && stderr.Len() > 0 {
        result = stderr.String()
    }

    if strings.HasPrefix(result, "SUCCESS:") {
        jsonData := strings.TrimPrefix(result, "SUCCESS:")
        var response CSTResponse
        if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
            return &CSTResult{
                CST:   "",
                Error: "Error parseando respuesta",
            }, nil
        }

        return &CSTResult{
            CST:   response.CST,
            Error: response.Error,
        }, nil

    } else if strings.HasPrefix(result, "ERROR:") {
        jsonData := strings.TrimPrefix(result, "ERROR:")
        var response CSTResponse
        if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
            return &CSTResult{
                CST:   "",
                Error: "Error parseando respuesta de error",
            }, nil
        }

        return &CSTResult{
            CST:   "",
            Error: response.Error,
        }, nil

    } else {
        return &CSTResult{
            CST:   "",
            Error: "Respuesta inesperada del generador",
        }, nil
    }
}