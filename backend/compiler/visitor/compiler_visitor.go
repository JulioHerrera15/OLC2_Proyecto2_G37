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
    
    return result
}

func (v *Visitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
    var value = ctx.GetText()
    intValue, _ := strconv.Atoi(value)
    
    log.Printf("=== INTEGER DEBUG ===")
    log.Printf("Value: %s (%d)", value, intValue)
    
    c.Comment("Constant: " + value)
    c.Mov(c.X0, intValue)
    c.Push(c.X0)

    return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
    
    log.Printf("=== ADDSUB DEBUG ===")
    log.Printf("Left: '%s'", ctx.GetChild(0).(antlr.ParseTree).GetText())      // ✅ Cast to ParseTree
    log.Printf("Operator: '%s'", op)
    log.Printf("Right: '%s'", ctx.GetChild(2).(antlr.ParseTree).GetText())     // ✅ Cast to ParseTree
    
    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand

    switch op {
    case "+":
        c.Comment("Addition operator")
        c.Add(c.X0, c.X0, c.X1)
    case "-":
        c.Comment("Subtraction operator")
        c.Sub(c.X0, c.X0, c.X1)
    }

    c.Push(c.X0)
    log.Printf("Finished addition/subtraction")
    return nil
}

func (v *Visitor) VisitMulDivMod(ctx *parser.MulDivModContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
    
    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

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
    log.Printf("=== PARENS DEBUG ===")
    log.Printf("Content: '%s'", ctx.GetText())
    
    // Exactamente igual que el analyzer: visitar la expressionStatement directamente
    result := v.Visit(ctx.ExpressionStatement())
    
    log.Printf("Finished visiting parens content")
    return result
}

func (v *Visitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
    c.Comment("Print statement")
    
    log.Printf("=== PRINT STATEMENT DEBUG ===")
    log.Printf("Total children: %d", ctx.GetChildCount())
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        log.Printf("Child %d: %T - '%s'", i, child, child.(antlr.ParseTree).GetText()) // ✅ Cast to ParseTree
    }
    
    expressions := ctx.AllExpressionStatement()
    log.Printf("Found %d expressions", len(expressions))
    
    if len(expressions) > 0 {
        log.Printf("Expression 0: %T - '%s'", expressions[0], expressions[0].GetText()) // ✅ Ya es ParseTree
        
        // Visitar la primera expresión
        log.Printf("About to visit expression...")
        v.Visit(expressions[0])
        log.Printf("Finished visiting expression")
        
        // Pop el resultado y imprimir
        c.Pop(c.X0)
        c.PrintInt(c.X0)
    }

    return nil
}


func (v *Visitor) VisitAnd(ctx *parser.AndContext) interface{} {
    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))
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
    
    log.Printf("=== NEGATE DEBUG ===")
    log.Printf("Operator: '%s'", op)
    log.Printf("Expression to negate: '%s'", ctx.GetChild(1).(antlr.ParseTree).GetText()) // ✅ Cast to ParseTree
    
    v.Visit(ctx.GetChild(1).(antlr.ParseTree))
    
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
    log.Printf("Finished negation")
    return nil
}