package arm

import (
	"strings"
)

var usedFunctions = make(map[string]bool)
var usedSymbols = make(map[string]bool)

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

`,

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

    "print_string_inline": `
print_string_inline:
    stp     x29, x30, [sp, #-16]!
    stp     x19, x20, [sp, #-16]!
    mov     x19, x0

print_loop_inline:
    ldrb    w20, [x19]
    cbz     w20, print_done_inline

    mov     x0, #1
    mov     x1, x19
    mov     x2, #1
    mov     x8, #64
    svc     #0

    add     x19, x19, #1
    b       print_loop_inline

print_done_inline:
    ldp     x19, x20, [sp], #16
    ldp     x29, x30, [sp], #16
    ret
`,

    "print_double": `
    //--------------------------------------------------------------
// print_double - Prints a double precision float to stdout
//
// Input:
//   d0 - The double value to print
//--------------------------------------------------------------
print_double:
    // Save context
    stp x29, x30, [sp, #-16]!    
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    
    // Check if number is negative
    fmov x19, d0
    tst x19, #(1 << 63)       // Comprueba el bit de signo
    beq skip_minus

    // Print minus sign
    mov x0, #1
    adr x1, minus_sign
    mov x2, #1
    mov x8, #64
    svc #0

    // Make value positive
    fneg d0, d0

skip_minus:
    // Convert integer part
    fcvtzs x0, d0             // x0 = int(d0)
    bl print_integer_inline

    // Print dot '.'
    mov x0, #1
    adr x1, dot_char
    mov x2, #1
    mov x8, #64
    svc #0

    // Get fractional part: frac = d0 - float(int(d0))
    frintm d4, d0             // d4 = floor(d0)
    fsub d2, d0, d4           // d2 = d0 - floor(d0) (exact fraction)

    // Para 2.5, d2 debe ser exactamente 0.5

    // Multiplicar por 1_000_000 (6 decimales)
    movz x1, #0x000F, lsl #16
    movk x1, #0x4240, lsl #0   // x1 = 1000000
    scvtf d3, x1              // d3 = 1000000.0
    fmul d2, d2, d3           // d2 = frac * 1_000_000
    
    // Redondear al entero más cercano para evitar errores de precisión
    frintn d2, d2             // d2 = round(d2)
    fcvtzs x0, d2             // x0 = int(d2)

    // Imprimir ceros a la izquierda si es necesario
    mov x20, x0               // x20 = fracción entera
    movz x21, #0x0001, lsl #16
    movk x21, #0x86A0, lsl #0  // x21 = 100000
    mov x22, #0               // inicializar contador de ceros
    mov x23, #10              // constante para división

leading_zero_loop:
    udiv x24, x20, x21        // x24 = x20 / x21
    cbnz x24, done_leading_zeros  // Si hay un dígito no cero, salir del bucle

    // Imprimir '0'
    mov x0, #1
    adr x1, zero_char
    mov x2, #1
    mov x8, #64
    svc #0

    udiv x21, x21, x23        // x21 /= 10
    add x22, x22, #1          // incrementar contador de ceros
    cmp x21, #0               // verificar si llegamos al final
    beq print_remaining       // si divisor es 0, saltar a imprimir el resto
    b leading_zero_loop

done_leading_zeros:
    // Print the remaining fractional part
    mov x0, x20
    bl print_integer_inline
    b exit_function

print_remaining:
    // Caso especial cuando la parte fraccionaria es 0 después de imprimir ceros
    cmp x20, #0
    bne exit_function
    
    // Ya imprimimos todos los ceros necesarios
    // No hace falta imprimir nada más

exit_function:
    // Print newline
    mov x0, #1
    adr x1, double_newline
    mov x2, #1
    mov x8, #64
    svc #0

    // Restore context
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret

`,

    "print_integer_inline": `
print_integer_inline:
    // Save registers
    stp x29, x30, [sp, #-16]!
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    
    // Save original number
    mov x19, x0
    
    // Check if negative
    cmp x19, #0
    bge positive_number_inline
    
    // Print minus sign for negative numbers
    mov x0, #1
    adr x1, minus_sign
    mov x2, #1
    mov x8, #64
    svc #0
    
    neg x19, x19

positive_number_inline:
    // Allocate buffer on stack
    sub sp, sp, #32
    mov x22, sp
    mov x23, #0
    
    // Handle zero case
    cmp x19, #0
    bne convert_loop_inline
    
    mov w24, #48
    strb w24, [x22]
    mov x23, #1
    b print_result_inline

convert_loop_inline:
    mov x24, #10
    udiv x25, x19, x24
    msub x26, x25, x24, x19
    add x26, x26, #48
    strb w26, [x22, x23]
    add x23, x23, #1
    mov x19, x25
    cbnz x19, convert_loop_inline
    
    // Reverse the string
    mov x24, #0
    sub x25, x23, #1
reverse_loop_inline:
    cmp x24, x25
    bge print_result_inline
    ldrb w26, [x22, x24]
    ldrb w27, [x22, x25]
    strb w27, [x22, x24]
    strb w26, [x22, x25]
    add x24, x24, #1
    sub x25, x25, #1
    b reverse_loop_inline

print_result_inline:
    // Print the number
    mov x0, #1
    mov x1, x22
    mov x2, x23
    mov x8, #64
    svc #0
    
    // Restore stack and registers
    add sp, sp, #32
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret
`,

"print_double_inline": `
print_double_inline:
    // Save context
    stp x29, x30, [sp, #-16]!    
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!
    stp x23, x24, [sp, #-16]!
    
    // Check if number is negative
    fmov x19, d0
    tst x19, #(1 << 63)       // Comprueba el bit de signo
    beq skip_minus_inline

    // Print minus sign
    mov x0, #1
    adr x1, minus_sign
    mov x2, #1
    mov x8, #64
    svc #0

    // Make value positive
    fneg d0, d0

skip_minus_inline:
    // Convert integer part
    fcvtzs x0, d0             // x0 = int(d0)
    bl print_integer_inline

    // Print dot '.'
    mov x0, #1
    adr x1, dot_char
    mov x2, #1
    mov x8, #64
    svc #0

    // Get fractional part: frac = d0 - float(int(d0))
    frintm d4, d0             // d4 = floor(d0)
    fsub d2, d0, d4           // d2 = d0 - floor(d0) (exact fraction)

    // Multiplicar por 1_000_000 (6 decimales)
    movz x1, #0x000F, lsl #16
    movk x1, #0x4240, lsl #0   // x1 = 1000000
    scvtf d3, x1              // d3 = 1000000.0
    fmul d2, d2, d3           // d2 = frac * 1_000_000
    
    // Redondear al entero más cercano para evitar errores de precisión
    frintn d2, d2             // d2 = round(d2)
    fcvtzs x0, d2             // x0 = int(d2)

    // Imprimir ceros a la izquierda si es necesario
    mov x20, x0               // x20 = fracción entera
    movz x21, #0x0001, lsl #16
    movk x21, #0x86A0, lsl #0  // x21 = 100000
    mov x22, #0               // inicializar contador de ceros
    mov x23, #10              // constante para división

leading_zero_loop_inline:
    udiv x24, x20, x21        // x24 = x20 / x21
    cbnz x24, done_leading_zeros_inline  // Si hay un dígito no cero, salir del bucle

    // Imprimir '0'
    mov x0, #1
    adr x1, zero_char
    mov x2, #1
    mov x8, #64
    svc #0

    udiv x21, x21, x23        // x21 /= 10
    add x22, x22, #1          // incrementar contador de ceros
    cmp x21, #0               // verificar si llegamos al final
    beq print_remaining_inline       // si divisor es 0, saltar a imprimir el resto
    b leading_zero_loop_inline

done_leading_zeros_inline:
    // Print the remaining fractional part
    mov x0, x20
    bl print_integer_inline
    b exit_function_inline

print_remaining_inline:
    // Caso especial cuando la parte fraccionaria es 0 después de imprimir ceros
    cmp x20, #0
    bne exit_function_inline
    
    // Ya imprimimos todos los ceros necesarios
    // No hace falta imprimir nada más

exit_function_inline:
    // NO imprime salto de línea aquí

    // Restore context
    ldp x23, x24, [sp], #16
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret
`,

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

    switch function {
    case "print_integer":
        usedSymbols["minus_sign"] = true
        usedSymbols["newline"] = true
    case "print_double":
        usedSymbols["dot_char"] = true
        usedSymbols["zero_char"] = true
        usedSymbols["double_newline"] = true
        usedSymbols["minus_sign"] = true
        usedFunctions["print_integer_inline"] = true
    case "print_integer_inline":
        usedSymbols["minus_sign"] = true
    case "print_double_inline":
        usedSymbols["dot_char"] = true
        usedSymbols["zero_char"] = true
        usedSymbols["minus_sign"] = true
        usedFunctions["print_integer_inline"] = true
    }
}

func GetFunctionDefinitions() string {
    var functions []string

    for function := range usedFunctions {
        if definition, exists := functionDefinitions[function]; exists {
            functions = append(functions, definition)
        }
    }

    return strings.Join(functions, "\n")
}

var Symbols = map[string]string{
    "minus_sign":    `.ascii "-"`,
    "dot_char":      `.ascii "."`,
    "zero_char":     `.ascii "0"`,
    "newline":       `.ascii "\n"`,
    "double_newline": `.ascii "\n"`,
    "char_91": `.ascii "["`,
    "char_93": `.ascii "]"`,
    "char_44": `.ascii ","`,
    "char_32": `.ascii " "`,
    "char_10": `.ascii "\n"`,
}