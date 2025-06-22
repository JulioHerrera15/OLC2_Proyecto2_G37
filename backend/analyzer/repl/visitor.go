package repl

import (
	parser "backend/parser"
	"fmt"
	"log"
	"strconv"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Estructura para representar un tipo de struct
type StructDefinition struct {
	Name   string
	Fields map[string]string
}

// Estructura para representar una instancia de struct
type StructInstance struct {
	TypeName string
	Fields   map[string]interface{}
}

type FunctionDefinition struct {
    Name       string
    Parameters []Parameter
    ReturnType string
    Body       antlr.ParseTree
}

type Parameter struct {
    Name string
    Type string
}

type ReplVisitor struct {
    parser.BaseLanguageVisitor 
    SymbolTable                map[string]interface{}
    ErrorTable                 *ErrorTable
    Console                    *Console
    ScopeTrace                 *ScopeTrace
    MainFunction               *parser.FunctionDeclarationContext
    SymbolsTable               *SymbolTable
    StructTypes                map[string]*StructDefinition
    Functions                  map[string]*FunctionDefinition
    recursionCount             int 
    shouldReturn               bool       
    returnValue                interface{}
}

// MODIFICAR NewReplVisitor para inicializar Functions:
func NewReplVisitor() *ReplVisitor {
    visitor := &ReplVisitor{
        BaseLanguageVisitor: parser.BaseLanguageVisitor{},
        SymbolTable:         make(map[string]interface{}),
        ErrorTable:          NewErrorTable(),
        Console:             NewConsole(),
        ScopeTrace:          NewScopeTrace(),
        MainFunction:        nil,
        SymbolsTable:        NewSymbolTable(),
        StructTypes:         make(map[string]*StructDefinition),
        Functions:           make(map[string]*FunctionDefinition),
    }
    return visitor
}

// Obtener el ámbito actual basado en el ScopeTrace
func (v *ReplVisitor) getCurrentScope() string {
	if len(v.ScopeTrace.ScopeStack) <= 1 {
		return "Global"
	}

	// Revisar todos los scopes para encontrar funciones (desde el más reciente)
	for i := len(v.ScopeTrace.ScopeStack) - 1; i >= 0; i-- {
		scope := v.ScopeTrace.ScopeStack[i]

		// Si encontramos un scope de función, devolver el nombre específico
		if scope.ScopeType == "function" {
			// Si estamos en función main, devolver "main"
			if v.MainFunction != nil {
				return "main"
			}
			// Para otras funciones, podrías devolver el nombre específico
			return "Función"
		}
	}

	// Si no hay función pero hay otros scopes (for, if, switch, etc.), es Local
	if len(v.ScopeTrace.ScopeStack) > 1 {
		return "Local"
	}

	return "Global"
}

func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {
    switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}

func (v *ReplVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	// Primera pasada: registrar todas las funciones (especialmente main)
	for _, stmt := range ctx.AllStatement() {
		if stmt.FunctionDeclaration() != nil {
			v.Visit(stmt.FunctionDeclaration())
		}
	}

	// Segunda pasada: ejecutar declaraciones globales (que no son funciones)
	for _, stmt := range ctx.AllStatement() {
		if stmt.FunctionDeclaration() == nil {
			v.Visit(stmt)
		}
	}

	// Tercera pasada: ejecutar función main si existe
	if v.MainFunction != nil {
		return v.ExecuteMainFunction()
	} else {
		return nil
	}
}

func (v *ReplVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
    // Verificar cada tipo de statement que tu gramática soporta
    if ctx.VariableDeclaration() != nil {
        return v.Visit(ctx.VariableDeclaration())
    }
    if ctx.FunctionDeclaration() != nil {
        return v.Visit(ctx.FunctionDeclaration())
    }
    if ctx.StructDeclaration() != nil {
        return v.Visit(ctx.StructDeclaration())
    }
    if ctx.NonDeclaration() != nil {
        return v.Visit(ctx.NonDeclaration())
    }

    // Si no coincide con ninguno, iterar por hijos manualmente
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        
        if child != nil {
            if parseTree, ok := child.(antlr.ParseTree); ok {
                result := v.Visit(parseTree)
                if result != nil {
                    return result
                }
            }
        }
    }

    return nil
}

func (v *ReplVisitor) VisitNonDeclaration(ctx *parser.NonDeclarationContext) interface{} {
    if ctx.PrintStatement() != nil {
        return v.Visit(ctx.PrintStatement())
    }
    if ctx.IfStatement() != nil {                    
        return v.Visit(ctx.IfStatement())
    }
    if ctx.AtoiStatement() != nil {
        return v.Visit(ctx.AtoiStatement())
    }
    if ctx.ParseFloatStatement() != nil {           
        return v.Visit(ctx.ParseFloatStatement())
    }
    if ctx.TypeOfStatement() != nil {               
        return v.Visit(ctx.TypeOfStatement())
    }
    if ctx.SwitchStatement() != nil {               
        return v.Visit(ctx.SwitchStatement())
    }
    if ctx.ForStatement() != nil {                  
        return v.Visit(ctx.ForStatement())
    }
    if ctx.TransferenceStatement() != nil {         
        return v.Visit(ctx.TransferenceStatement())
    }
    if ctx.BlockStatement() != nil {                
        return v.Visit(ctx.BlockStatement())
    }
    if ctx.ExpressionStatement() != nil {
        return v.Visit(ctx.ExpressionStatement())
    }
    
    return nil
}
func (v *ReplVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	var output strings.Builder

	// Obtener todas las expresiones
	expressions := ctx.AllExpressionStatement()

	for i, exprCtx := range expressions {
		expr := v.Visit(exprCtx)

		if expr != nil {
			if i > 0 {
				output.WriteString(" ")
			}

			switch val := expr.(type) {
			case string:
				output.WriteString(val)
			case int:
				output.WriteString(strconv.Itoa(val))
			case float64:
				output.WriteString(fmt.Sprintf("%f", val))
			case bool:
				output.WriteString(strconv.FormatBool(val))
			case []interface{}:
				output.WriteString(formatSlice(val))
			case [][]interface{}:
				output.WriteString(formatMatrix(val))
			default:
				output.WriteString(fmt.Sprintf("%v", val))
			}
		}
	}

	if output.Len() > 0 {
		v.Console.Print(output.String())
	}

	return nil
}

// Función auxiliar para formatear slices
func formatSlice(slice []interface{}) string {
	var elements []string

	for _, element := range slice {
		switch val := element.(type) {
		case string:
			elements = append(elements, fmt.Sprintf("\"%s\"", val))
		case int:
			elements = append(elements, strconv.Itoa(val))
		case float64:
			elements = append(elements, fmt.Sprintf("%.2f", val))
		case bool:
			elements = append(elements, strconv.FormatBool(val))
		default:
			elements = append(elements, fmt.Sprintf("%v", val))
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

// Función auxiliar para formatear matrices
func formatMatrix(matrix [][]interface{}) string {
	var rows []string

	for _, row := range matrix {
		rows = append(rows, formatSlice(row))
	}

	return fmt.Sprintf("[%s]", strings.Join(rows, ", "))
}

func (v *ReplVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
    
    child := ctx.GetChild(0).(antlr.ParseTree)
    // Usar Visit genérico para manejar cualquier tipo de expresión
    result := v.Visit(child)
    
    return result
}
// Método para manejar strings
func (v *ReplVisitor) VisitString(ctx *parser.StringContext) interface{} {
	// Obtener el texto completo incluyendo comillas
	text := ctx.GetText()

	// Verificar que tiene comillas
	if !((strings.HasPrefix(text, "\"") && strings.HasSuffix(text, "\"")) ||
		(strings.HasPrefix(text, "'") && strings.HasSuffix(text, "'"))) {
		return text
	}

	// Remover las comillas externas
	content := text[1 : len(text)-1]

	// Procesar secuencias de escape
	result := v.processEscapeSequences(content)

	return result
}

// Función para procesar secuencias de escape
func (v *ReplVisitor) processEscapeSequences(input string) string {
	var result strings.Builder

	i := 0
	for i < len(input) {
		if input[i] == '\\' && i+1 < len(input) {
			// Procesar secuencia de escape
			switch input[i+1] {
			case '"':
				result.WriteByte('"') // \"  → "
			case '\\':
				result.WriteByte('\\') // \\  → \
			case 'n':
				result.WriteByte('\n') // \n  → salto de línea
			case 'r':
				result.WriteByte('\r') // \r  → retorno de carro
			case 't':
				result.WriteByte('\t') // \t  → tabulación
			default:
				// Si no es una secuencia válida, mantener como está
				result.WriteByte('\\')
				result.WriteByte(input[i+1])
				v.ErrorTable.NewSemanticError(nil,
					fmt.Sprintf("Secuencia de escape no válida: \\%c", input[i+1]))
			}
			i += 2 // Saltar ambos caracteres
		} else {
			// Carácter normal
			result.WriteByte(input[i])
			i++
		}
	}

	return result.String()
}

func (v *ReplVisitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	text := ctx.GetText()
	num, _ := strconv.Atoi(text)
	return num
}

func (v *ReplVisitor) VisitFloat(ctx *parser.FloatContext) interface{} {
	text := ctx.GetText()
	num, _ := strconv.ParseFloat(text, 64)
	return num
}

func (v *ReplVisitor) VisitBoolean(ctx *parser.BooleanContext) interface{} {
	text := ctx.GetText()
	return text == "true"
}

func (v *ReplVisitor) VisitParens(ctx *parser.ParensContext) interface{} {
	result := v.Visit(ctx.ExpressionStatement())

	start := ctx.GetStart().GetText()
	end := ctx.GetStop().GetText()

	if (start == "(" && end == ")") || (start == "[" && end == "]") {

		return result
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Paréntesis no balanceados: '%s' y '%s'", start, end))
	return nil
}

func (v *ReplVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
    left := v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    op := ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
    right := v.Visit(ctx.GetChild(2).(antlr.ParseTree))
    
    // Validar valores no nulos
    if left == nil || right == nil {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación con valores nulos")
        return nil
    }
    
    if op == "+" {
        return add(left, right)
    } else {
        return sub(left, right)
    }
}

func (v *ReplVisitor) VisitMulDivMod(ctx *parser.MulDivModContext) interface{} {
	left := v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	right := v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	switch op {
	case "*":
		if left == nil || right == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación de multiplicación con valores nulos")
			return nil
		}
		return mul(left, right)
	case "/":
		if right == 0 || right == 0.0 {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "División por cero")
			return nil
		}
		if left == nil || right == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación de división con valores nulos")
			return nil
		}
		return div(left, right)
	case "%":
		if right == 0 {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Módulo por cero")
			return nil
		} else if _, ok := left.(int); !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El operador '%' solo es válido para enteros")
			return nil
		}
		if left == nil || right == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación de módulo con valores nulos")
			return nil
		}
		return mod(left, right)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Operador '%s' no válido", op))
	}
	return nil
}

func (v *ReplVisitor) VisitGreaterLess(ctx *parser.GreaterLessContext) interface{} {
    left := v.Visit(ctx.ExpressionStatement(0))
    operator := ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
    right := v.Visit(ctx.ExpressionStatement(1))

    // Debug temporal
    if left == nil || right == nil {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación de comparación con valores nulos")
        return false
    }

    leftFloat, leftOk := toFloat64(left)
    rightFloat, rightOk := toFloat64(right)

    if !leftOk || !rightOk {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipos incompatibles para comparación")
        return false
    }

    // LÓGICA DE COMPARACIÓN
    var result bool
    switch operator {
    case ">":
        result = leftFloat > rightFloat
    case "<":
        result = leftFloat < rightFloat
    case ">=":
        result = leftFloat >= rightFloat
    case "<=":
        result = leftFloat <= rightFloat
    default:
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operador de comparación no válido")
        return false
    }

    return result
}

func (v *ReplVisitor) VisitEqual(ctx *parser.EqualContext) interface{} {
	left := v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()
	right := v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	// CAMBIAR: En lugar de retornar nil, manejar nils correctamente
	if left == nil && right == nil {
		switch op {
		case "==":
			return true
		case "!=":
			return false
		}
	}

	if left == nil || right == nil {
		switch op {
		case "==":
			return false
		case "!=":
			return true
		}
	}

	switch op {
	case "==":
		return isEqual(left, right)
	case "!=":
		return !isEqual(left, right)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Operador '%s' no válido", op))
		return false
	}
}

func (v *ReplVisitor) VisitAddSubOperator(ctx *parser.AddSubOperatorContext) interface{} {
	varName := ctx.ID().GetText()
	operator := ""

	// Encontrar el operador += o -=
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if terminal, ok := child.(antlr.TerminalNode); ok {
			text := terminal.GetText()
			if text == "+=" || text == "-=" {
				operator = text
				break
			}
		}
	}

	// Evaluar la expresión del lado derecho
	rightValue := v.Visit(ctx.ExpressionStatement())

	// Obtener el valor actual de la variable
	currentValue, exists := v.GetVariable(varName)
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no está declarada", varName))
		return nil
	}

	// VALIDAR TIPOS ANTES DE LA OPERACIÓN
	if !v.areTypesCompatibleForArithmetic(currentValue, rightValue) {
		currentType := v.getTypeName(currentValue)
		rightType := v.getTypeName(rightValue)
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("No se puede aplicar '%s' entre tipos '%s' y '%s'", operator, currentType, rightType))
		return nil
	}

	var newValue interface{}

	switch operator {
	case "+=":
		newValue = v.performAddition(currentValue, rightValue)
	case "-=":
		newValue = v.performSubtraction(currentValue, rightValue)
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Operador '%s' no reconocido", operator))
		return nil
	}

	if newValue != nil {
		v.SetVariable(varName, newValue)
		return newValue
	}

	return nil
}

// Validar tipos para operaciones aritméticas
func (v *ReplVisitor) areTypesCompatibleForArithmetic(left, right interface{}) bool {
	leftType := v.getTypeName(left)
	rightType := v.getTypeName(right)

	// Números compatibles entre sí
	if (leftType == "int" || leftType == "float64") &&
		(rightType == "int" || rightType == "float64") {
		return true
	}

	// Strings para concatenación (+= solamente)
	if leftType == "string" && rightType == "string" {
		return true
	}

	return false
}

// Métodos auxiliares para las operaciones
func (v *ReplVisitor) performAddition(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int:
		if r, ok := right.(int); ok {
			return l + r
		}
		if r, ok := right.(float64); ok {
			return float64(l) + r
		}
	case float64:
		if r, ok := right.(float64); ok {
			return l + r
		}
		if r, ok := right.(int); ok {
			return l + float64(r)
		}
	case string:
		if r, ok := right.(string); ok {
			return l + r
		}
	}
	return nil
}

func (v *ReplVisitor) performSubtraction(left, right interface{}) interface{} {
	switch l := left.(type) {
	case int:
		if r, ok := right.(int); ok {
			return l - r
		}
		if r, ok := right.(float64); ok {
			return float64(l) - r
		}
	case float64:
		if r, ok := right.(float64); ok {
			return l - r
		}
		if r, ok := right.(int); ok {
			return l - float64(r)
		}
	}
	return nil
}

// BUSCAR VisitAnd y agregar manejo defensivo:
func (v *ReplVisitor) VisitAnd(ctx *parser.AndContext) interface{} {
	left := v.Visit(ctx.ExpressionStatement(0))
	right := v.Visit(ctx.ExpressionStatement(1))

	// Manejo defensivo de nils
	if left == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación AND con valores nulos")
		return false
	}
	if right == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación AND con valores nulos")
		return false
	}

	// Convertir a bool de forma segura
	leftBool := v.toBool(left)
	rightBool := v.toBool(right)

	return leftBool && rightBool
}

// Función helper:
func (v *ReplVisitor) toBool(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0.0
	case string:
		return v != ""
	default:
		return true
	}
}

func (v *ReplVisitor) VisitOr(ctx *parser.OrContext) interface{} {
	left := v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	right := v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	if left == nil || right == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede realizar la operación OR con valores nulos")
		return nil
	}

	leftBool, _ := left.(bool)
	rightBool, _ := right.(bool)

	return leftBool || rightBool
}

func (v *ReplVisitor) VisitNegate(ctx *parser.NegateContext) interface{} {
	op := ctx.GetChild(0).(antlr.TerminalNode).GetText()
	value := v.Visit(ctx.GetChild(1).(antlr.ParseTree))

	if value == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "No se puede aplicar el operador de negación a un valor nulo")
		return nil
	}

	switch op {
	case "-":
		switch val := value.(type) {
		case int:
			return -val
		case float64:
			return -val
		default:
			if num, ok := toFloat64(value); ok {
				return -num
			}
		}
	case "!":
		if b, ok := value.(bool); ok {
			return !b
		}
	}
	v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Operador '%s' no válido para tipo '%T'", op, value))

	return nil
}

func isEqual(a, b interface{}) bool {
	switch aVal := a.(type) {
	case int:
		if bVal, ok := b.(int); ok {
			return aVal == bVal
		}
		if bVal, ok := b.(float64); ok {
			return float64(aVal) == bVal
		}
	case float64:
		if bVal, ok := b.(float64); ok {
			return aVal == bVal
		}
		if bVal, ok := b.(int); ok {
			return aVal == float64(bVal)
		}
	case string:
		if bVal, ok := b.(string); ok {
			return aVal == bVal
		}
	case bool:
		if bVal, ok := b.(bool); ok {
			return aVal == bVal
		}
	}
	return false
}

func toFloat64(val interface{}) (float64, bool) {
	switch v := val.(type) {
	case int:
		return float64(v), true
	case float64:
		return v, true
	case string:
		f, _ := strconv.ParseFloat(v, 64)
		return f, true
	default:
		return 0, false
	}
}

func add(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a + b
		case float64:
			return float64(a) + b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a + float64(b)
		case float64:
			return a + b
		}
	case string:
		if bStr, ok := b.(string); ok {
			return a + bStr
		}
	}
	return 0
}

func sub(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a - b
		case float64:
			return float64(a) - b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a - float64(b)
		case float64:
			return a - b
		}
	}
	return 0
}

func mul(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a * b
		case float64:
			return float64(a) * b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a * float64(b)
		case float64:
			return a * b
		}
	}
	return 0
}

func div(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		switch b := b.(type) {
		case int:
			return a / b
		case float64:
			return float64(a) / b
		}
	case float64:
		switch b := b.(type) {
		case int:
			return a / float64(b)
		case float64:
			return a / b
		}
	}
	return 0
}

func mod(a, b interface{}) interface{} {
	aInt, aOk := a.(int)
	bInt, bOk := b.(int)
	if aOk && bOk {
		return aInt % bInt
	}
	return 0
}

func (v *ReplVisitor) VisitAtoiStatement(ctx *parser.AtoiStatementContext) interface{} {
    
    val := v.Visit(ctx.ExpressionStatement())
    
    strVal, ok := val.(string)
    if !ok {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "atoi: El argumento no es una cadena")
        return nil
    }
    
    intVal, err := strconv.Atoi(strVal)
    if err != nil {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("atoi: No se pudo convertir '%s' a entero", strVal))
        return nil
    }
    
    return intVal
}

func (v *ReplVisitor) VisitParseFloatStatement(ctx *parser.ParseFloatStatementContext) interface{} {
    val := v.Visit(ctx.ExpressionStatement())
    strVal, ok := val.(string)
    if !ok {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "parseFloat: El argumento no es una cadena")
        return nil
    }
    
    floatVal, err := strconv.ParseFloat(strVal, 64)
    if err != nil {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("parseFloat: No se pudo convertir '%s' a float64", strVal))
        return nil
    }
    
    return floatVal
}

func (v *ReplVisitor) VisitTypeOfStatement(ctx *parser.TypeOfStatementContext) interface{} {
    val := v.Visit(ctx.ExpressionStatement())
    
    if val == nil {
        return "nil"
    }
    
    switch val := val.(type) {
    case int:
        return "int"
    case float64:
        return "f64"
    case string:
        return "string"
    case bool:
        return "bool"
    case []interface{}:
        // Intentar inferir el tipo del slice basándose en el primer elemento
        if len(val) > 0 {
            firstElementType := v.getElementType(val[0])
            return fmt.Sprintf("[]%s", firstElementType)
        }
        return "[]unknown"
    case [][]interface{}:
        // Para matrices
        if len(val) > 0 && len(val[0]) > 0 {
            firstElementType := v.getElementType(val[0][0])
            return fmt.Sprintf("[][]%s", firstElementType)
        }
        return "[][]unknown"
    case *StructInstance:
        return val.TypeName
    default:
        return "unknown"
    }
}

func (v *ReplVisitor) VisitAtoiExpr(ctx *parser.AtoiExprContext) interface{} {
    return v.Visit(ctx.AtoiStatement())
}

func (v *ReplVisitor) VisitParseFloatExpr(ctx *parser.ParseFloatExprContext) interface{} {
    return v.Visit(ctx.ParseFloatStatement())
}

func (v *ReplVisitor) VisitTypeOfExpr(ctx *parser.TypeOfExprContext) interface{} {
    return v.Visit(ctx.TypeOfStatement())
}

// Función auxiliar para obtener el tipo de un elemento
func (v *ReplVisitor) getElementType(element interface{}) string {
	switch e := element.(type) {
	case int:
		return "int"
	case float64:
		return "f64"
	case string:
		return "string"
	case bool:
		return "bool"
	case *StructInstance:
		return e.TypeName
	default:
		return "unknown"
	}
}

func (v *ReplVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
    if ctx == nil {
        return nil
    }
    
    id := ctx.GetText()
    
    // Buscar en SymbolTable global
    if value, exists := v.SymbolTable[id]; exists {
        return value
    }
    
    // Buscar en ScopeTrace
    if value, exists := v.ScopeTrace.GetVariable(id); exists {
        return value
    }
    
    v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Variable '%s' no está declarada", id))
    return nil
}

func (v *ReplVisitor) VisitExplicitDeclaration(ctx *parser.ExplicitDeclarationContext) interface{} {
	var varName string
	var varValue interface{}
	var isMutable bool = false

	// Recorrer todos los hijos para encontrar ID, TYPE, mut, etc.
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		if terminal, ok := child.(antlr.TerminalNode); ok {
			tokenType := terminal.GetSymbol().GetTokenType()
			text := terminal.GetText()

			if text == "mut" {
				isMutable = true
				continue
			}

			if tokenType == parser.LanguageLexerID {
				varName = text
				continue
			}
		}
	}

	// Verificar redeclaración
	if varName != "" {
        if v.ScopeTrace.ExistsInCurrentScope(varName) {
            v.ErrorTable.NewSemanticError(ctx.GetStart(),
                fmt.Sprintf("Error de redeclaración: la variable '%s' ya está declarada en este scope", varName))
            return nil
        }
    }

	// Procesar valor
	if ctx.ExpressionStatement() != nil {
		varValue = v.Visit(ctx.ExpressionStatement())
	} else {
		// Valor por defecto según el tipo
		if ctx.TYPE() != nil {
			switch ctx.TYPE().GetText() {
			case "int":
				varValue = 0
			case "float64", "float":
				varValue = 0.0
			case "string":
				varValue = ""
			case "bool":
				varValue = false
			default:
				varValue = nil
			}
		}
	}

	// Registrar símbolo CON información de mutabilidad
	if varName != "" && ctx.TYPE() != nil {
		dataType := ctx.TYPE().GetText()
		if isMutable {
			dataType += " (mut)"
		}

		symbol := Symbol{
			ID:         varName,
			SymbolType: VariableSymbol,
			DataType:   dataType,
			Scope:      v.getCurrentScope(),
			Line:       ctx.GetStart().GetLine(),
			Column:     ctx.GetStart().GetColumn(),
		}
		v.SymbolsTable.AddSymbol(symbol)

		if varName != "" {
			v.SetVariable(varName, varValue)
		}
	}

	if varName != "" {
    // Si estamos en un scope local y la variable existe globalmente,
    // crear una variable local que "sombree" la global
    if len(v.ScopeTrace.ScopeStack) > 1 {
        if _, existsGlobally := v.SymbolTable[varName]; existsGlobally {
            // Crear variable local que sombrea la global
            v.ScopeTrace.CreateLocalVariable(varName, varValue)
            return nil
        }
    }
    
    v.SetVariable(varName, varValue)
}

	return nil
}

func (v *ReplVisitor) VisitImplicitDeclaration(ctx *parser.ImplicitDeclarationContext) interface{} {
    var varName string
    var varValue interface{}
    var isMutable bool = false

    // Buscar 'mut', ID y valor
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)

        if terminal, ok := child.(antlr.TerminalNode); ok {
            text := terminal.GetText()

            if text == "mut" {
                isMutable = true
                continue
            }

            if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
                varName = text
                continue
            }
        }
    }

    if varName != "" {
        if v.ScopeTrace.ExistsInCurrentScope(varName) {
            v.ErrorTable.NewSemanticError(ctx.GetStart(),
                fmt.Sprintf("Error de redeclaración: la variable '%s' ya está declarada en este scope", varName))
            return nil
        }
    }

    // Evaluar la expresión del lado derecho
    if ctx.ExpressionStatement() != nil {
        varValue = v.Visit(ctx.ExpressionStatement())
    }

    // Inferir el tipo y registrar el símbolo
    if varName != "" {
        dataType := "unknown"
        switch varValue.(type) {
        case int:
            dataType = IntType
        case float64:
            dataType = FloatType
        case string:
            dataType = StringType
        case bool:
            dataType = BoolType
        }

        if isMutable {
            dataType += " (mut)"
        }

        symbol := Symbol{
            ID:         varName,
            SymbolType: VariableSymbol,
            DataType:   dataType,
            Scope:      v.getCurrentScope(),
            Line:       ctx.GetStart().GetLine(),
            Column:     ctx.GetStart().GetColumn(),
        }
        v.SymbolsTable.AddSymbol(symbol)

        v.SetVariable(varName, varValue)
    }

    return nil
}



// Métodos para declaraciones de slices y matrices
func (v *ReplVisitor) VisitExplicitSliceDeclaration(ctx *parser.ExplicitSliceDeclarationContext) interface{} {
    var varName string
    var sliceValue []interface{}

    // Buscar el ID
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        if terminal, ok := child.(antlr.TerminalNode); ok {
            if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
                varName = terminal.GetText()
                break
            }
        }
    }

    // Registrar el símbolo como Slice
    if varName != "" && ctx.TYPE() != nil {
        symbol := Symbol{
            ID:         varName,
            SymbolType: SliceSymbol,
            DataType:   fmt.Sprintf("[]%s", ctx.TYPE().GetText()),
            Scope:      v.getCurrentScope(),
            Line:       ctx.GetStart().GetLine(),
            Column:     ctx.GetStart().GetColumn(),
        }
        v.SymbolsTable.AddSymbol(symbol)
    }

    // Manejar tanto SliceElements como ExpressionStatement
    if ctx.SliceElements() != nil {
        sliceValue = v.Visit(ctx.SliceElements()).([]interface{})
    } else if ctx.ExpressionStatement() != nil {
        // Manejar asignación directa (como matriz[0])
        assignedValue := v.Visit(ctx.ExpressionStatement())
        
        // Si es un slice (fila de matriz), usarlo directamente
        if slice, ok := assignedValue.([]interface{}); ok {
            sliceValue = slice
        } else {
            // Si es un valor individual, crear slice con ese valor
            sliceValue = []interface{}{assignedValue}
        }
    } else {
        sliceValue = make([]interface{}, 0)
    }

    if varName != "" {
        v.SetVariable(varName, sliceValue)
    }

    return nil
}

func (v *ReplVisitor) VisitImplicitSliceDeclaration(ctx *parser.ImplicitSliceDeclarationContext) interface{} {
	// Para declaraciones como: numbers := []int{1, 2, 3}

	var varName string
	var sliceValue []interface{}

	// Buscar el ID
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if terminal, ok := child.(antlr.TerminalNode); ok {
			if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
				varName = terminal.GetText()
				break
			}
		}
	}

	// Inferir el tipo y registrar el símbolo
	if varName != "" {
		dataType := "unknown"
		if ctx.TYPE() != nil {
			dataType = fmt.Sprintf("[]%s", ctx.TYPE().GetText())
			symbol := Symbol{
				ID:         varName,
				SymbolType: SliceSymbol,
				DataType:   dataType,
				Scope:      v.getCurrentScope(),
				Line:       ctx.GetStart().GetLine(),
				Column:     ctx.GetStart().GetColumn(),
			}
			v.SymbolsTable.AddSymbol(symbol)
		}
	}

	// Los elementos siempre están presentes en declaraciones implícitas
	if ctx.SliceElements() != nil {
		sliceValue = v.Visit(ctx.SliceElements()).([]interface{})
	}

	if varName != "" {
		v.SetVariable(varName, sliceValue)
	}

	return nil
}

// Método para manejar elementos de slice
func (v *ReplVisitor) VisitSliceElements(ctx *parser.SliceElementsContext) interface{} {
	var elements []interface{}

	// Obtener todas las expresiones dentro de los elementos
	expressions := ctx.AllExpressionStatement()

	for _, exprCtx := range expressions {
		value := v.Visit(exprCtx)
		elements = append(elements, value)
	}

	return elements
}

func (v *ReplVisitor) VisitExplicitMatrixDeclaration(ctx *parser.ExplicitMatrixDeclarationContext) interface{} {
	// Para declaraciones como: int[][] matrix = {{1, 2}, {3, 4}}

	var varName string
	var matrixValue [][]interface{}

	// Buscar el ID
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if terminal, ok := child.(antlr.TerminalNode); ok {
			if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
				varName = terminal.GetText()
				break
			}
		}
	}

	// Registrar el símbolo como Matrix
	if varName != "" && ctx.TYPE() != nil {
		symbol := Symbol{
			ID:         varName,
			SymbolType: MatrixSymbol,
			DataType:   fmt.Sprintf("[][]%s", ctx.TYPE().GetText()),
			Scope:      v.getCurrentScope(),
			Line:       ctx.GetStart().GetLine(),
			Column:     ctx.GetStart().GetColumn(),
		}
		v.SymbolsTable.AddSymbol(symbol)
	}

	// Si hay elementos iniciales
	if ctx.MatrixElements() != nil {
		matrixValue = v.Visit(ctx.MatrixElements()).([][]interface{})
	} else {
		matrixValue = make([][]interface{}, 0)
	}

	if varName != "" {
		v.SetVariable(varName, matrixValue)
	}

	return nil
}

func (v *ReplVisitor) VisitImplicitMatrixDeclaration(ctx *parser.ImplicitMatrixDeclarationContext) interface{} {
	var varName string
	var matrixValue [][]interface{}

	// Buscar el ID
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if terminal, ok := child.(antlr.TerminalNode); ok {
			if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
				varName = terminal.GetText()
				break
			}
		}
	}

	// Inferir el tipo y registrar el símbolo
	if varName != "" {
		dataType := "unknown"
		if ctx.TYPE() != nil {
			dataType = fmt.Sprintf("[][]%s", ctx.TYPE().GetText())
			symbol := Symbol{
				ID:         varName,
				SymbolType: MatrixSymbol,
				DataType:   dataType,
				Scope:      v.getCurrentScope(),
				Line:       ctx.GetStart().GetLine(),
				Column:     ctx.GetStart().GetColumn(),
			}
			v.SymbolsTable.AddSymbol(symbol)
		}
	}

	// Los elementos siempre están presentes
	if ctx.MatrixElements() != nil {
		matrixValue = v.Visit(ctx.MatrixElements()).([][]interface{})
	}

	if varName != "" {
		v.SetVariable(varName, matrixValue)
	}

	return nil
}

func (v *ReplVisitor) VisitMatrixElements(ctx *parser.MatrixElementsContext) interface{} {
	var rows [][]interface{}

	// Obtener todas las filas
	for _, rowCtx := range ctx.AllSliceElements() {
		row := v.Visit(rowCtx).([]interface{})
		rows = append(rows, row)
	}

	return rows
}

func (v *ReplVisitor) VisitSliceAccess(ctx *parser.SliceAccessContext) interface{} {
    sliceName := ctx.ID().GetText()
    indexValue := v.Visit(ctx.ExpressionStatement())

    index, ok := indexValue.(int)
    if !ok {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "El índice debe ser un entero")
        return nil
    }

    if slice, exists := v.GetVariable(sliceName); exists {
        // CASO 1: Es un slice normal
        if sliceData, ok := slice.([]interface{}); ok {
            if index >= 0 && index < len(sliceData) {
                result := sliceData[index]
                return result
            } else {
                v.ErrorTable.NewSemanticError(ctx.GetStart(),
                    fmt.Sprintf("Índice %d fuera de rango para slice [0:%d]", index, len(sliceData)-1))
                return nil
            }
        }

        // CASO 2: Es una matriz, devolver la fila como slice - CORREGIDO
        if matrixData, ok := slice.([][]interface{}); ok {
            if index >= 0 && index < len(matrixData) {
                // SIMPLIFICAR: matrixData[index] ya es []interface{}
                return matrixData[index]
            } else {
                v.ErrorTable.NewSemanticError(ctx.GetStart(),
                    fmt.Sprintf("Índice %d fuera de rango para matriz [0:%d]", index, len(matrixData)-1))
                return nil
            }
        }

        v.ErrorTable.NewSemanticError(ctx.GetStart(),
            fmt.Sprintf("'%s' no es un slice ni una matriz", sliceName))
        return nil
    }

    v.ErrorTable.NewSemanticError(ctx.GetStart(),
        fmt.Sprintf("Variable '%s' no declarada", sliceName))
    return nil
}

func (v *ReplVisitor) VisitMatrixAccess(ctx *parser.MatrixAccessContext) interface{} {
	matrixName := ctx.ID().GetText()

	// Validar que la variable existe
	matrix, exists := v.GetVariable(matrixName)
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Variable '%s' no declarada", matrixName))
		return nil
	}

	// Validar que es una matriz
	matrixData, ok := matrix.([][]interface{})
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("'%s' no es una matriz", matrixName))
		return nil
	}

	// Obtener índices con validación
	rowIndexValue := v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	colIndexValue := v.Visit(ctx.GetChild(5).(antlr.ParseTree))

	// Validar que los índices son enteros
	rowIndex, ok1 := rowIndexValue.(int)
	colIndex, ok2 := colIndexValue.(int)

	if !ok1 || !ok2 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Los índices deben ser enteros")
		return nil
	}

	// Validar rangos
	if rowIndex < 0 || rowIndex >= len(matrixData) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Índice de fila %d fuera de rango", rowIndex))
		return nil
	}

	if colIndex < 0 || colIndex >= len(matrixData[rowIndex]) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Índice de columna %d fuera de rango", colIndex))
		return nil
	}

	result := matrixData[rowIndex][colIndex]

	return result
}

func (v *ReplVisitor) VisitSliceAssignment(ctx *parser.SliceAssignmentContext) interface{} {
	sliceName := ctx.ID().GetText()
	index := v.Visit(ctx.GetChild(2).(antlr.ParseTree)).(int)
	value := v.Visit(ctx.GetChild(5).(antlr.ParseTree))

	if slice, exists := v.GetVariable(sliceName); exists {
		if slice, ok := slice.([]interface{}); ok {
			if index >= 0 && index < len(slice) {
				slice[index] = value
				v.UpdateVariable(sliceName, slice)
			} else {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Índice fuera de rango para el slice '"+sliceName+"'")
			}
		}
	}

	return nil
}

func (v *ReplVisitor) VisitMatrixAssignment(ctx *parser.MatrixAssignmentContext) interface{} {
	matrixName := ctx.ID().GetText()

	// Validar que la variable existe
	matrix, exists := v.GetVariable(matrixName)
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Variable '%s' no declarada", matrixName))
		return nil
	}

	// Validar que es una matriz
	matrixData, ok := matrix.([][]interface{})
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("'%s' no es una matriz", matrixName))
		return nil
	}

	// Obtener índices con validación
	rowIndexValue := v.Visit(ctx.GetChild(2).(antlr.ParseTree))
	colIndexValue := v.Visit(ctx.GetChild(5).(antlr.ParseTree))
	value := v.Visit(ctx.GetChild(8).(antlr.ParseTree))

	// Validar que los índices son enteros
	rowIndex, ok1 := rowIndexValue.(int)
	colIndex, ok2 := colIndexValue.(int)

	if !ok1 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Índice de fila debe ser entero, recibido: %T", rowIndexValue))
		return nil
	}

	if !ok2 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Índice de columna debe ser entero, recibido: %T", colIndexValue))
		return nil
	}

	// Validar rangos
	if rowIndex < 0 || rowIndex >= len(matrixData) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Índice de fila %d fuera de rango [0:%d]", rowIndex, len(matrixData)-1))
		return nil
	}

	if colIndex < 0 || colIndex >= len(matrixData[rowIndex]) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Índice de columna %d fuera de rango [0:%d]", colIndex, len(matrixData[rowIndex])-1))
		return nil
	}

	// Realizar la asignación
	matrixData[rowIndex][colIndex] = value

	// IMPORTANTE: Actualizar la referencia en SymbolTable
	v.UpdateVariable(matrixName, matrixData)

	return nil
}

func (v *ReplVisitor) VisitIncrementDecrement(ctx *parser.IncrementDecrementContext) interface{} {
	// Para expresiones como: i++, i--, j++, etc.

	varName := ctx.ID().GetText()
	operator := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	// Verificar que la variable existe
	currentValue, exists := v.GetVariable(varName)
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Variable '%s' no declarada", varName))
		return nil
	}

	// Verificar que es un número
	currentInt, ok := currentValue.(int)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Variable '%s' no es un entero", varName))
		return nil
	}

	// Aplicar el operador
	var newValue int
	switch operator {
	case "++":
		newValue = currentInt + 1
	case "--":
		newValue = currentInt - 1
	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Operador '%s' no válido", operator))
		return nil
	}

	// Actualizar la variable
	if !v.UpdateVariable(varName, newValue) {
		v.SetVariable(varName, newValue)
	}

	return newValue
}

func (v *ReplVisitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
    var varName string
    if identifier, ok := ctx.GetChild(0).(*parser.IdentifierContext); ok {
        varName = identifier.GetText()
    } else {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "Lado izquierdo de la asignación debe ser una variable")
        return nil
    }

    newValue := v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    // Verificar si la variable existe
    if _, exists := v.ScopeTrace.GetVariable(varName); exists {
        // Actualizar variable existente
        v.ScopeTrace.UpdateVariable(varName, newValue)
        v.SymbolTable[varName] = newValue
    } else {
        // Error: variable no declarada
        v.ErrorTable.NewSemanticError(ctx.GetStart(), 
            fmt.Sprintf("Variable '%s' no está declarada", varName))
        return nil
    }

    return newValue
}


// Obtener el nombre del tipo de un valor
func (v *ReplVisitor) getTypeName(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case bool:
		return "bool"
	case []interface{}:
		return "slice"
	case [][]interface{}:
		return "matrix"
	default:
		return "unknown"
	}
}

func (v *ReplVisitor) VisitForConditional(ctx *parser.ForConditionalContext) interface{} {
    condition := ctx.ExpressionStatement()
    block := ctx.BlockStatement()

    if condition == nil || block == nil {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "For requiere condición y bloque")
        return nil
    }

    v.ScopeTrace.EnterScope("for")
    defer v.ScopeTrace.ExitScope()

    for {
        // Evaluar condición
        conditionValue := v.Visit(condition)
        conditionBool, ok := conditionValue.(bool)
        if !ok {
            v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser un booleano")
            break
        }

        if !conditionBool {
            break
        }

        v.ScopeTrace.EnterScope("for-iteration")
        
        // Ejecutar statements y procesar continue inmediatamente
        shouldBreak := false
        shouldContinue := false
        
        for _, stmt := range block.AllStatement() {
            // Si ya decidimos continuar, no ejecutar más statements
            if shouldContinue {
                break
            }
            
            result := v.Visit(stmt)
            if result != nil {
                if resultStr, ok := result.(string); ok {
                    switch resultStr {
                    case "break":
                        shouldBreak = true
                    case "continue":
                        shouldContinue = true
                    }
                }
            }
        }
        
        v.ScopeTrace.ExitScope()
        
        if shouldBreak {
            break
        }
        
        if shouldContinue {
            continue  // Ir directamente a la siguiente iteración
        }
    }

    return nil
}

func (v *ReplVisitor) VisitForSimple(ctx *parser.ForSimpleContext) interface{} {
    v.ScopeTrace.EnterScope("for")
    defer v.ScopeTrace.ExitScope()

    // 1. INICIALIZACIÓN
    if ctx.VariableDeclaration() != nil {
        v.Visit(ctx.VariableDeclaration())
    } else if len(ctx.AllExpressionStatement()) > 0 {
        // Manejar primera expresión (puede ser declaración implícita)
        initExpr := ctx.ExpressionStatement(0)
        
        if assignment, ok := initExpr.(*parser.AssignmentContext); ok {
            if identifier, ok := assignment.GetChild(0).(*parser.IdentifierContext); ok {
                varName := identifier.GetText()
                value := v.Visit(assignment.GetChild(2).(antlr.ParseTree))
                
                if _, exists := v.GetVariable(varName); !exists {
                    // Crear variable nueva
                    dataType := "int"
                    switch value.(type) {
                    case int:
                        dataType = "int"
                    case float64:
                        dataType = "float64"
                    case string:
                        dataType = "string"
                    case bool:
                        dataType = "bool"
                    }
                    
                    symbol := Symbol{
                        ID:         varName,
                        SymbolType: VariableSymbol,
                        DataType:   dataType,
                        Scope:      v.getCurrentScope(),
                        Line:       initExpr.GetStart().GetLine(),
                        Column:     initExpr.GetStart().GetColumn(),
                    }
                    v.SymbolsTable.AddSymbol(symbol)
                    v.SetVariable(varName, value)
                } else {
                    v.Visit(initExpr)
                }
            }
        } else {
            v.Visit(initExpr)
        }
    }

    // 2. VERIFICAR BLOQUE
    block := ctx.BlockStatement()
    if block == nil {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "Bloque vacío en for")
        return nil
    }

    // 3. BUCLE PRINCIPAL - CORREGIR ÍNDICES
    iterationCount := 0
    for {
        iterationCount++
        
        if len(ctx.AllExpressionStatement()) > 0 {
            conditionExpr := ctx.ExpressionStatement(0)
            
            condition := v.Visit(conditionExpr)
            
            if condBool, ok := condition.(bool); !ok || !condBool {
                break
            }
        } else {
            break
        }

        // Prevenir bucles infinitos
        if iterationCount > 1000 {
            v.ErrorTable.NewSemanticError(ctx.GetStart(), "Bucle infinito detectado")
            break
        }

        // EJECUTAR BLOQUE
        v.ScopeTrace.EnterScope("for-iteration")
        
        shouldBreak := false
        shouldContinue := false
        
        for _, stmt := range block.AllStatement() {
            if shouldContinue {
                break
            }
            
            result := v.Visit(stmt)
            if result != nil {
                if resultStr, ok := result.(string); ok {
                    switch resultStr {
                    case "break":
                        shouldBreak = true
                    case "continue":
                        shouldContinue = true
                    }
                }
            }
        }
        
        v.ScopeTrace.ExitScope()
        
        if shouldBreak {
            break
        }

        // INCREMENTO - CAMBIAR: usar índice 1
        if len(ctx.AllExpressionStatement()) > 1 {
            v.Visit(ctx.ExpressionStatement(1))
        }
        
        if shouldContinue {
            continue
        }
    }

    return nil
}



func (v *ReplVisitor) VisitForSlice(ctx *parser.ForSliceContext) interface{} {
    // Entrar a scope del for
    v.ScopeTrace.EnterScope("for")
    defer v.ScopeTrace.ExitScope()

    // Obtener nombres de variables
    indexVarName := ctx.ID(0).GetText()
    valueVarName := ctx.ID(1).GetText()
    sliceName := ctx.ID(2).GetText()   

    // Registrar símbolos
    indexSymbol := Symbol{
        ID:         indexVarName,
        SymbolType: VariableSymbol,
        DataType:   "int",
        Scope:      v.getCurrentScope(),
        Line:       ctx.GetStart().GetLine(),
        Column:     ctx.GetStart().GetColumn(),
    }
    v.SymbolsTable.AddSymbol(indexSymbol)

    valueSymbol := Symbol{
        ID:         valueVarName,
        SymbolType: VariableSymbol,
        DataType:   "auto",
        Scope:      v.getCurrentScope(),
        Line:       ctx.GetStart().GetLine(),
        Column:     ctx.GetStart().GetColumn(),
    }
    v.SymbolsTable.AddSymbol(valueSymbol)

    // Obtener el slice
    slice, exists := v.GetVariable(sliceName)
    if !exists {
        v.ErrorTable.NewSemanticError(ctx.GetStart(),
            fmt.Sprintf("Slice '%s' no declarado", sliceName))
        return nil
    }

    sliceData, ok := slice.([]interface{})
    if !ok {
        v.ErrorTable.NewSemanticError(ctx.GetStart(),
            fmt.Sprintf("'%s' no es un slice", sliceName))
        return nil
    }

    // Obtener el bloque de instrucciones
    block := ctx.BlockStatement()
    if block == nil {
        v.ErrorTable.NewSemanticError(ctx.GetStart(),
            "Bloque de instrucciones vacío en el bucle for")
        return nil
    }

    // Iterar sobre el slice
    for index, value := range sliceData {
        // Crear scope para cada iteración
        v.ScopeTrace.EnterScope("for-iteration")
        
        // Actualizar ambas variables
        v.SetVariable(indexVarName, index)
        v.SetVariable(valueVarName, value)

        // Ejecutar el bloque de instrucciones
        shouldBreak := false 
        for _, stmt := range block.AllStatement() {
            result := v.Visit(stmt)
            if result != nil {
                if resultStr, ok := result.(string); ok {
                    switch resultStr {
                    case "break":
                        shouldBreak = true
                        v.ScopeTrace.ExitScope()
                        goto end
                    case "continue":
                        goto nextIteration
                    }
                }
            }
        }
        
        nextIteration:
        v.ScopeTrace.ExitScope()
        
        if shouldBreak {
            break
        }
    }

    end:
    return nil
}

// Establecer variable en ambos sistemas
func (v *ReplVisitor) SetVariable(name string, val interface{}) {
	// 1. ScopeTrace (scope correcto)
	v.ScopeTrace.SetVariable(name, val)

	// 2. SymbolTable global (para tabla de símbolos)
	v.SymbolTable[name] = val
}

// Obtener variable (scope-aware)
func (v *ReplVisitor) GetVariable(name string) (interface{}, bool) {
	if val, exists := v.SymbolTable[name]; exists {
		return val, true
	}
	if val, exists := v.ScopeTrace.GetVariable(name); exists {
		return val, true
	}
	return nil, false
}

// Actualizar variable existente
func (v *ReplVisitor) UpdateVariable(name string, val interface{}) bool {
	// 1. Actualizar en ScopeTrace
	updated := v.ScopeTrace.UpdateVariable(name, val)

	// 2. También actualizar SymbolTable
	v.SymbolTable[name] = val

	return updated
}

func (v *ReplVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
    
    if ctx == nil {
        return nil
    }
    
    // Evaluar la condición
    if ctx.ExpressionStatement() != nil {
        condition := v.Visit(ctx.ExpressionStatement())
        
        // Verificar si la condición es verdadera
        if condition != nil {
            conditionBool, ok := condition.(bool)
            if !ok {
                // Intentar convertir usando tu función isTruthy existente
                conditionBool = v.isTruthy(condition)
            }
            
            
            if conditionBool {
                // Entrar a scope local para el if
                v.ScopeTrace.EnterScope("if")
                
                // CAMBIO: Usar índice 0 para el primer BlockStatement
                if ctx.BlockStatement(0) != nil {
                    result := v.Visit(ctx.BlockStatement(0))
                    v.ScopeTrace.ExitScope()
                    return result
                } else {
                    v.ScopeTrace.ExitScope()
                }
            } else {
                // Verificar si hay else
                if len(ctx.AllBlockStatement()) > 1 {
                    // Scope para el else
                    v.ScopeTrace.EnterScope("if")
                    // CAMBIO: Usar índice 1 para el segundo BlockStatement
                    result := v.Visit(ctx.BlockStatement(1))
                    v.ScopeTrace.ExitScope()
                    return result
                } else if ctx.IfStatement() != nil {
                    // El else if manejará su propio scope
                    result := v.Visit(ctx.IfStatement())
                    return result
                }
            }
        }
    }
    
    return nil
}

