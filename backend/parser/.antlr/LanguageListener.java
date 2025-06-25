// Generated from /home/xavi-13/Escritorio/OLC2_Proyecto2_G37/backend/parser/Language.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link LanguageParser}.
 */
public interface LanguageListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link LanguageParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(LanguageParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(LanguageParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(LanguageParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(LanguageParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#nonDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterNonDeclaration(LanguageParser.NonDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#nonDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitNonDeclaration(LanguageParser.NonDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#blockStatement}.
	 * @param ctx the parse tree
	 */
	void enterBlockStatement(LanguageParser.BlockStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#blockStatement}.
	 * @param ctx the parse tree
	 */
	void exitBlockStatement(LanguageParser.BlockStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceAssignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterSliceAssignment(LanguageParser.SliceAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceAssignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitSliceAssignment(LanguageParser.SliceAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Or}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterOr(LanguageParser.OrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Or}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitOr(LanguageParser.OrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceLiteral}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterSliceLiteral(LanguageParser.SliceLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceLiteral}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitSliceLiteral(LanguageParser.SliceLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParseFloatExpr}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterParseFloatExpr(LanguageParser.ParseFloatExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParseFloatExpr}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitParseFloatExpr(LanguageParser.ParseFloatExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterParens(LanguageParser.ParensContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitParens(LanguageParser.ParensContext ctx);
	/**
	 * Enter a parse tree produced by the {@code String}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterString(LanguageParser.StringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code String}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitString(LanguageParser.StringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructAssignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterStructAssignment(LanguageParser.StructAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructAssignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitStructAssignment(LanguageParser.StructAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MatrixAccess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterMatrixAccess(LanguageParser.MatrixAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MatrixAccess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitMatrixAccess(LanguageParser.MatrixAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Assignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(LanguageParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Assignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(LanguageParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructAccess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterStructAccess(LanguageParser.StructAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructAccess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitStructAccess(LanguageParser.StructAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSubOperator}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterAddSubOperator(LanguageParser.AddSubOperatorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSubOperator}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitAddSubOperator(LanguageParser.AddSubOperatorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDivMod}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterMulDivMod(LanguageParser.MulDivModContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDivMod}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitMulDivMod(LanguageParser.MulDivModContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AtoiExpr}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterAtoiExpr(LanguageParser.AtoiExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AtoiExpr}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitAtoiExpr(LanguageParser.AtoiExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Identifier}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(LanguageParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Identifier}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(LanguageParser.IdentifierContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Equal}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterEqual(LanguageParser.EqualContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Equal}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitEqual(LanguageParser.EqualContext ctx);
	/**
	 * Enter a parse tree produced by the {@code GreaterLess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterGreaterLess(LanguageParser.GreaterLessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code GreaterLess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitGreaterLess(LanguageParser.GreaterLessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FunctionCall}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(LanguageParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FunctionCall}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(LanguageParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Boolean}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterBoolean(LanguageParser.BooleanContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Boolean}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitBoolean(LanguageParser.BooleanContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceAccess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterSliceAccess(LanguageParser.SliceAccessContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceAccess}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitSliceAccess(LanguageParser.SliceAccessContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(LanguageParser.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(LanguageParser.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MatrixAssignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterMatrixAssignment(LanguageParser.MatrixAssignmentContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MatrixAssignment}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitMatrixAssignment(LanguageParser.MatrixAssignmentContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructInstantiation}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterStructInstantiation(LanguageParser.StructInstantiationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructInstantiation}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitStructInstantiation(LanguageParser.StructInstantiationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Integer}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterInteger(LanguageParser.IntegerContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Integer}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitInteger(LanguageParser.IntegerContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Nil}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterNil(LanguageParser.NilContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Nil}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitNil(LanguageParser.NilContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TypeOfExpr}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterTypeOfExpr(LanguageParser.TypeOfExprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TypeOfExpr}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitTypeOfExpr(LanguageParser.TypeOfExprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Float}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterFloat(LanguageParser.FloatContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Float}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitFloat(LanguageParser.FloatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code And}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterAnd(LanguageParser.AndContext ctx);
	/**
	 * Exit a parse tree produced by the {@code And}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitAnd(LanguageParser.AndContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Negate}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterNegate(LanguageParser.NegateContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Negate}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitNegate(LanguageParser.NegateContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IncrementDecrement}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void enterIncrementDecrement(LanguageParser.IncrementDecrementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IncrementDecrement}
	 * labeled alternative in {@link LanguageParser#expressionStatement}.
	 * @param ctx the parse tree
	 */
	void exitIncrementDecrement(LanguageParser.IncrementDecrementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void enterIfStatement(LanguageParser.IfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#ifStatement}.
	 * @param ctx the parse tree
	 */
	void exitIfStatement(LanguageParser.IfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStatement(LanguageParser.SwitchStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#switchStatement}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStatement(LanguageParser.SwitchStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#switchCase}.
	 * @param ctx the parse tree
	 */
	void enterSwitchCase(LanguageParser.SwitchCaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#switchCase}.
	 * @param ctx the parse tree
	 */
	void exitSwitchCase(LanguageParser.SwitchCaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#defaultCase}.
	 * @param ctx the parse tree
	 */
	void enterDefaultCase(LanguageParser.DefaultCaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#defaultCase}.
	 * @param ctx the parse tree
	 */
	void exitDefaultCase(LanguageParser.DefaultCaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForConditional}
	 * labeled alternative in {@link LanguageParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForConditional(LanguageParser.ForConditionalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForConditional}
	 * labeled alternative in {@link LanguageParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForConditional(LanguageParser.ForConditionalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForSimple}
	 * labeled alternative in {@link LanguageParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForSimple(LanguageParser.ForSimpleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForSimple}
	 * labeled alternative in {@link LanguageParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForSimple(LanguageParser.ForSimpleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForSlice}
	 * labeled alternative in {@link LanguageParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void enterForSlice(LanguageParser.ForSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForSlice}
	 * labeled alternative in {@link LanguageParser#forStatement}.
	 * @param ctx the parse tree
	 */
	void exitForSlice(LanguageParser.ForSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BreakStatement}
	 * labeled alternative in {@link LanguageParser#transferenceStatement}.
	 * @param ctx the parse tree
	 */
	void enterBreakStatement(LanguageParser.BreakStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BreakStatement}
	 * labeled alternative in {@link LanguageParser#transferenceStatement}.
	 * @param ctx the parse tree
	 */
	void exitBreakStatement(LanguageParser.BreakStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ContinueStatement}
	 * labeled alternative in {@link LanguageParser#transferenceStatement}.
	 * @param ctx the parse tree
	 */
	void enterContinueStatement(LanguageParser.ContinueStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ContinueStatement}
	 * labeled alternative in {@link LanguageParser#transferenceStatement}.
	 * @param ctx the parse tree
	 */
	void exitContinueStatement(LanguageParser.ContinueStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ReturnStatement}
	 * labeled alternative in {@link LanguageParser#transferenceStatement}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(LanguageParser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ReturnStatement}
	 * labeled alternative in {@link LanguageParser#transferenceStatement}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(LanguageParser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#atoiStatement}.
	 * @param ctx the parse tree
	 */
	void enterAtoiStatement(LanguageParser.AtoiStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#atoiStatement}.
	 * @param ctx the parse tree
	 */
	void exitAtoiStatement(LanguageParser.AtoiStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#parseFloatStatement}.
	 * @param ctx the parse tree
	 */
	void enterParseFloatStatement(LanguageParser.ParseFloatStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#parseFloatStatement}.
	 * @param ctx the parse tree
	 */
	void exitParseFloatStatement(LanguageParser.ParseFloatStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#typeOfStatement}.
	 * @param ctx the parse tree
	 */
	void enterTypeOfStatement(LanguageParser.TypeOfStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#typeOfStatement}.
	 * @param ctx the parse tree
	 */
	void exitTypeOfStatement(LanguageParser.TypeOfStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#printStatement}.
	 * @param ctx the parse tree
	 */
	void enterPrintStatement(LanguageParser.PrintStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#printStatement}.
	 * @param ctx the parse tree
	 */
	void exitPrintStatement(LanguageParser.PrintStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExplicitDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterExplicitDeclaration(LanguageParser.ExplicitDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExplicitDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitExplicitDeclaration(LanguageParser.ExplicitDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExplicitSliceDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterExplicitSliceDeclaration(LanguageParser.ExplicitSliceDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExplicitSliceDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitExplicitSliceDeclaration(LanguageParser.ExplicitSliceDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExplicitMatrixDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterExplicitMatrixDeclaration(LanguageParser.ExplicitMatrixDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExplicitMatrixDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitExplicitMatrixDeclaration(LanguageParser.ExplicitMatrixDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExplicitStructDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterExplicitStructDeclaration(LanguageParser.ExplicitStructDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExplicitStructDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitExplicitStructDeclaration(LanguageParser.ExplicitStructDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ImplicitStructDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterImplicitStructDeclaration(LanguageParser.ImplicitStructDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ImplicitStructDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitImplicitStructDeclaration(LanguageParser.ImplicitStructDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ImplicitSliceDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterImplicitSliceDeclaration(LanguageParser.ImplicitSliceDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ImplicitSliceDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitImplicitSliceDeclaration(LanguageParser.ImplicitSliceDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ImplicitMatrixDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterImplicitMatrixDeclaration(LanguageParser.ImplicitMatrixDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ImplicitMatrixDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitImplicitMatrixDeclaration(LanguageParser.ImplicitMatrixDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ImplicitDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterImplicitDeclaration(LanguageParser.ImplicitDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ImplicitDeclaration}
	 * labeled alternative in {@link LanguageParser#variableDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitImplicitDeclaration(LanguageParser.ImplicitDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#sliceElements}.
	 * @param ctx the parse tree
	 */
	void enterSliceElements(LanguageParser.SliceElementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#sliceElements}.
	 * @param ctx the parse tree
	 */
	void exitSliceElements(LanguageParser.SliceElementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#matrixElements}.
	 * @param ctx the parse tree
	 */
	void enterMatrixElements(LanguageParser.MatrixElementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#matrixElements}.
	 * @param ctx the parse tree
	 */
	void exitMatrixElements(LanguageParser.MatrixElementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#structInitialization}.
	 * @param ctx the parse tree
	 */
	void enterStructInitialization(LanguageParser.StructInitializationContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#structInitialization}.
	 * @param ctx the parse tree
	 */
	void exitStructInitialization(LanguageParser.StructInitializationContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#structFieldInit}.
	 * @param ctx the parse tree
	 */
	void enterStructFieldInit(LanguageParser.StructFieldInitContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#structFieldInit}.
	 * @param ctx the parse tree
	 */
	void exitStructFieldInit(LanguageParser.StructFieldInitContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#structDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterStructDeclaration(LanguageParser.StructDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#structDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitStructDeclaration(LanguageParser.StructDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#structAtribute}.
	 * @param ctx the parse tree
	 */
	void enterStructAtribute(LanguageParser.StructAtributeContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#structAtribute}.
	 * @param ctx the parse tree
	 */
	void exitStructAtribute(LanguageParser.StructAtributeContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDeclaration(LanguageParser.FunctionDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#functionDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDeclaration(LanguageParser.FunctionDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#parameterList}.
	 * @param ctx the parse tree
	 */
	void enterParameterList(LanguageParser.ParameterListContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#parameterList}.
	 * @param ctx the parse tree
	 */
	void exitParameterList(LanguageParser.ParameterListContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#parameter}.
	 * @param ctx the parse tree
	 */
	void enterParameter(LanguageParser.ParameterContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#parameter}.
	 * @param ctx the parse tree
	 */
	void exitParameter(LanguageParser.ParameterContext ctx);
	/**
	 * Enter a parse tree produced by {@link LanguageParser#argumentList}.
	 * @param ctx the parse tree
	 */
	void enterArgumentList(LanguageParser.ArgumentListContext ctx);
	/**
	 * Exit a parse tree produced by {@link LanguageParser#argumentList}.
	 * @param ctx the parse tree
	 */
	void exitArgumentList(LanguageParser.ArgumentListContext ctx);
}