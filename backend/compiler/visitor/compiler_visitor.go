package visitor

import (
	c "backend/compiler/arm"
	parser "backend/parser"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Visitor para recorrer el árbol de sintaxis
type Visitor struct {
	*parser.BaseLanguageVisitor
}

// NewVisitor crea una nueva instancia del Visitor
func NewVisitor() *Visitor {
	return &Visitor{
		BaseLanguageVisitor: &parser.BaseLanguageVisitor{},
	}
}

func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Printf("Error en compilación: %s", val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
    //fmt.Printf("🔍 DEBUG: VisitProgram - Número de hijos: %d\n", ctx.GetChildCount())
    
    // Visitar todos los statements del programa
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        //fmt.Printf("🔍 DEBUG: Hijo %d: %T\n", i, child)
        v.Visit(child.(antlr.ParseTree))
    }
    return nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
    //fmt.Printf("🔍 DEBUG: VisitStatement - Número de hijos: %d\n", ctx.GetChildCount())
    
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        //fmt.Printf("🔍 DEBUG: Statement hijo %d: %T\n", i, child)
        result := v.Visit(child.(antlr.ParseTree))
		if result == "break" || result == "continue" {
			// Si es un break o continue, retornar inmediatamente
			return result
		}
    }
    return nil
}

func (v *Visitor) VisitNonDeclaration(ctx *parser.NonDeclarationContext) interface{} {
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		v.Visit(child.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	child := ctx.GetChild(0).(antlr.ParseTree)
	// Usar Visit genérico para manejar cualquier tipo de expresión
	result := v.Visit(child)

	// Hacer pop con el objeto del stack para no descuadrar el stack
	// Si ocurre algun fallo, hay que eliminar la línea de código que hace pop
	// y dejar que el stack se maneje de manera automática.

	return result
}

func (v *Visitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
    var value = ctx.GetText()
    intValue, _ := strconv.Atoi(value)
    c.Comment("Constant: " + value)
    var IntObject = c.IntObject()
    if !inFunctionArgs {
        c.PushConstant(intValue, IntObject) // Esto ya hace PushObject
    } else {
        c.Mov(c.X0, intValue)
        c.PushObject(IntObject) // Solo aquí
    }
    return nil
}

func (v *Visitor) VisitFloat(ctx *parser.FloatContext) interface{} {
    var value = ctx.GetText()
    floatValue, err := strconv.ParseFloat(value, 64)
    if err != nil {
        log.Printf("Error parsing float: %v", err)
        floatValue = 0.0
    }
    var floatObject = c.FloatObject()
    if !inFunctionArgs {
        c.PushConstant(floatValue, floatObject)
    } else {
        c.FMov(c.D0, fmt.Sprintf("%f", floatValue))
        c.PushObject(floatObject)
    }
    return nil
}

func (v *Visitor) VisitString(ctx *parser.StringContext) interface{} {
    var value = strings.Trim(ctx.GetText(), `"`)
    c.Comment("Constant String: " + value)
    var StringObject = c.StringObject()
    if !inFunctionArgs {
        c.PushConstant(value, StringObject)
    } else {
        c.PushStringNoStack(value)
        c.MovReg(c.X0, "x11")
        c.PushObject(StringObject)
    }
    return nil
}

func (v *Visitor) VisitBoolean(ctx *parser.BooleanContext) interface{} {
    var value = ctx.GetText()
    var boolValue bool
    if value == "true" {
        boolValue = true
    } else {
        boolValue = false
    }
    c.Comment("Constant Boolean: " + value)
    var BoolObject = c.BoolObject()
    if !inFunctionArgs {
        c.PushConstant(boolValue, BoolObject)
    } else {
        c.Mov(c.X0, map[bool]int{true: 1, false: 0}[boolValue])
        c.PushObject(BoolObject)
    }
    return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

    v.Visit(ctx.GetChild(0).(antlr.ParseTree)) // 'a' se pushea primero
    v.Visit(ctx.GetChild(2).(antlr.ParseTree)) // 'b' se pushea después

    // CORREGIR: El primer pop es el SEGUNDO operando (último pusheado)
    var isRightDouble = c.TopObject().Type == c.StackObjectType(c.Float)
    var rightReg interface{}
    if isRightDouble {
        rightReg = c.D1  // CAMBIAR: usar D1 para el segundo operando
    } else {
        rightReg = c.X1  // CAMBIAR: usar X1 para el segundo operando
    }
    c.PopObject(rightReg.(string)) // Pop 'b' (segundo operando)

    var isLeftDouble = c.TopObject().Type == c.StackObjectType(c.Float)
    var leftReg interface{}
    if isLeftDouble {
        leftReg = c.D0   // CAMBIAR: usar D0 para el primer operando
    } else {
        leftReg = c.X0   // CAMBIAR: usar X0 para el primer operando
    }
    var left = c.PopObject(leftReg.(string)) // Pop 'a' (primer operando)

    if isLeftDouble || isRightDouble {
        if !isLeftDouble {
            c.Scvtf(c.D0, c.X0)
        }
        if !isRightDouble {
            c.Scvtf(c.D1, c.X1)
        }

        switch op {
        case "+":
            c.Comment("Addition operator for double")
            c.FAdd(c.D0, c.D0, c.D1)  // D0 + D1
        case "-":
            c.Comment("Subtraction operator for double")
            c.FSub(c.D0, c.D0, c.D1)  // D0 - D1 (CORREGIR orden)
        }

        if !inReturn {
            c.Comment("Pushing result of double operation")
            c.Push(c.D0)
            c.PushObject(c.CloneObject(left))
        }
        return nil
    }

    switch op {
    case "+":
        c.Comment("Addition operator")
        c.Add(c.X0, c.X0, c.X1)  // X0 + X1 (a + b)
    case "-":
        c.Comment("Subtraction operator")
        c.Sub(c.X0, c.X0, c.X1)  // X0 - X1 (a - b)
    }

    if !inReturn {
        c.Push(c.X0)
        var objToPush interface{}
        if isLeftDouble {
            objToPush = left
        } else {
            objToPush = left  // CAMBIAR: usar left siempre para el resultado
        }
        c.PushObject(c.CloneObject(objToPush.(c.StackObject)))
    }

    return nil
}

func (v *Visitor) VisitMulDivMod(ctx *parser.MulDivModContext) interface{} {
	var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	c.PopObject(c.X1)            // Pop second operand
	var left = c.PopObject(c.X0) // Pop first operand

	switch op {
	case "*":
		c.Comment("Multiplication operator")
		c.Mul(c.X0, c.X0, c.X1)
	case "/":
		c.Comment("Division operator")
		c.SDiv(c.X0, c.X0, c.X1)
	}

	c.Push(c.X0)
	c.PushObject(c.CloneObject(left))
	return nil
}

func (v *Visitor) VisitIncrementDecrement(ctx *parser.IncrementDecrementContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

    // Obtener el nombre de la variable
    varName := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
    
    // Obtener el valor actual de la variable
    v.Visit(ctx.GetChild(0).(antlr.ParseTree)) 
    var left = c.PopObject(c.X0) 

    switch op {
    case "++":
        c.Comment("Increment operator")
        c.Addi(c.X0, c.X0, 1) 
    case "--":
        c.Comment("Decrement operator")
        c.Subi(c.X0, c.X0, 1) 
    }

    // Actualizar la variable en memoria
    offset, _ := c.GetObject(varName)
    c.Mov(c.X1, offset)
    c.Add(c.X1, c.SP, c.X1)
    c.Str(c.X0, c.X1, 0)

    // Push el resultado para expresiones
    c.Push(c.X0)                      
    c.PushObject(c.CloneObject(left)) 

    return nil
}

func (v *Visitor) VisitAddSubOperator(ctx *parser.AddSubOperatorContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
    varName := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()

    // Evalúa ambos operandos
    v.Visit(ctx.GetChild(0).(antlr.ParseTree)) // variable (ID)
    v.Visit(ctx.GetChild(2).(antlr.ParseTree)) // valor a sumar/restar

    c.PopObject(c.X1) // valor a sumar/restar
    c.PopObject(c.X0) // valor actual de la variable

    offset, varObject := c.GetObject(varName)

    if varObject.Type == c.Float {
        c.Scvtf(c.D0, c.X0)
        c.Scvtf(c.D1, c.X1)
        switch op {
        case "+=":
            c.Comment("Add assignment operator (float)")
            c.FAdd(c.D0, c.D0, c.D1)
        case "-=":
            c.Comment("Subtract assignment operator (float)")
            c.FSub(c.D0, c.D0, c.D1)
        }
        c.Mov(c.X1, offset)
        c.Add(c.X1, c.SP, c.X1)
        c.StrF(c.D0, c.X1, 0)
        c.Push(c.D0)
        floatObj := c.CloneObject(varObject)
        floatObj.Type = c.Float
        c.PushObject(floatObj)
    } else {
        switch op {
        case "+=":
            c.Comment("Add assignment operator")
            c.Add(c.X0, c.X0, c.X1)
        case "-=":
            c.Comment("Subtract assignment operator")
            c.Sub(c.X0, c.X0, c.X1)
        }
        c.Mov(c.X1, offset)
        c.Add(c.X1, c.SP, c.X1)
        c.Str(c.X0, c.X1, 0)
        c.Push(c.X0)
        intObj := c.CloneObject(varObject)
        intObj.Type = c.Int
        c.PushObject(intObj)
    }
    return nil
}

func (v *Visitor) VisitParens(ctx *parser.ParensContext) interface{} {

	result := v.Visit(ctx.ExpressionStatement())

	return result
}

func (v *Visitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	c.Comment("Print statement")

	expressions := ctx.AllExpressionStatement()
	n := len(expressions)

	for i, expr := range expressions {
		v.Visit(expr)
		obj := c.TopObject()

		// Si es slice, imprime todos los elementos
		if obj.IsSlice {
			c.Comment("Entrando a impresión de slice")
			c.Pop(c.X9)       // Pop real: dirección base del slice
			c.PopObject(c.X0) // Pop virtual: objeto del slice
			length := obj.Size
			elemType := obj.ElemType

			// Imprime '['
			c.PrintChar('[')

			for j := 0; j < length; j++ {
				c.Comment(fmt.Sprintf("DEBUG elemento #%d:", j))
				switch elemType {
				case c.Int:
					c.Ldr(c.X0, c.X9, j*8)
					c.PrintIntInline(c.X0)
				case c.String:
					c.Ldr(c.X0, c.X9, j*8)
					c.PrintStringInline(c.X0)
				case c.Float:
					c.LdrF(c.D0, c.X9, j*8)
					c.PrintFloatInline(c.D0)
				}
				if j < length-1 {
					c.PrintChar(',')
					c.PrintChar(' ')
				}
			}
			c.PrintChar(']')
			if i == n-1 {
				c.PrintChar('\n')
			}
		} else {
			// Comportamiento original para no-slices
			isDouble := obj.Type == c.StackObjectType(c.Float)
			var reg string
			if isDouble {
				reg = c.D0
			} else {
				reg = c.X0
			}
			value := c.PopObject(reg)

			if value.Type == c.StackObjectType(c.Int) {
				c.PrintInt(c.X0)
			} else if value.Type == c.StackObjectType(c.String) {
				if i == n-1 {
					c.PrintString(c.X0)
				} else {
					c.PrintStringInline(c.X0)
				}
			} else if value.Type == c.StackObjectType(c.Bool) {
				c.PrintInt(c.X0)
			} else if value.Type == c.StackObjectType(c.Float) {
				c.PrintFloat()
			}
		}
	}

	return nil
}

func (v *Visitor) VisitAnd(ctx *parser.AndContext) interface{} {
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	// TODO: Reemplazar por PopObject y agregar manejo de tipos

	c.PopObject(c.X1)                 // Pop second operand
	var left = c.PopObject(c.X0)      // Pop first operand
	c.And(c.X0, c.X0, c.X1)           // X0 = X0 AND X1
	c.Push(c.X0)                      // Push the result back
	c.PushObject(c.CloneObject(left)) // Push the left operand object to the stack
	c.Comment("Logical AND operation")
	return nil
}

func (v *Visitor) VisitOr(ctx *parser.OrContext) interface{} {
	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	// TODO: Reemplazar por PopObject y agregar manejo de tipos

	c.PopObject(c.X1)                 // Pop second operand
	var left = c.PopObject(c.X0)      // Pop first operand
	c.Orr(c.X0, c.X0, c.X1)           // X0 = X0 OR X1
	c.Push(c.X0)                      // Push the result back
	c.PushObject(c.CloneObject(left)) // Push the left operand object to the stack
	c.Comment("Logical OR operation")
	return nil
}

func (v *Visitor) VisitGreaterLess(ctx *parser.GreaterLessContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    c.PopObject(c.X1)            // Pop second operand
    c.PopObject(c.X0)            // Pop first operand

    c.Comment("Comparison operation: " + op)
    switch op {
    case ">":
        c.Cmp(c.X0, c.X1)
        c.Cset(c.X0, "gt")
    case "<":
        c.Cmp(c.X0, c.X1)
        c.Cset(c.X0, "lt")
    case ">=":
        c.Cmp(c.X0, c.X1)
        c.Cset(c.X0, "ge")
    case "<=":
        c.Cmp(c.X0, c.X1)
        c.Cset(c.X0, "le")
    }

    c.Push(c.X0)
    c.PushObject(c.BoolObject()) // <-- SIEMPRE push de booleano
    return nil
}

func (v *Visitor) VisitEqual(ctx *parser.EqualContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    right := c.PopObject(c.X1)
    left := c.PopObject(c.X0)
    c.Comment("Comparison operation: " + op)

    // Comparación de strings
    if left.Type == c.String && right.Type == c.String {
        c.UsedFunction("strcmp")
        c.MovReg("x2", "x0") // x2 = primer string
        c.MovReg("x3", "x1") // x3 = segundo string
        c.Call("strcmp")     // strcmp(x2, x3), resultado en x0
        if op == "==" {
            c.Cmp("x0", "#0")
            c.Cset("x0", "eq")
        } else {
            c.Cmp("x0", "#0")
            c.Cset("x0", "ne")
        }
    } else {
        // Comparación normal (int, bool, etc)
        switch op {
        case "==":
            c.Cmp(c.X0, c.X1)
            c.Cset(c.X0, "eq")
        case "!=":
            c.Cmp(c.X0, c.X1)
            c.Cset(c.X0, "ne")
        }
    }

    c.Push(c.X0)
    c.PushObject(c.BoolObject()) // <-- SIEMPRE push de booleano
    return nil
}

func (v *Visitor) VisitNegate(ctx *parser.NegateContext) interface{} {
	var op = ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()

	v.Visit(ctx.GetChild(1).(antlr.ParseTree))

	var operand = c.PopObject(c.X0) // Pop the operand
	c.Comment("Negation operation: " + op)

	switch op {
	case "!":
		c.CmpImm(c.X0, 0)
		c.Cset(c.X0, "eq")
	case "-":
		c.Neg(c.X0, c.X0)
	}

	c.Push(c.X0)
	c.PushObject(c.CloneObject(operand))
	return nil
}

func (v *Visitor) VisitExplicitDeclaration(ctx *parser.ExplicitDeclarationContext) interface{} {
    var varName string = ctx.ID().(*antlr.TerminalNodeImpl).GetText()

    if ctx.ExpressionStatement() != nil {
        v.Visit(ctx.ExpressionStatement())
        // Solo etiqueta, no hagas push extra
        c.TagObject(varName)
    } else {
        // Declaración sin inicialización: sí pushea un 0
        c.Mov(c.X0, 0)
        c.Push(c.X0)
        c.PushObject(c.CloneObject(c.IntObject()))
        c.TagObject(varName)
    }

    return nil
}
func (v *Visitor) VisitImplicitDeclaration(ctx *parser.ImplicitDeclarationContext) interface{} {
    var varName string = ctx.ID().GetText()
    log.Printf("🔍 DEBUG: ImplicitDeclaration para variable: %s", varName)

    // Evalúa la expresión y deja el resultado en la pila real y virtual
    v.Visit(ctx.ExpressionStatement())
    log.Printf("🔍 DEBUG: Expresión evaluada para %s", varName)

    // Forzar tipo int si la expresión es un entero literal
    obj := c.TopObject()
    log.Printf("🔍 DEBUG: TopObject tipo: %v", obj.Type)
    
    if _, ok := ctx.ExpressionStatement().GetChild(0).(antlr.ParseTree).(*parser.IntegerContext); ok {
        obj.Type = c.Int
        // CORREGIR: Actualizar el objeto en la pila
        c.PopObject("")  // Pop sin mover a registro
        c.PushObject(obj)
        log.Printf("🔍 DEBUG: Tipo forzado a Int para %s", varName)
    }

    // Solo etiqueta el objeto
    c.TagObject(varName)
    log.Printf("🔍 DEBUG: Variable %s etiquetada correctamente", varName)

    return nil
}

func (v *Visitor) VisitImplicitSliceDeclaration(ctx *parser.ImplicitSliceDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	var elemType c.StackObjectType

	// Determina el tipo declarado
	switch ctx.TYPE().GetText() {
	case "int":
		elemType = c.Int
	case "string":
		elemType = c.String
	case "float64":
		elemType = c.Float
	}

	// Visita los elementos del slice
	v.Visit(ctx.SliceElements())

	// Ajusta el tipo del slice si es necesario (por ejemplo, si está vacío)
	obj := c.TopObject()
	if obj.IsSlice && obj.Size == 0 {
		obj.ElemType = elemType
		obj.Type = elemType
		c.PopObject(c.X0)
		c.PushObject(obj)
	}

	// Etiqueta el objeto en la pila virtual con el nombre de la variable
	c.TagObject(varName)

	return nil
}

func (v *Visitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
	var assignee = ctx.ExpressionStatement(0)

	if idContext, ok := assignee.(*parser.IdentifierContext); ok {
		var varName string = idContext.GetText()
		c.Comment("Assignment to variable: " + varName)

		v.Visit(ctx.ExpressionStatement(1))
		c.PopObject(c.X0)

		offset, varObject := c.GetObject(varName)
		c.Mov(c.X1, offset)
		c.Add(c.X1, c.SP, c.X1)
		c.Str(c.X0, c.X1, 0)

		c.Push(c.X0)
		c.PushObject(c.CloneObject(varObject))
	}

	return nil
}

func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
    var id string = ctx.ID().GetText()
    
    // Verificar si es un parámetro de la función actual
    if paramIndex, isParam := currentFunctionParams[id]; isParam {
        // Para parámetros: HACER push desde el registro correspondiente
        reg := fmt.Sprintf("x%d", paramIndex)
        c.Push(reg)
        c.PushObject(c.IntObject())
        return nil
    }
    
    // Variables normales
    offset, object := c.GetObject(id)
    if offset == -1 {
        log.Printf("Variable no encontrada: %s", id)
        return nil
    }

    if object.Type == c.Float {
        c.LdrF(c.D0, c.SP, offset)
        c.Push(c.D0)
    } else {
        c.Ldr(c.X0, c.SP, offset)
        c.Push(c.X0)
    }

    c.PushObject(c.CloneObject(object))
    return nil
}

