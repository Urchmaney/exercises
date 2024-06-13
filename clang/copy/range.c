#include <limits.h>
#include <stdio.h>

enum months { JAN = 1, FEB, MAR, APR, MAY, JUN, JUL, AUG, SEP, OCT, NOV, DEC };

const int prime = 3;

int main() {
    long result;

    #define PLAUR 12

    result = INT_MAX - (unsigned int)INT_MIN;
    printf("Range for Int is  %ld %d\n", result, INT_MIN);
    printf("Last month is %d  %d\n", DEC, prime);
}