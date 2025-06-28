package arm

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type StackObjectType int

const (
	Int StackObjectType = iota
	Float
	String
	Bool
)

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
	Type     StackObjectType
	Length   int
	Depth    int
	Id       *string
	IsSlice  bool
	ElemType StackObjectType
	Size     int
}

func GetStackDebug() []StackObject {
    return stack
}

var instructions = []string{}
var stack = []StackObject{}
var depth = 0

func TopObject() StackObject {
    if len(stack) == 0 {
        log.Printf("⚠️  ADVERTENCIA: Intento de acceder a TopObject() con pila vacía")
        // Retornar un objeto entero por defecto
        return StackObject{Type: Int, Length: 8, Depth: 0, Id: nil}
    }
    return stack[len(stack)-1]
}

func PushObject(obj StackObject) {
	obj.Depth = depth // Asignar la profundidad actual
	stack = append(stack, obj)
}

func PushConstant(value interface{}, objType StackObject) {
	switch objType.Type {
	case Int:
		Mov(X0, value.(int))
		Push(X0)
	case Float:
		var floatBits uint64 = math.Float64bits(value.(float64))
		var floatParts = [4]uint16{}

		for i := 0; i < 4; i++ {
			floatParts[i] = uint16((floatBits >> (i * 16)) & 0xFFFF)
		}
		instructions = append(instructions, fmt.Sprintf("movz x0, #%d, lsl #0", floatParts[0]))
		for i := 1; i < 4; i++ {
			instructions = append(instructions, fmt.Sprintf("movk x0, #%d, lsl #%d", floatParts[i], i*16))
		}
		instructions = append(instructions, "fmov d0, x0")
		Push(D0)
	case String:
		var stringArray []byte = StringTo1ByteArray(value.(string))
		MovReg("x11", HP) // Guarda el inicio del string en x11
		for i := 0; i < len(stringArray); i++ {
			var charCode = stringArray[i]
			Comment(fmt.Sprintf("Pushing char %d to heap - (%c)", charCode, charCode))
			Mov(W0, int(charCode))
			StrB(W0, HP)
			Mov(X0, 1)
			Add(HP, HP, X0)
		}
		// Agregar terminador nulo
		Comment("Pushing null terminator to heap")
		Mov(W0, 0)
		StrB(W0, HP)
		Mov(X0, 1)
		Add(HP, HP, X0)
		// Ahora pon el puntero real del string en X0
		MovReg(X0, "x11")
		Push(X0)
	case Bool:
		if value.(bool) {
			Mov(X0, 1)
		} else {
			Mov(X0, 0)
		}
		Push(X0)
	}

	PushObject(objType)
}

func PushStringNoStack(value string) {
	var stringArray []byte = StringTo1ByteArray(value)
	MovReg("x11", HP) // Guarda el inicio del string en x11
	for i := 0; i < len(stringArray); i++ {
		var charCode = stringArray[i]
		Comment(fmt.Sprintf("Pushing char %d to heap - (%c)", charCode, charCode))
		Mov(W0, int(charCode))
		StrB(W0, HP)
		Mov(X0, 1)
		Add(HP, HP, X0)
	}
	// Agregar terminador nulo
	Comment("Pushing null terminator to heap")
	Mov(W0, 0)
	StrB(W0, HP)
	Mov(X0, 1)
	Add(HP, HP, X0)
}

func PopObject(rd string) StackObject {
    if len(stack) == 0 {
        log.Printf("⚠️  ADVERTENCIA: Intento de PopObject() con pila vacía")
        return StackObject{Type: Int, Length: 8, Depth: 0, Id: nil}
    }
    
    obj := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    
    // Solo hacer Pop si rd no está vacío
    if rd != "" {
        if obj.Type == StackObjectType(Float) {
            PopFloat(rd)
        } else {
            Pop(rd)
        }
    }
    
    return obj
}

// Agregar función PopFloat si no existe:
func PopFloat(rd string) {
    instructions = append(instructions, fmt.Sprintf("ldr %s, [sp], #8", rd))
}