func (v *Visitor) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
    c.Comment("Entering block statement")
    // NO crear nuevo scope aquí si ya estamos en una función
    // c.NewScope() - comentar esta línea

    for _, statement := range ctx.AllStatement() {
        result := v.Visit(statement.(antlr.ParseTree))
		if result == "break" || result == "continue" {
			return result
		}
    }

    // Comentar también el EndScope
    /* var bytesToRemove int = c.EndScope() // End the scope and get the bytes to remove from the stack
    if bytesToRemove > 0 {
        c.Mov(c.X0, bytesToRemove) // Move the bytes to remove to X0
        c.Add(c.SP, c.SP, c.X0)    // Adjust the stack pointer
        c.Comment("Stack pointer adjusted")
    } */
	

    return nil
}

var currentFunctionParams map[string]int = make(map[string]int)

func (v *Visitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
    funcName := ctx.ID(0).GetText()
    
    // Limpiar el mapeo anterior
    currentFunctionParams = make(map[string]int)
    
    c.StartFunction(funcName)

    var params []string
    if ctx.ParameterList() != nil {
        paramList := ctx.ParameterList()
        for _, paramCtx := range paramList.AllParameter() {
            paramIds := paramCtx.AllID()
            if len(paramIds) > 0 {
                paramName := paramIds[0].GetText()
                params = append(params, paramName)
            }
        }
    }
    
    // ARREGLAR: NO guardar parámetros en la pila, solo mapearlos
    for i, paramName := range params {
        if i < 4 {
            currentFunctionParams[paramName] = i // a=0, b=1, etc.
            // NO hacer PushObject ni TagObject aquí
        }
    }

    v.Visit(ctx.BlockStatement())
    
    c.EndFunction(funcName)
    
    // Limpiar después de la función
    currentFunctionParams = make(map[string]int)
    return nil
}

