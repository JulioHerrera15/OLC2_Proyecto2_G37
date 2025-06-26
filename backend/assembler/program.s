.text
.global _start

_start:
    // Inicializa el stack pointer
    LDR x1, =stack_top
    MOV SP, x1
    
    // Suma 2 + 3 = 5
    MOV x0, #2
    MOV x1, #3
    ADD x0, x0, x1        // x0 = 5
    
    // Llamar a print_integer(x0)
    BL print_integer
    
    // exit(0)
    MOV x0, #0
    MOV x8, #93
    SVC #0

print_integer:
    // Guardar registros
    stp x29, x30, [sp, #-16]!
    stp x19, x20, [sp, #-16]!
    stp x21, x22, [sp, #-16]!
    
    mov x19, x0            // Guardar número original
    
    // Verificar si es negativo
    cmp x19, #0
    bge positive_number
    
    // Si negativo, imprimir '-' y convertir a positivo
    mov x0, #1             // stdout
    adr x1, minus_sign
    mov x2, #1
    mov x8, #64            // sys_write
    svc #0
    neg x19, x19           // hacer positivo

positive_number:
    // Reservar espacio para buffer en el stack
    sub sp, sp, #32
    mov x20, sp            // x20 = puntero al buffer
    mov x21, #0            // x21 = contador de dígitos
    
    // Caso especial: si el número es 0
    cmp x19, #0
    bne convert_loop
    
    mov w22, #48           // ASCII '0'
    strb w22, [x20]
    mov x21, #1
    b print_digits

convert_loop:
    cbz x19, reverse_digits // Si x19 == 0, terminar
    
    mov x22, #10
    udiv x23, x19, x22     // x23 = x19 / 10
    msub x24, x23, x22, x19 // x24 = x19 % 10 (resto)
    
    add x24, x24, #48      // Convertir a ASCII
    strb w24, [x20, x21]   // Guardar dígito
    add x21, x21, #1       // Incrementar contador
    
    mov x19, x23           // x19 = cociente
    b convert_loop

reverse_digits:
    // Los dígitos están al revés, necesitamos invertirlos
    mov x22, #0            // índice izquierdo
    sub x23, x21, #1       // índice derecho

reverse_loop:
    cmp x22, x23
    bge print_digits
    
    ldrb w24, [x20, x22]   // cargar byte izquierdo
    ldrb w25, [x20, x23]   // cargar byte derecho
    strb w25, [x20, x22]   // guardar byte derecho en posición izquierda
    strb w24, [x20, x23]   // guardar byte izquierdo en posición derecha
    
    add x22, x22, #1
    sub x23, x23, #1
    b reverse_loop

print_digits:
    // Imprimir los dígitos
    mov x0, #1             // stdout
    mov x1, x20            // buffer
    mov x2, x21            // longitud
    mov x8, #64            // sys_write
    svc #0
    
    // Imprimir nueva línea
    mov x0, #1
    adr x1, newline
    mov x2, #1
    mov x8, #64
    svc #0
    
    // Restaurar stack y registros
    add sp, sp, #32
    ldp x21, x22, [sp], #16
    ldp x19, x20, [sp], #16
    ldp x29, x30, [sp], #16
    ret

.data
minus_sign:
    .ascii "-"
newline:
    .ascii "\n"

.section .bss
    .align 16
stack:
    .skip 0x40000
stack_top: