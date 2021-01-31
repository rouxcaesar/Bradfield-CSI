// Exercise 1-9 from K&R book.
// Write a program to copy its input to its output, replacing each string of one or more blanks by a single blank.

#include <stdio.h>

int main() {
  int c, b;
  b = 0;

  while ((c = getchar()) != EOF) {
    if (c == ' ') {
      if (b == 0) {
        b = 1;
      } else if (b > 0) {
        continue;
      }
    } else {
      b = 0;
    }

    putchar(c);
  }
}