func IntObject() StackObject {
	return StackObject{
		Type:   Int, // O como tengas definido el tipo entero
		Length: 8,   // 8 bytes para un entero de 64 bits
		Depth:  depth,
		Id:     nil,
	}
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

func SliceObject(elemType StackObjectType, size int) StackObject {
	return StackObject{
		Type:     elemType,
		Length:   8,
		Depth:    depth,
		Id:       nil,
		IsSlice:  true,
		ElemType: elemType,
		Size:     size,
	}
}

func CloneObject(obj StackObject) StackObject {
	return StackObject{
		Type:     obj.Type,
		Length:   obj.Length,
		Depth:    obj.Depth,
		Id:       obj.Id,
		IsSlice:  obj.IsSlice,
		ElemType: obj.ElemType,
		Size:     obj.Size,
	}
}

func NewScope() {
	depth++
}

func EndScope() int {
    var byteOffset = 0

    for i := len(stack) - 1; i >= 0; i-- {
        if stack[i].Depth == depth {
            byteOffset += stack[i].Length
            stack = stack[:len(stack)-1]
        } else {
            break
        }
    }

    depth--
    return byteOffset
}

func TagObject(id string) {
	if len(stack) == 0 {
		log.Fatal("Error: intento de etiquetar objeto en pila vacía")
	}
	stack[len(stack)-1].Id = &id
}

func GetObject(id string) (int, StackObject) {
    var byteOffset = 0
    for i := len(stack) - 1; i >= 0; i-- {
        if stack[i].Id != nil && *stack[i].Id == id {
            return byteOffset, stack[i]
        }
        byteOffset += stack[i].Length
    }
    return -1, StackObject{}
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
	instructions = append(instructions, fmt.Sprintf("add %s, %s, #%d", rd, rs1, imm))
}

func Subi(rd string, rs1 string, imm int) {
	instructions = append(instructions, fmt.Sprintf("sub %s, %s, #%d", rd, rs1, imm))
}

func Neg(rd string, rs string) {
	instructions = append(instructions, fmt.Sprintf("neg %s, %s", rd, rs))
}

func And(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("and %s, %s, %s", rd, rs1, rs2))
}
func Orr(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("orr %s, %s, %s", rd, rs1, rs2))
}

func Cmp(rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("cmp %s, %s", rs1, rs2))
}

func CmpImm(rs string, imm int) {
	instructions = append(instructions, fmt.Sprintf("cmp %s, #%d", rs, imm))
}

func Cset(rd string, condition string) {
	instructions = append(instructions, fmt.Sprintf("cset %s, %s", rd, condition))
}

func Str(rs1 string, rs2 string, offset int) {
	instructions = append(instructions, fmt.Sprintf("str %s, [%s, #%d]", rs1, rs2, offset))
}

func StrF(rd string, base string, offset int) {
	instructions = append(instructions, fmt.Sprintf("str %s, [%s, #%d]", rd, base, offset))
}

func LdrF(rd string, base string, offset int) {
	instructions = append(instructions, fmt.Sprintf("ldr %s, [%s, #%d]", rd, base, offset))
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

// Float operations

func Scvtf(rd string, rs string) {
	instructions = append(instructions, fmt.Sprintf("scvtf %s, %s", rd, rs))
}

func FMov(rd string, rs string) {
	instructions = append(instructions, fmt.Sprintf("fmov %s, %s", rd, rs))
}

func FAdd(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("fadd %s, %s, %s", rd, rs1, rs2))
}

func FSub(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("fsub %s, %s, %s", rd, rs1, rs2))
}

func FMul(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("fmul %s, %s, %s", rd, rs1, rs2))
}

func FDiv(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("fdiv %s, %s, %s", rd, rs1, rs2))
}

func Svc() {
	instructions = append(instructions, "svc #0")
}

func EndProgram() {
	Mov(X0, 0)
	Mov(X8, 93) // syscall number for exit
	Svc()       // make syscall
}

func PrintInt(rs string) {
	// Mover el valor al registro x0 (parámetro para print_integer)
	if rs != X0 {
		MovReg(X0, rs)
	}

	instructions = append(instructions, "bl print_integer")

	Use("print_integer")
}

func PrintFloat() {
	Use("print_integer")
	Use("print_double")
	instructions = append(instructions, "bl print_double")
}

func PrintString(rs string) {
    if rs != X0 {
        MovReg(X0, rs)
    }

    usedSymbols["newline"] = true

    instructions = append(instructions, "bl print_string")
    Use("print_string")
}

func PrintChar(char rune) {
	label := fmt.Sprintf("char_%d", int(char))
	instructions = append(instructions, fmt.Sprintf("adr x1, %s", label)) // x1 = dirección del carácter
	instructions = append(instructions, "mov x0, #1")                     // x0 = stdout
	instructions = append(instructions, "mov x2, #1")                     // x2 = longitud
	instructions = append(instructions, "mov x8, #64")
	instructions = append(instructions, "svc #0")
	usedSymbols[label] = true
}

