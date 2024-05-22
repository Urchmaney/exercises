#include <stdio.h>

#define MAXLINE 1000

int getlinen(char [], int);

int strindex(char [], char []);

static char pattern[] = "ould";

int main() {
    char line[MAXLINE];

    int found = 0;
    int p = -1;

    while(getlinen(line, MAXLINE) > 0) {
        if ((p = strindex(line, pattern)) >= 0) {
            printf("\n%s   %d \n", line, p);
            found++;
        }
    }

    printf("Found number is  %d\n", found);
    
    return found;
}


int getlinen(char s[], int lim) {
    int c, i;

    i = 0;

    while(--lim > 0 && (c = getchar()) != EOF && c != '\n')
        s[i++] = c;

    if (c == '\n') 
        s[i++] = c;

    s[i] = '\0';

    return i;
}

int strindex(char s[], char t[]) {
    int i, j, k, r;

    r = -1;

    for(i = 0; s[i] != '\0'; i++) {
        for (j = i, k = 0; t[k] != '\0' && s[j] == t[k]; j++, k++)
        ;
        if (k > 0 && t[k] == '\0')
            r = i;
    }

    return r;
}

