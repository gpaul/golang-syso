/* data.S */
.globl main.Awesome
main.Awesome: /* make a slice */
    .long data /* pointer */
    .long data_end - data /* len */
    .long data_end - data /* cap */
data:
.incbin "data"
data_end:
