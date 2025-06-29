aarch64-linux-gnu-as -mcpu=cortex-a57 program.s -o program.o
aarch64-linux-gnu-ld program.o -o program
qemu-aarch64 -g 1234 ./program
gdb-multiarch -q --nh \
    -ex 'set architecture aarch64' \
    -ex 'file program' \
    -ex 'target remote localhost:1234' \
    -ex 'layout split' \
    -ex 'layout regs'
qemu-aarch64 ./program