func (v *ReplVisitor) isTruthy(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case bool:
		return v
	case int:
		return v != 0
	case float64:
		return v != 0.0
	case string:
		return v != ""
	default:
		return true
	}
}

func (v *ReplVisitor) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
    var lastResult interface{} = nil
    
    // Recorrer todos los statements en el bloque
    for _, stmt := range ctx.AllStatement() {
        // Verificar el tipo específico de TransferenceStatement
        if stmt.NonDeclaration() != nil && 
           stmt.NonDeclaration().TransferenceStatement() != nil {
            
            transferenceStmt := stmt.NonDeclaration().TransferenceStatement()
            
            // USAR TYPE ASSERTION SEGURA para cada tipo posible
            switch transferenceCtx := transferenceStmt.(type) {
            case *parser.ReturnStatementContext:
                // Es un return statement
                result := v.Visit(transferenceCtx)
                return result
                
            case *parser.BreakStatementContext:
                // Es un break statement
                return "break" // Devolver string para indicar break
                
            case *parser.ContinueStatementContext:
                // Es un continue statement
                return "continue" // Devolver string para indicar continue
                
            default:
                // Otro tipo de transferenceStatement, visitar normalmente
                result := v.Visit(transferenceStmt)
                lastResult = result
            }
        } else {
            // Si no es transferenceStatement, ejecutar normalmente
            result := v.Visit(stmt)
            lastResult = result
        }
    }
    
    return lastResult
}

func (v *ReplVisitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
    varName := ctx.ID().GetText()
    
    switchValue, exists := v.GetVariable(varName)
    if !exists {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), 
            fmt.Sprintf("Variable '%s' no declarada", varName))
        return nil
    }
    
    // Usar AllCaseItem en lugar de AllSwitchCase
    for _, caseItem := range ctx.AllSwitchCase() {
        caseValue := v.Visit(caseItem.ExpressionStatement())
        
        if v.areValuesEqual(switchValue, caseValue) {
            // Ejecutar statements del case
            shouldBreak := false
            for _, stmt := range caseItem.AllStatement() {
                result := v.Visit(stmt)
                
                // Verificar break
                if result != nil {
                    if resultStr, ok := result.(string); ok && resultStr == "break" {
                        shouldBreak = true
                        break // Salir del bucle de statements
                    }
                }
            }
            
            if shouldBreak {
                return nil // Salir del switch completo
            }
            return nil
        }
    }
    
    // Usar DefaultCase
    if ctx.DefaultCase() != nil {
        shouldBreak := false
        for _, stmt := range ctx.DefaultCase().AllStatement() {
            result := v.Visit(stmt)
            
            // Verificar break en default también
            if result != nil {
                if resultStr, ok := result.(string); ok && resultStr == "break" {
                    shouldBreak = true
                    break
                }
            }
        }
        
        if shouldBreak {
            return nil
        }
    }
    
    return nil
}

func (v *ReplVisitor) areValuesEqual(a, b interface{}) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    
    // Usar la función isEqual existente
    return isEqual(a, b)
}

func (v *ReplVisitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
    functionName := ctx.ID(0).GetText()

    // Verificar que no esté ya declarada
    if _, exists := v.Functions[functionName]; exists {
        v.ErrorTable.NewSemanticError(ctx.GetStart(),
            fmt.Sprintf("Función '%s' ya está declarada", functionName))
        return nil
    }

    // Crear definición de función
    funcDef := &FunctionDefinition{
        Name:       functionName,
        Parameters: []Parameter{},
        ReturnType: "void",
        Body:       ctx.BlockStatement(),
    }

    // Procesar parámetros si existen
    if ctx.ParameterList() != nil {
        funcDef.Parameters = v.Visit(ctx.ParameterList()).([]Parameter)
    }

    // Guardar la función
    v.Functions[functionName] = funcDef

    // Procesar tipo de retorno si existe
    if ctx.TYPE() != nil {
        funcDef.ReturnType = ctx.TYPE().GetText()
    } else {
        // Verificar si hay un ID como tipo de retorno (struct)
        children := ctx.GetChildren()
        var ids []string
        
        for _, child := range children {
            if terminal, ok := child.(antlr.TerminalNode); ok {
                if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
                    ids = append(ids, terminal.GetText())
                }
            }
        }
        
        if len(ids) > 1 {
            funcDef.ReturnType = ids[1]
        }
    }

    // Guardar la función
    v.Functions[functionName] = funcDef

    // Registrar símbolo
    symbol := Symbol{
        ID:         functionName,
        SymbolType: FunctionSymbol,
        DataType:   FunctionType,
        Scope:      "Global",
        Line:       ctx.GetStart().GetLine(),
        Column:     ctx.GetStart().GetColumn(),
    }
    v.SymbolsTable.AddSymbol(symbol)

    // Si es main, marcarla especialmente
    if functionName == "main" {
        v.MainFunction = ctx
    }

    return nil
}


func (v *ReplVisitor) VisitSwitchCase(ctx *parser.SwitchCaseContext) interface{} {

    return v.Visit(ctx.ExpressionStatement())
}

func (v *ReplVisitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {

	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}
	
	return nil
}
// Ejecutar la función main
func (v *ReplVisitor) ExecuteMainFunction() interface{} {
    if v.MainFunction == nil {
        v.ErrorTable.NewSemanticError(nil, "No se encontró la función main()")
        return nil
    }
    
    // Crear scope específico para la función main
    v.ScopeTrace.EnterScope("function")
    
    result := v.Visit(v.MainFunction.BlockStatement())
    
    v.ScopeTrace.ExitScope()
    return result
}

func (v *ReplVisitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) interface{} {
	structName := ctx.ID().GetText()

	// Verificar que no esté ya declarado
	if _, exists := v.StructTypes[structName]; exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("El struct '%s' ya está declarado", structName))
		return nil
	}

	structType := &StructDefinition{
		Name:   structName,
		Fields: make(map[string]string),
	}

	// Procesar cada atributo
	for _, attrCtx := range ctx.AllStructAtribute() {
		var fieldType, fieldName string

		// Verificar si es tipo primitivo (TYPE) o tipo de struct (ID)
		if attrCtx.TYPE() != nil {
			// Tipo primitivo: TYPE ID
			fieldType = attrCtx.TYPE().GetText()

			// Verificar que hay al menos un ID para el nombre del campo
			if len(attrCtx.AllID()) > 0 {
				fieldName = attrCtx.ID(0).GetText()
			} else {
				v.ErrorTable.NewSemanticError(attrCtx.GetStart(),
					"Nombre de campo faltante")
				continue
			}
		} else {
			// Tipo de struct: ID ID (primer ID es el tipo, segundo es el nombre)
			if len(attrCtx.AllID()) >= 2 {
				fieldType = attrCtx.ID(0).GetText()
				fieldName = attrCtx.ID(1).GetText()
			} else {
				v.ErrorTable.NewSemanticError(attrCtx.GetStart(),
					"Declaración de campo inválida - se requieren tipo y nombre")
				continue
			}
		}

		// Verificar que el tipo de struct existe (si no es primitivo)
		if !v.isPrimitiveType(fieldType) {
			if _, exists := v.StructTypes[fieldType]; !exists {
				v.ErrorTable.NewSemanticError(attrCtx.GetStart(),
					fmt.Sprintf("Tipo de struct '%s' no declarado", fieldType))
				continue
			}
		}

		// Verificar campos duplicados
		if _, exists := structType.Fields[fieldName]; exists {
			v.ErrorTable.NewSemanticError(attrCtx.GetStart(),
				fmt.Sprintf("Campo '%s' ya declarado en struct '%s'", fieldName, structName))
			continue
		}

		structType.Fields[fieldName] = fieldType
	}

	// Registrar el tipo de struct
	v.StructTypes[structName] = structType

	// Registrar en tabla de símbolos
	symbol := Symbol{
		ID:         structName,
		SymbolType: StructSymbol,
		DataType:   "struct",
		Scope:      "Global",
		Line:       ctx.GetStart().GetLine(),
		Column:     ctx.GetStart().GetColumn(),
	}
	v.SymbolsTable.AddSymbol(symbol)

	return nil
}

func (v *ReplVisitor) isPrimitiveType(typeName string) bool {
	primitives := []string{"int", "string", "float64", "bool", "float", "nil"}
	for _, primitive := range primitives {
		if typeName == primitive {
			return true
		}
	}
	return false
}

func (v *ReplVisitor) VisitStructInstantiation(ctx *parser.StructInstantiationContext) interface{} {
	// Verificación básica del contexto
	if ctx == nil {
		v.ErrorTable.NewSemanticError(nil, "Contexto de struct nulo")
		return nil
	}

	// Verificación del ID del struct
	if ctx.ID() == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Nombre de struct faltante")
		return nil
	}

	structTypeName := ctx.ID().GetText()

	// Verificar que el tipo existe
	structType, exists := v.StructTypes[structTypeName]
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Tipo de struct '%s' no declarado", structTypeName))
		return nil
	}

	// Crear instancia
	instance := &StructInstance{
		TypeName: structTypeName,
		Fields:   make(map[string]interface{}),
	}

	// Inicializar campos con valores por defecto
	for fieldName, fieldType := range structType.Fields {
		switch fieldType {
		case "int":
			instance.Fields[fieldName] = 0
		case "float64":
			instance.Fields[fieldName] = 0.0
		case "string":
			instance.Fields[fieldName] = ""
		case "bool":
			instance.Fields[fieldName] = false
		default:
			instance.Fields[fieldName] = nil
		}
	}

	// Procesar inicialización si existe
	structInit := ctx.StructInitialization()
	if structInit != nil {
		for _, fieldInitCtx := range structInit.AllStructFieldInit() {
			if fieldInitCtx == nil {
				continue
			}

			// Verificar que el ID del campo existe
			if fieldInitCtx.ID() == nil {
				v.ErrorTable.NewSemanticError(fieldInitCtx.GetStart(), "Nombre de campo faltante")
				continue
			}

			fieldName := fieldInitCtx.ID().GetText()

			// Verificar que el campo existe en el struct
			if _, exists := structType.Fields[fieldName]; !exists {
				v.ErrorTable.NewSemanticError(fieldInitCtx.GetStart(),
					fmt.Sprintf("Campo '%s' no existe en struct '%s'", fieldName, structTypeName))
				continue
			}

			// Verificar que la expresión existe
			if fieldInitCtx.ExpressionStatement() == nil {
				v.ErrorTable.NewSemanticError(fieldInitCtx.GetStart(),
					fmt.Sprintf("Valor faltante para campo '%s'", fieldName))
				continue
			}

			// Evaluar el valor del campo
			fieldValue := v.Visit(fieldInitCtx.ExpressionStatement())
			instance.Fields[fieldName] = fieldValue
		}
	}

	return instance
}

