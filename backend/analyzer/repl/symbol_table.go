package repl

// Tipos de símbolos que podemos encontrar
const (
    VariableSymbol  = "Variable"
    FunctionSymbol  = "Función"
    ParameterSymbol = "Parámetro"
    SliceSymbol     = "Slice"
    MatrixSymbol    = "Matriz"
    StructSymbol   = "Struct"       
    StructTypeSymbol = "Tipo Struct" 
)

// Tipos de datos
const (
    IntType     = "int"
    FloatType   = "float64"
    StringType  = "string"
    BoolType    = "bool"
    SliceType   = "slice"
    MatrixType  = "matrix"
    FunctionType = "function"
    StructType  = "struct"
)

// Estructura para un símbolo individual
type Symbol struct {
    ID         string `json:"id"`         // Nombre del identificador
    SymbolType string `json:"symbolType"` // Variable, Función, etc.
    DataType   string `json:"dataType"`   // int, string, etc.
    Scope      string `json:"scope"`      // Global, Local, For, etc.
    Line       int    `json:"line"`       // Línea donde se declara
    Column     int    `json:"column"`     // Columna donde se declara
}

// Tabla de símbolos
type SymbolTable struct {
    Symbols []Symbol `json:"symbols"`
}

// Agregar un símbolo a la tabla
func (st *SymbolTable) AddSymbol(symbol Symbol) {
    st.Symbols = append(st.Symbols, symbol)
}

// Crear nueva tabla de símbolos
func NewSymbolTable() *SymbolTable {
    return &SymbolTable{
        Symbols: make([]Symbol, 0),
    }
}

// Obtener símbolo por nombre (útil para verificaciones)
func (st *SymbolTable) GetSymbol(name string) (*Symbol, bool) {
    for _, symbol := range st.Symbols {
        if symbol.ID == name {
            return &symbol, true
        }
    }
    return nil, false
}

// Verificar si un símbolo existe
func (st *SymbolTable) SymbolExists(name string) bool {
    _, exists := st.GetSymbol(name)
    return exists
}

// Obtener símbolos por tipo
func (st *SymbolTable) GetSymbolsByType(symbolType string) []Symbol {
    var result []Symbol
    for _, symbol := range st.Symbols {
        if symbol.SymbolType == symbolType {
            result = append(result, symbol)
        }
    }
    return result
}

// Obtener símbolos por ámbito
func (st *SymbolTable) GetSymbolsByScope(scope string) []Symbol {
    var result []Symbol
    for _, symbol := range st.Symbols {
        if symbol.Scope == scope {
            result = append(result, symbol)
        }
    }
    return result
}