func PrintIntInline(rs string) {
	if rs != X0 {
		MovReg(X0, rs)
	}
	instructions = append(instructions, "bl print_integer_inline")
	Use("print_integer_inline")
}

func PrintFloatInline(rs string) {
	if rs != D0 {
		FMov(D0, rs) // fmov d0, rs
	}
	instructions = append(instructions, "bl print_double_inline")
	Use("print_double_inline")
}

func Comment(comment string) {
	instructions = append(instructions, fmt.Sprintf("# %s", comment))
}

// Branch operations

func BranchEq(label string) {
	instructions = append(instructions, fmt.Sprintf("beq %s", label))
}
func Branch(label string) {
	instructions = append(instructions, fmt.Sprintf("b %s", label))
}
func Label(label string) {
	instructions = append(instructions, fmt.Sprintf("%s:", label))
}

// Function management
func Ret() {
    instructions = append(instructions, "ret")
}

var functionDepth = 0

func StartFunction(name string) {
    Label(name)
    Comment(fmt.Sprintf("Inicio de función %s", name))
    // Setup frame pointer
    instructions = append(instructions, "stp x29, x30, [sp, #-16]!")
    instructions = append(instructions, "mov x29, sp")
    
    // NO incrementar depth aquí - los parámetros deben tener depth 0
    // depth se incrementa en NewScope() cuando sea necesario
}

func EndFunction(name string) {
    Comment(fmt.Sprintf("Fin de función %s", name))
    // Restore frame pointer
    instructions = append(instructions, "ldp x29, x30, [sp], #16")
    Ret()
    
    // Salir del contexto de función
    functionDepth--
}

func StartMain() {
    // Llamar a main desde _start
    instructions = append(instructions, "bl main")
    EndProgram()
}

func ToString() string {
    var sb strings.Builder
    sb.WriteString(".data\n")
    sb.WriteString("heap: .space 4096\n")

    // Agrega los símbolos usados aquí
    if usedSymbols["minus_sign"] {
        sb.WriteString("minus_sign: .ascii \"-\"\n")
    }
    if usedSymbols["newline"] {
        sb.WriteString("newline: .ascii \"\\n\"\n")
    }
    if usedSymbols["dot_char"] {
        sb.WriteString("dot_char: .ascii \".\"\n")
    }
    if usedSymbols["zero_char"] {
        sb.WriteString("zero_char: .ascii \"0\"\n")
    }
    if usedSymbols["double_newline"] {
        sb.WriteString("double_newline: .ascii \"\\n\"\n")
    }
    if usedSymbols["atoi_error_msg"] {
        sb.WriteString("atoi_error_msg: .ascii \"Error: entrada inválida en Atoi\\n\"\n")
    }

    // Agregar los caracteres individuales usados por PrintChar
    for symbol := range usedSymbols {
        if strings.HasPrefix(symbol, "char_") {
            charCode := strings.TrimPrefix(symbol, "char_")
            sb.WriteString(fmt.Sprintf("%s: .byte %s\n", symbol, charCode))
        }
    }

    sb.WriteString(".text\n")
    sb.WriteString(".global _start\n")
    
    // PRIMERO: El punto de entrada _start
    sb.WriteString("_start:\n")
    sb.WriteString("    adr x10, heap\n")
    sb.WriteString("    bl main\n")
    sb.WriteString("    mov x0, #0\n")
    sb.WriteString("    mov x8, #93\n")
    sb.WriteString("    svc #0\n\n")
    
    // DESPUÉS: Las instrucciones (funciones definidas por el usuario)
    for _, instr := range instructions {
        sb.WriteString(fmt.Sprintf("    %s\n", instr))
    }

    // FINALMENTE: Las funciones de la biblioteca estándar
    standardFunctions := GetFunctionDefinitions()
    if standardFunctions != "" {
        sb.WriteString("\n// Standard Library Functions\n")
        sb.WriteString(standardFunctions)
    }

    return sb.String()
}

func PrintStringInline(rs string) {
	if rs != X0 {
		MovReg(X0, rs)
	}
	instructions = append(instructions, "bl print_string_inline")
	Use("print_string_inline")
}

func Call(name string) {
	instructions = append(instructions, fmt.Sprintf("bl %s", name))
}

func UsedFunction(name string) {
	usedFunctions[name] = true
}
