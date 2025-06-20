package repl

type ScopeTrace struct {
    GlobalScope *BaseScope
    LocalScope  *BaseScope
    ScopeStack  []*BaseScope // Stack de scopes locales
}

type BaseScope struct {
    Variables map[string]interface{} // CAMBIAR a interface{}
    Parent    *BaseScope
    ScopeType string // "global", "function", "for", "if", "block"
}

// Crear nuevo scope base
func NewBaseScope(scopeType string, parent *BaseScope) *BaseScope {
    return &BaseScope{
        Variables: make(map[string]interface{}), // CAMBIAR
        Parent:    parent,
        ScopeType: scopeType,
    }
}

// Crear nuevo ScopeTrace
func NewScopeTrace() *ScopeTrace {
    globalScope := NewBaseScope("global", nil)
    return &ScopeTrace{
        GlobalScope: globalScope,
        LocalScope:  globalScope, // Inicialmente apunta al global
        ScopeStack:  []*BaseScope{globalScope},
    }
}

// Entrar a un nuevo scope local
func (st *ScopeTrace) EnterScope(scopeType string) {
    newScope := NewBaseScope(scopeType, st.LocalScope)
    st.ScopeStack = append(st.ScopeStack, newScope)
    st.LocalScope = newScope
}

// Salir del scope local actual
func (st *ScopeTrace) ExitScope() {
    if len(st.ScopeStack) > 1 { // Mantener siempre el global
        st.ScopeStack = st.ScopeStack[:len(st.ScopeStack)-1]
        st.LocalScope = st.ScopeStack[len(st.ScopeStack)-1]
    }
}

// Buscar variable en todos los scopes (desde local hacia global)
func (st *ScopeTrace) GetVariable(varName string) (interface{}, bool) {
    // CORREGIDO: Usar st.ScopeStack directamente
    for i := len(st.ScopeStack) - 1; i >= 0; i-- {
        scope := st.ScopeStack[i]
        if value, exists := scope.Variables[varName]; exists {
            return value, true
        }
    }
    return nil, false
}

// Establecer variable en el scope local actual
func (st *ScopeTrace) SetVariable(name string, val interface{}) {
    st.LocalScope.Variables[name] = val
}

// Actualizar variable existente (buscar en todos los scopes)
func (st *ScopeTrace) UpdateVariable(varName string, newValue interface{}) bool {
    // Buscar la variable desde el scope más reciente hacia atrás
    for i := len(st.ScopeStack) - 1; i >= 0; i-- {
        scope := st.ScopeStack[i]
        if _, exists := scope.Variables[varName]; exists {
            // Encontrada - actualizar en este scope
            scope.Variables[varName] = newValue
            return true
        }
    }
    return false // No encontrada en ningún scope
}

// Obtener todas las variables del scope global (para tabla de símbolos)
func (st *ScopeTrace) GetGlobalVariables() map[string]interface{} {
    return st.GlobalScope.Variables
}

// Obtener todas las variables del scope local actual
func (st *ScopeTrace) GetLocalVariables() map[string]interface{} {
    return st.LocalScope.Variables
}

// Verificar si existe en scope actual
func (st *ScopeTrace) ExistsInCurrentScope(varName string) bool {
    if len(st.ScopeStack) == 0 {
        return false
    }
    
    // Solo buscar en el scope más reciente (actual)
    currentScope := st.ScopeStack[len(st.ScopeStack)-1]
    _, exists := currentScope.Variables[varName]
    return exists
}

func (st *ScopeTrace) CreateLocalVariable(varName string, value interface{}) {
    // Crear variable en el scope actual sin verificar scopes superiores
    st.LocalScope.Variables[varName] = value
}