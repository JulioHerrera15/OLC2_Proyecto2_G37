package visitor

import (
	parser "backend/parser"
	c "backend/compiler/arm"
	"strconv"
	"github.com/antlr4-go/antlr/v4"
    "log"
)

// Visitor para recorrer el árbol de sintaxis
type Visitor struct {
	*parser.BaseLanguageVisitor
};

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

    c.PopObject(c.X0) // Pop the result of the expression

    // Hacer pop con el objeto del stack para no descuadrar el stack
    // Si ocurre algun fallo, hay que eliminar la línea de código que hace pop
    // y dejar que el stack se maneje de manera automática.
    
    return result
}

func (v *Visitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
    var value = ctx.GetText()
    intValue, _ := strconv.Atoi(value)
    
    c.Comment("Constant: " + value)
    /* Reemplazado por el método de PushConstant
    c.Mov(c.X0, intValue)
    c.Push(c.X0) */

    var IntObject = c.IntObject()
    c.PushConstant(intValue, IntObject)

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
    
    /*var right  =*/ c.PopObject(c.X1)
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

    // TODO: Reemplazar por PopObject y agregar manejo de tipos

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand

    switch op {
    case "*":
        c.Comment("Multiplication operator")
        c.Mul(c.X0, c.X0, c.X1) // X0 = X0 * X1
    case "/":
        c.Comment("Division operator")
        c.SDiv(c.X0, c.X0, c.X1) // X0 = X0 / X1
    }

    c.Push(c.X0)

    return nil
}

func (v *Visitor) VisitIncrementDecrement(ctx *parser.IncrementDecrementContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
    
    v.Visit(ctx.GetChild(0).(antlr.ParseTree)) // Visit the variable

    // TODO: Reemplazar por PopObject y agregar manejo de tipos


    c.Pop(c.X0) // Pop the variable value

    switch op {
    case "++":
        c.Comment("Increment operator")
        c.Addi(c.X0, c.X0, 1) // Increment
    case "--":
        c.Comment("Decrement operator")
        c.Subi(c.X0, c.X0, 1) // Decrement
    }

    c.Push(c.X0) // Push the updated value back

    return nil
}

func (v *Visitor) VisitAddSubOperator(ctx *parser.AddSubOperatorContext) interface{} {
    var op = ctx.GetText()
    
    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    // TODO: Reemplazar por PopObject y agregar manejo de tipos

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand

    switch op {
    case "+=":
        c.Comment("Add assignment operator")
        c.Add(c.X0, c.X0, c.X1) // X0 = X0 + X1
    case "-=":
        c.Comment("Subtract assignment operator")
        c.Sub(c.X0, c.X0, c.X1) // X0 = X0 - X1
    }
    c.Push(c.X0) // Push the result back

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
        /*var value =*/ c.PopObject(c.X0)

        /*
        Agregar logica para imprimir el valor dependiendo de su tipo.
        Posiblemente sea con este if, aunque seria mejor un switch
        if value.Type == c.StackObjectType(c.Int) {
            
        }
        */
        c.PrintInt(c.X0)
    }

    return nil
}


func (v *Visitor) VisitAnd(ctx *parser.AndContext) interface{} {
    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    // TODO: Reemplazar por PopObject y agregar manejo de tipos

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand
    c.And(c.X0, c.X0, c.X1) // X0 = X0 AND X1
    c.Push(c.X0) // Push the result back
    c.Comment("Logical AND operation")
    return nil
}

func (v *Visitor) VisitOr(ctx *parser.OrContext) interface{} {
    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    // TODO: Reemplazar por PopObject y agregar manejo de tipos

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand
    c.Orr(c.X0, c.X0, c.X1) // X0 = X0 OR X1
    c.Push(c.X0) // Push the result back
    c.Comment("Logical OR operation")
    return nil
}

func (v *Visitor) VisitGreaterLess(ctx *parser.GreaterLessContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    // TODO: Reemplazar por PopObject y agregar manejo de tipos

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand
    c.Comment("Comparison operation: " + op)
    switch op {
    case ">":
        c.Cmp(c.X0, c.X1) // Compare X0 and X1
        c.Cset(c.X0, "gt") // Set X0 to 1 if X0 > X1, else 0
    case "<":
        c.Cmp(c.X0, c.X1) // Compare X0 and X1
        c.Cset(c.X0, "lt") // Set X0 to 1 if X0 < X1, else 0 
    case ">=":
        c.Cmp(c.X0, c.X1) // Compare X0 and X1
        c.Cset(c.X0, "ge") // Set X0 to 1 if X0 >= X1, else 0
    case "<=":
        c.Cmp(c.X0, c.X1) // Compare X0 and X1
        c.Cset(c.X0, "le") // Set X0 to 1 if X0 <= X1, else 0
    }

    c.Push(c.X0) // Push the result back
    return nil
}

