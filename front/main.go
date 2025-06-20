package main

import (
    "fmt"
    "front/interpreter"
    "image/color"
    "io"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
    "crypto/sha256"
    "runtime"
    "context"
    "sync"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
)

// Estructuras para manejar errores
type ErrorResponse struct {
	Success bool          `json:"success"`
	Output  string        `json:"output"`
	Errors  []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Variables globales para almacenar errores y símbolos
var currentErrors []interpreter.ErrorDetail
var currentSymbols []interpreter.SymbolDetail
var cstClient *interpreter.CSTClient
var currentCST string  
var lastExecutedCode string 

var (
    executionSemaphore = make(chan struct{}, 1) // Solo 1 ejecución a la vez
    executionMutex     sync.Mutex               // Proteger variables compartidas
    isExecuting        = false                  // Estado de ejecución
)

// AGREGAR: Sistema de cache DESPUÉS de las variables globales existentes
type ReportCache struct {
    mu           sync.RWMutex
    lastCodeHash string
    lastCode     string
    errorReport  string
    symbolReport string
    cstReport    string
    timestamp    time.Time
}

var (
    reportCache = &ReportCache{}
    
    // AGREGAR: Pool de workers para CST
    cstWorkerPool = make(chan struct{}, 2) // Máximo 2 CST simultáneos
)

// AGREGAR: Función para calcular hash del código
func calculateCodeHash(code string) string {
    if code == "" {
        return "empty"
    }
    hash := sha256.Sum256([]byte(code))
    return fmt.Sprintf("%x", hash[:8]) // Solo usar 8 bytes para eficiencia
}

func (r *ReportCache) getErrorReport(codeHash string) (string, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    if r.lastCodeHash == codeHash && 
       time.Since(r.timestamp) < 5*time.Minute && 
       r.errorReport != "" {
        return r.errorReport, true
    }
    return "", false
}

func (r *ReportCache) updateErrorReport(codeHash, code, errorReport string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    r.lastCodeHash = codeHash
    r.lastCode = code
    r.errorReport = errorReport
    r.timestamp = time.Now()
}

func (r *ReportCache) getCSTReport(codeHash string) (string, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    if r.lastCodeHash == codeHash && 
       time.Since(r.timestamp) < 10*time.Minute &&
       r.cstReport != "" {
        return r.cstReport, true
    }
    return "", false
}

func (r *ReportCache) updateCSTReport(codeHash, code, cstReport string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    r.lastCodeHash = codeHash
    r.lastCode = code
    r.cstReport = cstReport
    r.timestamp = time.Now()
}

func (r *ReportCache) getSymbolReport(codeHash string) (string, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    
    if r.lastCodeHash == codeHash && 
       time.Since(r.timestamp) < 5*time.Minute &&
       r.symbolReport != "" {
        return r.symbolReport, true
    }
    return "", false
}

func (r *ReportCache) updateSymbolReport(codeHash, code, symbolReport string) {
    r.mu.Lock()
    defer r.mu.Unlock()
    
    r.lastCodeHash = codeHash
    r.lastCode = code
    r.symbolReport = symbolReport
    r.timestamp = time.Now()
}

func processExecutionResult(result *interpreter.ExecutionResult, outputArea *ConsoleOutput, updateStatus func(string)) {
    // Limpiar datos anteriores
    currentErrors = result.Errors
    currentSymbols = result.Symbols

    // CONSTRUCCIÓN EFICIENTE DE OUTPUT
    var output strings.Builder
    output.Grow(len(result.Output) + 1000) // Pre-alocar memoria
    
    if result.Output != "" {
        output.WriteString(result.Output)
    } else {
        output.WriteString("✅ Código ejecutado exitosamente (sin salida)")
    }
    
    // AGREGAR INFORMACIÓN DETALLADA
    if len(result.Errors) > 0 || len(result.Symbols) > 0 {
        output.WriteString("\n" + strings.Repeat("=", 60))
        output.WriteString("\n📊 RESUMEN DE EJECUCIÓN")
        output.WriteString("\n" + strings.Repeat("=", 60))
        
        if len(result.Errors) > 0 {
            output.WriteString(fmt.Sprintf("\n❌ %d errores encontrados", len(result.Errors)))
            output.WriteString("\n💡 Consulta el 'Reporte de Errores' para detalles")
        } else {
            output.WriteString("\n✅ Sin errores detectados")
        }
        
        if len(result.Symbols) > 0 {
            output.WriteString(fmt.Sprintf("\n📋 %d símbolos en tabla", len(result.Symbols)))
            output.WriteString("\n💡 Consulta la 'Tabla de Símbolos' para detalles")
        }
        
        // MOSTRAR ESTADÍSTICAS SI ESTÁN DISPONIBLES
        if result.Stats.ExecutionTime > 0 {
            output.WriteString(fmt.Sprintf("\n⚡ Tiempo de ejecución: %dms", result.Stats.ExecutionTime))
            output.WriteString(fmt.Sprintf("\n📏 Tamaño de código: %d bytes", result.Stats.CodeSize))
        }
        
        output.WriteString("\n\n⚡ Ejecución completada exitosamente")
    }
    
    // CORRECCIÓN: ACTUALIZAR UI EN HILO PRINCIPAL
    finalOutput := output.String()
    var finalStatus string
    
    if len(result.Errors) > 0 {
        finalStatus = fmt.Sprintf("Completado con %d errores", len(result.Errors))
    } else {
        finalStatus = "Ejecución exitosa"
    }
    
    // USAR fyne.Do para actualizar UI desde goroutine
    fyne.Do(func() {
        outputArea.SetText(finalOutput)
        updateStatus(finalStatus)
    })
}

func updateExecutionStatus(executing bool) {
    executionMutex.Lock()
    defer executionMutex.Unlock()
    isExecuting = executing
}

func getExecutionStatus() bool {
    executionMutex.Lock()
    defer executionMutex.Unlock()
    return isExecuting
}


// Tema personalizado moderno
type modernTheme struct {
	fyne.Theme
}


func (m *modernTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		if variant == theme.VariantLight {
			return color.NRGBA{R: 248, G: 250, B: 252, A: 255} // Gris muy claro
		}
		return color.NRGBA{R: 17, G: 24, B: 39, A: 255} // Azul oscuro
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 59, G: 130, B: 246, A: 255} // Azul moderno
	case theme.ColorNameInputBackground:
		if variant == theme.VariantLight {
			return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		}
		return color.NRGBA{R: 30, G: 41, B: 59, A: 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m *modernTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *modernTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *modernTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 14
	}
	return theme.DefaultTheme().Size(name)
}

// Widget personalizado para entrada de código con fuente monospace
type CodeEntry struct {
	widget.Entry
}

func NewCodeEntry() *CodeEntry {
	entry := &CodeEntry{}
	entry.ExtendBaseWidget(entry)
	entry.MultiLine = true
	entry.Wrapping = fyne.TextWrapWord
	entry.SetPlaceHolder("Escribe tu código VLang aquí...")
	return entry
}

func (c *CodeEntry) CreateRenderer() fyne.WidgetRenderer {
	c.TextStyle = fyne.TextStyle{Monospace: true}
	return c.Entry.CreateRenderer()
}

// Widget personalizado para salida de consola con fuente monospace
type ConsoleOutput struct {
	widget.Entry
}

func NewConsoleOutput() *ConsoleOutput {
	entry := &ConsoleOutput{}
	entry.ExtendBaseWidget(entry)
	entry.MultiLine = true
	entry.Wrapping = fyne.TextWrapWord
	entry.SetPlaceHolder("Aquí se mostrará el resultado de la ejecución...")
	return entry
}

func (c *ConsoleOutput) CreateRenderer() fyne.WidgetRenderer {
	c.TextStyle = fyne.TextStyle{Monospace: true}
	return c.Entry.CreateRenderer()
}

func findCompiler() (string, string) {
    // Lista de posibles rutas para el compilador
    compilerPaths := []string{
        "./backend/compiler/compiler",
        "../backend/compiler/compiler",
        "backend/compiler/compiler",
        "../../backend/compiler/compiler",
    }
    
    // Lista de posibles rutas para el generador CST
    cstPaths := []string{
        "./backend/compiler/tools/cst_generator",
        "../backend/compiler/tools/cst_generator",
        "backend/compiler/tools/cst_generator",
        "../../backend/compiler/tools/cst_generator",
    }
    
    wd, _ := os.Getwd()
    fmt.Printf("🔍 Buscando ejecutables desde: %s\n", wd)
    
    var compilerPath, cstPath string
    
    // Buscar compilador
    for _, path := range compilerPaths {
        if _, err := os.Stat(path); err == nil {
            compilerPath = path
            fmt.Printf("✅ Compilador encontrado en: %s\n", path)
            break
        }
    }
    
    // Buscar generador CST
    for _, path := range cstPaths {
        if _, err := os.Stat(path); err == nil {
            cstPath = path
            fmt.Printf("✅ Generador CST encontrado en: %s\n", path)
            break
        }
    }
    
    return compilerPath, cstPath
}

