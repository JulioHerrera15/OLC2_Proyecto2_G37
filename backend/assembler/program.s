.data
heap: .space 4096
minus_sign: .ascii "-"
char_44: .byte 44
char_32: .byte 32
char_93: .byte 93
char_10: .byte 10
char_91: .byte 91
.text
.global _start
_start:
    adr x10, heap
    # Slice literal con 4 elementos
    mov x9, x10
    # Constant: 1
    mov x0, #1
    str x0, [sp, #-8]!
    ldr x0, [sp], #8
    str x0, [x10, #0]
    # Constant: 2
    mov x0, #2
    str x0, [sp, #-8]!
    ldr x0, [sp], #8
    str x0, [x10, #8]
    # Constant: 3
    mov x0, #3
    str x0, [sp, #-8]!
    ldr x0, [sp], #8
    str x0, [x10, #16]
    # Constant: 4
    mov x0, #4
    str x0, [sp, #-8]!
    ldr x0, [sp], #8
    str x0, [x10, #24]
    # Slice inicializado con 4 elementos de tipo Int
    add x10, x10, #32
    str x9, [sp, #-8]!
    # Print statement
    mov x0, #0
    add x0, sp, x0
    ldr x0, [x0, #0]
    str x0, [sp, #-8]!
    # Entrando a impresión de slice
    ldr x9, [sp], #8
    ldr x0, [sp], #8
    adr x0, char_91
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    ldr x0, [x9, #0]
    bl print_integer_inline
    adr x0, char_44
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    adr x0, char_32
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    ldr x0, [x9, #8]
    bl print_integer_inline
    adr x0, char_44
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    adr x0, char_32
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    ldr x0, [x9, #16]
    bl print_integer_inline
    adr x0, char_44
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    adr x0, char_32
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    ldr x0, [x9, #24]
    bl print_integer_inline
    adr x0, char_93
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    adr x0, char_10
    mov x1, #1
    mov x2, #1
    mov x8, #64
    svc #0
    mov x0, #0
    mov x8, #93
    svc #0
    mov x0, #0
    mov x8, #93
    svc #0

// Standard Library Functions

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