func (v *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
    //fmt.Printf("🔍 DEBUG: VisitChildren llamado para: %T\n", node)
    
    // Si es un contexto de función, manejarlo especialmente
    /* nodeStr := fmt.Sprintf("%T", node)
    if strings.Contains(nodeStr, "Function") {
        fmt.Printf("🔧 DEBUG: Posible función detectada: %s\n", nodeStr)
    } */
    
    // Llamar al comportamiento por defecto
    return v.BaseLanguageVisitor.VisitChildren(node)
}

var inFunctionArgs bool = false

func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
    fnName := ctx.ID().GetText()
    args := []antlr.ParseTree{}
    if ctx.ArgumentList() != nil {
        exprs := ctx.ArgumentList().AllExpressionStatement()
        for _, e := range exprs {
            args = append(args, e.(antlr.ParseTree))
        }
    }

    switch fnName {
    case "len":
        if len(args) != 1 {
            log.Fatal("len() espera un solo argumento")
        }
        v.Visit(args[0])
        obj := c.PopObject(c.X0)
        if obj.IsSlice {
            c.PushConstant(obj.Size, c.IntObject())
        } else {
            log.Fatal("len() solo soporta slices")
        }
        return nil
    default:
        // EVALUAR todos los argumentos PRIMERO (van a la pila)
        for _, expr := range args {
            v.Visit(expr)
        }
        // DESPUÉS moverlos a registros en ORDEN INVERSO
        for i := len(args) - 1; i >= 0 && i < 4; i-- {
            reg := fmt.Sprintf("x%d", i)
            c.PopObject(reg) // Pop en orden inverso
        }
        c.Call(fnName)
        c.Push(c.X0)         // El resultado está en x0
        c.PushObject(c.IntObject())
        return nil
    }
}

