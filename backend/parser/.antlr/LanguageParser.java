// Generated from /home/hp/Documentos/Proyecto2/backend/parser/Language.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class LanguageParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, T__35=36, T__36=37, T__37=38, 
		T__38=39, T__39=40, T__40=41, T__41=42, T__42=43, T__43=44, T__44=45, 
		T__45=46, T__46=47, T__47=48, T__48=49, T__49=50, T__50=51, TYPE=52, INTEGER=53, 
		BOOLEAN=54, FLOAT=55, STRING=56, ID=57, ONELINECOMMENT=58, MULTILINECOMMENT=59, 
		WS=60;
	public static final int
		RULE_program = 0, RULE_statement = 1, RULE_nonDeclaration = 2, RULE_blockStatement = 3, 
		RULE_expressionStatement = 4, RULE_ifStatement = 5, RULE_switchStatement = 6, 
		RULE_switchCase = 7, RULE_defaultCase = 8, RULE_forStatement = 9, RULE_transferenceStatement = 10, 
		RULE_atoiStatement = 11, RULE_parseFloatStatement = 12, RULE_typeOfStatement = 13, 
		RULE_printStatement = 14, RULE_variableDeclaration = 15, RULE_sliceElements = 16, 
		RULE_matrixElements = 17, RULE_structInitialization = 18, RULE_structFieldInit = 19, 
		RULE_structDeclaration = 20, RULE_structAtribute = 21, RULE_functionDeclaration = 22, 
		RULE_parameterList = 23, RULE_parameter = 24, RULE_argumentList = 25;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "statement", "nonDeclaration", "blockStatement", "expressionStatement", 
			"ifStatement", "switchStatement", "switchCase", "defaultCase", "forStatement", 
			"transferenceStatement", "atoiStatement", "parseFloatStatement", "typeOfStatement", 
			"printStatement", "variableDeclaration", "sliceElements", "matrixElements", 
			"structInitialization", "structFieldInit", "structDeclaration", "structAtribute", 
			"functionDeclaration", "parameterList", "parameter", "argumentList"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'", "'-'", "'!'", "'*'", "'/'", "'%'", "'+'", "'>'", 
			"'<'", "'>='", "'<='", "'=='", "'!='", "'&&'", "'||'", "'.'", "'('", 
			"')'", "'['", "','", "']'", "'='", "'+='", "'-='", "'++'", "'--'", "'nil'", 
			"'if'", "'else'", "'switch'", "'case'", "':'", "'default'", "'for'", 
			"';'", "'range'", "'break'", "'continue'", "'return'", "'atoi'", "'parseFloat'", 
			"'typeOf'", "'print'", "'mut'", "'[]'", "'[][]'", "':='", "'struct'", 
			"'fn'", "'void'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "TYPE", "INTEGER", "BOOLEAN", "FLOAT", "STRING", 
			"ID", "ONELINECOMMENT", "MULTILINECOMMENT", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Language.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public LanguageParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(55);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 280982157937344538L) != 0)) {
				{
				{
				setState(52);
				statement();
				}
				}
				setState(57);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public NonDeclarationContext nonDeclaration() {
			return getRuleContext(NonDeclarationContext.class,0);
		}
		public VariableDeclarationContext variableDeclaration() {
			return getRuleContext(VariableDeclarationContext.class,0);
		}
		public StructDeclarationContext structDeclaration() {
			return getRuleContext(StructDeclarationContext.class,0);
		}
		public FunctionDeclarationContext functionDeclaration() {
			return getRuleContext(FunctionDeclarationContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(62);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(58);
				nonDeclaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(59);
				variableDeclaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(60);
				structDeclaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(61);
				functionDeclaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NonDeclarationContext extends ParserRuleContext {
		public BlockStatementContext blockStatement() {
			return getRuleContext(BlockStatementContext.class,0);
		}
		public PrintStatementContext printStatement() {
			return getRuleContext(PrintStatementContext.class,0);
		}
		public AtoiStatementContext atoiStatement() {
			return getRuleContext(AtoiStatementContext.class,0);
		}
		public ParseFloatStatementContext parseFloatStatement() {
			return getRuleContext(ParseFloatStatementContext.class,0);
		}
		public TypeOfStatementContext typeOfStatement() {
			return getRuleContext(TypeOfStatementContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public SwitchStatementContext switchStatement() {
			return getRuleContext(SwitchStatementContext.class,0);
		}
		public ForStatementContext forStatement() {
			return getRuleContext(ForStatementContext.class,0);
		}
		public TransferenceStatementContext transferenceStatement() {
			return getRuleContext(TransferenceStatementContext.class,0);
		}
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public NonDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nonDeclaration; }
	}

	public final NonDeclarationContext nonDeclaration() throws RecognitionException {
		NonDeclarationContext _localctx = new NonDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_nonDeclaration);
		try {
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(64);
				blockStatement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(65);
				printStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(66);
				atoiStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(67);
				parseFloatStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(68);
				typeOfStatement();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(69);
				ifStatement();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(70);
				switchStatement();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(71);
				forStatement();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(72);
				transferenceStatement();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(73);
				expressionStatement(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockStatementContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockStatement; }
	}

	public final BlockStatementContext blockStatement() throws RecognitionException {
		BlockStatementContext _localctx = new BlockStatementContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_blockStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			match(T__0);
			setState(80);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 280982157937344538L) != 0)) {
				{
				{
				setState(77);
				statement();
				}
				}
				setState(82);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(83);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionStatementContext extends ParserRuleContext {
		public ExpressionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionStatement; }
	 
		public ExpressionStatementContext() { }
		public void copyFrom(ExpressionStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceAssignmentContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public SliceAssignmentContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OrContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public OrContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceLiteralContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public SliceLiteralContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParseFloatExprContext extends ExpressionStatementContext {
		public ParseFloatStatementContext parseFloatStatement() {
			return getRuleContext(ParseFloatStatementContext.class,0);
		}
		public ParseFloatExprContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParensContext extends ExpressionStatementContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public ParensContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ExpressionStatementContext {
		public TerminalNode STRING() { return getToken(LanguageParser.STRING, 0); }
		public StringContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAssignmentContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public StructAssignmentContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MatrixAccessContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public MatrixAccessContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public AssignmentContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAccessContext extends ExpressionStatementContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public StructAccessContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubOperatorContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public AddSubOperatorContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MulDivModContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public MulDivModContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AtoiExprContext extends ExpressionStatementContext {
		public AtoiStatementContext atoiStatement() {
			return getRuleContext(AtoiStatementContext.class,0);
		}
		public AtoiExprContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public IdentifierContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EqualContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public EqualContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GreaterLessContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public GreaterLessContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public ArgumentListContext argumentList() {
			return getRuleContext(ArgumentListContext.class,0);
		}
		public FunctionCallContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BooleanContext extends ExpressionStatementContext {
		public TerminalNode BOOLEAN() { return getToken(LanguageParser.BOOLEAN, 0); }
		public BooleanContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceAccessContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public SliceAccessContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public AddSubContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MatrixAssignmentContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public MatrixAssignmentContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInstantiationContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public StructInitializationContext structInitialization() {
			return getRuleContext(StructInitializationContext.class,0);
		}
		public StructInstantiationContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntegerContext extends ExpressionStatementContext {
		public TerminalNode INTEGER() { return getToken(LanguageParser.INTEGER, 0); }
		public IntegerContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilContext extends ExpressionStatementContext {
		public NilContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeOfExprContext extends ExpressionStatementContext {
		public TypeOfStatementContext typeOfStatement() {
			return getRuleContext(TypeOfStatementContext.class,0);
		}
		public TypeOfExprContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatContext extends ExpressionStatementContext {
		public TerminalNode FLOAT() { return getToken(LanguageParser.FLOAT, 0); }
		public FloatContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AndContext extends ExpressionStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public AndContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NegateContext extends ExpressionStatementContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public NegateContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncrementDecrementContext extends ExpressionStatementContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public IncrementDecrementContext(ExpressionStatementContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionStatementContext expressionStatement() throws RecognitionException {
		return expressionStatement(0);
	}

	private ExpressionStatementContext expressionStatement(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionStatementContext _localctx = new ExpressionStatementContext(_ctx, _parentState);
		ExpressionStatementContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_expressionStatement, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				_localctx = new NegateContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(86);
				_la = _input.LA(1);
				if ( !(_la==T__2 || _la==T__3) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(87);
				expressionStatement(29);
				}
				break;
			case 2:
				{
				_localctx = new FunctionCallContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(88);
				match(ID);
				setState(89);
				match(T__17);
				setState(91);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 279238570329505816L) != 0)) {
					{
					setState(90);
					argumentList();
					}
				}

				setState(93);
				match(T__18);
				}
				break;
			case 3:
				{
				_localctx = new SliceLiteralContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(94);
				match(T__19);
				setState(103);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 279238570329505816L) != 0)) {
					{
					setState(95);
					expressionStatement(0);
					setState(100);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==T__20) {
						{
						{
						setState(96);
						match(T__20);
						setState(97);
						expressionStatement(0);
						}
						}
						setState(102);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(105);
				match(T__21);
				}
				break;
			case 4:
				{
				_localctx = new MatrixAssignmentContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(106);
				match(ID);
				setState(107);
				match(T__19);
				setState(108);
				expressionStatement(0);
				setState(109);
				match(T__21);
				setState(110);
				match(T__19);
				setState(111);
				expressionStatement(0);
				setState(112);
				match(T__21);
				setState(113);
				match(T__22);
				setState(114);
				expressionStatement(19);
				}
				break;
			case 5:
				{
				_localctx = new SliceAssignmentContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(116);
				match(ID);
				setState(117);
				match(T__19);
				setState(118);
				expressionStatement(0);
				setState(119);
				match(T__21);
				setState(120);
				match(T__22);
				setState(121);
				expressionStatement(18);
				}
				break;
			case 6:
				{
				_localctx = new AddSubOperatorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(123);
				match(ID);
				setState(124);
				_la = _input.LA(1);
				if ( !(_la==T__23 || _la==T__24) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(125);
				expressionStatement(15);
				}
				break;
			case 7:
				{
				_localctx = new IncrementDecrementContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(126);
				match(ID);
				setState(127);
				_la = _input.LA(1);
				if ( !(_la==T__25 || _la==T__26) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 8:
				{
				_localctx = new AtoiExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(128);
				atoiStatement();
				}
				break;
			case 9:
				{
				_localctx = new ParseFloatExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(129);
				parseFloatStatement();
				}
				break;
			case 10:
				{
				_localctx = new TypeOfExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(130);
				typeOfStatement();
				}
				break;
			case 11:
				{
				_localctx = new MatrixAccessContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(131);
				match(ID);
				setState(132);
				match(T__19);
				setState(133);
				expressionStatement(0);
				setState(134);
				match(T__21);
				setState(135);
				match(T__19);
				setState(136);
				expressionStatement(0);
				setState(137);
				match(T__21);
				}
				break;
			case 12:
				{
				_localctx = new SliceAccessContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(139);
				match(ID);
				setState(140);
				match(T__19);
				setState(141);
				expressionStatement(0);
				setState(142);
				match(T__21);
				}
				break;
			case 13:
				{
				_localctx = new StructInstantiationContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(144);
				match(ID);
				setState(145);
				match(T__0);
				setState(147);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(146);
					structInitialization();
					}
				}

				setState(149);
				match(T__1);
				}
				break;
			case 14:
				{
				_localctx = new IntegerContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(150);
				match(INTEGER);
				}
				break;
			case 15:
				{
				_localctx = new BooleanContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(151);
				match(BOOLEAN);
				}
				break;
			case 16:
				{
				_localctx = new FloatContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(152);
				match(FLOAT);
				}
				break;
			case 17:
				{
				_localctx = new StringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(153);
				match(STRING);
				}
				break;
			case 18:
				{
				_localctx = new IdentifierContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(154);
				match(ID);
				}
				break;
			case 19:
				{
				_localctx = new NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(155);
				match(T__27);
				}
				break;
			case 20:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(164);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__17:
					{
					{
					setState(156);
					match(T__17);
					setState(157);
					expressionStatement(0);
					setState(158);
					match(T__18);
					}
					}
					break;
				case T__19:
					{
					{
					setState(160);
					match(T__19);
					setState(161);
					expressionStatement(0);
					setState(162);
					match(T__21);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(199);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(197);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new MulDivModContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(168);
						if (!(precpred(_ctx, 28))) throw new FailedPredicateException(this, "precpred(_ctx, 28)");
						setState(169);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 224L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(170);
						expressionStatement(29);
						}
						break;
					case 2:
						{
						_localctx = new AddSubContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(171);
						if (!(precpred(_ctx, 27))) throw new FailedPredicateException(this, "precpred(_ctx, 27)");
						setState(172);
						_la = _input.LA(1);
						if ( !(_la==T__2 || _la==T__7) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(173);
						expressionStatement(28);
						}
						break;
					case 3:
						{
						_localctx = new GreaterLessContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(174);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(175);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 7680L) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(176);
						expressionStatement(27);
						}
						break;
					case 4:
						{
						_localctx = new EqualContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(177);
						if (!(precpred(_ctx, 25))) throw new FailedPredicateException(this, "precpred(_ctx, 25)");
						setState(178);
						_la = _input.LA(1);
						if ( !(_la==T__12 || _la==T__13) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(179);
						expressionStatement(26);
						}
						break;
					case 5:
						{
						_localctx = new AndContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(180);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(181);
						match(T__14);
						setState(182);
						expressionStatement(25);
						}
						break;
					case 6:
						{
						_localctx = new OrContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(183);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(184);
						match(T__15);
						setState(185);
						expressionStatement(24);
						}
						break;
					case 7:
						{
						_localctx = new StructAssignmentContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(186);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(187);
						match(T__16);
						setState(188);
						match(ID);
						setState(189);
						match(T__22);
						setState(190);
						expressionStatement(18);
						}
						break;
					case 8:
						{
						_localctx = new AssignmentContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(191);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(192);
						match(T__22);
						setState(193);
						expressionStatement(17);
						}
						break;
					case 9:
						{
						_localctx = new StructAccessContext(new ExpressionStatementContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expressionStatement);
						setState(194);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(195);
						match(T__16);
						setState(196);
						match(ID);
						}
						break;
					}
					} 
				}
				setState(201);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends ParserRuleContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public List<BlockStatementContext> blockStatement() {
			return getRuleContexts(BlockStatementContext.class);
		}
		public BlockStatementContext blockStatement(int i) {
			return getRuleContext(BlockStatementContext.class,i);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_ifStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			match(T__28);
			setState(203);
			expressionStatement(0);
			setState(204);
			blockStatement();
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__29) {
				{
				setState(205);
				match(T__29);
				setState(208);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__28:
					{
					setState(206);
					ifStatement();
					}
					break;
				case T__0:
					{
					setState(207);
					blockStatement();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStatementContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public List<SwitchCaseContext> switchCase() {
			return getRuleContexts(SwitchCaseContext.class);
		}
		public SwitchCaseContext switchCase(int i) {
			return getRuleContext(SwitchCaseContext.class,i);
		}
		public DefaultCaseContext defaultCase() {
			return getRuleContext(DefaultCaseContext.class,0);
		}
		public SwitchStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchStatement; }
	}

	public final SwitchStatementContext switchStatement() throws RecognitionException {
		SwitchStatementContext _localctx = new SwitchStatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_switchStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(212);
			match(T__30);
			setState(213);
			match(ID);
			setState(214);
			match(T__0);
			setState(218);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__31) {
				{
				{
				setState(215);
				switchCase();
				}
				}
				setState(220);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__33) {
				{
				setState(221);
				defaultCase();
				}
			}

			setState(224);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchCaseContext extends ParserRuleContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public SwitchCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchCase; }
	}

	public final SwitchCaseContext switchCase() throws RecognitionException {
		SwitchCaseContext _localctx = new SwitchCaseContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_switchCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			match(T__31);
			setState(227);
			expressionStatement(0);
			setState(228);
			match(T__32);
			setState(232);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 280982157937344538L) != 0)) {
				{
				{
				setState(229);
				statement();
				}
				}
				setState(234);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefaultCaseContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public DefaultCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultCase; }
	}

	public final DefaultCaseContext defaultCase() throws RecognitionException {
		DefaultCaseContext _localctx = new DefaultCaseContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_defaultCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(T__33);
			setState(236);
			match(T__32);
			setState(240);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 280982157937344538L) != 0)) {
				{
				{
				setState(237);
				statement();
				}
				}
				setState(242);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForStatementContext extends ParserRuleContext {
		public ForStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forStatement; }
	 
		public ForStatementContext() { }
		public void copyFrom(ForStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForConditionalContext extends ForStatementContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public BlockStatementContext blockStatement() {
			return getRuleContext(BlockStatementContext.class,0);
		}
		public ForConditionalContext(ForStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForSimpleContext extends ForStatementContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public BlockStatementContext blockStatement() {
			return getRuleContext(BlockStatementContext.class,0);
		}
		public VariableDeclarationContext variableDeclaration() {
			return getRuleContext(VariableDeclarationContext.class,0);
		}
		public ForSimpleContext(ForStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForSliceContext extends ForStatementContext {
		public List<TerminalNode> ID() { return getTokens(LanguageParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(LanguageParser.ID, i);
		}
		public BlockStatementContext blockStatement() {
			return getRuleContext(BlockStatementContext.class,0);
		}
		public ForSliceContext(ForStatementContext ctx) { copyFrom(ctx); }
	}

	public final ForStatementContext forStatement() throws RecognitionException {
		ForStatementContext _localctx = new ForStatementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_forStatement);
		try {
			setState(266);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				_localctx = new ForConditionalContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(243);
				match(T__34);
				setState(244);
				expressionStatement(0);
				setState(245);
				blockStatement();
				}
				break;
			case 2:
				_localctx = new ForSimpleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(247);
				match(T__34);
				setState(250);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
				case 1:
					{
					setState(248);
					variableDeclaration();
					}
					break;
				case 2:
					{
					setState(249);
					expressionStatement(0);
					}
					break;
				}
				setState(252);
				match(T__35);
				setState(253);
				expressionStatement(0);
				setState(254);
				match(T__35);
				setState(255);
				expressionStatement(0);
				setState(256);
				blockStatement();
				}
				break;
			case 3:
				_localctx = new ForSliceContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(258);
				match(T__34);
				setState(259);
				match(ID);
				setState(260);
				match(T__20);
				setState(261);
				match(ID);
				setState(262);
				match(T__22);
				setState(263);
				match(T__36);
				setState(264);
				match(ID);
				setState(265);
				blockStatement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TransferenceStatementContext extends ParserRuleContext {
		public TransferenceStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_transferenceStatement; }
	 
		public TransferenceStatementContext() { }
		public void copyFrom(TransferenceStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStatementContext extends TransferenceStatementContext {
		public BreakStatementContext(TransferenceStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStatementContext extends TransferenceStatementContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public ReturnStatementContext(TransferenceStatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStatementContext extends TransferenceStatementContext {
		public ContinueStatementContext(TransferenceStatementContext ctx) { copyFrom(ctx); }
	}

	public final TransferenceStatementContext transferenceStatement() throws RecognitionException {
		TransferenceStatementContext _localctx = new TransferenceStatementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_transferenceStatement);
		try {
			setState(274);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__37:
				_localctx = new BreakStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(268);
				match(T__37);
				}
				break;
			case T__38:
				_localctx = new ContinueStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(269);
				match(T__38);
				}
				break;
			case T__39:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(270);
				match(T__39);
				setState(272);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
				case 1:
					{
					setState(271);
					expressionStatement(0);
					}
					break;
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AtoiStatementContext extends ParserRuleContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public AtoiStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atoiStatement; }
	}

	public final AtoiStatementContext atoiStatement() throws RecognitionException {
		AtoiStatementContext _localctx = new AtoiStatementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_atoiStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			match(T__40);
			setState(277);
			match(T__17);
			setState(278);
			expressionStatement(0);
			setState(279);
			match(T__18);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParseFloatStatementContext extends ParserRuleContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public ParseFloatStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parseFloatStatement; }
	}

	public final ParseFloatStatementContext parseFloatStatement() throws RecognitionException {
		ParseFloatStatementContext _localctx = new ParseFloatStatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_parseFloatStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(281);
			match(T__41);
			setState(282);
			match(T__17);
			setState(283);
			expressionStatement(0);
			setState(284);
			match(T__18);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeOfStatementContext extends ParserRuleContext {
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public TypeOfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeOfStatement; }
	}

	public final TypeOfStatementContext typeOfStatement() throws RecognitionException {
		TypeOfStatementContext _localctx = new TypeOfStatementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_typeOfStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			match(T__42);
			setState(287);
			match(T__17);
			setState(288);
			expressionStatement(0);
			setState(289);
			match(T__18);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintStatementContext extends ParserRuleContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public PrintStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printStatement; }
	}

	public final PrintStatementContext printStatement() throws RecognitionException {
		PrintStatementContext _localctx = new PrintStatementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_printStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(291);
			match(T__43);
			setState(292);
			match(T__17);
			setState(301);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 279238570329505816L) != 0)) {
				{
				setState(293);
				expressionStatement(0);
				setState(298);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__20) {
					{
					{
					setState(294);
					match(T__20);
					setState(295);
					expressionStatement(0);
					}
					}
					setState(300);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(303);
			match(T__18);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclarationContext extends ParserRuleContext {
		public VariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDeclaration; }
	 
		public VariableDeclarationContext() { }
		public void copyFrom(VariableDeclarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExplicitSliceDeclarationContext extends VariableDeclarationContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public SliceElementsContext sliceElements() {
			return getRuleContext(SliceElementsContext.class,0);
		}
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public ExplicitSliceDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ImplicitStructDeclarationContext extends VariableDeclarationContext {
		public List<TerminalNode> ID() { return getTokens(LanguageParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(LanguageParser.ID, i);
		}
		public StructInitializationContext structInitialization() {
			return getRuleContext(StructInitializationContext.class,0);
		}
		public ImplicitStructDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExplicitStructDeclarationContext extends VariableDeclarationContext {
		public List<TerminalNode> ID() { return getTokens(LanguageParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(LanguageParser.ID, i);
		}
		public StructInitializationContext structInitialization() {
			return getRuleContext(StructInitializationContext.class,0);
		}
		public ExplicitStructDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExplicitMatrixDeclarationContext extends VariableDeclarationContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public MatrixElementsContext matrixElements() {
			return getRuleContext(MatrixElementsContext.class,0);
		}
		public ExplicitMatrixDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ImplicitMatrixDeclarationContext extends VariableDeclarationContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public MatrixElementsContext matrixElements() {
			return getRuleContext(MatrixElementsContext.class,0);
		}
		public ImplicitMatrixDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ImplicitDeclarationContext extends VariableDeclarationContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public ImplicitDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExplicitDeclarationContext extends VariableDeclarationContext {
		public List<TerminalNode> ID() { return getTokens(LanguageParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(LanguageParser.ID, i);
		}
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public ExplicitDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ImplicitSliceDeclarationContext extends VariableDeclarationContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public SliceElementsContext sliceElements() {
			return getRuleContext(SliceElementsContext.class,0);
		}
		public ImplicitSliceDeclarationContext(VariableDeclarationContext ctx) { copyFrom(ctx); }
	}

	public final VariableDeclarationContext variableDeclaration() throws RecognitionException {
		VariableDeclarationContext _localctx = new VariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_variableDeclaration);
		int _la;
		try {
			setState(383);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				_localctx = new ExplicitDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(306);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(305);
					match(T__44);
					}
				}

				setState(308);
				match(ID);
				setState(309);
				_la = _input.LA(1);
				if ( !(_la==TYPE || _la==ID) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(312);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__22) {
					{
					setState(310);
					match(T__22);
					setState(311);
					expressionStatement(0);
					}
				}

				}
				break;
			case 2:
				_localctx = new ExplicitSliceDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(315);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(314);
					match(T__44);
					}
				}

				setState(317);
				match(ID);
				setState(318);
				match(T__45);
				setState(319);
				match(TYPE);
				setState(325);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__22) {
					{
					setState(320);
					match(T__22);
					setState(323);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case T__0:
						{
						setState(321);
						sliceElements();
						}
						break;
					case T__2:
					case T__3:
					case T__17:
					case T__19:
					case T__27:
					case T__40:
					case T__41:
					case T__42:
					case INTEGER:
					case BOOLEAN:
					case FLOAT:
					case STRING:
					case ID:
						{
						setState(322);
						expressionStatement(0);
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
				}

				}
				break;
			case 3:
				_localctx = new ExplicitMatrixDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(328);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(327);
					match(T__44);
					}
				}

				setState(330);
				match(ID);
				setState(331);
				match(T__46);
				setState(332);
				match(TYPE);
				setState(335);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__22) {
					{
					setState(333);
					match(T__22);
					setState(334);
					matrixElements();
					}
				}

				}
				break;
			case 4:
				_localctx = new ExplicitStructDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(338);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(337);
					match(T__44);
					}
				}

				setState(340);
				match(ID);
				setState(341);
				match(ID);
				setState(348);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__22) {
					{
					setState(342);
					match(T__22);
					setState(343);
					match(T__0);
					setState(345);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if (_la==ID) {
						{
						setState(344);
						structInitialization();
						}
					}

					setState(347);
					match(T__1);
					}
				}

				}
				break;
			case 5:
				_localctx = new ImplicitStructDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(351);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(350);
					match(T__44);
					}
				}

				setState(353);
				match(ID);
				setState(354);
				match(T__47);
				setState(355);
				match(ID);
				setState(356);
				match(T__0);
				setState(358);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(357);
					structInitialization();
					}
				}

				setState(360);
				match(T__1);
				}
				break;
			case 6:
				_localctx = new ImplicitSliceDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(362);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(361);
					match(T__44);
					}
				}

				setState(364);
				match(ID);
				setState(365);
				match(T__47);
				setState(366);
				match(T__45);
				setState(367);
				match(TYPE);
				setState(368);
				sliceElements();
				}
				break;
			case 7:
				_localctx = new ImplicitMatrixDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(370);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(369);
					match(T__44);
					}
				}

				setState(372);
				match(ID);
				setState(373);
				match(T__47);
				setState(374);
				match(T__46);
				setState(375);
				match(TYPE);
				setState(376);
				matrixElements();
				}
				break;
			case 8:
				_localctx = new ImplicitDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(378);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__44) {
					{
					setState(377);
					match(T__44);
					}
				}

				setState(380);
				match(ID);
				setState(381);
				match(T__47);
				setState(382);
				expressionStatement(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SliceElementsContext extends ParserRuleContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public SliceElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceElements; }
	}

	public final SliceElementsContext sliceElements() throws RecognitionException {
		SliceElementsContext _localctx = new SliceElementsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_sliceElements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(385);
			match(T__0);
			setState(394);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 279238570329505816L) != 0)) {
				{
				setState(386);
				expressionStatement(0);
				setState(391);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__20) {
					{
					{
					setState(387);
					match(T__20);
					setState(388);
					expressionStatement(0);
					}
					}
					setState(393);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(396);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MatrixElementsContext extends ParserRuleContext {
		public List<SliceElementsContext> sliceElements() {
			return getRuleContexts(SliceElementsContext.class);
		}
		public SliceElementsContext sliceElements(int i) {
			return getRuleContext(SliceElementsContext.class,i);
		}
		public MatrixElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrixElements; }
	}

	public final MatrixElementsContext matrixElements() throws RecognitionException {
		MatrixElementsContext _localctx = new MatrixElementsContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_matrixElements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(398);
			match(T__0);
			setState(407);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(399);
				sliceElements();
				setState(404);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__20) {
					{
					{
					setState(400);
					match(T__20);
					setState(401);
					sliceElements();
					}
					}
					setState(406);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(409);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructInitializationContext extends ParserRuleContext {
		public List<StructFieldInitContext> structFieldInit() {
			return getRuleContexts(StructFieldInitContext.class);
		}
		public StructFieldInitContext structFieldInit(int i) {
			return getRuleContext(StructFieldInitContext.class,i);
		}
		public StructInitializationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structInitialization; }
	}

	public final StructInitializationContext structInitialization() throws RecognitionException {
		StructInitializationContext _localctx = new StructInitializationContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_structInitialization);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(411);
			structFieldInit();
			setState(416);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__20) {
				{
				{
				setState(412);
				match(T__20);
				setState(413);
				structFieldInit();
				}
				}
				setState(418);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructFieldInitContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public StructFieldInitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structFieldInit; }
	}

	public final StructFieldInitContext structFieldInit() throws RecognitionException {
		StructFieldInitContext _localctx = new StructFieldInitContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_structFieldInit);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(419);
			match(ID);
			setState(420);
			match(T__32);
			setState(421);
			expressionStatement(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(LanguageParser.ID, 0); }
		public List<StructAtributeContext> structAtribute() {
			return getRuleContexts(StructAtributeContext.class);
		}
		public StructAtributeContext structAtribute(int i) {
			return getRuleContext(StructAtributeContext.class,i);
		}
		public StructDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDeclaration; }
	}

	public final StructDeclarationContext structDeclaration() throws RecognitionException {
		StructDeclarationContext _localctx = new StructDeclarationContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_structDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(423);
			match(T__48);
			setState(424);
			match(ID);
			setState(425);
			match(T__0);
			setState(429);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE || _la==ID) {
				{
				{
				setState(426);
				structAtribute();
				}
				}
				setState(431);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(432);
			match(T__1);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructAtributeContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(LanguageParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(LanguageParser.ID, i);
		}
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public StructAtributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structAtribute; }
	}

	public final StructAtributeContext structAtribute() throws RecognitionException {
		StructAtributeContext _localctx = new StructAtributeContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_structAtribute);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(434);
			_la = _input.LA(1);
			if ( !(_la==TYPE || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(435);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FunctionDeclarationContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(LanguageParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(LanguageParser.ID, i);
		}
		public BlockStatementContext blockStatement() {
			return getRuleContext(BlockStatementContext.class,0);
		}
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public FunctionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDeclaration; }
	}

	public final FunctionDeclarationContext functionDeclaration() throws RecognitionException {
		FunctionDeclarationContext _localctx = new FunctionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_functionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(437);
			match(T__49);
			setState(438);
			match(ID);
			setState(439);
			match(T__17);
			setState(441);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(440);
				parameterList();
				}
			}

			setState(443);
			match(T__18);
			setState(445);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 150870587516911616L) != 0)) {
				{
				setState(444);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 150870587516911616L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(447);
			blockStatement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterListContext extends ParserRuleContext {
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
		public ParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterList; }
	}

	public final ParameterListContext parameterList() throws RecognitionException {
		ParameterListContext _localctx = new ParameterListContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_parameterList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(449);
			parameter();
			setState(454);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__20) {
				{
				{
				setState(450);
				match(T__20);
				setState(451);
				parameter();
				}
				}
				setState(456);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(LanguageParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(LanguageParser.ID, i);
		}
		public TerminalNode TYPE() { return getToken(LanguageParser.TYPE, 0); }
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_parameter);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(457);
			match(ID);
			setState(458);
			_la = _input.LA(1);
			if ( !(_la==TYPE || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentListContext extends ParserRuleContext {
		public List<ExpressionStatementContext> expressionStatement() {
			return getRuleContexts(ExpressionStatementContext.class);
		}
		public ExpressionStatementContext expressionStatement(int i) {
			return getRuleContext(ExpressionStatementContext.class,i);
		}
		public ArgumentListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumentList; }
	}

	public final ArgumentListContext argumentList() throws RecognitionException {
		ArgumentListContext _localctx = new ArgumentListContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_argumentList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(460);
			expressionStatement(0);
			setState(465);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__20) {
				{
				{
				setState(461);
				match(T__20);
				setState(462);
				expressionStatement(0);
				}
				}
				setState(467);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return expressionStatement_sempred((ExpressionStatementContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expressionStatement_sempred(ExpressionStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 28);
		case 1:
			return precpred(_ctx, 27);
		case 2:
			return precpred(_ctx, 26);
		case 3:
			return precpred(_ctx, 25);
		case 4:
			return precpred(_ctx, 24);
		case 5:
			return precpred(_ctx, 23);
		case 6:
			return precpred(_ctx, 17);
		case 7:
			return precpred(_ctx, 16);
		case 8:
			return precpred(_ctx, 22);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001<\u01d5\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0001\u0000\u0005\u00006\b\u0000\n\u0000\f\u0000"+
		"9\t\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001"+
		"?\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"K\b\u0002\u0001\u0003\u0001\u0003\u0005\u0003O\b\u0003\n\u0003\f\u0003"+
		"R\t\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\\\b\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004c\b\u0004"+
		"\n\u0004\f\u0004f\t\u0004\u0003\u0004h\b\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u0094\b\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u00a5\b\u0004\u0003\u0004"+
		"\u00a7\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0005\u0004\u00c6\b\u0004\n\u0004\f\u0004\u00c9\t\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00d1"+
		"\b\u0005\u0003\u0005\u00d3\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0005\u0006\u00d9\b\u0006\n\u0006\f\u0006\u00dc\t\u0006\u0001"+
		"\u0006\u0003\u0006\u00df\b\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u00e7\b\u0007\n\u0007\f\u0007"+
		"\u00ea\t\u0007\u0001\b\u0001\b\u0001\b\u0005\b\u00ef\b\b\n\b\f\b\u00f2"+
		"\t\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u00fb"+
		"\b\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u010b\b\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0003\n\u0111\b\n\u0003\n\u0113\b\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u0129\b\u000e\n\u000e"+
		"\f\u000e\u012c\t\u000e\u0003\u000e\u012e\b\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0003\u000f\u0133\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u0139\b\u000f\u0001\u000f\u0003\u000f\u013c\b"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0003\u000f\u0144\b\u000f\u0003\u000f\u0146\b\u000f\u0001\u000f"+
		"\u0003\u000f\u0149\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u0150\b\u000f\u0001\u000f\u0003\u000f\u0153\b"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003"+
		"\u000f\u015a\b\u000f\u0001\u000f\u0003\u000f\u015d\b\u000f\u0001\u000f"+
		"\u0003\u000f\u0160\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u0167\b\u000f\u0001\u000f\u0001\u000f\u0003\u000f"+
		"\u016b\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u0173\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u017b\b\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0003\u000f\u0180\b\u000f\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0005\u0010\u0186\b\u0010\n\u0010\f\u0010\u0189"+
		"\t\u0010\u0003\u0010\u018b\b\u0010\u0001\u0010\u0001\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0005\u0011\u0193\b\u0011\n\u0011"+
		"\f\u0011\u0196\t\u0011\u0003\u0011\u0198\b\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0005\u0012\u019f\b\u0012\n\u0012"+
		"\f\u0012\u01a2\t\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u01ac\b\u0014"+
		"\n\u0014\f\u0014\u01af\t\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0003"+
		"\u0016\u01ba\b\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u01be\b\u0016"+
		"\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017"+
		"\u01c5\b\u0017\n\u0017\f\u0017\u01c8\t\u0017\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u01d0\b\u0019\n"+
		"\u0019\f\u0019\u01d3\t\u0019\u0001\u0019\u0000\u0001\b\u001a\u0000\u0002"+
		"\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e"+
		" \"$&(*,.02\u0000\t\u0001\u0000\u0003\u0004\u0001\u0000\u0018\u0019\u0001"+
		"\u0000\u001a\u001b\u0001\u0000\u0005\u0007\u0002\u0000\u0003\u0003\b\b"+
		"\u0001\u0000\t\f\u0001\u0000\r\u000e\u0002\u00004499\u0002\u00003499\u0217"+
		"\u00007\u0001\u0000\u0000\u0000\u0002>\u0001\u0000\u0000\u0000\u0004J"+
		"\u0001\u0000\u0000\u0000\u0006L\u0001\u0000\u0000\u0000\b\u00a6\u0001"+
		"\u0000\u0000\u0000\n\u00ca\u0001\u0000\u0000\u0000\f\u00d4\u0001\u0000"+
		"\u0000\u0000\u000e\u00e2\u0001\u0000\u0000\u0000\u0010\u00eb\u0001\u0000"+
		"\u0000\u0000\u0012\u010a\u0001\u0000\u0000\u0000\u0014\u0112\u0001\u0000"+
		"\u0000\u0000\u0016\u0114\u0001\u0000\u0000\u0000\u0018\u0119\u0001\u0000"+
		"\u0000\u0000\u001a\u011e\u0001\u0000\u0000\u0000\u001c\u0123\u0001\u0000"+
		"\u0000\u0000\u001e\u017f\u0001\u0000\u0000\u0000 \u0181\u0001\u0000\u0000"+
		"\u0000\"\u018e\u0001\u0000\u0000\u0000$\u019b\u0001\u0000\u0000\u0000"+
		"&\u01a3\u0001\u0000\u0000\u0000(\u01a7\u0001\u0000\u0000\u0000*\u01b2"+
		"\u0001\u0000\u0000\u0000,\u01b5\u0001\u0000\u0000\u0000.\u01c1\u0001\u0000"+
		"\u0000\u00000\u01c9\u0001\u0000\u0000\u00002\u01cc\u0001\u0000\u0000\u0000"+
		"46\u0003\u0002\u0001\u000054\u0001\u0000\u0000\u000069\u0001\u0000\u0000"+
		"\u000075\u0001\u0000\u0000\u000078\u0001\u0000\u0000\u00008\u0001\u0001"+
		"\u0000\u0000\u000097\u0001\u0000\u0000\u0000:?\u0003\u0004\u0002\u0000"+
		";?\u0003\u001e\u000f\u0000<?\u0003(\u0014\u0000=?\u0003,\u0016\u0000>"+
		":\u0001\u0000\u0000\u0000>;\u0001\u0000\u0000\u0000><\u0001\u0000\u0000"+
		"\u0000>=\u0001\u0000\u0000\u0000?\u0003\u0001\u0000\u0000\u0000@K\u0003"+
		"\u0006\u0003\u0000AK\u0003\u001c\u000e\u0000BK\u0003\u0016\u000b\u0000"+
		"CK\u0003\u0018\f\u0000DK\u0003\u001a\r\u0000EK\u0003\n\u0005\u0000FK\u0003"+
		"\f\u0006\u0000GK\u0003\u0012\t\u0000HK\u0003\u0014\n\u0000IK\u0003\b\u0004"+
		"\u0000J@\u0001\u0000\u0000\u0000JA\u0001\u0000\u0000\u0000JB\u0001\u0000"+
		"\u0000\u0000JC\u0001\u0000\u0000\u0000JD\u0001\u0000\u0000\u0000JE\u0001"+
		"\u0000\u0000\u0000JF\u0001\u0000\u0000\u0000JG\u0001\u0000\u0000\u0000"+
		"JH\u0001\u0000\u0000\u0000JI\u0001\u0000\u0000\u0000K\u0005\u0001\u0000"+
		"\u0000\u0000LP\u0005\u0001\u0000\u0000MO\u0003\u0002\u0001\u0000NM\u0001"+
		"\u0000\u0000\u0000OR\u0001\u0000\u0000\u0000PN\u0001\u0000\u0000\u0000"+
		"PQ\u0001\u0000\u0000\u0000QS\u0001\u0000\u0000\u0000RP\u0001\u0000\u0000"+
		"\u0000ST\u0005\u0002\u0000\u0000T\u0007\u0001\u0000\u0000\u0000UV\u0006"+
		"\u0004\uffff\uffff\u0000VW\u0007\u0000\u0000\u0000W\u00a7\u0003\b\u0004"+
		"\u001dXY\u00059\u0000\u0000Y[\u0005\u0012\u0000\u0000Z\\\u00032\u0019"+
		"\u0000[Z\u0001\u0000\u0000\u0000[\\\u0001\u0000\u0000\u0000\\]\u0001\u0000"+
		"\u0000\u0000]\u00a7\u0005\u0013\u0000\u0000^g\u0005\u0014\u0000\u0000"+
		"_d\u0003\b\u0004\u0000`a\u0005\u0015\u0000\u0000ac\u0003\b\u0004\u0000"+
		"b`\u0001\u0000\u0000\u0000cf\u0001\u0000\u0000\u0000db\u0001\u0000\u0000"+
		"\u0000de\u0001\u0000\u0000\u0000eh\u0001\u0000\u0000\u0000fd\u0001\u0000"+
		"\u0000\u0000g_\u0001\u0000\u0000\u0000gh\u0001\u0000\u0000\u0000hi\u0001"+
		"\u0000\u0000\u0000i\u00a7\u0005\u0016\u0000\u0000jk\u00059\u0000\u0000"+
		"kl\u0005\u0014\u0000\u0000lm\u0003\b\u0004\u0000mn\u0005\u0016\u0000\u0000"+
		"no\u0005\u0014\u0000\u0000op\u0003\b\u0004\u0000pq\u0005\u0016\u0000\u0000"+
		"qr\u0005\u0017\u0000\u0000rs\u0003\b\u0004\u0013s\u00a7\u0001\u0000\u0000"+
		"\u0000tu\u00059\u0000\u0000uv\u0005\u0014\u0000\u0000vw\u0003\b\u0004"+
		"\u0000wx\u0005\u0016\u0000\u0000xy\u0005\u0017\u0000\u0000yz\u0003\b\u0004"+
		"\u0012z\u00a7\u0001\u0000\u0000\u0000{|\u00059\u0000\u0000|}\u0007\u0001"+
		"\u0000\u0000}\u00a7\u0003\b\u0004\u000f~\u007f\u00059\u0000\u0000\u007f"+
		"\u00a7\u0007\u0002\u0000\u0000\u0080\u00a7\u0003\u0016\u000b\u0000\u0081"+
		"\u00a7\u0003\u0018\f\u0000\u0082\u00a7\u0003\u001a\r\u0000\u0083\u0084"+
		"\u00059\u0000\u0000\u0084\u0085\u0005\u0014\u0000\u0000\u0085\u0086\u0003"+
		"\b\u0004\u0000\u0086\u0087\u0005\u0016\u0000\u0000\u0087\u0088\u0005\u0014"+
		"\u0000\u0000\u0088\u0089\u0003\b\u0004\u0000\u0089\u008a\u0005\u0016\u0000"+
		"\u0000\u008a\u00a7\u0001\u0000\u0000\u0000\u008b\u008c\u00059\u0000\u0000"+
		"\u008c\u008d\u0005\u0014\u0000\u0000\u008d\u008e\u0003\b\u0004\u0000\u008e"+
		"\u008f\u0005\u0016\u0000\u0000\u008f\u00a7\u0001\u0000\u0000\u0000\u0090"+
		"\u0091\u00059\u0000\u0000\u0091\u0093\u0005\u0001\u0000\u0000\u0092\u0094"+
		"\u0003$\u0012\u0000\u0093\u0092\u0001\u0000\u0000\u0000\u0093\u0094\u0001"+
		"\u0000\u0000\u0000\u0094\u0095\u0001\u0000\u0000\u0000\u0095\u00a7\u0005"+
		"\u0002\u0000\u0000\u0096\u00a7\u00055\u0000\u0000\u0097\u00a7\u00056\u0000"+
		"\u0000\u0098\u00a7\u00057\u0000\u0000\u0099\u00a7\u00058\u0000\u0000\u009a"+
		"\u00a7\u00059\u0000\u0000\u009b\u00a7\u0005\u001c\u0000\u0000\u009c\u009d"+
		"\u0005\u0012\u0000\u0000\u009d\u009e\u0003\b\u0004\u0000\u009e\u009f\u0005"+
		"\u0013\u0000\u0000\u009f\u00a5\u0001\u0000\u0000\u0000\u00a0\u00a1\u0005"+
		"\u0014\u0000\u0000\u00a1\u00a2\u0003\b\u0004\u0000\u00a2\u00a3\u0005\u0016"+
		"\u0000\u0000\u00a3\u00a5\u0001\u0000\u0000\u0000\u00a4\u009c\u0001\u0000"+
		"\u0000\u0000\u00a4\u00a0\u0001\u0000\u0000\u0000\u00a5\u00a7\u0001\u0000"+
		"\u0000\u0000\u00a6U\u0001\u0000\u0000\u0000\u00a6X\u0001\u0000\u0000\u0000"+
		"\u00a6^\u0001\u0000\u0000\u0000\u00a6j\u0001\u0000\u0000\u0000\u00a6t"+
		"\u0001\u0000\u0000\u0000\u00a6{\u0001\u0000\u0000\u0000\u00a6~\u0001\u0000"+
		"\u0000\u0000\u00a6\u0080\u0001\u0000\u0000\u0000\u00a6\u0081\u0001\u0000"+
		"\u0000\u0000\u00a6\u0082\u0001\u0000\u0000\u0000\u00a6\u0083\u0001\u0000"+
		"\u0000\u0000\u00a6\u008b\u0001\u0000\u0000\u0000\u00a6\u0090\u0001\u0000"+
		"\u0000\u0000\u00a6\u0096\u0001\u0000\u0000\u0000\u00a6\u0097\u0001\u0000"+
		"\u0000\u0000\u00a6\u0098\u0001\u0000\u0000\u0000\u00a6\u0099\u0001\u0000"+
		"\u0000\u0000\u00a6\u009a\u0001\u0000\u0000\u0000\u00a6\u009b\u0001\u0000"+
		"\u0000\u0000\u00a6\u00a4\u0001\u0000\u0000\u0000\u00a7\u00c7\u0001\u0000"+
		"\u0000\u0000\u00a8\u00a9\n\u001c\u0000\u0000\u00a9\u00aa\u0007\u0003\u0000"+
		"\u0000\u00aa\u00c6\u0003\b\u0004\u001d\u00ab\u00ac\n\u001b\u0000\u0000"+
		"\u00ac\u00ad\u0007\u0004\u0000\u0000\u00ad\u00c6\u0003\b\u0004\u001c\u00ae"+
		"\u00af\n\u001a\u0000\u0000\u00af\u00b0\u0007\u0005\u0000\u0000\u00b0\u00c6"+
		"\u0003\b\u0004\u001b\u00b1\u00b2\n\u0019\u0000\u0000\u00b2\u00b3\u0007"+
		"\u0006\u0000\u0000\u00b3\u00c6\u0003\b\u0004\u001a\u00b4\u00b5\n\u0018"+
		"\u0000\u0000\u00b5\u00b6\u0005\u000f\u0000\u0000\u00b6\u00c6\u0003\b\u0004"+
		"\u0019\u00b7\u00b8\n\u0017\u0000\u0000\u00b8\u00b9\u0005\u0010\u0000\u0000"+
		"\u00b9\u00c6\u0003\b\u0004\u0018\u00ba\u00bb\n\u0011\u0000\u0000\u00bb"+
		"\u00bc\u0005\u0011\u0000\u0000\u00bc\u00bd\u00059\u0000\u0000\u00bd\u00be"+
		"\u0005\u0017\u0000\u0000\u00be\u00c6\u0003\b\u0004\u0012\u00bf\u00c0\n"+
		"\u0010\u0000\u0000\u00c0\u00c1\u0005\u0017\u0000\u0000\u00c1\u00c6\u0003"+
		"\b\u0004\u0011\u00c2\u00c3\n\u0016\u0000\u0000\u00c3\u00c4\u0005\u0011"+
		"\u0000\u0000\u00c4\u00c6\u00059\u0000\u0000\u00c5\u00a8\u0001\u0000\u0000"+
		"\u0000\u00c5\u00ab\u0001\u0000\u0000\u0000\u00c5\u00ae\u0001\u0000\u0000"+
		"\u0000\u00c5\u00b1\u0001\u0000\u0000\u0000\u00c5\u00b4\u0001\u0000\u0000"+
		"\u0000\u00c5\u00b7\u0001\u0000\u0000\u0000\u00c5\u00ba\u0001\u0000\u0000"+
		"\u0000\u00c5\u00bf\u0001\u0000\u0000\u0000\u00c5\u00c2\u0001\u0000\u0000"+
		"\u0000\u00c6\u00c9\u0001\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000"+
		"\u0000\u00c7\u00c8\u0001\u0000\u0000\u0000\u00c8\t\u0001\u0000\u0000\u0000"+
		"\u00c9\u00c7\u0001\u0000\u0000\u0000\u00ca\u00cb\u0005\u001d\u0000\u0000"+
		"\u00cb\u00cc\u0003\b\u0004\u0000\u00cc\u00d2\u0003\u0006\u0003\u0000\u00cd"+
		"\u00d0\u0005\u001e\u0000\u0000\u00ce\u00d1\u0003\n\u0005\u0000\u00cf\u00d1"+
		"\u0003\u0006\u0003\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000\u00d0\u00cf"+
		"\u0001\u0000\u0000\u0000\u00d1\u00d3\u0001\u0000\u0000\u0000\u00d2\u00cd"+
		"\u0001\u0000\u0000\u0000\u00d2\u00d3\u0001\u0000\u0000\u0000\u00d3\u000b"+
		"\u0001\u0000\u0000\u0000\u00d4\u00d5\u0005\u001f\u0000\u0000\u00d5\u00d6"+
		"\u00059\u0000\u0000\u00d6\u00da\u0005\u0001\u0000\u0000\u00d7\u00d9\u0003"+
		"\u000e\u0007\u0000\u00d8\u00d7\u0001\u0000\u0000\u0000\u00d9\u00dc\u0001"+
		"\u0000\u0000\u0000\u00da\u00d8\u0001\u0000\u0000\u0000\u00da\u00db\u0001"+
		"\u0000\u0000\u0000\u00db\u00de\u0001\u0000\u0000\u0000\u00dc\u00da\u0001"+
		"\u0000\u0000\u0000\u00dd\u00df\u0003\u0010\b\u0000\u00de\u00dd\u0001\u0000"+
		"\u0000\u0000\u00de\u00df\u0001\u0000\u0000\u0000\u00df\u00e0\u0001\u0000"+
		"\u0000\u0000\u00e0\u00e1\u0005\u0002\u0000\u0000\u00e1\r\u0001\u0000\u0000"+
		"\u0000\u00e2\u00e3\u0005 \u0000\u0000\u00e3\u00e4\u0003\b\u0004\u0000"+
		"\u00e4\u00e8\u0005!\u0000\u0000\u00e5\u00e7\u0003\u0002\u0001\u0000\u00e6"+
		"\u00e5\u0001\u0000\u0000\u0000\u00e7\u00ea\u0001\u0000\u0000\u0000\u00e8"+
		"\u00e6\u0001\u0000\u0000\u0000\u00e8\u00e9\u0001\u0000\u0000\u0000\u00e9"+
		"\u000f\u0001\u0000\u0000\u0000\u00ea\u00e8\u0001\u0000\u0000\u0000\u00eb"+
		"\u00ec\u0005\"\u0000\u0000\u00ec\u00f0\u0005!\u0000\u0000\u00ed\u00ef"+
		"\u0003\u0002\u0001\u0000\u00ee\u00ed\u0001\u0000\u0000\u0000\u00ef\u00f2"+
		"\u0001\u0000\u0000\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f0\u00f1"+
		"\u0001\u0000\u0000\u0000\u00f1\u0011\u0001\u0000\u0000\u0000\u00f2\u00f0"+
		"\u0001\u0000\u0000\u0000\u00f3\u00f4\u0005#\u0000\u0000\u00f4\u00f5\u0003"+
		"\b\u0004\u0000\u00f5\u00f6\u0003\u0006\u0003\u0000\u00f6\u010b\u0001\u0000"+
		"\u0000\u0000\u00f7\u00fa\u0005#\u0000\u0000\u00f8\u00fb\u0003\u001e\u000f"+
		"\u0000\u00f9\u00fb\u0003\b\u0004\u0000\u00fa\u00f8\u0001\u0000\u0000\u0000"+
		"\u00fa\u00f9\u0001\u0000\u0000\u0000\u00fb\u00fc\u0001\u0000\u0000\u0000"+
		"\u00fc\u00fd\u0005$\u0000\u0000\u00fd\u00fe\u0003\b\u0004\u0000\u00fe"+
		"\u00ff\u0005$\u0000\u0000\u00ff\u0100\u0003\b\u0004\u0000\u0100\u0101"+
		"\u0003\u0006\u0003\u0000\u0101\u010b\u0001\u0000\u0000\u0000\u0102\u0103"+
		"\u0005#\u0000\u0000\u0103\u0104\u00059\u0000\u0000\u0104\u0105\u0005\u0015"+
		"\u0000\u0000\u0105\u0106\u00059\u0000\u0000\u0106\u0107\u0005\u0017\u0000"+
		"\u0000\u0107\u0108\u0005%\u0000\u0000\u0108\u0109\u00059\u0000\u0000\u0109"+
		"\u010b\u0003\u0006\u0003\u0000\u010a\u00f3\u0001\u0000\u0000\u0000\u010a"+
		"\u00f7\u0001\u0000\u0000\u0000\u010a\u0102\u0001\u0000\u0000\u0000\u010b"+
		"\u0013\u0001\u0000\u0000\u0000\u010c\u0113\u0005&\u0000\u0000\u010d\u0113"+
		"\u0005\'\u0000\u0000\u010e\u0110\u0005(\u0000\u0000\u010f\u0111\u0003"+
		"\b\u0004\u0000\u0110\u010f\u0001\u0000\u0000\u0000\u0110\u0111\u0001\u0000"+
		"\u0000\u0000\u0111\u0113\u0001\u0000\u0000\u0000\u0112\u010c\u0001\u0000"+
		"\u0000\u0000\u0112\u010d\u0001\u0000\u0000\u0000\u0112\u010e\u0001\u0000"+
		"\u0000\u0000\u0113\u0015\u0001\u0000\u0000\u0000\u0114\u0115\u0005)\u0000"+
		"\u0000\u0115\u0116\u0005\u0012\u0000\u0000\u0116\u0117\u0003\b\u0004\u0000"+
		"\u0117\u0118\u0005\u0013\u0000\u0000\u0118\u0017\u0001\u0000\u0000\u0000"+
		"\u0119\u011a\u0005*\u0000\u0000\u011a\u011b\u0005\u0012\u0000\u0000\u011b"+
		"\u011c\u0003\b\u0004\u0000\u011c\u011d\u0005\u0013\u0000\u0000\u011d\u0019"+
		"\u0001\u0000\u0000\u0000\u011e\u011f\u0005+\u0000\u0000\u011f\u0120\u0005"+
		"\u0012\u0000\u0000\u0120\u0121\u0003\b\u0004\u0000\u0121\u0122\u0005\u0013"+
		"\u0000\u0000\u0122\u001b\u0001\u0000\u0000\u0000\u0123\u0124\u0005,\u0000"+
		"\u0000\u0124\u012d\u0005\u0012\u0000\u0000\u0125\u012a\u0003\b\u0004\u0000"+
		"\u0126\u0127\u0005\u0015\u0000\u0000\u0127\u0129\u0003\b\u0004\u0000\u0128"+
		"\u0126\u0001\u0000\u0000\u0000\u0129\u012c\u0001\u0000\u0000\u0000\u012a"+
		"\u0128\u0001\u0000\u0000\u0000\u012a\u012b\u0001\u0000\u0000\u0000\u012b"+
		"\u012e\u0001\u0000\u0000\u0000\u012c\u012a\u0001\u0000\u0000\u0000\u012d"+
		"\u0125\u0001\u0000\u0000\u0000\u012d\u012e\u0001\u0000\u0000\u0000\u012e"+
		"\u012f\u0001\u0000\u0000\u0000\u012f\u0130\u0005\u0013\u0000\u0000\u0130"+
		"\u001d\u0001\u0000\u0000\u0000\u0131\u0133\u0005-\u0000\u0000\u0132\u0131"+
		"\u0001\u0000\u0000\u0000\u0132\u0133\u0001\u0000\u0000\u0000\u0133\u0134"+
		"\u0001\u0000\u0000\u0000\u0134\u0135\u00059\u0000\u0000\u0135\u0138\u0007"+
		"\u0007\u0000\u0000\u0136\u0137\u0005\u0017\u0000\u0000\u0137\u0139\u0003"+
		"\b\u0004\u0000\u0138\u0136\u0001\u0000\u0000\u0000\u0138\u0139\u0001\u0000"+
		"\u0000\u0000\u0139\u0180\u0001\u0000\u0000\u0000\u013a\u013c\u0005-\u0000"+
		"\u0000\u013b\u013a\u0001\u0000\u0000\u0000\u013b\u013c\u0001\u0000\u0000"+
		"\u0000\u013c\u013d\u0001\u0000\u0000\u0000\u013d\u013e\u00059\u0000\u0000"+
		"\u013e\u013f\u0005.\u0000\u0000\u013f\u0145\u00054\u0000\u0000\u0140\u0143"+
		"\u0005\u0017\u0000\u0000\u0141\u0144\u0003 \u0010\u0000\u0142\u0144\u0003"+
		"\b\u0004\u0000\u0143\u0141\u0001\u0000\u0000\u0000\u0143\u0142\u0001\u0000"+
		"\u0000\u0000\u0144\u0146\u0001\u0000\u0000\u0000\u0145\u0140\u0001\u0000"+
		"\u0000\u0000\u0145\u0146\u0001\u0000\u0000\u0000\u0146\u0180\u0001\u0000"+
		"\u0000\u0000\u0147\u0149\u0005-\u0000\u0000\u0148\u0147\u0001\u0000\u0000"+
		"\u0000\u0148\u0149\u0001\u0000\u0000\u0000\u0149\u014a\u0001\u0000\u0000"+
		"\u0000\u014a\u014b\u00059\u0000\u0000\u014b\u014c\u0005/\u0000\u0000\u014c"+
		"\u014f\u00054\u0000\u0000\u014d\u014e\u0005\u0017\u0000\u0000\u014e\u0150"+
		"\u0003\"\u0011\u0000\u014f\u014d\u0001\u0000\u0000\u0000\u014f\u0150\u0001"+
		"\u0000\u0000\u0000\u0150\u0180\u0001\u0000\u0000\u0000\u0151\u0153\u0005"+
		"-\u0000\u0000\u0152\u0151\u0001\u0000\u0000\u0000\u0152\u0153\u0001\u0000"+
		"\u0000\u0000\u0153\u0154\u0001\u0000\u0000\u0000\u0154\u0155\u00059\u0000"+
		"\u0000\u0155\u015c\u00059\u0000\u0000\u0156\u0157\u0005\u0017\u0000\u0000"+
		"\u0157\u0159\u0005\u0001\u0000\u0000\u0158\u015a\u0003$\u0012\u0000\u0159"+
		"\u0158\u0001\u0000\u0000\u0000\u0159\u015a\u0001\u0000\u0000\u0000\u015a"+
		"\u015b\u0001\u0000\u0000\u0000\u015b\u015d\u0005\u0002\u0000\u0000\u015c"+
		"\u0156\u0001\u0000\u0000\u0000\u015c\u015d\u0001\u0000\u0000\u0000\u015d"+
		"\u0180\u0001\u0000\u0000\u0000\u015e\u0160\u0005-\u0000\u0000\u015f\u015e"+
		"\u0001\u0000\u0000\u0000\u015f\u0160\u0001\u0000\u0000\u0000\u0160\u0161"+
		"\u0001\u0000\u0000\u0000\u0161\u0162\u00059\u0000\u0000\u0162\u0163\u0005"+
		"0\u0000\u0000\u0163\u0164\u00059\u0000\u0000\u0164\u0166\u0005\u0001\u0000"+
		"\u0000\u0165\u0167\u0003$\u0012\u0000\u0166\u0165\u0001\u0000\u0000\u0000"+
		"\u0166\u0167\u0001\u0000\u0000\u0000\u0167\u0168\u0001\u0000\u0000\u0000"+
		"\u0168\u0180\u0005\u0002\u0000\u0000\u0169\u016b\u0005-\u0000\u0000\u016a"+
		"\u0169\u0001\u0000\u0000\u0000\u016a\u016b\u0001\u0000\u0000\u0000\u016b"+
		"\u016c\u0001\u0000\u0000\u0000\u016c\u016d\u00059\u0000\u0000\u016d\u016e"+
		"\u00050\u0000\u0000\u016e\u016f\u0005.\u0000\u0000\u016f\u0170\u00054"+
		"\u0000\u0000\u0170\u0180\u0003 \u0010\u0000\u0171\u0173\u0005-\u0000\u0000"+
		"\u0172\u0171\u0001\u0000\u0000\u0000\u0172\u0173\u0001\u0000\u0000\u0000"+
		"\u0173\u0174\u0001\u0000\u0000\u0000\u0174\u0175\u00059\u0000\u0000\u0175"+
		"\u0176\u00050\u0000\u0000\u0176\u0177\u0005/\u0000\u0000\u0177\u0178\u0005"+
		"4\u0000\u0000\u0178\u0180\u0003\"\u0011\u0000\u0179\u017b\u0005-\u0000"+
		"\u0000\u017a\u0179\u0001\u0000\u0000\u0000\u017a\u017b\u0001\u0000\u0000"+
		"\u0000\u017b\u017c\u0001\u0000\u0000\u0000\u017c\u017d\u00059\u0000\u0000"+
		"\u017d\u017e\u00050\u0000\u0000\u017e\u0180\u0003\b\u0004\u0000\u017f"+
		"\u0132\u0001\u0000\u0000\u0000\u017f\u013b\u0001\u0000\u0000\u0000\u017f"+
		"\u0148\u0001\u0000\u0000\u0000\u017f\u0152\u0001\u0000\u0000\u0000\u017f"+
		"\u015f\u0001\u0000\u0000\u0000\u017f\u016a\u0001\u0000\u0000\u0000\u017f"+
		"\u0172\u0001\u0000\u0000\u0000\u017f\u017a\u0001\u0000\u0000\u0000\u0180"+
		"\u001f\u0001\u0000\u0000\u0000\u0181\u018a\u0005\u0001\u0000\u0000\u0182"+
		"\u0187\u0003\b\u0004\u0000\u0183\u0184\u0005\u0015\u0000\u0000\u0184\u0186"+
		"\u0003\b\u0004\u0000\u0185\u0183\u0001\u0000\u0000\u0000\u0186\u0189\u0001"+
		"\u0000\u0000\u0000\u0187\u0185\u0001\u0000\u0000\u0000\u0187\u0188\u0001"+
		"\u0000\u0000\u0000\u0188\u018b\u0001\u0000\u0000\u0000\u0189\u0187\u0001"+
		"\u0000\u0000\u0000\u018a\u0182\u0001\u0000\u0000\u0000\u018a\u018b\u0001"+
		"\u0000\u0000\u0000\u018b\u018c\u0001\u0000\u0000\u0000\u018c\u018d\u0005"+
		"\u0002\u0000\u0000\u018d!\u0001\u0000\u0000\u0000\u018e\u0197\u0005\u0001"+
		"\u0000\u0000\u018f\u0194\u0003 \u0010\u0000\u0190\u0191\u0005\u0015\u0000"+
		"\u0000\u0191\u0193\u0003 \u0010\u0000\u0192\u0190\u0001\u0000\u0000\u0000"+
		"\u0193\u0196\u0001\u0000\u0000\u0000\u0194\u0192\u0001\u0000\u0000\u0000"+
		"\u0194\u0195\u0001\u0000\u0000\u0000\u0195\u0198\u0001\u0000\u0000\u0000"+
		"\u0196\u0194\u0001\u0000\u0000\u0000\u0197\u018f\u0001\u0000\u0000\u0000"+
		"\u0197\u0198\u0001\u0000\u0000\u0000\u0198\u0199\u0001\u0000\u0000\u0000"+
		"\u0199\u019a\u0005\u0002\u0000\u0000\u019a#\u0001\u0000\u0000\u0000\u019b"+
		"\u01a0\u0003&\u0013\u0000\u019c\u019d\u0005\u0015\u0000\u0000\u019d\u019f"+
		"\u0003&\u0013\u0000\u019e\u019c\u0001\u0000\u0000\u0000\u019f\u01a2\u0001"+
		"\u0000\u0000\u0000\u01a0\u019e\u0001\u0000\u0000\u0000\u01a0\u01a1\u0001"+
		"\u0000\u0000\u0000\u01a1%\u0001\u0000\u0000\u0000\u01a2\u01a0\u0001\u0000"+
		"\u0000\u0000\u01a3\u01a4\u00059\u0000\u0000\u01a4\u01a5\u0005!\u0000\u0000"+
		"\u01a5\u01a6\u0003\b\u0004\u0000\u01a6\'\u0001\u0000\u0000\u0000\u01a7"+
		"\u01a8\u00051\u0000\u0000\u01a8\u01a9\u00059\u0000\u0000\u01a9\u01ad\u0005"+
		"\u0001\u0000\u0000\u01aa\u01ac\u0003*\u0015\u0000\u01ab\u01aa\u0001\u0000"+
		"\u0000\u0000\u01ac\u01af\u0001\u0000\u0000\u0000\u01ad\u01ab\u0001\u0000"+
		"\u0000\u0000\u01ad\u01ae\u0001\u0000\u0000\u0000\u01ae\u01b0\u0001\u0000"+
		"\u0000\u0000\u01af\u01ad\u0001\u0000\u0000\u0000\u01b0\u01b1\u0005\u0002"+
		"\u0000\u0000\u01b1)\u0001\u0000\u0000\u0000\u01b2\u01b3\u0007\u0007\u0000"+
		"\u0000\u01b3\u01b4\u00059\u0000\u0000\u01b4+\u0001\u0000\u0000\u0000\u01b5"+
		"\u01b6\u00052\u0000\u0000\u01b6\u01b7\u00059\u0000\u0000\u01b7\u01b9\u0005"+
		"\u0012\u0000\u0000\u01b8\u01ba\u0003.\u0017\u0000\u01b9\u01b8\u0001\u0000"+
		"\u0000\u0000\u01b9\u01ba\u0001\u0000\u0000\u0000\u01ba\u01bb\u0001\u0000"+
		"\u0000\u0000\u01bb\u01bd\u0005\u0013\u0000\u0000\u01bc\u01be\u0007\b\u0000"+
		"\u0000\u01bd\u01bc\u0001\u0000\u0000\u0000\u01bd\u01be\u0001\u0000\u0000"+
		"\u0000\u01be\u01bf\u0001\u0000\u0000\u0000\u01bf\u01c0\u0003\u0006\u0003"+
		"\u0000\u01c0-\u0001\u0000\u0000\u0000\u01c1\u01c6\u00030\u0018\u0000\u01c2"+
		"\u01c3\u0005\u0015\u0000\u0000\u01c3\u01c5\u00030\u0018\u0000\u01c4\u01c2"+
		"\u0001\u0000\u0000\u0000\u01c5\u01c8\u0001\u0000\u0000\u0000\u01c6\u01c4"+
		"\u0001\u0000\u0000\u0000\u01c6\u01c7\u0001\u0000\u0000\u0000\u01c7/\u0001"+
		"\u0000\u0000\u0000\u01c8\u01c6\u0001\u0000\u0000\u0000\u01c9\u01ca\u0005"+
		"9\u0000\u0000\u01ca\u01cb\u0007\u0007\u0000\u0000\u01cb1\u0001\u0000\u0000"+
		"\u0000\u01cc\u01d1\u0003\b\u0004\u0000\u01cd\u01ce\u0005\u0015\u0000\u0000"+
		"\u01ce\u01d0\u0003\b\u0004\u0000\u01cf\u01cd\u0001\u0000\u0000\u0000\u01d0"+
		"\u01d3\u0001\u0000\u0000\u0000\u01d1\u01cf\u0001\u0000\u0000\u0000\u01d1"+
		"\u01d2\u0001\u0000\u0000\u0000\u01d23\u0001\u0000\u0000\u0000\u01d3\u01d1"+
		"\u0001\u0000\u0000\u000027>JP[dg\u0093\u00a4\u00a6\u00c5\u00c7\u00d0\u00d2"+
		"\u00da\u00de\u00e8\u00f0\u00fa\u010a\u0110\u0112\u012a\u012d\u0132\u0138"+
		"\u013b\u0143\u0145\u0148\u014f\u0152\u0159\u015c\u015f\u0166\u016a\u0172"+
		"\u017a\u017f\u0187\u018a\u0194\u0197\u01a0\u01ad\u01b9\u01bd\u01c6\u01d1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}