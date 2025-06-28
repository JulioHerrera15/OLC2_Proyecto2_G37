package visitor

import (
	c "backend/compiler/arm"
	parser "backend/parser"
	"log"
	"strconv"
	"strings"
	"fmt"
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
	// Visitar todos los statements del programa
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		v.Visit(child.(antlr.ParseTree))
	}
	return nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		v.Visit(child.(antlr.ParseTree))
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
	c.PushConstant(intValue, IntObject)

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
	c.PushConstant(floatValue, floatObject)
	
	return nil
}

func (v *Visitor) VisitString(ctx *parser.StringContext) interface{} {
	var value = strings.Trim(ctx.GetText(), `"`) // Eliminar comillas
	c.Comment("Constant String: " + value)
	var StringObject = c.StringObject()
	c.PushConstant(value, StringObject)

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
	c.PushConstant(boolValue, BoolObject)

	return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	/* Reemplazado por el método de PopObject
	   c.Pop(c.X1) // Pop second operand
	   c.Pop(c.X0) // Pop first operand*/

	// Se pueden obtener los operandos directamente de la pila

	var isRightDouble = c.TopObject().Type == c.StackObjectType(c.Float)
	var rightReg interface{}
	if isRightDouble {
		rightReg = c.D0
	} else {
		rightReg = c.X0
	}
	var right = c.PopObject(rightReg.(string)) // Pop second operand

	var isLeftDouble = c.TopObject().Type == c.StackObjectType(c.Float)
	var leftReg interface{}
	if isLeftDouble {
		leftReg = c.D1
	} else {
		leftReg = c.X1
	}
	var left = c.PopObject(leftReg.(string)) // Pop first operand

	// Esto permite que los operandos sean objetos del stack
	// y se manejen de manera más flexible.

	// Se pueden utilizar para obtener el tipo de dato de la siguiente forma:
	// right.Type
	// Lo cual permite saber que tipo de operaciones se pueden realizar o no
	// con los operandos, por ejemplo, si la operación es de entero con entero
	// o de entero con flotante, etc.

	if isLeftDouble || isRightDouble {
		if !isLeftDouble {c.Scvtf(c.D1, c.X1)}
		if !isRightDouble {c.Scvtf(c.D0, c.X0)}

		switch op {
		case "+":
			c.Comment("Addition operator for double")
			c.FAdd(c.D0, c.D0, c.D1) // D0 = D0 + D1
		case "-":
			c.Comment("Subtraction operator for double")
			c.FSub(c.D0, c.D1, c.D0) // D0 = D1 - D0
		}

		c.Comment("Pushing result of double operation")
		c.Push(c.D0) // Push the result of the double operation
		c.PushObject(c.CloneObject(left)) // Push the left operand object to the stack

		return nil
	}

	switch op {
	case "+":
		c.Comment("Addition operator")
		c.Add(c.X0, c.X0, c.X1)
	case "-":
		c.Comment("Subtraction operator")
		c.Sub(c.X0, c.X1, c.X0)
	}

	c.Push(c.X0)
	// Pushear tambien en la pila virtual
	var objToPush interface{}
	if isLeftDouble {
		objToPush = left
	} else {
		objToPush = right
	}
	c.PushObject(c.CloneObject(objToPush.(c.StackObject)))

	// Se hace de momento con left a forma de demostración,
	// pero se debería hacer con el tipo de mayor predominancia.
	// Por ejemplo, si se hace una suma de un entero con un flotante,
	// se debería pushear un objeto flotante en la pila virtual.

	// TODO: Agregar lógica para pushear el tipo de dato correcto
	// dependiendo de la operación y los operandos.

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

	v.Visit(ctx.GetChild(0).(antlr.ParseTree)) // Visit the variable

	// TODO: Reemplazar por PopObject y agregar manejo de tipos

	var left = c.PopObject(c.X0) // Pop the variable value

	switch op {
	case "++":
		c.Comment("Increment operator")
		c.Addi(c.X0, c.X0, 1) // Increment
	case "--":
		c.Comment("Decrement operator")
		c.Subi(c.X0, c.X0, 1) // Decrement
	}

	c.Push(c.X0)                      // Push the updated value back
	c.PushObject(c.CloneObject(left)) // Push the original object to the stack

	return nil
}

