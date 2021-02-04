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
#include <sysexits.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/file.h>
#include <unistd.h>
#include "dirent.h"

// flags struct will track which flags the user has passed in
struct flags {
  bool all;
  bool time;
};

int process_args(int argc, char *argv[], struct flags *f, char *dir);

int print_files(DIR *folder, struct flags *f, char *dir);

int main(int argc, char *argv[]) {
  char dir[100];
  DIR *folder;
  struct flags f;

  if (argc > 1) {
    // If user has passed the "--help" option, print out the help message
    // and then immediately exit.
    if (strcmp(argv[1], "--help") == 0) {
      puts("");
      printf("./ls-clone [flags] [target-directory]\n");
      puts("");
      printf("ls-clone - A minimal clone of the ls tool.\n");
      puts("");
      printf("Flags:\n");
      puts("");
      printf("-a - Output all files in target directory including dotfiles\n");
      printf("-t - Output all files sorted by time of last modification in descending order\n");
      puts("");
      exit(EXIT_SUCCESS);
    } else {
      // If user has passed in more than one argument, process the arguments to
      // properly consider flags and target directory.
      if (process_args(argc, argv, &f, dir) != 0) {
        exit(EXIT_FAILURE);
      }
    }
  } else {
    // In the case of no arguments other than the program name ("./ls-clone"),
    // set dir to the current working directory.
    strcpy(dir, ".");
  }

  folder = opendir(dir);
  if (folder == NULL)  {
    puts("Couldn't open directory\n");
    exit(EXIT_FAILURE);
  }

  if ((print_files(folder, &f, dir) != 0)) {
    exit(EXIT_FAILURE);
  }
  closedir(folder);

  exit(EXIT_SUCCESS);
}

// process_args will process each of the arguments provided to the program.
// Based on the arguments, the function will properly assign the name of
// the directory to the *dir variable and set the values in the flags struct *f.
int process_args(int argc, char *argv[], struct flags *f, char *dir) {
  int i;
  char arg[100];

  for (i = 1; i < argc; i++) {
    strcpy(arg, argv[i]);

    if (arg[0] == '-') {
      switch (arg[1]) {
        case 'a':
          f->all = true;
          break;
        case 't':
          f->time = true;
          break;
        default:
          return 1;
      } 
    } else if ((arg[0] == '.') || (arg[0] == '/')) {
      strcpy(dir, arg);
    }

    memset(arg, 0, sizeof arg);
  }

  return 0;
}

// print_files takes a DIR instance and considers each file in the DIR.
// For each file, the function prints its the filesize and name.
int print_files(DIR *folder, struct flags *f, char *dir) {
  struct dirent *entry;
  int fd;
  struct stat buf;
  off_t size;
  bool use_path;
  char rel_path[100];

  if (strcmp(dir, ".") != 0) {
    use_path = true;
  }

  while ((entry = readdir(folder))) {
    if (!(f->all)) {
      if (((strcmp(entry->d_name, ".")) == 0) || ((strcmp(entry->d_name, "..")) == 0)) {
        continue;
      } else if (entry->d_name[0] == '.') {
        continue;
      }
    }

    lstat(entry->d_name, &buf);
    if (!S_ISREG(buf.st_mode)) {
      // Need to properly determine size of directories.
      // Likely need to sum the size of all files in subdirectories recursively.
      // Best to extract that logic into a separate function.
      size = buf.st_size;
    } else {
      if (use_path) {
        strcat(rel_path, dir);
        strcat(rel_path, entry->d_name);

        if ((fd = open(rel_path, O_RDONLY, 0)) == -1) {
          printf("Can't open %s\n", rel_path);
          return 1;
        }

      } else {
        if ((fd = open(entry->d_name, O_RDONLY, 0)) == -1) {
          printf("Can't open %s\n", entry->d_name);
          return 1;
        }
      }

      fstat(fd, &buf);
      size = buf.st_size;
    }

    printf("%lld\t%s\n", size, entry->d_name);
    memset(rel_path, 0, sizeof rel_path);
  }

  return 0;
}
