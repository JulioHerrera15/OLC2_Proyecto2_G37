package arm

import (
	"fmt"
	"strings"

)

var instructions = []string{}

func Add(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("ADD %s, %s, %s", rd, rs1, rs2))
}

func Sub(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("SUB %s, %s, %s", rd, rs1, rs2))
}

func Mul(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("MUL %s, %s, %s", rd, rs1, rs2))
}

func Div(rd string, rs1 string, rs2 string) {
	instructions = append(instructions, fmt.Sprintf("DIV %s, %s, %s", rd, rs1, rs2))
}

func Addi(rd string, rs1 string, imm int) {
	instructions = append(instructions, fmt.Sprintf("ADDI %s, %s, #%d", rd, rs1, imm))
}

// Memory operations

func Str(rs1 string, rs2 string, offset int) {
	instructions = append(instructions, fmt.Sprintf("STR %s, [%s, #%d]", rs1, rs2, offset))
}

func Ldr(rd string, rs1 string, offset int) {
	instructions = append(instructions, fmt.Sprintf("LDR %s, [%s, #%d]", rd, rs1, offset))
}

func Mov(rd string, imm int) {
	instructions = append(instructions, fmt.Sprintf("MOV %s, #%d", rd, imm))
}

func MovReg(rd string, rs string) {
    instructions = append(instructions, fmt.Sprintf("MOV %s, %s", rd, rs))
}

func Push(rs string) {
    instructions = append(instructions, fmt.Sprintf("STR %s, [SP, #-8]!", rs))
}

func Pop(rd string) {
	instructions = append(instructions, fmt.Sprintf("LDR %s, [SP], #8", rd))
}

func Svc() {
	instructions = append(instructions, "SVC #0")
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
    instructions = append(instructions, "BL print_integer")
    
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