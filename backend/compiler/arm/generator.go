package arm

import (
    "fmt"
    "strings"

)

type StackObjectType int

const (
    Int StackObjectType = iota
    Float
    String
    Bool
)

// String method para convertir el enum a string (útil para debugging)
func (s StackObjectType) String() string {
    switch s {
    case Int:
        return "Int"
    case Float:
        return "Float"
    case String:
        return "String"
    case Bool:
        return "Bool"
    default:
        return "Unknown"
    }
}

// Estructura StackObject corregida
type StackObject struct {
    Type   StackObjectType
    Length int
    Depth  int
    Id     *string
}

var instructions = []string{}
var stack = []StackObject{}
var depth = 0

// Stack Operations

func PushObject(obj StackObject) {
    stack = append(stack, obj)
}

func PushConstant(value interface{}, objType StackObject) {
    switch objType.Type { 
    case Int:             
        Mov(X0, value.(int))
    case Float:
        //TODO    
    case String:
        var stringArray []byte = StringTo1ByteArray(value.(string))
        Push(HP)

        for i:=0; i<len(stringArray); i++{
            var charCode = stringArray[i]
            Comment(fmt.Sprintf("Pushing char %d to heap - (%c)", charCode, charCode))
            Mov(W0, int(charCode))
            StrB(W0, HP) // Guardar el byte en la posición actual del heap
            Mov(X0, 1)
            Add(HP, HP, X0) // Incrementar el puntero del heap
        }
    case Bool:
        if value.(bool) {
            Mov(X0, 1) // Verdadero
        } else {
            Mov(X0, 0) // Falso
        }
        Push(X0) // Guardar el booleano en el stack
    }

    PushObject(objType)
}


func PopObject(rd string) StackObject {
    var obj = stack[len(stack)-1]
    stack = stack[:len(stack)-1] // Eliminar el último objeto del stack
    Pop(rd)
    return obj
}

func IntObject() StackObject {
    return StackObject{Type: Int, Length: 8, Depth: depth, Id: nil}
}

func FloatObject() StackObject {
    return StackObject{Type: Float, Length: 8, Depth: depth, Id: nil}
}

func StringObject() StackObject {
    return StackObject{Type: String, Length: 8, Depth: depth, Id: nil}
}

func BoolObject() StackObject {
    return StackObject{Type: Bool, Length: 8, Depth: depth, Id: nil}
}

func CloneObject(obj StackObject) StackObject {
    // Clonar un objeto del stack
    return StackObject{
        Type:   obj.Type,
        Length: obj.Length,
        Depth:  obj.Depth,
        Id:     obj.Id,
    }
}

func NewScope() {
    depth++
}

func EndScope() int {
    var byteOffset = 0

    for i := len(stack) - 1; i >= 0; i-- {
        if stack[i].Depth == depth {
            // Si el objeto es del mismo nivel de profundidad, lo eliminamos
            byteOffset += stack[i].Length
            stack = stack[:len(stack)-1] // Eliminar el último objeto del stack  
        } else {
            // Si encontramos un objeto de menor profundidad, salimos del bucle
            break
        }
    }

    depth-- // Reducir la profundidad al salir del scope
    return byteOffset // Retornar el tamaño en bytes del scope eliminado
}

func TagObject(id string) {
    stack[len(stack)-1].Id = &id // Asignar un ID al último objeto del stack
}

func GetObject(id string) (int, StackObject) {

    var byteOffset = 0
    for i := len(stack) - 1; i >= 0; i-- {
        if stack[i].Id != nil && *stack[i].Id == id {
            return byteOffset, stack[i] // Retornar el índice y el objeto si encontramos el ID
        }

        byteOffset += stack[i].Length // Sumar el tamaño del objeto al byteOffset
    }

    // Posiblemente sea necesario corregir para que en vez de hacer un retorno, se arroje una excepción
    // como se hace en la grabacion del 10/4 en el minuto 42:00
    return -1, StackObject{} // Retornar -1 y un objeto vacío si no se encuentra
}

func Add(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("add %s, %s, %s", rd, rs1, rs2))
}

func Sub(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("sub %s, %s, %s", rd, rs1, rs2))
}

func Mul(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("mul %s, %s, %s", rd, rs1, rs2))
}