func (v *ReplVisitor) VisitStructAccess(ctx *parser.StructAccessContext) interface{} {
	// Obtener el nombre del campo (siempre es el último ID)
	fieldName := ctx.ID().GetText()

	// La parte izquierda puede ser un ID directo o una expresión
	var leftSide interface{}

	// Obtener la expresión del lado izquierdo
	if ctx.ExpressionStatement() != nil {
		leftSide = v.Visit(ctx.ExpressionStatement())
	} else {
		for i := 0; i < ctx.GetChildCount(); i++ {
			child := ctx.GetChild(i)
			if terminal, ok := child.(antlr.TerminalNode); ok {
				if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
					// Este es el primer ID (la variable)
					varName := terminal.GetText()

					// Si es diferente del fieldName, es la variable base
					if varName != fieldName {
						var exists bool
						leftSide, exists = v.GetVariable(varName)
						if !exists {
							v.ErrorTable.NewSemanticError(ctx.GetStart(),
								fmt.Sprintf("Variable '%s' no declarada", varName))
							return nil
						}
						break
					}
				}
			}
		}
	}

	// Verificar que tenemos algo en leftSide
	if leftSide == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Acceso a struct inválido")
		return nil
	}

	// Si leftSide es una instancia de struct
	if structInstance, ok := leftSide.(*StructInstance); ok {
		// Verificar que el campo existe
		if fieldValue, exists := structInstance.Fields[fieldName]; exists {
			return fieldValue
		} else {
			v.ErrorTable.NewSemanticError(ctx.GetStart(),
				fmt.Sprintf("Campo '%s' no existe en struct '%s'", fieldName, structInstance.TypeName))
			return nil
		}
	}

	v.ErrorTable.NewSemanticError(ctx.GetStart(), "Intento de acceso a campo en variable que no es struct")
	return nil
}

func (v *ReplVisitor) VisitExplicitStructDeclaration(ctx *parser.ExplicitStructDeclarationContext) interface{} {
	var varName string
	var structTypeName string
	var isMutable bool = false

	// Procesar elementos
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		if terminal, ok := child.(antlr.TerminalNode); ok {
			text := terminal.GetText()

			if text == "mut" {
				isMutable = true
				continue
			}

			if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
				if varName == "" {
					varName = text
				} else {
					structTypeName = text
				}
				continue
			}
		}
	}

	// Verificar redeclaración
	if varName != "" {
		if _, exists := v.GetVariable(varName); exists {
			v.ErrorTable.NewSemanticError(ctx.GetStart(),
				fmt.Sprintf("Error de redeclaración: la variable '%s' ya está declarada", varName))
			return nil
		}
	}

	// Verificar que el tipo de struct existe
	structType, exists := v.StructTypes[structTypeName]
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Tipo de struct '%s' no declarado", structTypeName))
		return nil
	}

	// Crear instancia
	structInstance := &StructInstance{
		TypeName: structTypeName,
		Fields:   make(map[string]interface{}),
	}

	// Inicializar con valores por defecto
	for fieldName, fieldType := range structType.Fields {
		switch fieldType {
		case "int":
			structInstance.Fields[fieldName] = 0
		case "float64":
			structInstance.Fields[fieldName] = 0.0
		case "string":
			structInstance.Fields[fieldName] = ""
		case "bool":
			structInstance.Fields[fieldName] = false
		default:
			structInstance.Fields[fieldName] = nil
		}
	}

	// Procesar inicialización si existe
	if ctx.StructInitialization() != nil {
		for _, fieldInitCtx := range ctx.StructInitialization().AllStructFieldInit() {
			fieldName := fieldInitCtx.ID().GetText()
			fieldValue := v.Visit(fieldInitCtx.ExpressionStatement())

			if _, exists := structType.Fields[fieldName]; exists {
				structInstance.Fields[fieldName] = fieldValue
			} else {
				v.ErrorTable.NewSemanticError(fieldInitCtx.GetStart(),
					fmt.Sprintf("Campo '%s' no existe en struct '%s'", fieldName, structTypeName))
			}
		}
	}

	// Registrar símbolo
	if varName != "" {
		dataType := structTypeName
		if isMutable {
			dataType += " (mut)"
		}

		symbol := Symbol{
			ID:         varName,
			SymbolType: StructSymbol,
			DataType:   dataType,
			Scope:      v.getCurrentScope(),
			Line:       ctx.GetStart().GetLine(),
			Column:     ctx.GetStart().GetColumn(),
		}
		v.SymbolsTable.AddSymbol(symbol)

		v.SetVariable(varName, structInstance)
	}

	return nil
}

func (v *ReplVisitor) VisitImplicitStructDeclaration(ctx *parser.ImplicitStructDeclarationContext) interface{} {

	var varName string
	var structTypeName string
	var isMutable bool = false

	// Analizar los hijos para extraer información
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		if terminal, ok := child.(antlr.TerminalNode); ok {
			text := terminal.GetText()

			if text == "mut" {
				isMutable = true
				continue
			}

			if terminal.GetSymbol().GetTokenType() == parser.LanguageLexerID {
				if varName == "" {
					varName = text
				} else if structTypeName == "" {
					structTypeName = text
				}
			}
		}
	}

	// Verificar redeclaración
	if varName != "" {
		if _, exists := v.GetVariable(varName); exists {
			v.ErrorTable.NewSemanticError(ctx.GetStart(),
				fmt.Sprintf("Error de redeclaración: la variable '%s' ya está declarada", varName))
			return nil
		}
	}

	// Verificar que el tipo de struct existe
	structType, exists := v.StructTypes[structTypeName]
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Tipo de struct '%s' no declarado", structTypeName))
		return nil
	}

	// Crear la instancia del struct
	structInstance := &StructInstance{
		TypeName: structTypeName,
		Fields:   make(map[string]interface{}),
	}

	// Inicializar campos con valores por defecto
	for fieldName, fieldType := range structType.Fields {
		structInstance.Fields[fieldName] = v.getDefaultValue(fieldType)
	}

	// Procesar la inicialización si existe
	if ctx.StructInitialization() != nil {
		if initCtx, ok := ctx.StructInitialization().(*parser.StructInitializationContext); ok {
			v.processStructInitialization(initCtx, structInstance, structType)
		}
	}

	// Registrar símbolo
	if varName != "" {
		dataType := structInstance.TypeName
		if isMutable {
			dataType += " (mut)"
		}

		symbol := Symbol{
			ID:         varName,
			SymbolType: StructSymbol,
			DataType:   dataType,
			Scope:      v.getCurrentScope(),
			Line:       ctx.GetStart().GetLine(),
			Column:     ctx.GetStart().GetColumn(),
		}
		v.SymbolsTable.AddSymbol(symbol)

		v.SetVariable(varName, structInstance)
	}

	return nil
}

func (v *ReplVisitor) getDefaultValue(dataType string) interface{} {
	switch dataType {
	case "int":
		return 0
	case "float64", "float":
		return 0.0
	case "string":
		return ""
	case "bool":
		return false
	default:
		// Para tipos de structs, crear instancia vacía
		if structType, exists := v.StructTypes[dataType]; exists {
			instance := &StructInstance{
				TypeName: dataType,
				Fields:   make(map[string]interface{}),
			}
			// Inicializar recursivamente
			for fieldName, fieldType := range structType.Fields {
				instance.Fields[fieldName] = v.getDefaultValue(fieldType)
			}
			return instance
		}
		return nil
	}
}

func (v *ReplVisitor) processStructInitialization(ctx *parser.StructInitializationContext, instance *StructInstance, structType *StructDefinition) {
	for _, fieldInitCtx := range ctx.AllStructFieldInit() {
		fieldName := fieldInitCtx.ID().GetText()
		fieldValue := v.Visit(fieldInitCtx.ExpressionStatement())

		// Verificar que el campo existe
		if _, exists := structType.Fields[fieldName]; exists {
			instance.Fields[fieldName] = fieldValue
		} else {
			v.ErrorTable.NewSemanticError(fieldInitCtx.GetStart(),
				fmt.Sprintf("Campo '%s' no existe en struct '%s'", fieldName, instance.TypeName))
		}
	}
}

