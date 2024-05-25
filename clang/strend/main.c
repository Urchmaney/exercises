#include <stdio.h>

#define MAXLINE 20

int strend(char *, char *);

int main() {
    char c;
    char firstline[MAXLINE], secondline[MAXLINE];
    int fcount = 0, scount = 0, count = 0;
    int result = 0;

    while((c = getchar()) != EOF) {
        if (c == '\n') {
            if(++count > 1) {
                secondline[scount] = '\0';
                count = 0;
                fcount = 0;
                scount = 0;
                result = strend(firstline, secondline);
                printf("\n%d    :  %s ends the first sentence. \n\n", result, result ? "Yes it" : "No it doesn't");
            }
            else{
                firstline[fcount] = '\0';
            }

        }
        else if (count == 0) {
            firstline[fcount++] = c;
        }
        else if (count == 1) {
            secondline[scount++] = c;
        }
    }
    
    return 0;
}

int strend(char *s, char *t) {
    char *temp = t;
    for(; *s != '\0'; s++) {
        if (*s == *temp)
            temp++;
        else if (temp > t) {
            temp = t;
        }
    }
    return *temp == '\0';
}