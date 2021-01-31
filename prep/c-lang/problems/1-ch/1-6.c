// If EOF is signaled with either Control + D or Control + Z,
// then "getchar() = 0" will be printed.
// Otherwise, any other key including Enter will print
// "getchar() =  1".
#include <stdio.h>

int main() {
  int c;
  printf("getchar() = %d\n", getchar() != EOF);
}