func main() {
    compilerPath, cstGeneratorPath := findCompiler()
    if compilerPath == "" {
        fmt.Println("❌ Error: No se pudo encontrar el compilador")
        return
    }
    if cstGeneratorPath == "" {
        fmt.Println("❌ Error: No se pudo encontrar el generador CST")
        return
    }
    
    interpreterClient := interpreter.NewClient(compilerPath)
    cstClient = interpreter.NewCSTClient(cstGeneratorPath)

    startMemoryCleanup()
	
    a := app.NewWithID("com.vlang.compiler.interpreter")
	a.Settings().SetTheme(&modernTheme{})
	w := a.NewWindow("VLang Interpreter Pro")
	w.SetIcon(theme.ComputerIcon())

	// Colores personalizados para una interfaz moderna
	primaryColor := color.NRGBA{R: 59, G: 130, B: 246, A: 255}

	// Título principal con estilo
	titleText := canvas.NewText("VLang Interpreter", primaryColor)
	titleText.TextSize = 24
	titleText.TextStyle.Bold = true
	titleText.Alignment = fyne.TextAlignCenter

	// Subtítulo
	subtitleText := canvas.NewText("Editor de código y consola integrada", color.NRGBA{R: 107, G: 114, B: 128, A: 255})
	subtitleText.TextSize = 12
	subtitleText.Alignment = fyne.TextAlignCenter

	// Área de entrada de código con fuente monospace
	codeInput := NewCodeEntry()

	// Área de salida con fuente monospace estilo consola
	outputArea := NewConsoleOutput()

	// Etiquetas de entrada y salida (sin cambios)
	inputLabel := container.NewHBox(
		widget.NewIcon(theme.DocumentIcon()),
		widget.NewLabelWithStyle("Editor de Código", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	)

	outputLabel := container.NewHBox(
		widget.NewIcon(theme.InfoIcon()),
		widget.NewLabelWithStyle("Consola de Salida", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
	)

	// ===== BARRA DE ESTADO DINÁMICA =====
	statusLabel := widget.NewLabel("Listo para ejecutar")
	linesLabel := widget.NewLabel("Líneas: 0")
	charsLabel := widget.NewLabel("Caracteres: 0")

	// Función para actualizar las estadísticas
	updateStats := func() {
		text := codeInput.Text
		charCount := len(text)
		lineCount := len(strings.Split(text, "\n"))
		if text == "" {
			lineCount = 0
		}

		linesLabel.SetText(fmt.Sprintf("Líneas: %d", lineCount))
		charsLabel.SetText(fmt.Sprintf("Caracteres: %d", charCount))
	}

	// Función para actualizar el estado
	updateStatus := func(status string) {
		statusLabel.SetText(status)
	}

	// Conectar el evento OnChanged del editor de código
	codeInput.OnChanged = func(text string) {
		updateStats()
		if text == "" {
			updateStatus("Listo para ejecutar")
		} else {
			updateStatus("Código modificado")
		}
	}

	executeBtn := widget.NewButtonWithIcon("Ejecutar Código", theme.MediaPlayIcon(), func() {
        code := codeInput.Text
        if code == "" {
            outputArea.SetText("❌ Error: No hay código para ejecutar")
            updateStatus("Error: Sin código")
            return
        }

        // PREVENIR ejecuciones múltiples
        if getExecutionStatus() {
            outputArea.SetText("⚠️ Ya hay una ejecución en progreso...\n" +
                            "🔄 Por favor espera a que termine.")
            return
        }

        // INTENTAR ADQUIRIR SEMÁFORO (no bloqueante)
        select {
        case executionSemaphore <- struct{}{}:
            // Continuar con la ejecución
        default:
            outputArea.SetText("⚠️ Sistema ocupado. Intenta de nuevo en un momento.")
            return
        }

        // MARCAR COMO EJECUTANDO
        updateExecutionStatus(true)
        lastExecutedCode = code
        
        // FEEDBACK INMEDIATO CON ESTIMACIONES
        codeLines := len(strings.Split(strings.TrimSpace(code), "\n"))
        estimatedTime := "1-3 segundos"
        if codeLines > 50 {
            estimatedTime = "3-10 segundos"
        } else if codeLines > 100 {
            estimatedTime = "10-30 segundos"
        }
        
        outputArea.SetText("🔄 Ejecutando código...\n\n" +
                        fmt.Sprintf("📝 Procesando %d líneas\n", codeLines) +
                        fmt.Sprintf("⏱️ Tiempo estimado: %s\n", estimatedTime) +
                        "⚡ Optimizaciones activas\n\n" +
                        "💡 La interfaz permanece responsive...")
        
        updateStatus("⚡ Ejecutando...")

        // LIMPIAR datos anteriores antes de ejecutar
        currentErrors = []interpreter.ErrorDetail{}
        currentSymbols = []interpreter.SymbolDetail{}
        currentCST = ""

        // EJECUCIÓN ASÍNCRONA CON TIMEOUT INTELIGENTE
        go func() {
            defer func() {
                // CLEANUP: Liberar recursos SIEMPRE
                <-executionSemaphore
                updateExecutionStatus(false)
            }()
            
            // TIMEOUT DINÁMICO basado en tamaño del código
            timeoutDuration := 30 * time.Second
            if len(code) > 10000 {
                timeoutDuration = 60 * time.Second
            } else if len(code) < 1000 {
                timeoutDuration = 15 * time.Second
            }
            
            ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
            defer cancel()
            
            // CANALES para manejar resultado, error o timeout
            resultChan := make(chan *interpreter.ExecutionResult, 1)
            errorChan := make(chan error, 1)
            
            // EJECUTAR EN GOROUTINE SEPARADA
            go func() {
                start := time.Now()
                result, err := interpreterClient.Execute(code)
                
                if err != nil {
                    errorChan <- err
                } else {
                    // Debug opcional
                    executionTime := time.Since(start)
                    fmt.Printf("⚡ Ejecución completada en: %v\n", executionTime)
                    resultChan <- result
                }
            }()
            
            // ESPERAR RESULTADO, ERROR O TIMEOUT
            select {
            case result := <-resultChan:
                // EJECUCIÓN EXITOSA - procesar resultado
                processExecutionResult(result, outputArea, updateStatus)
                
            case err := <-errorChan:
                // ERROR EN EJECUCIÓN
                errorMessage := "❌ Error al ejecutar el código:\n\n" + 
                            err.Error() + "\n\n" +
                            "💡 Posibles causas:\n" +
                            "• Error en la sintaxis del código\n" +
                            "• Problema con el compilador backend\n" +
                            "• Archivo temporal corrupto\n\n" +
                            "🔧 Verifica tu código e intenta de nuevo."
                
                fyne.Do(func() {
                    outputArea.SetText(errorMessage)
                    updateStatus("❌ Error en ejecución")
                })
                
            case <-ctx.Done():
                // TIMEOUT - CORREGIR con fyne.Do()
                timeoutMessage := fmt.Sprintf("⏰ Timeout: La ejecución tomó más de %v\n\n", timeoutDuration) +
                                "🔧 Posibles causas:\n" +
                                "• Bucle infinito en tu código (for sin condición de salida)\n" +
                                "• Código muy complejo que requiere mucho procesamiento\n" +
                                "• Recursión infinita en funciones\n" +
                                "• Operaciones muy intensivas\n\n" +
                                "💡 Recomendaciones:\n" +
                                "• Revisa bucles for y while\n" +
                                "• Verifica condiciones de parada\n" +
                                "• Simplifica el código si es muy complejo\n" +
                                "• Intenta ejecutar partes más pequeñas"
                
                fyne.Do(func() {
                    outputArea.SetText(timeoutMessage)
                    updateStatus("⏰ Timeout en ejecución")
                })
            }
        }()
    })
	executeBtn.Importance = widget.HighImportance

	clearBtn := widget.NewButtonWithIcon("Limpiar", theme.DeleteIcon(), func() {
		codeInput.SetText("")
		outputArea.SetText("")
		updateStatus("Listo para ejecutar")
		updateStats()
	})

	// Barra de herramientas
	toolbar := container.NewPadded(
		container.NewHBox(
			executeBtn,
			widget.NewSeparator(),
			clearBtn,
		),
	)

	// Paneles
	leftPanel := container.NewBorder(
		container.NewVBox(inputLabel, widget.NewSeparator()),
		nil, nil, nil,
		container.NewPadded(codeInput),
	)

	rightPanel := container.NewBorder(
		container.NewVBox(outputLabel, widget.NewSeparator()),
		nil, nil, nil,
		container.NewPadded(outputArea),
	)

	split := container.NewHSplit(leftPanel, rightPanel)
	split.Offset = 0.5

	// ===== BARRA DE ESTADO DINÁMICA ACTUALIZADA =====
	statusBar := container.NewPadded(
		container.NewHBox(
			widget.NewIcon(theme.InfoIcon()),
			statusLabel,
			widget.NewSeparator(),
			linesLabel,
			widget.NewSeparator(),
			charsLabel,
		),
	)

	// Funciones para los menús con estado dinámico
	newFile := func() {
		codeInput.SetText("")
		outputArea.SetText("")
		updateStatus("Archivo nuevo")
	}

	openFile := func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			defer reader.Close()
			data, err := io.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			codeInput.SetText(string(data))
			updateStatus("Archivo cargado")
			updateStats()
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".vch"})) // Solo archivos .vch
		fd.Resize(fyne.NewSize(600, 400))
		fd.Show()
	}

	mainMenu := fyne.NewMainMenu(
		fyne.NewMenu("Archivo",
			fyne.NewMenuItem("Nuevo", newFile),
			fyne.NewMenuItem("Abrir", openFile),
			fyne.NewMenuItem("Guardar", func() {
				saveFile(w, &codeInput.Entry)
				updateStatus("Archivo guardado")
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Salir", func() { a.Quit() }),
		),
		fyne.NewMenu("Herramientas",
			fyne.NewMenuItem("Ejecutar", func() { executeBtn.OnTapped() }),
			fyne.NewMenuItemSeparator(),
		),
		fyne.NewMenu("Reportes",
			fyne.NewMenuItem("Reporte de Errores", func() {
				showErrorReport(w)
			}),
			fyne.NewMenuItem("Tabla de Símbolos", func() {
				showSymbolsReport(w)
			}),
			fyne.NewMenuItem("Árbol CST", func() {
                showCSTReport(w)
            }),
		),
	)

	w.SetMainMenu(mainMenu)

	header := container.NewPadded(
		container.NewVBox(
			container.NewCenter(titleText),
			container.NewCenter(subtitleText),
			widget.NewSeparator(),
		),
	)

	content := container.NewBorder(
		container.NewVBox(header, toolbar, widget.NewSeparator()),
		statusBar,
		nil, nil,
		container.NewPadded(split),
	)

	// Inicializar estadísticas
	updateStats()

	w.SetContent(content)
	w.Resize(fyne.NewSize(1200, 800))
	w.CenterOnScreen()
	w.ShowAndRun()
}

// Función auxiliar para guardar archivos
func saveFile(w fyne.Window, codeInput *widget.Entry) {
	fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil || writer == nil {
			return
		}
		defer writer.Close()
		_, err = writer.Write([]byte(codeInput.Text))
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		dialog.ShowInformation("Guardado", "Archivo guardado correctamente", w)
	}, w)
	fd.SetFilter(storage.NewExtensionFileFilter([]string{".vch"}))
	fd.SetFileName("codigo.vch") // Nombre por defecto con extensión .vch
	fd.Resize(fyne.NewSize(600, 400))
	fd.Show()
}

