grammar Language;

program: statement*;

statement: nonDeclaration
         | variableDeclaration
         | structDeclaration
         | functionDeclaration
;

nonDeclaration: blockStatement
                | printStatement          
                | atoiStatement           
                | parseFloatStatement     
                | typeOfStatement         
                | ifStatement             
                | switchStatement         
                | forStatement            
                | transferenceStatement   
                | expressionStatement     
;

blockStatement: '{' statement* '}';

expressionStatement: 
    ( '-' | '!' ) expressionStatement                                                       # Negate
    | expressionStatement ( '*' | '/' | '%') expressionStatement                            # MulDivMod
    | expressionStatement ( '+' | '-' ) expressionStatement                                 # AddSub
    | expressionStatement ( '>' | '<' | '>=' | '<=' ) expressionStatement                   # GreaterLess
    | expressionStatement ( '==' | '!=' ) expressionStatement                               # Equal
    | expressionStatement '&&' expressionStatement                                          # And
    | expressionStatement '||' expressionStatement                                          # Or
    | expressionStatement '.' ID                                                           # StructAccess
    
    | ID '(' argumentList? ')'                                                              # FunctionCall
    | '[' ( expressionStatement ( ',' expressionStatement )* )? ']'                        # SliceLiteral

    // MOVER LAS ASIGNACIONES ESPECÍFICAS ANTES DE LA GENERAL:
    | ID '[' expressionStatement ']' '[' expressionStatement ']' '=' expressionStatement   # MatrixAssignment
    | ID '[' expressionStatement ']' '=' expressionStatement                               # SliceAssignment
    | expressionStatement '.' ID '=' expressionStatement                                  # StructAssignment
    | expressionStatement '=' expressionStatement                                          # Assignment
    
    | ID ( '+=' | '-=' ) expressionStatement                                                # AddSubOperator
    | ID ( '++' | '--' )                                                                    # IncrementDecrement
    | atoiStatement                                                                         # AtoiExpr
    | parseFloatStatement                                                                   # ParseFloatExpr
    | typeOfStatement                                                                       # TypeOfExpr
    | ID '[' expressionStatement ']' '[' expressionStatement ']'                            # MatrixAccess
    | ID '[' expressionStatement ']'                                                        # SliceAccess    
    
    | ID '{' structInitialization? '}'                                                     # StructInstantiation
    | INTEGER                                                                               # Integer
    | BOOLEAN                                                                               # Boolean
    | FLOAT                                                                                 # Float
    | STRING                                                                                # String
    | ID                                                                                    # Identifier
    | 'nil'                                                                                  # Nil
    | (('(' expressionStatement ')') | ('[' expressionStatement ']'))                       # Parens
;

ifStatement: 'if' expressionStatement blockStatement ('else' (  ifStatement | blockStatement    ))?;

switchStatement: 'switch' ID '{' switchCase* defaultCase? '}';

switchCase: 'case' expressionStatement ':' statement*;

defaultCase: 'default' ':' statement*;

forStatement: 'for' (variableDeclaration | expressionStatement) ';' expressionStatement ';' expressionStatement blockStatement      # ForSimple
            | 'for' expressionStatement blockStatement                                                                              # ForConditional
            | 'for' ID ',' ID '=' 'range' ID blockStatement                                                                         # ForSlice
;

transferenceStatement: 'break'                          # BreakStatement
                     | 'continue'                       # ContinueStatement
                     | 'return' expressionStatement?    # ReturnStatement
;

// Funciones embedidas
atoiStatement: 'atoi' '(' expressionStatement ')';
parseFloatStatement: 'parseFloat' '(' expressionStatement ')';
typeOfStatement: 'typeOf' '(' expressionStatement ')';
printStatement: 'print' '(' ( expressionStatement ( ',' expressionStatement )* )? ')'
;

variableDeclaration: ('mut')? ID TYPE  ( '=' expressionStatement )?                             # ExplicitDeclaration
                   | ('mut')? ID '[]' TYPE ( '=' ( sliceElements | expressionStatement ) )?     # ExplicitSliceDeclaration
                   | ('mut')? ID '[][]' TYPE ( '=' matrixElements )?                            # ExplicitMatrixDeclaration
                   | ('mut')? ID ID ( '=' '{' structInitialization? '}' )?                      # ExplicitStructDeclaration
                   
                   // MOVER LAS ESPECÍFICAS ANTES DE LAS GENERALES:
                   | ('mut')? ID ':=' ID '{' structInitialization? '}'                          # ImplicitStructDeclaration
                   | ('mut')? ID ':=' '[]' TYPE sliceElements                                   # ImplicitSliceDeclaration
                   | ('mut')? ID ':=' '[][]' TYPE matrixElements                                # ImplicitMatrixDeclaration
                   | ('mut')? ID ':=' expressionStatement                                       # ImplicitDeclaration
;

sliceElements: '{' ( expressionStatement ( ',' expressionStatement )* )? '}';

matrixElements: '{' ( sliceElements ( ',' sliceElements )* )? '}';

structInitialization: structFieldInit (',' structFieldInit)*;
structFieldInit: ID ':' expressionStatement;

structDeclaration: 'struct' ID '{' structAtribute* '}' ;

structAtribute: (TYPE | ID) ID;

functionDeclaration: 'fn' ID '(' parameterList? ')' (TYPE | ID | 'void')? blockStatement ;

parameterList: parameter (',' parameter)* ;

parameter: ID (TYPE | ID) ;

argumentList: expressionStatement (',' expressionStatement)* ;

TYPE: 'int' | 'float64' | 'string' | 'bool' | 'nil';

INTEGER: [0-9]+;
BOOLEAN: 'true' | 'false';
FLOAT: [0-9]+ '.' [0-9]*;
STRING: ('"' ( '\\' . | ~[\\"\r\n] )* '"') | ('\'' ( '\\' . | ~[\\'\r\n] )* '\'');
ID: [a-zA-Z_][a-zA-Z_0-9]*;

ONELINECOMMENT: '//' ~[\r\n]* -> skip;
MULTILINECOMMENT: '/*' .*? '*/' -> skip;
WS: [ \t\r\n]+ -> skip;