func (v *ReplVisitor) VisitStructAssignment(ctx *parser.StructAssignmentContext) interface{} {
	// Obtener el nuevo valor
	newValue := v.Visit(ctx.ExpressionStatement(1))

	// Obtener el nombre del campo (el ID)
	fieldName := ctx.ID().GetText()

	// Obtener la estructura del lado izquierdo
	leftSide := v.Visit(ctx.ExpressionStatement(0))

	if leftSide == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Estructura no válida para asignación")
		return nil
	}

	// Verificar que es una instancia de struct
	structInstance, ok := leftSide.(*StructInstance)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Intento de asignación a campo en variable que no es struct")
		return nil
	}

	// Verificar que el campo existe
	if _, exists := structInstance.Fields[fieldName]; !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Campo '%s' no existe en struct '%s'", fieldName, structInstance.TypeName))
		return nil
	}

	// Obtener el tipo esperado del campo
	structType, typeExists := v.StructTypes[structInstance.TypeName]
	if !typeExists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Tipo de struct '%s' no encontrado", structInstance.TypeName))
		return nil
	}

	expectedType, fieldExists := structType.Fields[fieldName]
	if !fieldExists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("Campo '%s' no existe en struct '%s'", fieldName, structInstance.TypeName))
		return nil
	}

	// Validar compatibilidad de tipos
	if !v.isValueCompatibleWithType(newValue, expectedType) {
		actualType := v.getTypeName(newValue)
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			fmt.Sprintf("No se puede asignar valor de tipo '%s' a campo '%s' de tipo '%s'",
				actualType, fieldName, expectedType))
		return nil
	}

	// Asignar el nuevo valor
	structInstance.Fields[fieldName] = newValue

	return newValue
}

func (v *ReplVisitor) isValueCompatibleWithType(value interface{}, expectedType string) bool {
	actualType := v.getTypeName(value)

	// Tipos exactamente iguales
	if actualType == expectedType {
		return true
	}

	// Mapear float a float64 para compatibilidad
	if expectedType == "float" && actualType == "float64" {
		return true
	}
	if expectedType == "float64" && actualType == "float64" {
		return true
	}

	// Sin conversiones automáticas para structs (más estricto)
	return false
}


func (v *ReplVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	if ctx == nil || ctx.ID() == nil {
		return nil
	}

	functionName := ctx.ID().GetText()

	// Procesar argumentos
	var args []interface{}
	if ctx.ArgumentList() != nil {
		args = v.processArgumentList(ctx.ArgumentList())
	}

	// Funciones builtin
	switch functionName {
	case "indexOf":
		if len(args) != 2 {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "indexOf espera exactamente 2 argumentos")
			return nil
		}
		slice, ok := args[0].([]interface{})
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El primer argumento de indexOf debe ser un slice")
			return nil
		}
		target := args[1]
		for i, val := range slice {
			if reflect.DeepEqual(val, target) {
				return i
			}
		}
		return -1

	case "join":
		if len(args) != 2 {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "join espera exactamente 2 argumentos: slice y separador")
			return nil
		}

		// Verificar que el primer argumento es slice de strings
		slice, ok := args[0].([]interface{})
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El primer argumento de join debe ser un slice")
			return nil
		}

		sep, ok := args[1].(string)
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "El segundo argumento de join debe ser un string")
			return nil
		}

		var parts []string
		for _, elem := range slice {
			strElem, ok := elem.(string)
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Todos los elementos del slice para join deben ser strings")
				return nil
			}
			parts = append(parts, strElem)
		}

		// Usamos strings.Join para unirlos
		return strings.Join(parts, sep)

	case "len":
		if len(args) != 1 {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "len espera exactamente 1 argumento")
			return nil
		}

		arg := args[0]

		switch val := arg.(type) {
		case string:
			return len(val)
		case []interface{}:
			// MEJORAR: Es un slice (incluyendo filas de matriz)
			return len(val)
		case [][]interface{}:
			// Es una matriz, devolver número de filas
			return len(val)
		default:
			// Verificar si es un resultado de slice access
			// Esto maneja casos como matrizIrregular[1] que devuelve []interface{}
			if reflect.TypeOf(arg).Kind() == reflect.Slice {
				// Use reflection para obtener la longitud de cualquier tipo de slice
				v := reflect.ValueOf(arg)
				return v.Len()
			}
			
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "len sólo soporta slices o strings")
			return nil
		}

	case "append":
    if len(args) < 2 {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "append espera al menos 2 argumentos: un slice y al menos un valor")
        return nil
    }

    var sliceArg []interface{}
    
    // Verificar si el primer argumento es slice o matriz
    switch firstArg := args[0].(type) {
    case []interface{}:
        sliceArg = firstArg
    case [][]interface{}:
        // Es una matriz, mantener como matriz
        newMatrix := append([][]interface{}{}, firstArg...)
        
        // Agregar los nuevos elementos (deben ser slices)
        for _, arg := range args[1:] {
            if rowSlice, ok := arg.([]interface{}); ok {
                newMatrix = append(newMatrix, rowSlice)
            } else {
                v.ErrorTable.NewSemanticError(ctx.GetStart(), "Al agregar a una matriz, los elementos deben ser slices")
                return nil
            }
        }
        return newMatrix
    default:
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "El primer argumento de append debe ser un slice o matriz")
        return nil
    }

    // Los valores siguientes se agregan al slice
    newSlice := append([]interface{}{}, sliceArg...)
    newSlice = append(newSlice, args[1:]...)

    return newSlice
	}

	function, exists := v.Functions[functionName]
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), fmt.Sprintf("Función '%s' no encontrada", functionName))
		return nil
	}

	result := v.executeFunction(function, args)
	return result
}
func (v *ReplVisitor) VisitParameterList(ctx *parser.ParameterListContext) interface{} {
    var parameters []Parameter
    
    for _, paramCtx := range ctx.AllParameter() {
        param := v.Visit(paramCtx).(Parameter)
        parameters = append(parameters, param)
    }
    
    return parameters
}

func (v *ReplVisitor) VisitParameter(ctx *parser.ParameterContext) interface{} {
    var paramType, paramName string
    
    // Obtener todos los IDs
    ids := ctx.AllID()
    
    if ctx.TYPE() != nil {
        // Caso: TYPE ID (ej: int a)
        paramType = ctx.TYPE().GetText()
        paramName = ids[0].GetText()
    } else if len(ids) >= 2 {
        // Caso: ID ID (ej: Persona p)
        paramType = ids[0].GetText()  // Tipo struct
        paramName = ids[1].GetText()  // Nombre parámetro
    } else {
        v.ErrorTable.NewSemanticError(ctx.GetStart(), "Parámetro mal formado")
        return Parameter{}
    }
    
    return Parameter{
        Type: paramType,
        Name: paramName,
    }
}

func (v *ReplVisitor) processArgumentList(ctx parser.IArgumentListContext) []interface{} {
    var args []interface{}
    
    // Cast a tipo concreto
    if argListCtx, ok := ctx.(*parser.ArgumentListContext); ok {
        expressions := argListCtx.AllExpressionStatement()
        
        for _, exprCtx := range expressions {
            arg := v.Visit(exprCtx)
            args = append(args, arg)
        }
    } 
    return args
}

func (v *ReplVisitor) executeFunction(function *FunctionDefinition, args []interface{}) interface{} {
    if function == nil || function.Body == nil {
        return nil
    }
    
    const MAX_RECURSION = 1000
    v.recursionCount++

    if v.recursionCount > MAX_RECURSION {
        v.recursionCount--
        return nil
    }
    
    scopeName := fmt.Sprintf("function_%s_%d_%p", function.Name, v.recursionCount, &args)
    v.ScopeTrace.EnterScope(scopeName)
    
    // Resetear flag de return al inicio
    v.shouldReturn = false
    v.returnValue = nil
    
    // ASIGNAR PARÁMETROS CON COPIA PROFUNDA
    for i, param := range function.Parameters {
        if i < len(args) && param.Name != "" {
            var argValue interface{}
            switch val := args[i].(type) {
            case int:
                argValue = int(val)
            case string:
                argValue = string(val)
            case float64:
                argValue = float64(val)
            case bool:
                argValue = bool(val)
            default:
                argValue = val
            }
            
            v.ScopeTrace.SetVariable(param.Name, argValue)
        }
    }
    
    var returnValue interface{} = nil
    
    blockCtx, ok := function.Body.(*parser.BlockStatementContext)
    if !ok {
        v.ScopeTrace.ExitScope()
        v.recursionCount--
        return nil
    }
    
    statements := blockCtx.AllStatement()
    if statements == nil {
        v.ScopeTrace.ExitScope()
        v.recursionCount--
        return nil
    }
    
    // EJECUCIÓN DE STATEMENTS CON VERIFICACIÓN DE RETURN
    for _, stmt := range statements {
        if stmt == nil {
            continue
        }
        
        result := v.Visit(stmt)
        
        // Verificar flag de return DESPUÉS de cada statement
        if v.shouldReturn {
            returnValue = v.returnValue
            v.shouldReturn = false  // Reset flag
            v.returnValue = nil     // Reset value
            break
        }
        
        // Para funciones void (como hanoi), no romper el bucle
        if result != nil && function.ReturnType != "void" {
            returnValue = result
            break
        }
    }
    
    v.ScopeTrace.ExitScope()
    v.recursionCount--
    
    return returnValue
}

func (v *ReplVisitor) DeleteVariable(name string) {
    delete(v.SymbolTable, name)
}

func (v *ReplVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
    var returnValue interface{} = nil
    
    if ctx.ExpressionStatement() != nil {
        returnValue = v.Visit(ctx.ExpressionStatement())
    }
    
    // Flag para indicar que se debe retornar
    v.shouldReturn = true
    v.returnValue = returnValue
    
    return returnValue
}

func (v *ReplVisitor) VisitSliceLiteral(ctx *parser.SliceLiteralContext) interface{} {
    var elements []interface{}
    
    // Obtener todas las expresiones dentro del slice literal
    expressions := ctx.AllExpressionStatement()
    
    for _, exprCtx := range expressions {
        value := v.Visit(exprCtx)
        elements = append(elements, value)
    }
    
    return elements
}

func (v *ReplVisitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
    return "break"
}

func (v *ReplVisitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
    return "continue"
}

func (v *ReplVisitor) VisitNil(ctx *parser.NilContext) interface{} {
	return nil
}