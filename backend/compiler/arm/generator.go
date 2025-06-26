package arm

import (
	"fmt"
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
	Type   StackObjectType
	Length int
	Depth  int
	Id     *string
}

var instructions = []string{}
var stack = []StackObject{}
var depth = 0

func TopObject() StackObject {
	return stack[len(stack)-1]
}


func PushObject(obj StackObject) {
	stack = append(stack, obj)
}

func PushConstant(value interface{}, objType StackObject) {
	switch objType.Type {
	case Int:
		Mov(X0, value.(int))
		Push(X0)
	case Float:
		var floatBits int64 = int64(math.Float64bits(value.(float64)))
		var floatParts = [4]int16{}

		for i := 0; i < 4; i++ {
			floatParts[i] = int16((floatBits >> (i * 16)) & 0xFFFF)
		}
		instructions = append(instructions, fmt.Sprintf("movz x0, #%d, lsl #0", floatParts[0]))

		for i := 1; i < 4; i++ {
			instructions = append(instructions, fmt.Sprintf("movk x0, #%d, lsl #%d", floatParts[i], i*16))
		}

		Push(X0)


	case String:
		var stringArray []byte = StringTo1ByteArray(value.(string))
		Push(HP)
		for i := 0; i < len(stringArray); i++ {
			var charCode = stringArray[i]
			Comment(fmt.Sprintf("Pushing char %d to heap - (%c)", charCode, charCode))
			Mov(W0, int(charCode))
			StrB(W0, HP) // Guardar el byte en la posición actual del heap
			Mov(X0, 1)
			Add(HP, HP, X0)
		}
		// Agregar terminador nulo
		Comment("Pushing null terminator to heap")
		Mov(W0, 0)
		StrB(W0, HP)
		Mov(X0, 1)
		Add(HP, HP, X0)
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

func PopObject(rd string) StackObject {
	var obj = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
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
	instructions = append(instructions, fmt.Sprintf("addi %s, %s, #%d", rd, rs1, imm))
}

func Subi(rd string, rs1 string, imm int) {
	instructions = append(instructions, fmt.Sprintf("subi %s, %s, #%d", rd, rs1, imm))
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

func FAdd(rd string, rs1 string, rs2 string ) {
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

	instructions = append(instructions, "bl print_string")

	Use("print_string")
}

func Comment(comment string) {
	instructions = append(instructions, fmt.Sprintf("# %s", comment))
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

    sb.WriteString(".text\n")
    sb.WriteString(".global _start\n")
    sb.WriteString("_start:\n")
    sb.WriteString("    adr x10, heap\n")

    for _, instr := range instructions {
        sb.WriteString(fmt.Sprintf("    %s\n", instr))
    }

    sb.WriteString("    mov x0, #0\n")
    sb.WriteString("    mov x8, #93\n")
    sb.WriteString("    svc #0\n")

    standardFunctions := GetFunctionDefinitions()
    if standardFunctions != "" {
        sb.WriteString("\n// Standard Library Functions\n")
        sb.WriteString(standardFunctions)
    }

    return sb.String()
}

func Bl(label string) {
	instructions = append(instructions, fmt.Sprintf("bl %s", label))
}
