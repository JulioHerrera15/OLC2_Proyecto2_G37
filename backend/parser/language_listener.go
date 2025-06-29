// Code generated from parser/Language.g4 by ANTLR 4.13.2. DO NOT EDIT.

package interpreter // Language
import "github.com/antlr4-go/antlr/v4"

// LanguageListener is a complete listener for a parse tree produced by LanguageParser.
type LanguageListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterNonDeclaration is called when entering the nonDeclaration production.
	EnterNonDeclaration(c *NonDeclarationContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterSliceAssignment is called when entering the SliceAssignment production.
	EnterSliceAssignment(c *SliceAssignmentContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterSliceLiteral is called when entering the SliceLiteral production.
	EnterSliceLiteral(c *SliceLiteralContext)

	// EnterParseFloatExpr is called when entering the ParseFloatExpr production.
	EnterParseFloatExpr(c *ParseFloatExprContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterString is called when entering the String production.
	EnterString(c *StringContext)

	// EnterStructAssignment is called when entering the StructAssignment production.
	EnterStructAssignment(c *StructAssignmentContext)

	// EnterMatrixAccess is called when entering the MatrixAccess production.
	EnterMatrixAccess(c *MatrixAccessContext)

	// EnterAssignment is called when entering the Assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterStructAccess is called when entering the StructAccess production.
	EnterStructAccess(c *StructAccessContext)

	// EnterAddSubOperator is called when entering the AddSubOperator production.
	EnterAddSubOperator(c *AddSubOperatorContext)

	// EnterMulDivMod is called when entering the MulDivMod production.
	EnterMulDivMod(c *MulDivModContext)

	// EnterAtoiExpr is called when entering the AtoiExpr production.
	EnterAtoiExpr(c *AtoiExprContext)

	// EnterIdentifier is called when entering the Identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterEqual is called when entering the Equal production.
	EnterEqual(c *EqualContext)

	// EnterGreaterLess is called when entering the GreaterLess production.
	EnterGreaterLess(c *GreaterLessContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterBoolean is called when entering the Boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterSliceAccess is called when entering the SliceAccess production.
	EnterSliceAccess(c *SliceAccessContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterMatrixAssignment is called when entering the MatrixAssignment production.
	EnterMatrixAssignment(c *MatrixAssignmentContext)

	// EnterStructInstantiation is called when entering the StructInstantiation production.
	EnterStructInstantiation(c *StructInstantiationContext)

	// EnterInteger is called when entering the Integer production.
	EnterInteger(c *IntegerContext)

	// EnterNil is called when entering the Nil production.
	EnterNil(c *NilContext)

	// EnterTypeOfExpr is called when entering the TypeOfExpr production.
	EnterTypeOfExpr(c *TypeOfExprContext)

	// EnterFloat is called when entering the Float production.
	EnterFloat(c *FloatContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterNegate is called when entering the Negate production.
	EnterNegate(c *NegateContext)

	// EnterIncrementDecrement is called when entering the IncrementDecrement production.
	EnterIncrementDecrement(c *IncrementDecrementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterSwitchCase is called when entering the switchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the defaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterForSimple is called when entering the ForSimple production.
	EnterForSimple(c *ForSimpleContext)

	// EnterForConditional is called when entering the ForConditional production.
	EnterForConditional(c *ForConditionalContext)

	// EnterForSlice is called when entering the ForSlice production.
	EnterForSlice(c *ForSliceContext)

	// EnterBreakStatement is called when entering the BreakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the ContinueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the ReturnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterAtoiStatement is called when entering the atoiStatement production.
	EnterAtoiStatement(c *AtoiStatementContext)

	// EnterParseFloatStatement is called when entering the parseFloatStatement production.
	EnterParseFloatStatement(c *ParseFloatStatementContext)

	// EnterTypeOfStatement is called when entering the typeOfStatement production.
	EnterTypeOfStatement(c *TypeOfStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterExplicitDeclaration is called when entering the ExplicitDeclaration production.
	EnterExplicitDeclaration(c *ExplicitDeclarationContext)

	// EnterExplicitSliceDeclaration is called when entering the ExplicitSliceDeclaration production.
	EnterExplicitSliceDeclaration(c *ExplicitSliceDeclarationContext)

	// EnterExplicitMatrixDeclaration is called when entering the ExplicitMatrixDeclaration production.
	EnterExplicitMatrixDeclaration(c *ExplicitMatrixDeclarationContext)

	// EnterExplicitStructDeclaration is called when entering the ExplicitStructDeclaration production.
	EnterExplicitStructDeclaration(c *ExplicitStructDeclarationContext)

	// EnterImplicitStructDeclaration is called when entering the ImplicitStructDeclaration production.
	EnterImplicitStructDeclaration(c *ImplicitStructDeclarationContext)

	// EnterImplicitSliceDeclaration is called when entering the ImplicitSliceDeclaration production.
	EnterImplicitSliceDeclaration(c *ImplicitSliceDeclarationContext)

	// EnterImplicitMatrixDeclaration is called when entering the ImplicitMatrixDeclaration production.
	EnterImplicitMatrixDeclaration(c *ImplicitMatrixDeclarationContext)

	// EnterImplicitDeclaration is called when entering the ImplicitDeclaration production.
	EnterImplicitDeclaration(c *ImplicitDeclarationContext)

	// EnterSliceElements is called when entering the sliceElements production.
	EnterSliceElements(c *SliceElementsContext)

	// EnterMatrixElements is called when entering the matrixElements production.
	EnterMatrixElements(c *MatrixElementsContext)

	// EnterStructInitialization is called when entering the structInitialization production.
	EnterStructInitialization(c *StructInitializationContext)

	// EnterStructFieldInit is called when entering the structFieldInit production.
	EnterStructFieldInit(c *StructFieldInitContext)

	// EnterStructDeclaration is called when entering the structDeclaration production.
	EnterStructDeclaration(c *StructDeclarationContext)

	// EnterStructAtribute is called when entering the structAtribute production.
	EnterStructAtribute(c *StructAtributeContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitNonDeclaration is called when exiting the nonDeclaration production.
	ExitNonDeclaration(c *NonDeclarationContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitSliceAssignment is called when exiting the SliceAssignment production.
	ExitSliceAssignment(c *SliceAssignmentContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitSliceLiteral is called when exiting the SliceLiteral production.
	ExitSliceLiteral(c *SliceLiteralContext)

	// ExitParseFloatExpr is called when exiting the ParseFloatExpr production.
	ExitParseFloatExpr(c *ParseFloatExprContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitString is called when exiting the String production.
	ExitString(c *StringContext)

	// ExitStructAssignment is called when exiting the StructAssignment production.
	ExitStructAssignment(c *StructAssignmentContext)

	// ExitMatrixAccess is called when exiting the MatrixAccess production.
	ExitMatrixAccess(c *MatrixAccessContext)

	// ExitAssignment is called when exiting the Assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitStructAccess is called when exiting the StructAccess production.
	ExitStructAccess(c *StructAccessContext)

	// ExitAddSubOperator is called when exiting the AddSubOperator production.
	ExitAddSubOperator(c *AddSubOperatorContext)

	// ExitMulDivMod is called when exiting the MulDivMod production.
	ExitMulDivMod(c *MulDivModContext)

	// ExitAtoiExpr is called when exiting the AtoiExpr production.
	ExitAtoiExpr(c *AtoiExprContext)

	// ExitIdentifier is called when exiting the Identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitEqual is called when exiting the Equal production.
	ExitEqual(c *EqualContext)

	// ExitGreaterLess is called when exiting the GreaterLess production.
	ExitGreaterLess(c *GreaterLessContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitBoolean is called when exiting the Boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitSliceAccess is called when exiting the SliceAccess production.
	ExitSliceAccess(c *SliceAccessContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitMatrixAssignment is called when exiting the MatrixAssignment production.
	ExitMatrixAssignment(c *MatrixAssignmentContext)

	// ExitStructInstantiation is called when exiting the StructInstantiation production.
	ExitStructInstantiation(c *StructInstantiationContext)

	// ExitInteger is called when exiting the Integer production.
	ExitInteger(c *IntegerContext)

	// ExitNil is called when exiting the Nil production.
	ExitNil(c *NilContext)

	// ExitTypeOfExpr is called when exiting the TypeOfExpr production.
	ExitTypeOfExpr(c *TypeOfExprContext)

	// ExitFloat is called when exiting the Float production.
	ExitFloat(c *FloatContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitNegate is called when exiting the Negate production.
	ExitNegate(c *NegateContext)

	// ExitIncrementDecrement is called when exiting the IncrementDecrement production.
	ExitIncrementDecrement(c *IncrementDecrementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitSwitchCase is called when exiting the switchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the defaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitForSimple is called when exiting the ForSimple production.
	ExitForSimple(c *ForSimpleContext)

	// ExitForConditional is called when exiting the ForConditional production.
	ExitForConditional(c *ForConditionalContext)

	// ExitForSlice is called when exiting the ForSlice production.
	ExitForSlice(c *ForSliceContext)

	// ExitBreakStatement is called when exiting the BreakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the ContinueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the ReturnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitAtoiStatement is called when exiting the atoiStatement production.
	ExitAtoiStatement(c *AtoiStatementContext)

	// ExitParseFloatStatement is called when exiting the parseFloatStatement production.
	ExitParseFloatStatement(c *ParseFloatStatementContext)

	// ExitTypeOfStatement is called when exiting the typeOfStatement production.
	ExitTypeOfStatement(c *TypeOfStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitExplicitDeclaration is called when exiting the ExplicitDeclaration production.
	ExitExplicitDeclaration(c *ExplicitDeclarationContext)

	// ExitExplicitSliceDeclaration is called when exiting the ExplicitSliceDeclaration production.
	ExitExplicitSliceDeclaration(c *ExplicitSliceDeclarationContext)

	// ExitExplicitMatrixDeclaration is called when exiting the ExplicitMatrixDeclaration production.
	ExitExplicitMatrixDeclaration(c *ExplicitMatrixDeclarationContext)

	// ExitExplicitStructDeclaration is called when exiting the ExplicitStructDeclaration production.
	ExitExplicitStructDeclaration(c *ExplicitStructDeclarationContext)

	// ExitImplicitStructDeclaration is called when exiting the ImplicitStructDeclaration production.
	ExitImplicitStructDeclaration(c *ImplicitStructDeclarationContext)

	// ExitImplicitSliceDeclaration is called when exiting the ImplicitSliceDeclaration production.
	ExitImplicitSliceDeclaration(c *ImplicitSliceDeclarationContext)

	// ExitImplicitMatrixDeclaration is called when exiting the ImplicitMatrixDeclaration production.
	ExitImplicitMatrixDeclaration(c *ImplicitMatrixDeclarationContext)

	// ExitImplicitDeclaration is called when exiting the ImplicitDeclaration production.
	ExitImplicitDeclaration(c *ImplicitDeclarationContext)

	// ExitSliceElements is called when exiting the sliceElements production.
	ExitSliceElements(c *SliceElementsContext)

	// ExitMatrixElements is called when exiting the matrixElements production.
	ExitMatrixElements(c *MatrixElementsContext)

	// ExitStructInitialization is called when exiting the structInitialization production.
	ExitStructInitialization(c *StructInitializationContext)

	// ExitStructFieldInit is called when exiting the structFieldInit production.
	ExitStructFieldInit(c *StructFieldInitContext)

	// ExitStructDeclaration is called when exiting the structDeclaration production.
	ExitStructDeclaration(c *StructDeclarationContext)

	// ExitStructAtribute is called when exiting the structAtribute production.
	ExitStructAtribute(c *StructAtributeContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)
}