var inReturn bool = false

func (v *Visitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
    if ctx.ExpressionStatement() != nil {
        v.Visit(ctx.ExpressionStatement())
        c.PopObject(c.X0) // El resultado debe estar en X0
    }
    return nil
}

func (v *Visitor) VisitAtoiExpr(ctx *parser.AtoiExprContext) interface{} {
    // Obtenemos el nodo hijo que es el atoiStatement
    atoiStmt := ctx.AtoiStatement()
    if atoiStmt == nil {
        log.Fatal("Error: AtoiExprContext sin atoiStatement")
    }

    // Visitamos el argumento dentro de atoi(...)
    argExpr := atoiStmt.ExpressionStatement()
    v.Visit(argExpr)

    // Sacamos el string de la pila virtual
    arg := c.PopObject(c.X0)
    if arg.Type != c.String {
        log.Fatal("atoi solo acepta strings como argumento")
    }

    c.UsedFunction("atoi")
    c.Comment("Llamando a función atoi")
    c.Call("atoi")
    c.Comment("Valor en x0 tras atoi (debe ser entero válido)")

    // Empujamos el resultado tanto en la pila real como virtual
    c.Push(c.X0)
    c.PushObject(c.IntObject())

    return nil
}

//---------------------------------------------------------------------------------------

