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

func PushObject(obj StackObject) {
	stack = append(stack, obj)
}

func PushConstant(value interface{}, objType StackObject) {
	switch objType.Type {
	case Int:
		Mov(X0, value.(int))
		Push(X0)
	case Float:
		//TODO
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
	sb.WriteString(".text\n")
	sb.WriteString(".global _start\n")
	sb.WriteString("_start:\n")
	sb.WriteString("    adr x10, heap\n")

	for _, instr := range instructions {
		sb.WriteString(fmt.Sprintf("    %s\n", instr))
	}

	// Agregar llamada a EndProgram si no se ha agregado manualmente
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
