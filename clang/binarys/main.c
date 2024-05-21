
#include <stdio.h>

unsigned getbits(unsigned, int, int);

unsigned setbits(unsigned, int, int, unsigned);

int main() {
    
    unsigned int val = (unsigned int)~0 >> 1;

    unsigned s = ~63 >> 1;

    printf("Largest int value is %d  \n", setbits(49, 7, 5, 77));
    return 0;
}

unsigned getbits(unsigned x, int p, int n) {
    return (x >> (p + 1 - n)) & ~(~0 << n);
}

unsigned setbits(unsigned x, int p, int n, unsigned y) {
    unsigned savedx = x & ~(~0 << (p + 1 - n));
    unsigned pully = y & ~(~0 << n);
    unsigned clearx = ((x >> p) << n) & pully;
    return (clearx << (p + 1 - n)) & savedx;
}