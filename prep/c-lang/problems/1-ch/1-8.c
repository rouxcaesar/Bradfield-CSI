// Exercise 1-8 from K&R book.
// Write a program to count blanks, tabs, and newlines.

#include <stdio.h>

int main() {
  int c, nl, b, t;
  c = nl = b = t = 0;
  //printf("c: %d, nl: %d, t: %d, b: %d\n", c, nl, t, b);

  nl = 0;
  while ((c = getchar()) != EOF) {
    if (c == '\n') {
      ++nl;
    } else if (c == '\t') {
      ++t;
    } else if (c == ' ') {
      ++b;
    }
  }
  
  printf("\n");
  printf("Newlines: %d\n", nl);
  printf("Tabs: %d\n", t);
  printf("Blanks: %d\n", b);
}