func (v *Visitor) VisitSliceLiteral(ctx *parser.SliceLiteralContext) interface{} {
	exprs := ctx.AllExpressionStatement()
	n := len(exprs)

	c.Comment(fmt.Sprintf("Slice literal con %d elementos", n))

	// Reserva espacio para el slice en el heap y avanza el heap pointer
	c.MovReg(c.X9, c.HP) // x9 = base del slice
	if n > 0 {
		c.Addi(c.HP, c.HP, n*8) // Avanza el heap pointer para NO pisar el slice
	}

	var elemType c.StackObjectType

	for i, expr := range exprs {
		// Detecta si el elemento es un string
		if strCtx, ok := expr.(*parser.StringContext); ok {
			value := strCtx.GetText()
			value = strings.Trim(value, `"`)
			c.Comment("Constant String (slice): " + value)
			c.PushStringNoStack(value)
			c.Str("x11", c.X9, i*8)
			if i == 0 {
				elemType = c.String
			}
		} else {
			v.Visit(expr)
			obj := c.TopObject()
			if i == 0 {
				elemType = obj.Type
			}
			if obj.Type == c.Float {
				c.PopObject(c.D0)
				c.StrF(c.D0, c.X9, i*8)
			} else {
				c.PopObject(c.X0)
				c.Str(c.X0, c.X9, i*8)
			}
		}
	}

	if n == 0 {
		c.Comment("Slice vacío, no se inicializan elementos")
		elemType = c.StackObjectType(c.Int) // Por defecto, puedes cambiarlo
	} else {
		c.Comment(fmt.Sprintf("Slice inicializado con %d elementos de tipo %s", n, elemType))
	}

	// Push la dirección base como el slice
	c.Push(c.X9)
	c.PushObject(c.SliceObject(elemType, n))

	return nil
}

