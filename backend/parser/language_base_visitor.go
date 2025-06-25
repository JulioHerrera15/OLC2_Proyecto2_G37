// Code generated from Language.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Language
import "github.com/antlr4-go/antlr/v4"

type BaseLanguageVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLanguageVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitNonDeclaration(ctx *NonDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitBlockStatement(ctx *BlockStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSliceAssignment(ctx *SliceAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSliceLiteral(ctx *SliceLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitParseFloatExpr(ctx *ParseFloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStructAssignment(ctx *StructAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitMatrixAccess(ctx *MatrixAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStructAccess(ctx *StructAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAddSubOperator(ctx *AddSubOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitMulDivMod(ctx *MulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAtoiExpr(ctx *AtoiExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitEqual(ctx *EqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitGreaterLess(ctx *GreaterLessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSliceAccess(ctx *SliceAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitMatrixAssignment(ctx *MatrixAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStructInstantiation(ctx *StructInstantiationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitNil(ctx *NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitTypeOfExpr(ctx *TypeOfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitNegate(ctx *NegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitIncrementDecrement(ctx *IncrementDecrementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSwitchStatement(ctx *SwitchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitForConditional(ctx *ForConditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitForSimple(ctx *ForSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitForSlice(ctx *ForSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitAtoiStatement(ctx *AtoiStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitParseFloatStatement(ctx *ParseFloatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitTypeOfStatement(ctx *TypeOfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitExplicitDeclaration(ctx *ExplicitDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitExplicitSliceDeclaration(ctx *ExplicitSliceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitExplicitMatrixDeclaration(ctx *ExplicitMatrixDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitExplicitStructDeclaration(ctx *ExplicitStructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitImplicitStructDeclaration(ctx *ImplicitStructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitImplicitSliceDeclaration(ctx *ImplicitSliceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitImplicitMatrixDeclaration(ctx *ImplicitMatrixDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitImplicitDeclaration(ctx *ImplicitDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitSliceElements(ctx *SliceElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitMatrixElements(ctx *MatrixElementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStructInitialization(ctx *StructInitializationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStructFieldInit(ctx *StructFieldInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitStructAtribute(ctx *StructAtributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitParameterList(ctx *ParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}
