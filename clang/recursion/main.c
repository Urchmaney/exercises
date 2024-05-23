#include <stdio.h>

void printd(int n) {
    if (n < 0) {
        putchar('-');
        n = -n;
    }
    if (n / 10) 
        printd(n / 10);
    putchar(n % 10 + '0');
}

void reverse(char* start, int size) {
    if (size <= 0)
        return;
    char temp;
    temp = *start;
    *start = *(start + size);
    *(start + size) = temp;
    reverse(start + 1, size - 2);
}

int main() {
    
    char string[] = "Civic";
    int count = 0;

    while(string[count++] != '\0');
    printf("Welcome to recursion \n");
    reverse(string, count - 2);
    
    printf("%s", string);
    printf("\n");
    return 0;
}