// This is a minimal close of the `ls` Unix program.
// Minimally, it will list the contents of a directory including
// some information about each file, such as file size.

// As a stretch goal, use man ls to identify any interesting
// flags you may wish to support, and implement them.

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
#include "dirent.h"

int main() {
  // Minimal ls constraints:
  //   - only consider current directory (none passed in)
  //   - no flags at this time
  //   - print to stdout the name of each file in directory
  //       - print each filename on a newline for now
  //   - next, include the file size of each file

  DIR *folder;
  struct dirent *entry;
  int files = 0;
  
  folder = opendir(".");

  if (folder == NULL)  {
    puts("Couldn't open directory\n");
    exit(1);
  }

  puts("Directory can be opened!");

  while ((entry = readdir(folder))) {
    files++;
    printf("File: %3d: %s\n", files, entry->d_name);
  }

  closedir(folder);
  return 0;
}