func (v *Visitor) VisitEqual(ctx *parser.EqualContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()

    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    // TODO: Reemplazar por PopObject y agregar manejo de tipos

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand
    c.Comment("Comparison operation: " + op)
    
    switch op {
    case "==":
        c.Cmp(c.X0, c.X1) // Compare X0 and X1
        c.Cset(c.X0, "eq") // Set X0 to 1 if X0 == X1, else 0
    case "!=":
        c.Cmp(c.X0, c.X1) // Compare X0 and X1
        c.Cset(c.X0, "ne") // Set X0 to 1 if X0 != X1, else 0
    }

    c.Push(c.X0) // Push the result back
    return nil
}

func (v *Visitor) VisitNegate(ctx *parser.NegateContext) interface{} {
    var op = ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText()
    
    v.Visit(ctx.GetChild(1).(antlr.ParseTree))

    // TODO: Reemplazar por PopObject y agregar manejo de tipos
    
    c.Pop(c.X0) // Pop the operand
    c.Comment("Negation operation: " + op)

    switch op {
    case "!":
        c.CmpImm(c.X0, 0)
        c.Cset(c.X0, "eq")
    case "-":
        c.Neg(c.X0, c.X0)
    }

    c.Push(c.X0) // Push the result back
    return nil
}

func (v *Visitor) VisitExplicitDeclaration(ctx *parser.ExplicitDeclarationContext) interface{} {
    
    // TODO: Agregar manejo de mutabilidad
    /*var isMutable bool = false

    if ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetText() == "mut" {
        isMutable = true // La variable es mutable
    }*/
    
    var varName string = ctx.ID().(*antlr.TerminalNodeImpl).GetText()
    
    if ctx.ExpressionStatement() != nil {
        // Si hay una expresión, visitarla para obtener el valor inicial
        v.Visit(ctx.ExpressionStatement())
        
        c.TagObject(varName)

    } else {
        // Si no hay expresión, inicializar con cero
        c.Mov(c.X0, 0)
        c.PushObject(c.CloneObject(c.IntObject())) // Aquí se debería usar el tipo correcto
    }
    
    return nil
}

func (v *Visitor) VisitAssignment(ctx *parser.AssignmentContext) interface{} {
    var assignee = ctx.ExpressionStatement(0)

    if idContext, ok := assignee.(*parser.IdentifierContext); ok {
        var varName string = idContext.GetText()

        c.Comment("Assignment to variable: " + varName)

        v.Visit(ctx.ExpressionStatement(1)) 

        // Pop the value to assign
        /*var value =*/ c.PopObject(c.X0)

        // Push the value to the variable
        offset, varObject := c.GetObject(varName)

        c.Mov(c.X1, offset)

        c.Add(c.X1, c.SP, c.X1) // Calculate the address of the variable

        c.Str(c.X0, c.X1, 0) // Store the value in the variable's address

        // TODO: Agregar manejo de tipos para el objeto
        // Esto se hace con la variable value que esta comentada
        // Y comparando con el tipo del varObject

        c.Push(c.X0)
        c.PushObject(c.CloneObject(varObject)) // Push the variable object to the stack
        
    }

    return nil
}

func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
    var id string = ctx.ID().GetText()

    offset, object := c.GetObject(id)

    c.Mov(c.X0, offset) // Move the offset to X0
    c.Add(c.X0, c.SP, c.X0) // Calculate the address of

    c.Ldr(c.X0, c.X0, 0) // Load the value from the variable's address
    c.Push(c.X0) // Push the value to the stack

    var newObject = c.CloneObject(object)
    newObject.Id = nil // Set the ID of the object
    c.PushObject(newObject) // Push the object to the stack


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
        c.Add(c.SP, c.SP, c.X0) // Adjust the stack pointer
        c.Comment("Stack pointer adjusted")
    }
    
    return nil
}