package visitor

import (
	c "backend/compiler/arm"
	parser "backend/parser"
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

	c.PopObject(c.X1)
	var left = c.PopObject(c.X0)

	// Esto permite que los operandos sean objetos del stack
	// y se manejen de manera más flexible.

	// Se pueden utilizar para obtener el tipo de dato de la siguiente forma:
	// right.Type
	// Lo cual permite saber que tipo de operaciones se pueden realizar o no
	// con los operandos, por ejemplo, si la operación es de entero con entero
	// o de entero con flotante, etc.

	switch op {
	case "+":
		c.Comment("Addition operator")
		c.Add(c.X0, c.X0, c.X1)
	case "-":
		c.Comment("Subtraction operator")
		c.Sub(c.X0, c.X0, c.X1)
	}

	c.Push(c.X0)
	// Pushear tambien en la pila virtual
	c.PushObject(c.CloneObject(left))

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

// Agregar al final de compiler_visitor.go
func (v *Visitor) VisitParens(ctx *parser.ParensContext) interface{} {

	result := v.Visit(ctx.ExpressionStatement())

	return result
}

func (v *Visitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	c.Comment("Print statement")

	expressions := ctx.AllExpressionStatement()

	if len(expressions) > 0 {

		// Visitar la primera expresión
		v.Visit(expressions[0])

		// Pop el resultado y imprimir
		var value = c.PopObject(c.X0)

		if value.Type == c.StackObjectType(c.Int) {
			c.PrintInt(c.X0) // Imprimir entero

		} else if value.Type == c.StackObjectType(c.String) {
			c.PrintString(c.X0) // Imprimir cadena
		} else if value.Type == c.StackObjectType(c.Bool) {
			c.PrintInt(c.X0) // Imprimir booleano como entero (1 o 0)
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
	functionName := ctx.ID().GetText()

	if functionName == "Atoi" {
		if ctx.ArgumentList() == nil || len(ctx.ArgumentList().AllExpressionStatement()) != 1 {
			log.Fatalf("Atoi espera 1 argumento.")
		}

		v.Visit(ctx.ArgumentList().ExpressionStatement(0))
		arg := c.PopObject(c.X0)

		if arg.Type != c.StackObjectType(c.String) {
			log.Fatalf("Atoi espera un argumento de tipo string.")
		}

		c.Comment("Llamada a Atoi")
		c.MovReg(c.X0, c.X0)
		c.Bl("atoi")
		c.Push(c.X0)
		c.PushObject(c.IntObject())

		c.Use("atoi") // <<< IMPORTANTE para que se incluya en el binario final
		return nil
	}

	log.Fatalf("Función no reconocida: %s", functionName)
	return nil
}
