// Exercise 7-1 from K&R book

#include <stdio.h>
#include <ctype.h>
#include <stdbool.h>
#include <string.h>

int main(int argc, char *argv[]) { 
  bool to_lower = false;
  int c;

  if (strcmp(argv[0], "./lower") == 0) {
    to_lower = true;
  }

  while ((c = getchar()) != EOF) {
    if (to_lower) {
      putchar(tolower(c));
    } else {
      putchar(toupper(c));
    }
  }
  
  return 0;
}
