#include <stdio.h>

int main() {
  printf("hello, ");
  printf("world");
  printf("\n");   // prints a newline
  printf("\t");   // prints a tab
  printf("\c");   // prints the char "c"
  printf("\e");   // doesn't print anything; compiler issues warning
  printf("\w");   // doesn't print anything; compiler issues warning
}