func (v *Visitor) VisitSliceElements(ctx *parser.SliceElementsContext) interface{} {
	exprs := ctx.AllExpressionStatement()
	n := len(exprs)

	c.Comment(fmt.Sprintf("Slice literal con %d elementos", n))

	// Reserva espacio para el slice en el heap y avanza el heap pointer
	c.MovReg(c.X9, c.HP) // x9 = base del slice
	if n > 0 {
		c.Addi(c.HP, c.HP, n*8) // Avanza el heap pointer para NO pisar el slice
	}

	var elemType c.StackObjectType

	for i, expr := range exprs {
		v.Visit(expr)
		obj := c.TopObject()
		if i == 0 {
			elemType = obj.Type
		}
		switch obj.Type {
		case c.Float:
			c.PopObject(c.D0)
			c.StrF(c.D0, c.X9, i*8)
		case c.String:
			c.PopObject(c.X0)
			c.Str(c.X0, c.X9, i*8)
		default:
			c.PopObject(c.X0)
			c.Str(c.X0, c.X9, i*8)
		}
	}

	if n == 0 {
		c.Comment("Slice vacío, no se inicializan elementos")
		elemType = c.StackObjectType(c.Int) // Por defecto, puedes cambiarlo
	} else {
		c.Comment(fmt.Sprintf("Slice inicializado con %d elementos de tipo %s", n, elemType))
	}

	// Push la dirección base como el slice
	c.Push(c.X9)
	c.PushObject(c.SliceObject(elemType, n))

	return nil
}

func (v *Visitor) VisitSliceAccess(ctx *parser.SliceAccessContext) interface{} {
	// ID '[' expressionStatement ']'
	id := ctx.ID().GetText()
	v.Visit(ctx.ExpressionStatement())
	c.PopObject(c.X1) // índice

	offset, obj := c.GetObject(id)
	c.Mov(c.X0, offset)
	c.Add(c.X0, c.SP, c.X0)
	c.Ldr(c.X0, c.X0, 0) // x0 = dirección base del slice

	// Calcula la dirección del elemento: base + índice*8
	c.MovReg(c.X2, c.X1)
	c.Mov(c.X3, 8)
	c.Mul(c.X2, c.X2, c.X3)
	c.Add(c.X0, c.X0, c.X2)

	if obj.ElemType == c.Float {
		c.LdrF(c.D0, c.X0, 0)
		c.Push(c.D0)
		c.PushObject(c.StackObject{
			Type:     c.Float,
			Length:   8,
			Depth:    obj.Depth,
			Id:       nil,
			IsSlice:  false,
			ElemType: c.Float,
			Size:     1,
		})
	} else {
		c.Ldr(c.X0, c.X0, 0)
		c.Push(c.X0)
		c.PushObject(c.StackObject{
			Type:     obj.ElemType,
			Length:   8,
			Depth:    obj.Depth,
			Id:       nil,
			IsSlice:  false,
			ElemType: obj.ElemType,
			Size:     1,
		})
	}

	return nil
}

