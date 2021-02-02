// This is a minimal close of the `ls` Unix program.
// Minimally, it will list the contents of a directory including
// some information about each file, such as file size.

// Approach:
// 1) Process argc and argv values to determine correct action(s).
// 2) If (argc == 1), then we handle the current directory.
// 3) If (argc > 1), then we have flags or a different directory to handle.
// 4) For flags, we can process all the flags at once and set
//    corresponding variables to 0 or 1 (bitfields), then
//    reference them later in the program.
// 5) Once specified directory is known, begin by opening the
//    directory (it's just a file).
// 6) Read the directory contents (contained files/directories).

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/file.h>
#include <unistd.h>
#include "dirent.h"

struct flags {
  bool a;
};

void print_files(DIR *folder, struct flags f);

int main(int argc, char *argv[]) {
  char *dir;
  DIR *folder;
  struct flags f;

  // This can be a separate function.
  // Extend into parsing of command args (flags and dir name).
  if (argc > 1) {
    dir = argv[1];
  } else {
    dir = ".";
  }

  //if (argc > 2) {
  //}

  folder = opendir(dir);
  if (folder == NULL)  {
    puts("Couldn't open directory\n");
    exit(1);
  }

  print_files(folder, f);
  closedir(folder);

  return 0;
}

// print_files takes a DIR instance and considers each file in the DIR.
// For each file, the function prints its the filesize and name.
void print_files(DIR *folder, struct flags f) {
  struct dirent *entry;
  int fd;
  struct stat buf;
  off_t size;

  while ((entry = readdir(folder))) {
    if (!(f.a)) {
      if (((strcmp(entry->d_name, ".")) == 0) || ((strcmp(entry->d_name, "..")) == 0)) {
        continue;
      } else if (entry->d_name[0] == '.') {
        continue;
      }
    }

    if ((fd = open(entry->d_name, O_RDONLY, 0)) == -1) {
      printf("Can't open %s\n", entry->d_name);
    }

    fstat(fd, &buf);
    size = buf.st_size;
    
    printf("%lld\t%s\n", size, entry->d_name);
  }
}
