#include <stdio.h>
#include <math.h>

int htoi(char []);


int main() {
  char string[1024];
  char c;
  int count = 0;
  while((c = getchar()) != EOF) {
    if (c == '\n') {
      string[count] = '\0';
      printf("Decimal value for hexadecimal   %s   is     %d \n\n\n", string, htoi(string));
      count = 0;
      continue;
    }
    string[count++] = c;
  }

  return 0;
}


int htoi(char s[]) {
  int i, result;
  char c;

  result = 0;
  for(i = 0; (c = s[i]) != '\0'; i++) {
    result *= 16;
    if (i == 1 && (c == 'x' || c == 'X'))
      ;
    else if (c >= '0' && c <= '9')
      result += (c - '0');
    else if(c >= 'a' && c <= 'f')
      result += (10 + (c - 'a'));
    else if(c >= 'A' && c <= 'F')
      result += (10 + (c - 'A'));
    else
      return -1;

  }
  return result;
}
