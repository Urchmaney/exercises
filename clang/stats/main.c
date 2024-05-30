#include <stdio.h>

#define MAXNUMBERS 1000

double numbers[MAXNUMBERS];
int indexcount;

int fetchvalues(void);

int main() {
    int result = fetchvalues();
    int c = 0;

    while (c < indexcount)
        printf("%f\n", numbers[c++]);
    return result;
}

int fetchvalues() {
    double decimal = 0, dprecision = 0.1, number = 0, precision = 10;
    int afterpoint = 0;
    char c;

    while((c = getchar()) != EOF) {
        if (c <= '0' && c >= '9' && c != '.' && c != '\n')
            return -1;

        if (c == '.') {
            afterpoint = 1;
            continue;
        }
            
        if (c == '\n') {
            numbers[indexcount] = number + decimal;
            indexcount++;
            decimal = 0;
            dprecision = 0.1;
            number = afterpoint = 0;
            continue;
        }

        if (afterpoint < 1){
            number *= precision;
            number += (c - '0');
            // printf("%d\n", number);
        }
        
        if (afterpoint == 1) {
            decimal += dprecision * (c - '0');
            dprecision /= 10;
            // printf("%f\n", decimal);
        }
            
    }

    return 0;
}