func SDiv(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("sdiv %s, %s, %s", rd, rs1, rs2))
}

func UDiv(rd string, rs1 string, rs2 string) {
    instructions = append(instructions, fmt.Sprintf("udiv %s, %s, %s", rd, rs1, rs2))
}

func Addi(rd string, rs1 string, imm int) {
	instructions = append(instructions, fmt.Sprintf("addi %s, %s, #%d", rd, rs1, imm))
}

func Subi(rd string, rs1 string, imm int) {
    instructions = append(instructions, fmt.Sprintf("subi %s, %s, #%d", rd, rs1, imm))
}

func Neg(rd string, rs string) {
    instructions = append(instructions, fmt.Sprintf("neg %s, %s", rd, rs))
}

// Logical operations
func And(rd string, rs1 string, rs2 string) {
    instructions = append(instructions, fmt.Sprintf("and %s, %s, %s", rd, rs1, rs2))
}
func Orr(rd string, rs1 string, rs2 string) {
    instructions = append(instructions, fmt.Sprintf("orr %s, %s, %s", rd, rs1, rs2))
}

// Comparison operations
func Cmp(rs1 string, rs2 string) {
    instructions = append(instructions, fmt.Sprintf("cmp %s, %s", rs1, rs2))
}

func CmpImm(rs string, imm int) {
    instructions = append(instructions, fmt.Sprintf("cmp %s, #%d", rs, imm))
}

func Cset(rd string, condition string) {
    instructions = append(instructions, fmt.Sprintf("cset %s, %s", rd, condition))
}

// Memory operations

func Str(rs1 string, rs2 string, offset int) {
	instructions = append(instructions, fmt.Sprintf("str %s, [%s, #%d]", rs1, rs2, offset))
}

func StrB(rs1 string, rs2 string) {
    instructions = append(instructions, fmt.Sprintf("strb %s, [%s]", rs1, rs2))
}

func Ldr(rd string, rs1 string, offset int) {
	instructions = append(instructions, fmt.Sprintf("ldr %s, [%s, #%d]", rd, rs1, offset))
}

func Mov(rd string, imm int) {
	instructions = append(instructions, fmt.Sprintf("mov %s, #%d", rd, imm))
}

func MovReg(rd string, rs string) {
    instructions = append(instructions, fmt.Sprintf("mov %s, %s", rd, rs))
}

func Push(rs string) {
    instructions = append(instructions, fmt.Sprintf("str %s, [sp, #-8]!", rs))
}

func Pop(rd string) {
	instructions = append(instructions, fmt.Sprintf("ldr %s, [sp], #8", rd))
}

func Svc() {
	instructions = append(instructions, "svc #0")
}

func EndProgram() {
	Mov(X0, 0)
	Mov(X8, 93) // syscall number for exit
	Svc() // make syscall
}

func PrintInt(rs string) {
    // Mover el valor al registro x0 (parámetro para print_integer)
    if rs != X0 {  // Solo mover si no es ya X0
        MovReg(X0, rs)
    }
    
    // Llamar a la función print_integer
    instructions = append(instructions, "bl print_integer")
    
    // Marcar que usamos esta función estándar
    Use("print_integer")
}

func PrintString(rs string) {
    // Mover el puntero de la cadena al registro x0 (parámetro para print_string)
    if rs != X0 {  // Solo mover si no es ya X0
        MovReg(X0, rs)
    }
    
    // Llamar a la función print_string
    instructions = append(instructions, "bl print_string")
    
    // Marcar que usamos esta función estándar
    Use("print_string")
}

func Comment(comment string) {
	instructions = append(instructions, fmt.Sprintf("# %s", comment))
}

func ToString() string {
    var sb strings.Builder
    sb.WriteString(".data\n")
    sb.WriteString("heap: .space 4096\n")
    sb.WriteString(".text\n")
    sb.WriteString(".global _start\n")
    sb.WriteString("_start:\n")
    sb.WriteString("    adr x10, heap\n")

    for _, instr := range instructions {
        sb.WriteString(fmt.Sprintf("    %s\n", instr))
    }
    
    standardFunctions := GetFunctionDefinitions()
    if standardFunctions != "" {
        sb.WriteString("\n\n\n// Standard Library\n")
        sb.WriteString(standardFunctions)
    }

    return sb.String()
}