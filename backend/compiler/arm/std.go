package arm

import (
	"strings"
)

var usedFunctions = make(map[string]bool)

var functionDefinitions = map[string]string{
	"print_integer": `
print_integer:
    // Save registers
    stp x29, x30, [sp, #-16]!
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    
    // Save original number
    mov x19, x0
    
    // Check if negative
    cmp x19, #0
    bge positive_number
    
    // Print minus sign for negative numbers
    mov x0, #1
    adr x1, minus_sign
    mov x2, #1
    mov x8, #64
    svc #0
    
    neg x19, x19

positive_number:
    // Allocate buffer on stack
    sub sp, sp, #32
    mov x22, sp
    mov x23, #0
    
    // Handle zero case
    cmp x19, #0
    bne convert_loop
    
    mov w24, #48
    strb w24, [x22]
    mov x23, #1
    b print_result

convert_loop:
    mov x24, #10
    udiv x25, x19, x24
    msub x26, x25, x24, x19
    add x26, x26, #48
    strb w26, [x22, x23]
    add x23, x23, #1
    mov x19, x25
    cbnz x19, convert_loop
    
    // Reverse the string
    mov x24, #0
    sub x25, x23, #1
reverse_loop:
    cmp x24, x25
    bge print_result
    ldrb w26, [x22, x24]
    ldrb w27, [x22, x25]
    strb w27, [x22, x24]
    strb w26, [x22, x25]
    add x24, x24, #1
    sub x25, x25, #1
    b reverse_loop

print_result:
    // Print the number
    mov x0, #1
    mov x1, x22
    mov x2, x23
    mov x8, #64
    svc #0
    
    // Print newline
    mov x0, #1              // stdout
    adr x1, newline         // address of newline
    mov x2, #1              // length
    mov x8, #64             // sys_write
    svc #0
    
    // Restore stack and registers
    add sp, sp, #32
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret

minus_sign:
    .ascii "-"
newline:
    .ascii "\n"`,

	"print_string": `
print_string:
    // Save link register and other registers we'll use
    stp     x29, x30, [sp, #-16]!
    stp     x19, x20, [sp, #-16]!
    
    // x19 will hold the string address
    mov     x19, x0
    
print_loop:
    // Load a byte from the string
    ldrb    w20, [x19]
    
    // Check if it's the null terminator (0)
    cbz     w20, print_done
    
    // Prepare for write syscall
    mov     x0, #1              // File descriptor: 1 for stdout
    mov     x1, x19             // Address of the character to print
    mov     x2, #1              // Length: 1 byte
    mov     x8, #64             // syscall: write (64 on ARM64)
    svc     #0                  // Make the syscall
    
    // Move to the next character
    add     x19, x19, #1
    
    // Continue the loop
    b       print_loop
    
print_done:
    mov     x0, #1
    adr     x1, string_newline
    mov     x2, #1
    mov     x8, #64
    svc     #0
    
    // Restore saved registers
    ldp     x19, x20, [sp], #16
    ldp     x29, x30, [sp], #16
    ret

string_newline:
    .ascii "\n"`,

	"atoi": `
atoi:
    // x0 = dirección de string
    mov x1, #0              // acumulador

atoi_loop:
    ldrb w2, [x0], #1       // cargar byte, avanzar puntero
    cmp w2, #0
    beq atoi_done           // fin si es nulo

    cmp w2, #'.'
    beq atoi_error

    cmp w2, #'0'
    blt atoi_error
    cmp w2, #'9'
    bgt atoi_error

    uxtb x2, w2             // extender byte sin signo a 64 bits
    sub x2, x2, #'0'        // convertir ASCII a número

    mov x3, #10
    mul x1, x1, x3
    add x1, x1, x2

    b atoi_loop

atoi_done:
    mov x0, x1              // devolver resultado en x0
    ret

atoi_error:
    mov x0, #1
    adr x1, atoi_error_msg
    mov x2, #27
    mov x8, #64
    svc #0

    mov x0, #1
    mov x8, #93
    svc #0

atoi_error_msg:
    .ascii "Error: entrada inválida en Atoi\n"
`,
}

func Use(function string) {
	usedFunctions[function] = true
}

func GetFunctionDefinitions() string {
	var functions []string

	for function := range usedFunctions {
		if definition, exists := functionDefinitions[function]; exists {
			functions = append(functions, definition)
		}
	}

	return strings.Join(functions, "\n\n")
}