// Función específica para mostrar la tabla de símbolos
func showSymbolsReport(w fyne.Window) {
    if len(currentSymbols) == 0 {
        // Mensaje cuando no hay símbolos
        emptyIcon := widget.NewIcon(theme.InfoIcon())
        emptyLabel := widget.NewLabelWithStyle("No se encontraron símbolos", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
        emptyDesc := widget.NewLabel("Ejecuta código para generar la tabla de símbolos.")
        emptyDesc.Alignment = fyne.TextAlignCenter

        content := container.NewVBox(
            widget.NewSeparator(),
            container.NewCenter(emptyIcon),
            widget.NewSeparator(),
            container.NewCenter(emptyLabel),
            container.NewCenter(emptyDesc),
            widget.NewSeparator(),
        )

        dialog.ShowCustom("Tabla de Símbolos", "Cerrar", content, w)
        return
    }

    // Generar y abrir reporte HTML
    generateHTMLSymbolsReport()

    // Mostrar confirmación simple
    dialog.ShowInformation("Tabla Generada",
        fmt.Sprintf("Se generó una tabla HTML con %d símbolos y se abrió en tu navegador.", len(currentSymbols)), w)
}

func generateHTMLSymbolsReport() {
    if lastExecutedCode == "" {
        fmt.Println("⚠️ No hay código ejecutado para generar tabla de símbolos")
        return
    }
    
    codeHash := calculateCodeHash(lastExecutedCode)
    
    // VERIFICAR CACHE PRIMERO
    if cachedReport, found := reportCache.getSymbolReport(codeHash); found {
        fmt.Println("📊 Usando tabla de símbolos en cache (⚡ instantáneo)")
        
        // Escribir el reporte cacheado al archivo
        reportsDir := filepath.Join(".", "Reportes")
        if err := os.MkdirAll(reportsDir, 0755); err == nil {
            fileName := "tabla_simbolos.html"
            filePath := filepath.Join(reportsDir, fileName)
            if err := os.WriteFile(filePath, []byte(cachedReport), 0644); err == nil {
                absPath, _ := filepath.Abs(filePath)
                fmt.Printf("📊 Tabla de símbolos de cache escrita: %s\n", absPath)
                openInBrowser(absPath)
            }
        }
        return
    }
    
    // GENERAR NUEVA TABLA
    start := time.Now()
    
    htmlContent := generateSymbolsHTMLContent()

    // Crear directorio de reportes en front/Reportes
    reportsDir := filepath.Join(".", "Reportes")
    err := os.MkdirAll(reportsDir, 0755)
    if err != nil {
        fmt.Printf("Error creando directorio de reportes: %v\n", err)
        return
    }

    // Usar nombre específico para tabla de símbolos
    fileName := "tabla_simbolos.html"
    filePath := filepath.Join(reportsDir, fileName)

    // Escribir archivo
    err = os.WriteFile(filePath, []byte(htmlContent), 0644)
    if err != nil {
        fmt.Printf("Error creando archivo HTML: %v\n", err)
        return
    }

    // GUARDAR EN CACHE DESPUÉS DE GENERAR
    reportCache.updateSymbolReport(codeHash, lastExecutedCode, htmlContent)
    
    generationTime := time.Since(start)
    fmt.Printf("📊 Nueva tabla de símbolos generada en: %v\n", generationTime)

    // Obtener ruta absoluta para mejor compatibilidad
    absPath, _ := filepath.Abs(filePath)
    fmt.Printf("📊 Tabla de símbolos guardada en: %s\n", absPath)
    
    // Abrir en navegador
    openInBrowser(absPath)
}

// Función para generar el contenido HTML de la tabla de símbolos
func generateSymbolsHTMLContent() string {
    var symbolRows strings.Builder

    for i, symbol := range currentSymbols {
        // Determinar colores según el tipo de símbolo
        var badgeClass, iconColor, rowBg string
        
        switch symbol.SymbolType {
        case "Variable":
            badgeClass = "bg-green-100 text-green-800 border-green-200 dark:bg-green-900 dark:text-green-300 dark:border-green-700"
            iconColor = "text-green-500 dark:text-green-400"
            if i%2 == 0 {
                rowBg = "bg-green-25 dark:bg-green-950"
            } else {
                rowBg = "bg-green-50 dark:bg-green-900"
            }
        case "Función":
            badgeClass = "bg-blue-100 text-blue-800 border-blue-200 dark:bg-blue-900 dark:text-blue-300 dark:border-blue-700"
            iconColor = "text-blue-500 dark:text-blue-400"
            if i%2 == 0 {
                rowBg = "bg-blue-25 dark:bg-blue-950"
            } else {
                rowBg = "bg-blue-50 dark:bg-blue-900"
            }
        case "Slice":
            badgeClass = "bg-purple-100 text-purple-800 border-purple-200 dark:bg-purple-900 dark:text-purple-300 dark:border-purple-700"
            iconColor = "text-purple-500 dark:text-purple-400"
            if i%2 == 0 {
                rowBg = "bg-purple-25 dark:bg-purple-950"
            } else {
                rowBg = "bg-purple-50 dark:bg-purple-900"
            }
        case "Matriz":
            badgeClass = "bg-indigo-100 text-indigo-800 border-indigo-200 dark:bg-indigo-900 dark:text-indigo-300 dark:border-indigo-700"
            iconColor = "text-indigo-500 dark:text-indigo-400"
            if i%2 == 0 {
                rowBg = "bg-indigo-25 dark:bg-indigo-950"
            } else {
                rowBg = "bg-indigo-50 dark:bg-indigo-900"
            }
        case "Parámetro":
            badgeClass = "bg-yellow-100 text-yellow-800 border-yellow-200 dark:bg-yellow-900 dark:text-yellow-300 dark:border-yellow-700"
            iconColor = "text-yellow-500 dark:text-yellow-400"
            if i%2 == 0 {
                rowBg = "bg-yellow-25 dark:bg-yellow-950"
            } else {
                rowBg = "bg-yellow-50 dark:bg-yellow-900"
            }
        default:
            badgeClass = "bg-gray-100 text-gray-800 border-gray-200 dark:bg-gray-700 dark:text-gray-300 dark:border-gray-600"
            iconColor = "text-gray-500 dark:text-gray-400"
            if i%2 == 0 {
                rowBg = "bg-white dark:bg-gray-800"
            } else {
                rowBg = "bg-gray-50 dark:bg-gray-700"
            }
        }

        // Determinar icono específico por tipo de símbolo
        var symbolIcon string
        switch symbol.SymbolType {
        case "Variable":
            symbolIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                        </svg>`
        case "Función":
            symbolIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"/>
                        </svg>`
        case "Slice":
            symbolIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M3 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zm0 4a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clip-rule="evenodd"/>
                        </svg>`
        case "Matriz":
            symbolIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11 4a1 1 0 10-2 0v4a1 1 0 102 0V7zm-3 1a1 1 0 10-2 0v3a1 1 0 102 0V8zM8 9a1 1 0 00-2 0v2a1 1 0 102 0V9z" clip-rule="evenodd"/>
                        </svg>`
        default:
            symbolIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"/>
                        </svg>`
        }

        symbolRows.WriteString(fmt.Sprintf(`
            <tr class="%s transition-colors duration-200">
                <td class="px-6 py-4 text-sm text-gray-900 dark:text-gray-100">
                    <div class="flex items-center space-x-3">
                        <div class="flex-shrink-0">
                            %s
                        </div>
                        <div class="flex-1">
                            <p class="font-medium text-gray-900 dark:text-gray-100">%s</p>
                        </div>
                    </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-center">
                    <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium border %s">
                        %s
                    </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-gray-100 text-center">
                    <code class="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded text-xs">%s</code>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-gray-100 text-center">
                    <span class="inline-flex items-center px-2 py-1 rounded-md text-xs font-medium bg-gray-100 dark:bg-gray-700 text-gray-800 dark:text-gray-200">
                        %s
                    </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-gray-100 text-center">
                    <span class="inline-flex items-center justify-center w-8 h-8 bg-blue-100 dark:bg-blue-800 text-blue-800 dark:text-blue-200 rounded-full text-xs font-bold">
                        %d
                    </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-gray-100 text-center">
                    <span class="inline-flex items-center justify-center w-8 h-8 bg-indigo-100 dark:bg-indigo-800 text-indigo-800 dark:text-indigo-200 rounded-full text-xs font-bold">
                        %d
                    </span>
                </td>
            </tr>`,
            rowBg,
            fmt.Sprintf(symbolIcon, iconColor),
            symbol.ID,
            badgeClass,
            symbol.SymbolType,
            symbol.DataType,
            symbol.Scope,
            symbol.Line,
            symbol.Column))
    }

    return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="es" class="h-full">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tabla de Símbolos - VLang Compiler</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            darkMode: 'media',
            theme: {
                extend: {
                    fontFamily: {
                        'mono': ['JetBrains Mono', 'Fira Code', 'Monaco', 'Consolas', 'monospace'],
                    },
                    colors: {
                        'green-25': '#f0fdf4',
                        'blue-25': '#eff6ff',
                        'purple-25': '#fdf4ff',
                        'indigo-25': '#eef2ff',
                        'yellow-25': '#fffbeb',
                        'green-950': '#14532d',
                        'blue-950': '#1e3a8a',
                        'purple-950': '#581c87',
                        'indigo-950': '#312e81',
                        'yellow-950': '#92400e'
                    }
                }
            }
        }
    </script>
</head>
<body class="h-full bg-gradient-to-br from-green-50 via-white to-blue-50 dark:from-gray-900 dark:via-gray-800 dark:to-green-900">
    <div class="min-h-full">
        <!-- Header -->
        <header class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div class="flex justify-between items-center py-6">
                    <div class="flex items-center space-x-4">
                        <div class="flex-shrink-0">
                            <div class="w-12 h-12 bg-gradient-to-r from-green-500 to-blue-600 dark:from-green-600 dark:to-blue-700 rounded-lg flex items-center justify-center">
                                <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11 4a1 1 0 10-2 0v4a1 1 0 102 0V7zm-3 1a1 1 0 10-2 0v3a1 1 0 102 0V8zM8 9a1 1 0 00-2 0v2a1 1 0 102 0V9z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div>
                            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">VLang Compiler</h1>
                            <p class="text-sm text-gray-500 dark:text-gray-400">Tabla de Símbolos</p>
                        </div>
                    </div>
                    <div class="flex items-center space-x-4">
                        <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300">
                            %d símbolos encontrados
                        </span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">%s</span>
                    </div>
                </div>
            </div>
        </header>

        <!-- Main Content -->
        <main class="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
            <!-- Stats -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-green-100 dark:bg-green-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Total de Símbolos</p>
                            <p class="text-2xl font-bold text-gray-900 dark:text-white">%d</p>
                        </div>
                    </div>
                </div>
                
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-blue-100 dark:bg-blue-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11 4a1 1 0 10-2 0v4a1 1 0 102 0V7zm-3 1a1 1 0 10-2 0v3a1 1 0 102 0V8zM8 9a1 1 0 00-2 0v2a1 1 0 102 0V9z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Tipos Únicos</p>
                            <p class="text-2xl font-bold text-gray-900 dark:text-white">%d</p>
                        </div>
                    </div>
                </div>
                
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-purple-100 dark:bg-purple-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-purple-600 dark:text-purple-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M5 2a1 1 0 011 1v1h1a1 1 0 010 2H6v1a1 1 0 01-2 0V6H3a1 1 0 010-2h1V3a1 1 0 011-1zm0 10a1 1 0 011 1v1h1a1 1 0 110 2H6v1a1 1 0 11-2 0v-1H3a1 1 0 110-2h1v-1a1 1 0 011-1zM12 2a1 1 0 01.967.744L14.146 7.2 17.5 9.134a1 1 0 010 1.732L14.146 12.8l-1.179 4.456a1 1 0 01-1.934 0L9.854 12.8 6.5 10.866a1 1 0 010-1.732L9.854 7.2l1.179-4.456A1 1 0 0112 2z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Estado</p>
                            <p class="text-lg font-bold text-green-600 dark:text-green-400">Análisis Completo</p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Symbols Table -->
            <div class="bg-white dark:bg-gray-800 shadow-sm rounded-lg border border-gray-200 dark:border-gray-700 overflow-hidden">
                <div class="px-6 py-4 border-b border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-700">
                    <h3 class="text-lg font-medium text-gray-900 dark:text-white">Tabla de Símbolos</h3>
                    <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Lista completa de símbolos encontrados durante el análisis</p>
                </div>
                
                <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                        <thead class="bg-gray-50 dark:bg-gray-700">
                            <tr>
                                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Identificador
                                </th>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Tipo
                                </th>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Tipo de Dato
                                </th>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Ámbito
                                </th>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Línea
                                </th>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Columna
                                </th>
                            </tr>
                        </thead>
                        <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
                            %s
                        </tbody>
                    </table>
                </div>
            </div>

            <!-- Footer -->
            <div class="mt-8 bg-green-50 dark:bg-green-900 rounded-lg p-6 border border-green-200 dark:border-green-700">
                <div class="flex items-start space-x-3">
                    <div class="flex-shrink-0">
                        <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="currentColor" viewBox="0 0 20 20">
                            <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                        </svg>
                    </div>
                    <div>
                        <h4 class="text-sm font-medium text-green-900 dark:text-green-100">Información</h4>
                        <p class="mt-1 text-sm text-green-700 dark:text-green-200">
                            Esta tabla muestra todos los símbolos encontrados en tu código, incluyendo variables, funciones, slices y matrices.
                            Los símbolos están ordenados por línea de declaración.
                        </p>
                    </div>
                </div>
            </div>
        </main>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', function() {
            console.log('Tabla de símbolos cargada exitosamente');
            
            const rows = document.querySelectorAll('tbody tr');
            rows.forEach(row => {
                row.addEventListener('mouseenter', function() {
                    this.style.transform = 'scale(1.01)';
                    this.style.transition = 'transform 0.2s ease';
                });
                
                row.addEventListener('mouseleave', function() {
                    this.style.transform = 'scale(1)';
                });
            });
        });
    </script>
</body>
</html>`,
        len(currentSymbols),
        time.Now().Format("02/01/2006 15:04:05"),
        len(currentSymbols),
        getUniqueSymbolTypes(),
        symbolRows.String())
}

// Función para obtener tipos únicos de símbolos
func getUniqueSymbolTypes() int {
    typeMap := make(map[string]bool)
    for _, symbol := range currentSymbols {
        typeMap[symbol.SymbolType] = true
    }
    return len(typeMap)
}

// Función específica para mostrar el reporte de errores
func showErrorReport(w fyne.Window) {
	if len(currentErrors) == 0 {
		// Mensaje de éxito simple para cuando no hay errores
		successIcon := widget.NewIcon(theme.ConfirmIcon())
		successLabel := widget.NewLabelWithStyle("No se encontraron errores", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
		successDesc := widget.NewLabel("El código está libre de errores.")
		successDesc.Alignment = fyne.TextAlignCenter

		content := container.NewVBox(
			widget.NewSeparator(),
			container.NewCenter(successIcon),
			widget.NewSeparator(),
			container.NewCenter(successLabel),
			container.NewCenter(successDesc),
			widget.NewSeparator(),
		)

		dialog.ShowCustom("Reporte de Errores", "Cerrar", content, w)
		return
	}

	// Generar y abrir reporte HTML
	generateHTMLErrorReport()

	// Mostrar confirmación simple
	dialog.ShowInformation("Reporte Generado",
		fmt.Sprintf("Se generó un reporte HTML con %d errores y se abrió en tu navegador.", len(currentErrors)), w)
}

// REEMPLAZAR tu función generateHTMLErrorReport() completa:
func generateHTMLErrorReport() {
    if lastExecutedCode == "" {
        fmt.Println("⚠️ No hay código ejecutado para generar reporte")
        return
    }
    
    codeHash := calculateCodeHash(lastExecutedCode)
    
    // VERIFICAR CACHE PRIMERO
    if cachedReport, found := reportCache.getErrorReport(codeHash); found {
        fmt.Println("📋 Usando reporte de errores en cache (⚡ instantáneo)")
        
        // Escribir el reporte cacheado al archivo
        reportsDir := filepath.Join(".", "Reportes")
        if err := os.MkdirAll(reportsDir, 0755); err == nil {
            fileName := "reporte_errores.html"
            filePath := filepath.Join(reportsDir, fileName)
            if err := os.WriteFile(filePath, []byte(cachedReport), 0644); err == nil {
                absPath, _ := filepath.Abs(filePath)
                fmt.Printf("📄 Reporte de cache escrito: %s\n", absPath)
                openInBrowser(absPath)
            }
        }
        return
    }
    
    // GENERAR NUEVO REPORTE
    start := time.Now()
    
    // Generar contenido HTML usando tu función existente
    htmlContent := generateHTMLContent()
    
    // Crear directorio de reportes
    reportsDir := filepath.Join(".", "Reportes")
    err := os.MkdirAll(reportsDir, 0755)
    if err != nil {
        fmt.Printf("Error creando directorio de reportes: %v\n", err)
        return
    }

    // Usar nombre fijo para que se reemplace automáticamente
    fileName := "reporte_errores.html"
    filePath := filepath.Join(reportsDir, fileName)

    // Escribir archivo
    err = os.WriteFile(filePath, []byte(htmlContent), 0644)
    if err != nil {
        fmt.Printf("Error creando archivo HTML: %v\n", err)
        return
    }

    // GUARDAR EN CACHE DESPUÉS DE GENERAR
    reportCache.updateErrorReport(codeHash, lastExecutedCode, htmlContent)
    
    generationTime := time.Since(start)
    fmt.Printf("📄 Nuevo reporte generado en: %v\n", generationTime)

    // Obtener ruta absoluta y abrir
    absPath, _ := filepath.Abs(filePath)
    fmt.Printf("📄 Reporte guardado en: %s\n", absPath)
    
    openInBrowser(absPath)
}

// MODIFICAR tu función generateHTMLContent() para agregar estadísticas opcionales:
func generateHTMLContent() string {
    var errorRows strings.Builder

    for i, err := range currentErrors {
        // Determinar clases CSS según el tipo de error con colores distintivos
        var badgeClass, iconColor string
        switch err.Type {
        case "Error sintáctico":
            badgeClass = "bg-red-100 text-red-800 border-red-200 dark:bg-red-900 dark:text-red-300 dark:border-red-700"
            iconColor = "text-red-500 dark:text-red-400"
        case "Error léxico":
            badgeClass = "bg-orange-100 text-orange-800 border-orange-200 dark:bg-orange-900 dark:text-orange-300 dark:border-orange-700"
            iconColor = "text-orange-500 dark:text-orange-400"
        case "Error semántico":
            badgeClass = "bg-purple-100 text-purple-800 border-purple-200 dark:bg-purple-900 dark:text-purple-300 dark:border-purple-700"
            iconColor = "text-purple-500 dark:text-purple-400"
        default:
            badgeClass = "bg-gray-100 text-gray-800 border-gray-200 dark:bg-gray-700 dark:text-gray-300 dark:border-gray-600"
            iconColor = "text-gray-500 dark:text-gray-400"
        }

        // Alternar color de fondo de filas con tinte sutil del tipo de error
        var rowBg string
        switch err.Type {
        case "Error sintáctico":
            if i%2 == 0 {
                rowBg = "bg-red-25 dark:bg-red-950"
            } else {
                rowBg = "bg-red-50 dark:bg-red-900"
            }
        case "Error léxico":
            if i%2 == 0 {
                rowBg = "bg-orange-25 dark:bg-orange-950"
            } else {
                rowBg = "bg-orange-50 dark:bg-orange-900"
            }
        case "Error semántico":
            if i%2 == 0 {
                rowBg = "bg-purple-25 dark:bg-purple-950"
            } else {
                rowBg = "bg-purple-50 dark:bg-purple-900"
            }
        default:
            if i%2 == 0 {
                rowBg = "bg-white dark:bg-gray-800"
            } else {
                rowBg = "bg-gray-50 dark:bg-gray-700"
            }
        }

        // Determinar icono específico por tipo de error
        var errorIcon string
        switch err.Type {
        case "Error sintáctico":
            errorIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"/>
                        </svg>`
        case "Error léxico":
            errorIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"/>
                        </svg>`
        case "Error semántico":
            errorIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
                        </svg>`
        default:
            errorIcon = `<svg class="w-5 h-5 %s mt-0.5" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"/>
                        </svg>`
        }

        errorRows.WriteString(fmt.Sprintf(`
            <tr class="%s transition-colors duration-200">
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-gray-100 text-center">
                    <span class="inline-flex items-center justify-center w-8 h-8 bg-blue-100 dark:bg-blue-800 text-blue-800 dark:text-blue-200 rounded-full text-xs font-bold">
                        %d
                    </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 dark:text-gray-100 text-center">
                    <span class="inline-flex items-center justify-center w-8 h-8 bg-indigo-100 dark:bg-indigo-800 text-indigo-800 dark:text-indigo-200 rounded-full text-xs font-bold">
                        %d
                    </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-center">
                    <span class="inline-flex items-center px-3 py-1 rounded-full text-xs font-medium border %s">
                        %s
                    </span>
                </td>
                <td class="px-6 py-4 text-sm text-gray-900 dark:text-gray-100">
                    <div class="flex items-start space-x-3">
                        <div class="flex-shrink-0">
                            %s
                        </div>
                        <div class="flex-1">
                            <p class="font-medium text-gray-900 dark:text-gray-100">%s</p>
                        </div>
                    </div>
                </td>
            </tr>`,
            rowBg,
            err.Line,
            err.Column,
            badgeClass,
            err.Type,
            fmt.Sprintf(errorIcon, iconColor),
            err.Message))
    }

    // AGREGAR: Calcular información de cache para las estadísticas
    codeHash := calculateCodeHash(lastExecutedCode)
    codeLines := len(strings.Split(strings.TrimSpace(lastExecutedCode), "\n"))

    return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="es" class="h-full">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Reporte de Errores - VLang Compiler</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            darkMode: 'media',
            theme: {
                extend: {
                    fontFamily: {
                        'mono': ['JetBrains Mono', 'Fira Code', 'Monaco', 'Consolas', 'monospace'],
                    },
                    colors: {
                        'red-25': '#fef7f7',
                        'orange-25': '#fffbf5',
                        'purple-25': '#fdf7ff',
                        'yellow-25': '#fffef0',
                        'red-950': '#4c1d1d',
                        'orange-950': '#4c2a1a',
                        'purple-950': '#3b1a4c',
                        'yellow-950': '#4c4a1a'
                    }
                }
            }
        }
    </script>
</head>
<body class="h-full bg-gradient-to-br from-blue-50 via-white to-indigo-50 dark:from-gray-900 dark:via-gray-800 dark:to-blue-900">
    <div class="min-h-full">
        <!-- Header -->
        <header class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div class="flex justify-between items-center py-6">
                    <div class="flex items-center space-x-4">
                        <div class="flex-shrink-0">
                            <div class="w-12 h-12 bg-gradient-to-r from-blue-500 to-indigo-600 dark:from-blue-600 dark:to-indigo-700 rounded-lg flex items-center justify-center">
                                <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.707a1 1 0 00-1.414-1.414L10 9.586 8.707 8.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div>
                            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">VLang Compiler</h1>
                            <p class="text-sm text-gray-500 dark:text-gray-400">Reporte de Errores</p>
                        </div>
                    </div>
                    <div class="flex items-center space-x-4">
                        <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300">
                            %d errores encontrados
                        </span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">%s</span>
                    </div>
                </div>
            </div>
        </header>

        <!-- Main Content -->
        <main class="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
            <!-- Stats -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-red-100 dark:bg-red-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-red-600 dark:text-red-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Total de Errores</p>
                            <p class="text-2xl font-bold text-gray-900 dark:text-white">%d</p>
                        </div>
                    </div>
                </div>
                
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-orange-100 dark:bg-orange-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-orange-600 dark:text-orange-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Tipos Únicos</p>
                            <p class="text-2xl font-bold text-gray-900 dark:text-white">%d</p>
                        </div>
                    </div>
                </div>
                
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-blue-100 dark:bg-blue-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Estado</p>
                            <p class="text-lg font-bold text-red-600 dark:text-red-400">Requiere Corrección</p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Error Table -->
            <div class="bg-white dark:bg-gray-800 shadow-sm rounded-lg border border-gray-200 dark:border-gray-700 overflow-hidden">
                <div class="px-6 py-4 border-b border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-700">
                    <h3 class="text-lg font-medium text-gray-900 dark:text-white">Detalle de Errores</h3>
                    <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Lista completa de errores encontrados durante la compilación</p>
                </div>
                
                <div class="overflow-x-auto">
                    <table class="min-w-full divide-y divide-gray-200 dark:divide-gray-700">
                        <thead class="bg-gray-50 dark:bg-gray-700">
                            <tr>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Línea
                                </th>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Columna
                                </th>
                                <th scope="col" class="px-6 py-3 text-center text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Tipo
                                </th>
                                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 dark:text-gray-300 uppercase tracking-wider">
                                    Mensaje
                                </th>
                            </tr>
                        </thead>
                        <tbody class="bg-white dark:bg-gray-800 divide-y divide-gray-200 dark:divide-gray-700">
                            %s
                        </tbody>
                    </table>
                </div>
            </div>

            <!-- AGREGAR: Información adicional de rendimiento (opcional) -->
            <div class="mt-8 bg-blue-50 dark:bg-blue-900 rounded-lg p-6 border border-blue-200 dark:border-blue-700">
                <div class="flex items-start space-x-3">
                    <div class="flex-shrink-0">
                        <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="currentColor" viewBox="0 0 20 20">
                            <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"/>
                        </svg>
                    </div>
                    <div class="flex-1">
                        <h4 class="text-sm font-medium text-blue-900 dark:text-blue-100">Información del Análisis</h4>
                        <div class="mt-2 grid grid-cols-1 md:grid-cols-3 gap-4 text-sm text-blue-700 dark:text-blue-200">
                            <div>
                                <strong>Hash del código:</strong><br>
                                <code class="text-xs bg-blue-100 dark:bg-blue-800 px-2 py-1 rounded">%s</code>
                            </div>
                            <div>
                                <strong>Líneas procesadas:</strong> %d
                            </div>
                            <div>
                                
                            </div>
                        </div>
                        <p class="mt-3 text-sm text-blue-700 dark:text-blue-200">
                            Corrige los errores mostrados en la tabla antes de ejecutar tu código. 
                            Los errores están ordenados por línea para facilitar su corrección.
                        </p>
                    </div>
                </div>
            </div>
        </main>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', function() {
            console.log('Reporte de errores cargado exitosamente');
            
            const rows = document.querySelectorAll('tbody tr');
            rows.forEach(row => {
                row.addEventListener('mouseenter', function() {
                    this.style.transform = 'scale(1.01)';
                    this.style.transition = 'transform 0.2s ease';
                });
                
                row.addEventListener('mouseleave', function() {
                    this.style.transform = 'scale(1)';
                });
            });
        });
    </script>
</body>
</html>`,
        len(currentErrors),
        time.Now().Format("02/01/2006 15:04:05"),
        len(currentErrors),
        getUniqueErrorTypes(),
        errorRows.String(),
        codeHash,
        codeLines)
}

func showCSTReport(w fyne.Window) {
    if lastExecutedCode == "" {
        dialog.ShowInformation("Sin código", 
            "Primero debes ejecutar algún código para generar el CST.", w)
        return
    }

    codeHash := calculateCodeHash(lastExecutedCode)
    
    // VERIFICAR CACHE PRIMERO
    if cachedCST, found := reportCache.getCSTReport(codeHash); found {
        fmt.Println("🌳 Usando árbol CST en cache (⚡ instantáneo)")
        
        // Escribir CST cacheado al archivo
        reportsDir := filepath.Join(".", "Reportes")
        if err := os.MkdirAll(reportsDir, 0755); err == nil {
            fileName := "arbol_cst.html"
            filePath := filepath.Join(reportsDir, fileName)
            if err := os.WriteFile(filePath, []byte(cachedCST), 0644); err == nil {
                absPath, _ := filepath.Abs(filePath)
                fmt.Printf("🌳 CST de cache escrito: %s\n", absPath)
                openInBrowser(absPath)
                
                dialog.ShowInformation("Árbol CST (Cache)", 
                    "Se utilizó el árbol CST en cache y se abrió en tu navegador.\n⚡ Tiempo: instantáneo", w)
            }
        }
        return
    }

    // VERIFICAR TAMAÑO ANTES DE GENERAR (optimización de recursos)
    codeSize := len(lastExecutedCode)
    codeLines := len(strings.Split(strings.TrimSpace(lastExecutedCode), "\n"))
    
    var estimatedTime string
    var warningMessage string
    
    if codeSize > 50000 { // 50KB
        estimatedTime = "5-15 minutos"
        warningMessage = fmt.Sprintf("⚠️ CÓDIGO GRANDE DETECTADO\n\n"+
            "Tamaño: %d KB (%d líneas)\n"+
            "Tiempo estimado: %s\n\n"+
            "El proceso puede ser intensivo para el sistema.\n"+
            "¿Deseas continuar?", codeSize/1024, codeLines, estimatedTime)
    } else if codeSize > 20000 { // 20KB
        estimatedTime = "2-5 minutos"
        warningMessage = fmt.Sprintf("⚠️ Código mediano detectado\n\n"+
            "Tamaño: %d KB (%d líneas)\n"+
            "Tiempo estimado: %s\n\n"+
            "¿Continuar con la generación?", codeSize/1024, codeLines, estimatedTime)
    } else if codeSize > 5000 { // 5KB
        estimatedTime = "30-120 segundos"
    } else {
        estimatedTime = "10-30 segundos"
    }
    
    // MOSTRAR CONFIRMACIÓN PARA CÓDIGO GRANDE
    if warningMessage != "" {
        dialog.ShowConfirm("Confirmación de Procesamiento", warningMessage,
            func(confirmed bool) {
                if confirmed {
                    generateCSTWithWorkerPool(w, codeHash, estimatedTime, codeSize, codeLines)
                }
            }, w)
        return
    }
    
    // PROCEDER DIRECTAMENTE PARA CÓDIGO PEQUEÑO
    generateCSTWithWorkerPool(w, codeHash, estimatedTime, codeSize, codeLines)
}

func generateCSTWithWorkerPool(w fyne.Window, codeHash, estimatedTime string, codeSize, codeLines int) {
    // VERIFICAR DISPONIBILIDAD DEL POOL DE WORKERS
    select {
    case cstWorkerPool <- struct{}{}:
        // Worker disponible - continuar
    default:
        dialog.ShowInformation("Sistema Ocupado", 
            fmt.Sprintf("Ya se están generando %d CST simultáneamente.\n"+
                       "⏳ Intenta de nuevo en un momento.", cap(cstWorkerPool)), w)
        return
    }
    
    // MOSTRAR PROGRESO OPTIMIZADO CON ESTADÍSTICAS
    progress := widget.NewProgressBarInfinite()
    
    statsText := fmt.Sprintf("📊 Estadísticas del proceso:\n"+
        "• Tamaño del código: %d bytes (%.1f KB)\n"+
        "• Líneas a procesar: %d\n"+
        "• Tiempo estimado: %s\n"+
        "• Worker pool activo: ✅\n"+
        "• Cache habilitado: ✅", 
        codeSize, float64(codeSize)/1024, codeLines, estimatedTime)
    
    progressContent := container.NewVBox(
        widget.NewLabelWithStyle("🌳 Generando Árbol de Sintaxis Concreta", 
            fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
        widget.NewSeparator(),
        progress,
        widget.NewSeparator(),
        widget.NewLabel(statsText),
        widget.NewSeparator(),
        widget.NewLabelWithStyle("💡 El proceso se ejecuta en segundo plano", 
            fyne.TextAlignCenter, fyne.TextStyle{Italic: true}),
    )
    
    progressDialog := dialog.NewCustom("Generando CST", "Cancelar", progressContent, w)
    
    progress.Start()
    progressDialog.Show()

    // EJECUTAR EN WORKER POOL CON TIMEOUT DINÁMICO Y OPTIMIZACIONES
    go func() {
        defer func() {
            <-cstWorkerPool // LIBERAR WORKER SIEMPRE
            progress.Stop()
            progressDialog.Hide()
        }()
        
        start := time.Now()
        
        // TIMEOUT INTELIGENTE BASADO EN TAMAÑO
        var timeoutDuration time.Duration
        if codeSize > 50000 {
            timeoutDuration = 15 * time.Minute // Código muy grande
        } else if codeSize > 20000 {
            timeoutDuration = 8 * time.Minute  // Código grande
        } else if codeSize > 5000 {
            timeoutDuration = 3 * time.Minute  // Código mediano
        } else {
            timeoutDuration = 90 * time.Second // Código pequeño
        }
        
        ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
        defer cancel()
        
        // CANALES PARA MANEJO ASÍNCRONO
        resultChan := make(chan *interpreter.CSTResult, 1)
        errorChan := make(chan error, 1)
        
        // EJECUTAR GENERACIÓN CON OPTIMIZACIONES DE MEMORIA
        go func() {
            // Limpiar memoria antes de comenzar
            runtime.GC()
            runtime.GOMAXPROCS(2) // Limitar a 2 CPUs para no saturar
            
            result, err := cstClient.GenerateCST(lastExecutedCode)
            
            if err != nil {
                errorChan <- err
            } else {
                resultChan <- result
            }
            
            // Limpiar memoria después
            runtime.GC()
        }()
        
        // ESPERAR RESULTADO, ERROR O TIMEOUT
        select {
        case result := <-resultChan:
            executionTime := time.Since(start)
            
            if result.Error != "" {
                fyne.Do(func() {
                    dialog.ShowError(fmt.Errorf("❌ Error generando CST: %s", result.Error), w)
                })
                return
            }
            
            if result.CST == "" {
                fyne.Do(func() {
                    dialog.ShowInformation("Sin CST", 
                        "No se pudo generar el árbol CST para este código.", w)
                })
                return
            }
            
            // ASIGNAR CST Y GENERAR REPORTE
            currentCST = result.CST
            htmlContent := generateCSTHTMLContent()
            
            // GUARDAR EN CACHE
            reportCache.updateCSTReport(codeHash, lastExecutedCode, htmlContent)
            
            // ESCRIBIR ARCHIVO
            reportsDir := filepath.Join(".", "Reportes")
            if err := os.MkdirAll(reportsDir, 0755); err == nil {
                fileName := "arbol_cst.html"
                filePath := filepath.Join(reportsDir, fileName)
                if err := os.WriteFile(filePath, []byte(htmlContent), 0644); err == nil {
                    absPath, _ := filepath.Abs(filePath)
                    fmt.Printf("🌳 Nuevo CST generado en: %v (%s)\n", executionTime, absPath)
                    
                    fyne.Do(func() {
                        openInBrowser(absPath)
                        
                        dialog.ShowInformation("✅ Árbol CST Generado", 
                            fmt.Sprintf("CST generado exitosamente!\n\n"+
                                       "⚡ Tiempo: %v\n"+
                                       "💾 Tamaño: ~%d KB\n"+
                                       "🔄 Guardado en cache\n"+
                                       "🌐 Abierto en navegador", 
                                       executionTime, len(htmlContent)/1024), w)
                    })
                } else {
                    fyne.Do(func() {
                        dialog.ShowError(fmt.Errorf("error escribiendo archivo: %v", err), w)
                    })
                }
            } else {
                fyne.Do(func() {
                    dialog.ShowError(fmt.Errorf("error creando directorio: %v", err), w)
                })
            }
            
        case err := <-errorChan:
            fyne.Do(func() {
                dialog.ShowError(fmt.Errorf("❌ Error en generación CST: %v", err), w)
            })
            
        case <-ctx.Done():
            fyne.Do(func() {
                dialog.ShowError(fmt.Errorf("⏰ Timeout: CST tardó más de %v\n\n"+
                    "💡 Posibles causas:\n"+
                    "• Código demasiado complejo\n"+
                    "• Sistema sobrecargado\n"+
                    "• Insuficiente memoria disponible\n\n"+
                    "🔧 Recomendaciones:\n"+
                    "• Intenta con código más pequeño\n"+
                    "• Cierra otras aplicaciones\n"+
                    "• Intenta de nuevo en un momento", timeoutDuration), w)
            })
        }
    }()
}

func generateCSTHTMLContent() string {
    timestamp := time.Now().Format("02/01/2006 15:04:05")
    
    return fmt.Sprintf(`
<!DOCTYPE html>
<html lang="es" class="h-full">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Árbol CST - VLang Compiler</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <script>
        tailwind.config = {
            darkMode: 'media',
            theme: {
                extend: {
                    fontFamily: {
                        'mono': ['JetBrains Mono', 'Fira Code', 'Monaco', 'Consolas', 'monospace'],
                    },
                    colors: {
                        'green-25': '#f0fdf4',
                        'emerald-25': '#ecfdf5',
                        'green-950': '#14532d',
                        'emerald-950': '#064e3b'
                    }
                }
            }
        }
    </script>
    <style>
        .draggable-container {
            cursor: grab;
            user-select: none;
        }
        .draggable-container:active {
            cursor: grabbing;
        }
        .draggable-container.dragging {
            cursor: grabbing;
        }
    </style>
</head>
<body class="h-full bg-gradient-to-br from-green-50 via-white to-emerald-50 dark:from-gray-900 dark:via-gray-800 dark:to-green-900">
    <div class="min-h-full">
        <!-- Header -->
        <header class="bg-white dark:bg-gray-800 shadow-sm border-b border-gray-200 dark:border-gray-700">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div class="flex justify-between items-center py-6">
                    <div class="flex items-center space-x-4">
                        <div class="flex-shrink-0">
                            <div class="w-12 h-12 bg-gradient-to-r from-green-500 to-emerald-600 dark:from-green-600 dark:to-emerald-700 rounded-lg flex items-center justify-center">
                                <svg class="w-8 h-8 text-white" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.707a1 1 0 00-1.414-1.414L10 9.586 8.707 8.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div>
                            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">VLang Compiler</h1>
                            <p class="text-sm text-gray-500 dark:text-gray-400">Árbol de Sintaxis Concreta (CST)</p>
                        </div>
                    </div>
                    <div class="flex items-center space-x-4">
                        <span class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300 border border-green-200 dark:border-green-700">
                            Generado por ANTLR
                        </span>
                        <span class="text-sm text-gray-500 dark:text-gray-400">%s</span>
                    </div>
                </div>
            </div>
        </header>

        <!-- Main Content -->
        <main class="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
            <!-- Stats -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-green-100 dark:bg-green-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11.707 4.707a1 1 0 00-1.414-1.414L10 9.586 8.707 8.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Árbol CST</p>
                            <p class="text-2xl font-bold text-gray-900 dark:text-white">Generado</p>
                        </div>
                    </div>
                </div>
                
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-emerald-100 dark:bg-emerald-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-emerald-600 dark:text-emerald-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path fill-rule="evenodd" d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Analizador</p>
                            <p class="text-2xl font-bold text-gray-900 dark:text-white">ANTLR</p>
                        </div>
                    </div>
                </div>
                
                <div class="bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                    <div class="flex items-center">
                        <div class="flex-shrink-0">
                            <div class="w-10 h-10 bg-blue-100 dark:bg-blue-900 rounded-lg flex items-center justify-center">
                                <svg class="w-6 h-6 text-blue-600 dark:text-blue-400" fill="currentColor" viewBox="0 0 20 20">
                                    <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                                </svg>
                            </div>
                        </div>
                        <div class="ml-4">
                            <p class="text-sm font-medium text-gray-500 dark:text-gray-400">Estado</p>
                            <p class="text-lg font-bold text-green-600 dark:text-green-400">Exitoso</p>
                        </div>
                    </div>
                </div>
            </div>

            <!-- CST Tree -->
            <div class="bg-white dark:bg-gray-800 shadow-sm rounded-lg border border-gray-200 dark:border-gray-700 overflow-hidden">
                <div class="px-6 py-4 border-b border-gray-200 dark:border-gray-700 bg-gray-50 dark:bg-gray-700">
                    <h3 class="text-lg font-medium text-gray-900 dark:text-white">Árbol de Sintaxis Concreta</h3>
                    <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Representación visual del árbol de parsing generado por ANTLR</p>
                </div>
                
                <div id="tree-viewport" class="overflow-hidden bg-white dark:bg-gray-800 relative" style="height: 600px;">
                    <div id="svg-container" class="draggable-container absolute" style="left: 50%%; top: 50%%; transform: translate(-50%%, -50%%);">
                        %s
                    </div>
                </div>
            </div>

            <!-- Controls -->
            <div class="mt-6 bg-white dark:bg-gray-800 rounded-lg shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                <div class="flex items-center justify-between">
                    <div>
                        <h4 class="text-sm font-medium text-gray-900 dark:text-white">Controles de Visualización</h4>
                        <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Usa el mouse para arrastrar y los botones para hacer zoom</p>
                    </div>
                    <div class="flex space-x-3">
                        <button onclick="zoomIn()" class="inline-flex items-center px-3 py-2 border border-gray-300 dark:border-gray-600 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                            <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"/>
                            </svg>
                            Zoom +
                        </button>
                        <button onclick="zoomOut()" class="inline-flex items-center px-3 py-2 border border-gray-300 dark:border-gray-600 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                            <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM13 10H7"/>
                            </svg>
                            Zoom -
                        </button>
                        <button onclick="resetView()" class="inline-flex items-center px-3 py-2 border border-gray-300 dark:border-gray-600 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                            <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
                            </svg>
                            Reset
                        </button>
                        <button onclick="centerView()" class="inline-flex items-center px-3 py-2 border border-gray-300 dark:border-gray-600 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 dark:text-gray-300 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500">
                            <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z"/>
                            </svg>
                            Centrar
                        </button>
                    </div>
                </div>
            </div>

            <!-- Footer -->
            <div class="mt-8 bg-green-50 dark:bg-green-900 rounded-lg p-6 border border-green-200 dark:border-green-700">
                <div class="flex items-start space-x-3">
                    <div class="flex-shrink-0">
                        <svg class="w-6 h-6 text-green-600 dark:text-green-400" fill="currentColor" viewBox="0 0 20 20">
                            <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
                        </svg>
                    </div>
                    <div>
                        <h4 class="text-sm font-medium text-green-900 dark:text-green-100">Navegación</h4>
                        <p class="mt-1 text-sm text-green-700 dark:text-green-200">
                            📌 <strong>Arrastrar:</strong> Mantén presionado el mouse y arrastra para mover el árbol<br>
                            🔍 <strong>Zoom:</strong> Usa los botones + / - para acercar o alejar<br>
                            🎯 <strong>Centrar:</strong> Usa el botón "Centrar" para volver al centro del árbol
                        </p>
                    </div>
                </div>
            </div>
        </main>
    </div>


    <script>
        let currentZoom = 1;
        let isDragging = false;
        let lastX = 0;
        let lastY = 0;
        let translateX = 0;
        let translateY = 0;
        
        function zoomIn() {
            currentZoom += 0.2;
            updateTransformInstant();
            console.log('Zoom level:', Math.round(currentZoom * 100) + '%%');
        }
        
        function zoomOut() {
            if (currentZoom > 0.2) {
                currentZoom -= 0.2;
                updateTransformInstant();
                console.log('Zoom level:', Math.round(currentZoom * 100) + '%%');
            }
        }
        
        function resetView() {
            currentZoom = 1;
            translateX = 0;
            translateY = 0;
            updateTransformWithTransition();
        }
        
        function centerView() {
            translateX = 0;
            translateY = 0;
            updateTransformWithTransition();
        }
        
        // Actualización inmediata sin transición
        function updateTransformInstant() {
            const svgContainer = document.getElementById('svg-container');
            if (svgContainer) {
                svgContainer.style.transition = 'none';
                svgContainer.style.transform = 'translate(' + translateX + 'px, ' + translateY + 'px) scale(' + currentZoom + ')';
            }
        }
        
        // Actualización con transición suave (solo para botones)
        function updateTransformWithTransition() {
            const svgContainer = document.getElementById('svg-container');
            if (svgContainer) {
                svgContainer.style.transition = 'transform 0.3s ease-out';
                svgContainer.style.transform = 'translate(' + translateX + 'px, ' + translateY + 'px) scale(' + currentZoom + ')';
            }
        }

        document.addEventListener('DOMContentLoaded', function() {
            console.log('Árbol CST cargado exitosamente');
            
            const svgContainer = document.getElementById('svg-container');
            const viewport = document.getElementById('tree-viewport');
            
            if (svgContainer && viewport) {
                const svg = svgContainer.querySelector('svg');
                
                // Optimizar rendimiento del SVG
                if (svg) {
                    svg.style.pointerEvents = 'none'; // Evitar interferencia con eventos
                    
                    const isDarkMode = window.matchMedia('(prefers-color-scheme: dark)').matches;
                    if (isDarkMode) {
                        svg.style.filter = 'invert(0.9) hue-rotate(180deg)';
                    }
                    
                    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', function(e) {
                        if (e.matches) {
                            svg.style.filter = 'invert(0.9) hue-rotate(180deg)';
                        } else {
                            svg.style.filter = 'none';
                        }
                    });
                }
                
                // Optimizar el contenedor para mejor rendimiento
                svgContainer.style.willChange = 'transform';
                
                // Event listeners optimizados para desplazamiento inmediato
                viewport.addEventListener('mousedown', function(e) {
                    if (e.button === 0) {
                        isDragging = true;
                        lastX = e.clientX;
                        lastY = e.clientY;
                        svgContainer.classList.add('dragging');
                        viewport.style.cursor = 'grabbing';
                        e.preventDefault();
                        e.stopPropagation();
                    }
                });
                
                // Usar mousemove en el documento para capturar movimientos fuera del viewport
                document.addEventListener('mousemove', function(e) {
                    if (!isDragging) return;
                    
                    // Calcular desplazamiento inmediato
                    const deltaX = e.clientX - lastX;
                    const deltaY = e.clientY - lastY;
                    
                    // Aplicar movimiento instantáneo
                    translateX += deltaX;
                    translateY += deltaY;
                    
                    lastX = e.clientX;
                    lastY = e.clientY;
                    
                    // Actualización inmediata sin delay
                    updateTransformInstant();
                    
                    e.preventDefault();
                }, { passive: false });
                
                document.addEventListener('mouseup', function(e) {
                    if (isDragging) {
                        isDragging = false;
                        svgContainer.classList.remove('dragging');
                        viewport.style.cursor = 'grab';
                    }
                });
                
                // Cursor inicial
                viewport.style.cursor = 'grab';
                
                // Prevenir selección optimizada
                viewport.addEventListener('selectstart', function(e) {
                    e.preventDefault();
                });
                
                viewport.addEventListener('dragstart', function(e) {
                    e.preventDefault();
                });
                
                // Zoom optimizado con rueda del mouse
                viewport.addEventListener('wheel', function(e) {
                    if (e.ctrlKey) {
                        e.preventDefault();
                        
                        const zoomSpeed = 0.1;
                        const oldZoom = currentZoom;
                        
                        if (e.deltaY < 0) {
                            currentZoom = Math.min(currentZoom + zoomSpeed, 3);
                        } else {
                            currentZoom = Math.max(currentZoom - zoomSpeed, 0.2);
                        }
                        
                        if (oldZoom !== currentZoom) {
                            const rect = viewport.getBoundingClientRect();
                            const mouseX = e.clientX - rect.left;
                            const mouseY = e.clientY - rect.top;
                            
                            const zoomRatio = currentZoom / oldZoom;
                            translateX = mouseX - (mouseX - translateX) * zoomRatio;
                            translateY = mouseY - (mouseY - translateY) * zoomRatio;
                        }
                        
                        updateTransformInstant();
                        console.log('Zoom level:', Math.round(currentZoom * 100) + '%%');
                    }
                }, { passive: false });
                
                // Doble clic para centrar (con transición)
                viewport.addEventListener('dblclick', function(e) {
                    centerView();
                });
                
                // Limpiar will-change cuando no se está arrastrando para optimizar memoria
                document.addEventListener('mouseup', function() {
                    if (!isDragging) {
                        setTimeout(() => {
                            if (!isDragging) {
                                svgContainer.style.willChange = 'auto';
                            }
                        }, 100);
                    }
                });
                
                viewport.addEventListener('mousedown', function() {
                    svgContainer.style.willChange = 'transform';
                });
            }
        });
    </script>
</body>
</html>`,
        timestamp,
        currentCST)
}

// Función para obtener tipos únicos de errores
func getUniqueErrorTypes() int {
	typeMap := make(map[string]bool)
	for _, err := range currentErrors {
		typeMap[err.Type] = true
	}
	return len(typeMap)
}

func openInBrowser(filePath string) {
    // Comando simple para Ubuntu
    cmd := exec.Command("xdg-open", filePath)
    
    // Asegurar que tiene las variables de entorno básicas
    cmd.Env = os.Environ()
    
    err := cmd.Start()
    if err != nil {
        // Si xdg-open falla, intentar con firefox directamente
        fmt.Printf("xdg-open falló, intentando con firefox...\n")
        firefoxCmd := exec.Command("firefox", filePath)
        firefoxCmd.Env = os.Environ()
        
        if firefoxErr := firefoxCmd.Start(); firefoxErr != nil {
            fmt.Printf("❌ No se pudo abrir el navegador automáticamente\n")
            fmt.Printf("📁 Archivo guardado en: %s\n", filePath)
            fmt.Printf("💡 Abre manualmente con: firefox %s\n", filePath)
        } else {
            fmt.Printf("✅ Archivo abierto con Firefox\n")
        }
    } else {
        fmt.Printf("✅ Archivo abierto con navegador predeterminado\n")
    }
}

func startMemoryCleanup() {
    ticker := time.NewTicker(3 * time.Minute) // Cada 3 minutos
    
    go func() {
        defer ticker.Stop()
        
        for range ticker.C {
            // LIMPIAR MEMORIA
            runtime.GC()
            
            // MOSTRAR ESTADÍSTICAS DE CACHE
            printCacheStats()
            
            // LIMPIAR CACHE ANTIGUO
            reportCache.mu.Lock()
            if time.Since(reportCache.timestamp) > 15*time.Minute {
                oldEntries := 0
                if reportCache.errorReport != "" { oldEntries++ }
                if reportCache.symbolReport != "" { oldEntries++ }
                if reportCache.cstReport != "" { oldEntries++ }
                
                reportCache.lastCodeHash = ""
                reportCache.errorReport = ""
                reportCache.symbolReport = ""
                reportCache.cstReport = ""
                
                if oldEntries > 0 {
                    fmt.Printf("🧹 Cache automático limpiado (%d entradas expiradas)\n", oldEntries)
                }
            }
            reportCache.mu.Unlock()
            
            // LIMPIAR ARCHIVOS TEMPORALES ANTIGUOS
            cleanupTempFiles()
            
            // ESTADÍSTICAS DE MEMORIA MEJORADAS
            var m runtime.MemStats
            runtime.ReadMemStats(&m)
            if m.Alloc > 50*1024*1024 { // Si usa más de 50MB
                fmt.Printf("💾 Memoria en uso: %.1f MB (GC forzado)\n", float64(m.Alloc)/1024/1024)
                runtime.GC()
            }
        }
    }()
}

func cleanupTempFiles() {
    reportsDir := "Reportes"
    if entries, err := os.ReadDir(reportsDir); err == nil {
        cutoff := time.Now().Add(-2 * time.Hour) // Limpiar archivos de 2+ horas
        cleaned := 0
        totalSize := int64(0)
        
        for _, entry := range entries {
            if info, err := entry.Info(); err == nil {
                if info.ModTime().Before(cutoff) {
                    filePath := filepath.Join(reportsDir, entry.Name())
                    totalSize += info.Size()
                    if err := os.Remove(filePath); err == nil {
                        cleaned++
                    }
                }
            }
        }
        
        if cleaned > 0 {
            fmt.Printf("🧹 Limpiados %d archivos temporales (%.1f MB liberados)\n", 
                      cleaned, float64(totalSize)/1024/1024)
        }
    }
}

// AGREGAR: Función para estadísticas de cache
func printCacheStats() {
    reportCache.mu.RLock()
    defer reportCache.mu.RUnlock()
    
    cacheEntries := 0
    if reportCache.errorReport != "" {
        cacheEntries++
    }
    if reportCache.symbolReport != "" {
        cacheEntries++
    }
    if reportCache.cstReport != "" {
        cacheEntries++
    }
    
    if cacheEntries > 0 {
        age := time.Since(reportCache.timestamp)
        fmt.Printf("📊 Cache activo: %d entradas, edad: %v\n", cacheEntries, age.Round(time.Second))
    }
}