func (v *Visitor) VisitAddSubOperator(ctx *parser.AddSubOperatorContext) interface{} {
	var op = ctx.GetText()

	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	// TODO: Reemplazar por PopObject y agregar manejo de tipos

	c.PopObject(c.X1)            // Pop second operand
	var left = c.PopObject(c.X0) // Pop first operand

	switch op {
	case "+=":
		c.Comment("Add assignment operator")
		c.Add(c.X0, c.X0, c.X1) // X0 = X0 + X1
	case "-=":
		c.Comment("Subtract assignment operator")
		c.Sub(c.X0, c.X0, c.X1) // X0 = X0 - X1
	}
	c.Push(c.X0)                      // Push the result back
	c.PushObject(c.CloneObject(left)) // Push the left operand object to the stack

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
			c.Pop(c.X9)           // Pop real: dirección base del slice
			c.PopObject(c.X0)     // Pop virtual: objeto del slice
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
	var left = c.PopObject(c.X0) // Pop first operand

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
	c.PushObject(c.CloneObject(left))
	return nil
}

func (v *Visitor) VisitEqual(ctx *parser.EqualContext) interface{} {
	var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

	v.Visit(ctx.GetChild(0).(antlr.ParseTree))
	v.Visit(ctx.GetChild(2).(antlr.ParseTree))

	// TODO: Reemplazar por PopObject y agregar manejo de tipos

	c.PopObject(c.X1)            // Pop second operand
	var left = c.PopObject(c.X0) // Pop first operand
	c.Comment("Comparison operation: " + op)

	switch op {
	case "==":
		c.Cmp(c.X0, c.X1)  // Compare X0 and X1
		c.Cset(c.X0, "eq") // Set X0 to 1 if X0 == X1, else 0
	case "!=":
		c.Cmp(c.X0, c.X1)  // Compare X0 and X1
		c.Cset(c.X0, "ne") // Set X0 to 1 if X0 != X1, else 0
	}

	c.Push(c.X0)                      // Push the result back
	c.PushObject(c.CloneObject(left)) // Push the left operand object to the stack
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
		c.TagObject(varName)
	} else {
		c.Mov(c.X0, 0)
		c.Push(c.X0)
		c.PushObject(c.CloneObject(c.IntObject()))
	}

	return nil
}

func (v *Visitor) VisitImplicitDeclaration(ctx *parser.ImplicitDeclarationContext) interface{} {
    var varName string = ctx.ID().GetText()

    // Evalúa la expresión y deja el resultado en la pila real y virtual
    v.Visit(ctx.ExpressionStatement())

    // Etiqueta el objeto en la pila virtual con el nombre de la variable
    c.TagObject(varName)

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

	offset, object := c.GetObject(id)
	c.Mov(c.X0, offset)
	c.Add(c.X0, c.SP, c.X0)
	c.Ldr(c.X0, c.X0, 0)

	c.Push(c.X0)

	var newObject = c.CloneObject(object)
	newObject.Id = nil
	c.PushObject(newObject)

	return nil
}
func (v *Visitor) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
	c.Comment("Entering block statement")
	c.NewScope()

	for _, statement := range ctx.AllStatement() {
		v.Visit(statement.(antlr.ParseTree))
	}

	var bytesToRemove int = c.EndScope() // End the scope and get the bytes to remove from the stack
	if bytesToRemove > 0 {
		c.Mov(c.X0, bytesToRemove) // Move the bytes to remove to X0
		c.Add(c.SP, c.SP, c.X0)    // Adjust the stack pointer
		c.Comment("Stack pointer adjusted")
	}

	return nil
}

//---------------------------------------------------------------------------------------
func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	functionName := strings.ToLower(ctx.ID().GetText())

	if functionName == "atoi" {
		if ctx.ArgumentList() == nil || len(ctx.ArgumentList().AllExpressionStatement()) != 1 {
			log.Fatalf("atoi espera 1 argumento.")
		}

		v.Visit(ctx.ArgumentList().ExpressionStatement(0))
		arg := c.PopObject(c.X0)

		if arg.Type != c.StackObjectType(c.String) {
			log.Fatalf("atoi espera un argumento de tipo string.")
		}

		c.Comment("Llamada a atoi")
		c.MovReg(c.X0, c.X0)
		c.Use("atoi")
		c.Bl("atoi")
		c.PushObject(c.IntObject())
		c.Push(c.X0)
		return nil
	}

	log.Fatalf("Función no reconocida: %s", functionName)
	return nil
}

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
    v.Visit(ctx.BlockStatement(0))

    c.Branch(labelEnd)

    // 3. Else (si existe)
    c.Label(labelElse)
    // Si hay al menos 1 hijo más después del bloque if
    if ctx.GetChildCount() > 3 {
        // El hijo 3 es 'else', el hijo 4 es ifStatement o blockStatement
        elseNode := ctx.GetChild(4)
        if elseIfCtx, ok := elseNode.(*parser.IfStatementContext); ok {
            v.Visit(elseIfCtx)
        } else if blockCtx, ok := elseNode.(*parser.BlockStatementContext); ok {
            v.Visit(blockCtx)
        }
    }
    c.Label(labelEnd)

    return nil
}