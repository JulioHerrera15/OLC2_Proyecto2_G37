package arm

import (
	"fmt"
	"strings"

)

var instructions = []string{}

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

func Comment(comment string) {
	instructions = append(instructions, fmt.Sprintf("// %s", comment))
}

func ToString() string {
    var sb strings.Builder
    sb.WriteString(".text\n")
    sb.WriteString(".global _start\n")
    sb.WriteString("_start:\n")

    // ✅ Simplemente imprimir todas las instrucciones
    for _, instr := range instructions {
        sb.WriteString(fmt.Sprintf("    %s\n", instr))
    }
    
    // ✅ Funciones estándar
    standardFunctions := GetFunctionDefinitions()
    if standardFunctions != "" {
        sb.WriteString("\n\n\n// Standard Library\n")
        sb.WriteString(standardFunctions)
    }

    return sb.String()
}