func (v *Visitor) VisitSliceAssignment(ctx *parser.SliceAssignmentContext) interface{} {
	// ID '[' expressionStatement ']' '=' expressionStatement
	id := ctx.ID().GetText()

	// Evalúa el índice
	v.Visit(ctx.ExpressionStatement(0))
	c.PopObject(c.X1) // índice

	// Evalúa el valor a asignar
	v.Visit(ctx.ExpressionStatement(1))
	c.PopObject(c.X2) // valor

	// Obtiene la dirección base del slice
	offset, obj := c.GetObject(id)
	c.Mov(c.X0, offset)
	c.Add(c.X0, c.SP, c.X0)
	c.Ldr(c.X0, c.X0, 0) // x0 = dirección base del slice

	// Calcula la dirección del elemento: base + índice*8
	c.MovReg(c.X3, c.X1)
	c.Mov(c.X4, 8)
	c.Mul(c.X3, c.X3, c.X4)
	c.Add(c.X0, c.X0, c.X3)

	// Guarda el valor en la posición calculada
	c.Str(c.X2, c.X0, 0)

	// Push el valor asignado
	c.Push(c.X2)
	c.PushObject(c.StackObject{
		Type:     obj.ElemType,
		Length:   8,
		Depth:    obj.Depth,
		Id:       nil,
		IsSlice:  false,
		ElemType: obj.ElemType,
		Size:     1,
	})

	return nil
}

func (v *Visitor) VisitExplicitSliceDeclaration(ctx *parser.ExplicitSliceDeclarationContext) interface{} {
	varName := ctx.ID().GetText()
	var elemType c.StackObjectType

	// Determina el tipo declarado
	switch ctx.TYPE().GetText() {
	case "int":
		elemType = c.Int
	case "string":
		elemType = c.String
	case "float64":
		elemType = c.Float
	}

	if ctx.ExpressionStatement() != nil {
		v.Visit(ctx.ExpressionStatement())
		obj := c.TopObject()
		// Si el slice es vacío, ajusta el tipo aquí:
		if obj.IsSlice && obj.Size == 0 {
			obj.ElemType = elemType
			obj.Type = elemType
			c.PopObject(c.X0)
			c.PushObject(obj)
		}
		c.TagObject(varName)
	} else {
		// Declaración sin inicialización: crea slice vacío del tipo correcto
		c.Mov(c.X0, 0) // Dirección nula o como manejes slices vacíos
		c.Push(c.X0)
		c.PushObject(c.SliceObject(elemType, 0))
		c.TagObject(varName)
	}
	return nil
}

func (v *Visitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	c.Comment("If statement")

	labelElse := fmt.Sprintf("else_%p", ctx)
	labelEnd := fmt.Sprintf("endif_%p", ctx)

	// 1. Evalúa la condición (hijo 1)
	v.Visit(ctx.ExpressionStatement())

	c.PopObject(c.X0)
	c.CmpImm(c.X0, 0)
	c.BranchEq(labelElse)

	// 2. Bloque del if (hijo 2)
	result := v.Visit(ctx.BlockStatement(0))

	if result == "break" || result == "continue" {
		// Si el bloque if devolvió "break" o "continue", salimos
		log.Print("Se encontró un break o continue en el bloque if")
		c.Branch(labelEnd) 
		c.Label(labelEnd)
		return result
	}

	c.Branch(labelEnd)

	// 3. Else (si existe)
	c.Label(labelElse)
	// Si hay al menos 1 hijo más después del bloque if
	if ctx.GetChildCount() > 3 {
		// El hijo 3 es 'else', el hijo 4 es ifStatement o blockStatement
		elseNode := ctx.GetChild(4)
		if elseIfCtx, ok := elseNode.(*parser.IfStatementContext); ok {
			result := v.Visit(elseIfCtx)
			if result == "break" || result == "continue" {
				// Si el bloque else-if devolvió "break" o "continue", salimos
				c.Branch(labelEnd)
				c.Label(labelEnd)
				return result
			}
		} else if blockCtx, ok := elseNode.(*parser.BlockStatementContext); ok {
			result := v.Visit(blockCtx)
			if result == "break" || result == "continue" {
				// Si el bloque else devolvió "break" o "continue", salimos
				c.Branch(labelEnd)
				c.Label(labelEnd)
				return result
			}
		}
	}
	c.Label(labelEnd)

	return nil
}

func (v *Visitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	c.Comment("Switch statement")

	// 1. Evalúa el valor del switch (ID)
	id := ctx.ID().GetText()
	offset, _ := c.GetObject(id)
	c.Mov(c.X0, offset)
	c.Add(c.X0, c.SP, c.X0)
	c.Ldr(c.X0, c.X0, 0) // x0 = valor del switch

	// 2. Prepara labels
	nCases := len(ctx.AllSwitchCase())
	labelEnd := fmt.Sprintf("end_switch_%p", ctx)
	labelDefault := fmt.Sprintf("default_%p", ctx)
	caseLabels := make([]string, nCases)
	for i := range caseLabels {
		caseLabels[i] = fmt.Sprintf("case_%d_%p", i, ctx)
	}

	// 3. Compara con cada case
	c.MovReg("x2", "x0") // Guarda el valor del switch en x2
	for i, caseCtx := range ctx.AllSwitchCase() {
		v.Visit(caseCtx.ExpressionStatement())
		c.PopObject(c.X1) // valor del case en x1
		c.Cmp("x2", "x1") // compara el valor original del switch con el del case
		c.BranchEq(caseLabels[i])
	}
	// Si no hay match, salta a default (si existe) o al final
	if ctx.DefaultCase() != nil {
		c.Branch(labelDefault)
	} else {
		c.Branch(labelEnd)
	}

	// 4. Genera el código de cada case
	for i, caseCtx := range ctx.AllSwitchCase() {
		c.Label(caseLabels[i])
		for _, stmt := range caseCtx.AllStatement() {
			result := v.Visit(stmt)
			if result == "break"{
				// Si el bloque case devolvió "break", salimos salimos del switch
				c.Branch(labelEnd)
				c.Label(labelEnd)
				return nil
			} else if result == "continue" {
				// Si el bloque case devolvió "continue" salimos del switch
				c.Branch(labelEnd)
				c.Label(labelEnd)
				return result
			}

		}
		c.Branch(labelEnd)
	}

	// 5. Default case (si existe)
	if ctx.DefaultCase() != nil {
		c.Label(labelDefault)
		for _, stmt := range ctx.DefaultCase().AllStatement() {
			result := v.Visit(stmt)
			if result == "break" {
				// Si el bloque default devolvió "break", salimos del switch
				c.Branch(labelEnd)
				c.Label(labelEnd)
				return nil
			}
			if result == "continue" {
				// Si el bloque default devolvió "continue", salimos del switch
				c.Branch(labelEnd)
				c.Label(labelEnd)
				return result
			}
		}
	}

	// 6. Fin del switch
	c.Label(labelEnd)
	return nil
}

func (v *Visitor) VisitForConditional(ctx *parser.ForConditionalContext) interface{} {
	c.Comment("While statement")
	labelStart := fmt.Sprintf("while_start_%p", ctx)
	labelEnd := fmt.Sprintf("while_end_%p", ctx)
	c.Label(labelStart)

	// 1. Evalúa la condición (hijo 1)
	v.Visit(ctx.ExpressionStatement())

	c.PopObject(c.X0) // Sacamos el resultado de la condición
	c.CmpImm(c.X0, 0)
	c.BranchEq(labelEnd) // Si la condición es falsa, salta al final

	// 2. Bloque del while (hijo 2)
	result := v.Visit(ctx.BlockStatement())
    if result == "break" {
		log.Print("Se encontró un break en el bloque while")
        c.Branch(labelEnd) // Break sale del while
    } else if result == "continue" {
        c.Branch(labelStart) // Continue vuelve al inicio del while
    } else {
        c.Branch(labelStart) // Comportamiento normal, vuelve al inicio
    }
    
    c.Label(labelEnd) // Marca el final del while

    return nil
}

func (v *Visitor) VisitForSimple(ctx *parser.ForSimpleContext) interface{} {
    c.Comment("For loop statement")
    
    // Crear labels únicos para este for
    labelStart := fmt.Sprintf("for_start_%p", ctx)
    labelUpdate := fmt.Sprintf("for_update_%p", ctx)
    labelEnd := fmt.Sprintf("for_end_%p", ctx)
    
    // Crear nuevo scope para las variables del for
    c.NewScope()
    
    // 1. Inicialización
    c.Comment("For initialization")
    if ctx.VariableDeclaration() != nil {
        v.Visit(ctx.VariableDeclaration())
    } else if len(ctx.AllExpressionStatement()) >= 1 {
        v.Visit(ctx.AllExpressionStatement()[0])
    }
    
    // 2. Label del inicio del loop
    c.Label(labelStart)
    
    // 3. Condición (PRIMER expressionStatement) ⚡ CORREGIR AQUÍ
    c.Comment("For condition")
    if len(ctx.AllExpressionStatement()) >= 1 {
        v.Visit(ctx.AllExpressionStatement()[0])  // ⚡ CAMBIAR [1] por [0]
        c.PopObject(c.X0)
        c.CmpImm(c.X0, 0)
        c.BranchEq(labelEnd)
    }
    
    // 4. Cuerpo del for
    c.Comment("For body")
    result := v.Visit(ctx.BlockStatement())
    if result == "break" {
        c.Branch(labelEnd)
    } else if result == "continue" {
        c.Branch(labelUpdate)
    } else {
        c.Branch(labelUpdate)
    }
    
    // 5. Actualización
    c.Label(labelUpdate)
    c.Comment("For update")
    if len(ctx.AllExpressionStatement()) >= 2 {
        v.Visit(ctx.AllExpressionStatement()[1])
        c.PopObject(c.X0)
    }
    
    // 6. Volver al inicio
    c.Branch(labelStart)
    
    // 7. Salida
    c.Label(labelEnd)
    
    // Limpiar scope
    var bytesToRemove int = c.EndScope()
    if bytesToRemove > 0 {
        c.Mov(c.X0, bytesToRemove)
        c.Add(c.SP, c.SP, c.X0)
    }
    
    return nil
}

func (v *Visitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	c.Comment("Break statement")

	return "break"
}

func (v *Visitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	c.Comment("Continue statement")

	return "continue"
}