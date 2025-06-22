// Code generated from parser/Language.g4 by ANTLR 4.13.2. DO NOT EDIT.

package interpreter // Language
import "github.com/antlr4-go/antlr/v4"

// BaseLanguageListener is a complete listener for a parse tree produced by LanguageParser.
type BaseLanguageListener struct{}

var _ LanguageListener = &BaseLanguageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLanguageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLanguageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLanguageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLanguageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLanguageListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLanguageListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseLanguageListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseLanguageListener) ExitStatement(ctx *StatementContext) {}

// EnterNonDeclaration is called when production nonDeclaration is entered.
func (s *BaseLanguageListener) EnterNonDeclaration(ctx *NonDeclarationContext) {}

// ExitNonDeclaration is called when production nonDeclaration is exited.
func (s *BaseLanguageListener) ExitNonDeclaration(ctx *NonDeclarationContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseLanguageListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseLanguageListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterSliceAssignment is called when production SliceAssignment is entered.
func (s *BaseLanguageListener) EnterSliceAssignment(ctx *SliceAssignmentContext) {}

// ExitSliceAssignment is called when production SliceAssignment is exited.
func (s *BaseLanguageListener) ExitSliceAssignment(ctx *SliceAssignmentContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseLanguageListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseLanguageListener) ExitOr(ctx *OrContext) {}

// EnterSliceLiteral is called when production SliceLiteral is entered.
func (s *BaseLanguageListener) EnterSliceLiteral(ctx *SliceLiteralContext) {}

// ExitSliceLiteral is called when production SliceLiteral is exited.
func (s *BaseLanguageListener) ExitSliceLiteral(ctx *SliceLiteralContext) {}

// EnterParseFloatExpr is called when production ParseFloatExpr is entered.
func (s *BaseLanguageListener) EnterParseFloatExpr(ctx *ParseFloatExprContext) {}

// ExitParseFloatExpr is called when production ParseFloatExpr is exited.
func (s *BaseLanguageListener) ExitParseFloatExpr(ctx *ParseFloatExprContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseLanguageListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseLanguageListener) ExitParens(ctx *ParensContext) {}

// EnterString is called when production String is entered.
func (s *BaseLanguageListener) EnterString(ctx *StringContext) {}

// ExitString is called when production String is exited.
func (s *BaseLanguageListener) ExitString(ctx *StringContext) {}

// EnterStructAssignment is called when production StructAssignment is entered.
func (s *BaseLanguageListener) EnterStructAssignment(ctx *StructAssignmentContext) {}

// ExitStructAssignment is called when production StructAssignment is exited.
func (s *BaseLanguageListener) ExitStructAssignment(ctx *StructAssignmentContext) {}

// EnterMatrixAccess is called when production MatrixAccess is entered.
func (s *BaseLanguageListener) EnterMatrixAccess(ctx *MatrixAccessContext) {}

// ExitMatrixAccess is called when production MatrixAccess is exited.
func (s *BaseLanguageListener) ExitMatrixAccess(ctx *MatrixAccessContext) {}

// EnterAssignment is called when production Assignment is entered.
func (s *BaseLanguageListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production Assignment is exited.
func (s *BaseLanguageListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterStructAccess is called when production StructAccess is entered.
func (s *BaseLanguageListener) EnterStructAccess(ctx *StructAccessContext) {}

// ExitStructAccess is called when production StructAccess is exited.
func (s *BaseLanguageListener) ExitStructAccess(ctx *StructAccessContext) {}

// EnterAddSubOperator is called when production AddSubOperator is entered.
func (s *BaseLanguageListener) EnterAddSubOperator(ctx *AddSubOperatorContext) {}

// ExitAddSubOperator is called when production AddSubOperator is exited.
func (s *BaseLanguageListener) ExitAddSubOperator(ctx *AddSubOperatorContext) {}

// EnterMulDivMod is called when production MulDivMod is entered.
func (s *BaseLanguageListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production MulDivMod is exited.
func (s *BaseLanguageListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterAtoiExpr is called when production AtoiExpr is entered.
func (s *BaseLanguageListener) EnterAtoiExpr(ctx *AtoiExprContext) {}

// ExitAtoiExpr is called when production AtoiExpr is exited.
func (s *BaseLanguageListener) ExitAtoiExpr(ctx *AtoiExprContext) {}

// EnterIdentifier is called when production Identifier is entered.
func (s *BaseLanguageListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production Identifier is exited.
func (s *BaseLanguageListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterEqual is called when production Equal is entered.
func (s *BaseLanguageListener) EnterEqual(ctx *EqualContext) {}

// ExitEqual is called when production Equal is exited.
func (s *BaseLanguageListener) ExitEqual(ctx *EqualContext) {}

// EnterGreaterLess is called when production GreaterLess is entered.
func (s *BaseLanguageListener) EnterGreaterLess(ctx *GreaterLessContext) {}

// ExitGreaterLess is called when production GreaterLess is exited.
func (s *BaseLanguageListener) ExitGreaterLess(ctx *GreaterLessContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseLanguageListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseLanguageListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterBoolean is called when production Boolean is entered.
func (s *BaseLanguageListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production Boolean is exited.
func (s *BaseLanguageListener) ExitBoolean(ctx *BooleanContext) {}

// EnterSliceAccess is called when production SliceAccess is entered.
func (s *BaseLanguageListener) EnterSliceAccess(ctx *SliceAccessContext) {}

// ExitSliceAccess is called when production SliceAccess is exited.
func (s *BaseLanguageListener) ExitSliceAccess(ctx *SliceAccessContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseLanguageListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseLanguageListener) ExitAddSub(ctx *AddSubContext) {}

// EnterMatrixAssignment is called when production MatrixAssignment is entered.
func (s *BaseLanguageListener) EnterMatrixAssignment(ctx *MatrixAssignmentContext) {}

// ExitMatrixAssignment is called when production MatrixAssignment is exited.
func (s *BaseLanguageListener) ExitMatrixAssignment(ctx *MatrixAssignmentContext) {}

// EnterStructInstantiation is called when production StructInstantiation is entered.
func (s *BaseLanguageListener) EnterStructInstantiation(ctx *StructInstantiationContext) {}

// ExitStructInstantiation is called when production StructInstantiation is exited.
func (s *BaseLanguageListener) ExitStructInstantiation(ctx *StructInstantiationContext) {}

// EnterInteger is called when production Integer is entered.
func (s *BaseLanguageListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production Integer is exited.
func (s *BaseLanguageListener) ExitInteger(ctx *IntegerContext) {}

// EnterNil is called when production Nil is entered.
func (s *BaseLanguageListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production Nil is exited.
func (s *BaseLanguageListener) ExitNil(ctx *NilContext) {}

// EnterTypeOfExpr is called when production TypeOfExpr is entered.
func (s *BaseLanguageListener) EnterTypeOfExpr(ctx *TypeOfExprContext) {}

// ExitTypeOfExpr is called when production TypeOfExpr is exited.
func (s *BaseLanguageListener) ExitTypeOfExpr(ctx *TypeOfExprContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseLanguageListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseLanguageListener) ExitFloat(ctx *FloatContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseLanguageListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseLanguageListener) ExitAnd(ctx *AndContext) {}

// EnterNegate is called when production Negate is entered.
func (s *BaseLanguageListener) EnterNegate(ctx *NegateContext) {}

// ExitNegate is called when production Negate is exited.
func (s *BaseLanguageListener) ExitNegate(ctx *NegateContext) {}

// EnterIncrementDecrement is called when production IncrementDecrement is entered.
func (s *BaseLanguageListener) EnterIncrementDecrement(ctx *IncrementDecrementContext) {}

// ExitIncrementDecrement is called when production IncrementDecrement is exited.
func (s *BaseLanguageListener) ExitIncrementDecrement(ctx *IncrementDecrementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseLanguageListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseLanguageListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseLanguageListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseLanguageListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterSwitchCase is called when production switchCase is entered.
func (s *BaseLanguageListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production switchCase is exited.
func (s *BaseLanguageListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production defaultCase is entered.
func (s *BaseLanguageListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production defaultCase is exited.
func (s *BaseLanguageListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterForConditional is called when production ForConditional is entered.
func (s *BaseLanguageListener) EnterForConditional(ctx *ForConditionalContext) {}

// ExitForConditional is called when production ForConditional is exited.
func (s *BaseLanguageListener) ExitForConditional(ctx *ForConditionalContext) {}

// EnterForSimple is called when production ForSimple is entered.
func (s *BaseLanguageListener) EnterForSimple(ctx *ForSimpleContext) {}

// ExitForSimple is called when production ForSimple is exited.
func (s *BaseLanguageListener) ExitForSimple(ctx *ForSimpleContext) {}

// EnterForSlice is called when production ForSlice is entered.
func (s *BaseLanguageListener) EnterForSlice(ctx *ForSliceContext) {}

// ExitForSlice is called when production ForSlice is exited.
func (s *BaseLanguageListener) ExitForSlice(ctx *ForSliceContext) {}

// EnterBreakStatement is called when production BreakStatement is entered.
func (s *BaseLanguageListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production BreakStatement is exited.
func (s *BaseLanguageListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production ContinueStatement is entered.
func (s *BaseLanguageListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production ContinueStatement is exited.
func (s *BaseLanguageListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production ReturnStatement is entered.
func (s *BaseLanguageListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production ReturnStatement is exited.
func (s *BaseLanguageListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterAtoiStatement is called when production atoiStatement is entered.
func (s *BaseLanguageListener) EnterAtoiStatement(ctx *AtoiStatementContext) {}

// ExitAtoiStatement is called when production atoiStatement is exited.
func (s *BaseLanguageListener) ExitAtoiStatement(ctx *AtoiStatementContext) {}

// EnterParseFloatStatement is called when production parseFloatStatement is entered.
func (s *BaseLanguageListener) EnterParseFloatStatement(ctx *ParseFloatStatementContext) {}

// ExitParseFloatStatement is called when production parseFloatStatement is exited.
func (s *BaseLanguageListener) ExitParseFloatStatement(ctx *ParseFloatStatementContext) {}

// EnterTypeOfStatement is called when production typeOfStatement is entered.
func (s *BaseLanguageListener) EnterTypeOfStatement(ctx *TypeOfStatementContext) {}

// ExitTypeOfStatement is called when production typeOfStatement is exited.
func (s *BaseLanguageListener) ExitTypeOfStatement(ctx *TypeOfStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseLanguageListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseLanguageListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterExplicitDeclaration is called when production ExplicitDeclaration is entered.
func (s *BaseLanguageListener) EnterExplicitDeclaration(ctx *ExplicitDeclarationContext) {}

// ExitExplicitDeclaration is called when production ExplicitDeclaration is exited.
func (s *BaseLanguageListener) ExitExplicitDeclaration(ctx *ExplicitDeclarationContext) {}

// EnterExplicitSliceDeclaration is called when production ExplicitSliceDeclaration is entered.
func (s *BaseLanguageListener) EnterExplicitSliceDeclaration(ctx *ExplicitSliceDeclarationContext) {}

// ExitExplicitSliceDeclaration is called when production ExplicitSliceDeclaration is exited.
func (s *BaseLanguageListener) ExitExplicitSliceDeclaration(ctx *ExplicitSliceDeclarationContext) {}

// EnterExplicitMatrixDeclaration is called when production ExplicitMatrixDeclaration is entered.
func (s *BaseLanguageListener) EnterExplicitMatrixDeclaration(ctx *ExplicitMatrixDeclarationContext) {
}

// ExitExplicitMatrixDeclaration is called when production ExplicitMatrixDeclaration is exited.
func (s *BaseLanguageListener) ExitExplicitMatrixDeclaration(ctx *ExplicitMatrixDeclarationContext) {}

// EnterExplicitStructDeclaration is called when production ExplicitStructDeclaration is entered.
func (s *BaseLanguageListener) EnterExplicitStructDeclaration(ctx *ExplicitStructDeclarationContext) {
}

// ExitExplicitStructDeclaration is called when production ExplicitStructDeclaration is exited.
func (s *BaseLanguageListener) ExitExplicitStructDeclaration(ctx *ExplicitStructDeclarationContext) {}

// EnterImplicitStructDeclaration is called when production ImplicitStructDeclaration is entered.
func (s *BaseLanguageListener) EnterImplicitStructDeclaration(ctx *ImplicitStructDeclarationContext) {
}

// ExitImplicitStructDeclaration is called when production ImplicitStructDeclaration is exited.
func (s *BaseLanguageListener) ExitImplicitStructDeclaration(ctx *ImplicitStructDeclarationContext) {}

// EnterImplicitSliceDeclaration is called when production ImplicitSliceDeclaration is entered.
func (s *BaseLanguageListener) EnterImplicitSliceDeclaration(ctx *ImplicitSliceDeclarationContext) {}

// ExitImplicitSliceDeclaration is called when production ImplicitSliceDeclaration is exited.
func (s *BaseLanguageListener) ExitImplicitSliceDeclaration(ctx *ImplicitSliceDeclarationContext) {}

// EnterImplicitMatrixDeclaration is called when production ImplicitMatrixDeclaration is entered.
func (s *BaseLanguageListener) EnterImplicitMatrixDeclaration(ctx *ImplicitMatrixDeclarationContext) {
}

// ExitImplicitMatrixDeclaration is called when production ImplicitMatrixDeclaration is exited.
func (s *BaseLanguageListener) ExitImplicitMatrixDeclaration(ctx *ImplicitMatrixDeclarationContext) {}

// EnterImplicitDeclaration is called when production ImplicitDeclaration is entered.
func (s *BaseLanguageListener) EnterImplicitDeclaration(ctx *ImplicitDeclarationContext) {}

// ExitImplicitDeclaration is called when production ImplicitDeclaration is exited.
func (s *BaseLanguageListener) ExitImplicitDeclaration(ctx *ImplicitDeclarationContext) {}

// EnterSliceElements is called when production sliceElements is entered.
func (s *BaseLanguageListener) EnterSliceElements(ctx *SliceElementsContext) {}

// ExitSliceElements is called when production sliceElements is exited.
func (s *BaseLanguageListener) ExitSliceElements(ctx *SliceElementsContext) {}

// EnterMatrixElements is called when production matrixElements is entered.
func (s *BaseLanguageListener) EnterMatrixElements(ctx *MatrixElementsContext) {}

// ExitMatrixElements is called when production matrixElements is exited.
func (s *BaseLanguageListener) ExitMatrixElements(ctx *MatrixElementsContext) {}

// EnterStructInitialization is called when production structInitialization is entered.
func (s *BaseLanguageListener) EnterStructInitialization(ctx *StructInitializationContext) {}

// ExitStructInitialization is called when production structInitialization is exited.
func (s *BaseLanguageListener) ExitStructInitialization(ctx *StructInitializationContext) {}

// EnterStructFieldInit is called when production structFieldInit is entered.
func (s *BaseLanguageListener) EnterStructFieldInit(ctx *StructFieldInitContext) {}

// ExitStructFieldInit is called when production structFieldInit is exited.
func (s *BaseLanguageListener) ExitStructFieldInit(ctx *StructFieldInitContext) {}

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *BaseLanguageListener) EnterStructDeclaration(ctx *StructDeclarationContext) {}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *BaseLanguageListener) ExitStructDeclaration(ctx *StructDeclarationContext) {}

// EnterStructAtribute is called when production structAtribute is entered.
func (s *BaseLanguageListener) EnterStructAtribute(ctx *StructAtributeContext) {}

// ExitStructAtribute is called when production structAtribute is exited.
func (s *BaseLanguageListener) ExitStructAtribute(ctx *StructAtributeContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseLanguageListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseLanguageListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseLanguageListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseLanguageListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseLanguageListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseLanguageListener) ExitParameter(ctx *ParameterContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseLanguageListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseLanguageListener) ExitArgumentList(ctx *ArgumentListContext) {}
