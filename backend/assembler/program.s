.text
.global _start
_start:
    // Print statement
    // Constant: 2
    MOV x0, #2
    STR x0, [SP, #-8]!
    // Constant: 1
    MOV x0, #1
    STR x0, [SP, #-8]!
    LDR x1, [SP], #8
    LDR x0, [SP], #8
    SUB x0, x0, x1
    STR x0, [SP, #-8]!
    LDR x0, [SP], #8
    BL print_integer
    MOV x0, #0
    MOV x8, #93
    SVC #0



// Standard Library
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
    .ascii "\n"