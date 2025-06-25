// Code generated from Language.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Language
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by LanguageParser.
type LanguageVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LanguageParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by LanguageParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#nonDeclaration.
	VisitNonDeclaration(ctx *NonDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#SliceAssignment.
	VisitSliceAssignment(ctx *SliceAssignmentContext) interface{}

	// Visit a parse tree produced by LanguageParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by LanguageParser#SliceLiteral.
	VisitSliceLiteral(ctx *SliceLiteralContext) interface{}

	// Visit a parse tree produced by LanguageParser#ParseFloatExpr.
	VisitParseFloatExpr(ctx *ParseFloatExprContext) interface{}

	// Visit a parse tree produced by LanguageParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by LanguageParser#String.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by LanguageParser#StructAssignment.
	VisitStructAssignment(ctx *StructAssignmentContext) interface{}

	// Visit a parse tree produced by LanguageParser#MatrixAccess.
	VisitMatrixAccess(ctx *MatrixAccessContext) interface{}

	// Visit a parse tree produced by LanguageParser#Assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by LanguageParser#StructAccess.
	VisitStructAccess(ctx *StructAccessContext) interface{}

	// Visit a parse tree produced by LanguageParser#AddSubOperator.
	VisitAddSubOperator(ctx *AddSubOperatorContext) interface{}

	// Visit a parse tree produced by LanguageParser#MulDivMod.
	VisitMulDivMod(ctx *MulDivModContext) interface{}

	// Visit a parse tree produced by LanguageParser#AtoiExpr.
	VisitAtoiExpr(ctx *AtoiExprContext) interface{}

	// Visit a parse tree produced by LanguageParser#Identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by LanguageParser#Equal.
	VisitEqual(ctx *EqualContext) interface{}

	// Visit a parse tree produced by LanguageParser#GreaterLess.
	VisitGreaterLess(ctx *GreaterLessContext) interface{}

	// Visit a parse tree produced by LanguageParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by LanguageParser#Boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by LanguageParser#SliceAccess.
	VisitSliceAccess(ctx *SliceAccessContext) interface{}

	// Visit a parse tree produced by LanguageParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by LanguageParser#MatrixAssignment.
	VisitMatrixAssignment(ctx *MatrixAssignmentContext) interface{}

	// Visit a parse tree produced by LanguageParser#StructInstantiation.
	VisitStructInstantiation(ctx *StructInstantiationContext) interface{}

	// Visit a parse tree produced by LanguageParser#Integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by LanguageParser#Nil.
	VisitNil(ctx *NilContext) interface{}

	// Visit a parse tree produced by LanguageParser#TypeOfExpr.
	VisitTypeOfExpr(ctx *TypeOfExprContext) interface{}

	// Visit a parse tree produced by LanguageParser#Float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by LanguageParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by LanguageParser#Negate.
	VisitNegate(ctx *NegateContext) interface{}

	// Visit a parse tree produced by LanguageParser#IncrementDecrement.
	VisitIncrementDecrement(ctx *IncrementDecrementContext) interface{}

	// Visit a parse tree produced by LanguageParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#switchStatement.
	VisitSwitchStatement(ctx *SwitchStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#switchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by LanguageParser#defaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by LanguageParser#ForConditional.
	VisitForConditional(ctx *ForConditionalContext) interface{}

	// Visit a parse tree produced by LanguageParser#ForSimple.
	VisitForSimple(ctx *ForSimpleContext) interface{}

	// Visit a parse tree produced by LanguageParser#ForSlice.
	VisitForSlice(ctx *ForSliceContext) interface{}

	// Visit a parse tree produced by LanguageParser#BreakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#ContinueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#ReturnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#atoiStatement.
	VisitAtoiStatement(ctx *AtoiStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#parseFloatStatement.
	VisitParseFloatStatement(ctx *ParseFloatStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#typeOfStatement.
	VisitTypeOfStatement(ctx *TypeOfStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by LanguageParser#ExplicitDeclaration.
	VisitExplicitDeclaration(ctx *ExplicitDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#ExplicitSliceDeclaration.
	VisitExplicitSliceDeclaration(ctx *ExplicitSliceDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#ExplicitMatrixDeclaration.
	VisitExplicitMatrixDeclaration(ctx *ExplicitMatrixDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#ExplicitStructDeclaration.
	VisitExplicitStructDeclaration(ctx *ExplicitStructDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#ImplicitStructDeclaration.
	VisitImplicitStructDeclaration(ctx *ImplicitStructDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#ImplicitSliceDeclaration.
	VisitImplicitSliceDeclaration(ctx *ImplicitSliceDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#ImplicitMatrixDeclaration.
	VisitImplicitMatrixDeclaration(ctx *ImplicitMatrixDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#ImplicitDeclaration.
	VisitImplicitDeclaration(ctx *ImplicitDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#sliceElements.
	VisitSliceElements(ctx *SliceElementsContext) interface{}

	// Visit a parse tree produced by LanguageParser#matrixElements.
	VisitMatrixElements(ctx *MatrixElementsContext) interface{}

	// Visit a parse tree produced by LanguageParser#structInitialization.
	VisitStructInitialization(ctx *StructInitializationContext) interface{}

	// Visit a parse tree produced by LanguageParser#structFieldInit.
	VisitStructFieldInit(ctx *StructFieldInitContext) interface{}

	// Visit a parse tree produced by LanguageParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#structAtribute.
	VisitStructAtribute(ctx *StructAtributeContext) interface{}

	// Visit a parse tree produced by LanguageParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) interface{}

	// Visit a parse tree produced by LanguageParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by LanguageParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}
}
