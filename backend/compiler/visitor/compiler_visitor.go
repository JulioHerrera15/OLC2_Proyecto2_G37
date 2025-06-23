package visitor

import (
	parser "backend/parser"
	c "backend/compiler/arm"
	"strconv"
	"github.com/antlr4-go/antlr/v4"
	"fmt"
	"os"
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

// Agregar al final del archivo si no lo tienes
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
    switch t := tree.(type) {
    case *parser.ProgramContext:
        return v.VisitProgram(t)
    case *parser.StatementContext:
        return v.VisitStatement(t)
    case *parser.NonDeclarationContext:
        return v.VisitNonDeclaration(t)
    case *parser.PrintStatementContext:
        return v.VisitPrintStatement(t)
    case *parser.ExpressionStatementContext:
        return v.VisitExpressionStatement(t)
    case *parser.AddSubContext:
        return v.VisitAddSub(t)
    case *parser.IntegerContext:
        return v.VisitInteger(t)
    default:
        fmt.Fprintf(os.Stderr, "⚠️ Nodo no manejado: %T\n", t)
        // Para nodos no implementados, visita sus hijos
        if t.GetChildCount() > 0 {
            for i := 0; i < t.GetChildCount(); i++ {
                child := t.GetChild(i)
                if child != nil {
                    v.Visit(child.(antlr.ParseTree))
                }
            }
        }
        return nil
    }
}

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
    fmt.Fprintf(os.Stderr, "🔍 VisitProgram - Hijos: %d\n", ctx.GetChildCount())
    
    // Visitar todos los statements del programa
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        fmt.Fprintf(os.Stderr, "🔍 Hijo %d: %T\n", i, child)
        v.Visit(child.(antlr.ParseTree))
    }
    return nil
}

func (v *Visitor) VisitStatement(ctx *parser.StatementContext) interface{} {
    fmt.Fprintf(os.Stderr, "📄 VisitStatement - Hijos: %d\n", ctx.GetChildCount())
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        fmt.Fprintf(os.Stderr, "📄 Statement hijo %d: %T\n", i, child)
        v.Visit(child.(antlr.ParseTree))
    }
    return nil
}

func (v *Visitor) VisitNonDeclaration(ctx *parser.NonDeclarationContext) interface{} {
    fmt.Fprintf(os.Stderr, "🎯 VisitNonDeclaration - Hijos: %d\n", ctx.GetChildCount())
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        fmt.Fprintf(os.Stderr, "🎯 NonDeclaration hijo %d: %T\n", i, child)
        v.Visit(child.(antlr.ParseTree))
    }
    return nil
}

// También necesitas este para manejar expressionStatement dentro de print:
func (v *Visitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
    // Este método debe manejar todas las reglas de expressionStatement
    // Por ahora, simplemente visita los hijos
    for i := 0; i < ctx.GetChildCount(); i++ {
        child := ctx.GetChild(i)
        if child != nil {
            v.Visit(child.(antlr.ParseTree))
        }
    }
    return nil
}

func (v *Visitor) VisitInteger(ctx *parser.IntegerContext) interface{} {
	var value = ctx.GetText()
	intValue, _ := strconv.Atoi(value)
	c.Comment("Constant: " + value)
	c.Mov(c.X0, intValue)
	c.Push(c.X0)

	return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
    var op = ctx.GetChild(1).(*antlr.TerminalNodeImpl).GetText()
    
    v.Visit(ctx.GetChild(0).(antlr.ParseTree))
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))

    c.Pop(c.X1) // Pop second operand
    c.Pop(c.X0) // Pop first operand

    switch op {
    case "+":
        c.Add(c.X0, c.X0, c.X1) // X0 = X0 + X1
    case "-":
        c.Sub(c.X0, c.X0, c.X1) // X0 = X0 - X1
    }

    c.Push(c.X0)

    return nil
}

// Y agrega el método VisitPrintStatement si no lo tienes:
func (v *Visitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
    fmt.Fprintf(os.Stderr, "🖨️ VisitPrintStatement - ¡ENCONTRADO!\n")
    c.Comment("Print statement")
    
    // Visitar la expresión dentro de print()
    v.Visit(ctx.GetChild(2).(antlr.ParseTree))
    
    // Pop el valor a imprimir
    c.Pop(c.X0)
    
    // Llamar a la función de imprimir
    c.PrintInt(c.X0)

    return nil
}