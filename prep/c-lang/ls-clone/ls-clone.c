// This is a minimal close of the `ls` Unix program.
// Minimally, it will list the contents of a directory including
// some information about each file, such as file size.

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

// flags struct will track which flags the user has passed in.
struct flags {
  bool all;
  bool file_size;
};

// process_args will process any flags and target directory provided by the user.
int process_args(int argc, char *argv[], struct flags *f, char *dir);

// print_files will open the provided directory file and print each of the files
// contained in the directory.
// Output will be modified depending on values set in the flags struct argument.
int print_files(DIR *folder, struct flags *f, char *dir);

int main(int argc, char *argv[]) {
  char dir[100];
  DIR *folder;
  struct flags f;
  f.all = f.file_size = false;

  if (argc > 1) {
    // If user has passed the "--help" option, print out the help message
    // and then immediately exit.
    if (strcmp(argv[1], "--help") == 0) {
      puts("");
      puts("./ls-clone [flags] [target-directory]");
      puts("");
      puts("ls-clone - A minimal clone of the ls tool.");
      puts("");
      puts("Flags: Each separated by a space and prefixed with a '-'");
      puts("");
      puts("-a - Output all files in target directory including dotfiles");
      puts("-t - Output all files sorted by time of last modification in descending order");
      puts("");
      puts("Example: ./ls-clone -a -f ./");
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
        case 'f':
          f->file_size = true;
          break;
        default:
          return 1;
      } 
    } else if ((arg[0] == '.') || (arg[0] == '/')) {
      strcpy(dir, arg);
    }

    memset(arg, 0, sizeof arg);
  }

  if (strchr(dir, '.') == NULL) {
    strcpy(dir, ".");
  }

  return 0;
}

// print_files takes an opened directory instance and considers each file in the directory.
// For each file, the function will print out the file in accordance to the values of
// any flags provided by the user.
int print_files(DIR *folder, struct flags *f, char *dir) {
  struct dirent *entry;
  int fd;
  struct stat buf;
  off_t size;
  bool use_path;
  char rel_path[100];

  // If user has provided a target directory, we set a feature flag to ensure control
  // flow down a branch that sets up the relative path for each file to be printed.
  // The relative path needs to be constructed for proper execution of opening each file.
  if (strcmp(dir, ".") != 0) {
    use_path = true;
  }

  while ((entry = readdir(folder))) {
    // If user has not passed in the '-a' flag, then we ignore all hidden or dotfiles files.
    if (!(f->all)) {
      if (((strcmp(entry->d_name, ".")) == 0) || ((strcmp(entry->d_name, "..")) == 0)) {
        continue;
      } else if (entry->d_name[0] == '.') {
        continue;
      }
    }

    lstat(entry->d_name, &buf);
    // If entry is not a regular file, then it's a directory and should follow it's own branch.
    if (!S_ISREG(buf.st_mode)) {
      // Need to properly determine size of directories.
      // Likely need to sum the size of all files in subdirectories recursively.
      // Best to extract that logic into a separate function.
      size = buf.st_size;
    } else {
      // If target directory was provided, create needed relative path to entry
      // for proper execution of open().
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

    // If user passed in the '-f' flag, print out the file size of each file.
    if (f->file_size) {
    // (TODO): Currently, program does not properly compute the file size of directories.
    // To do so, we'll need to recursively find the file_size of all files in the
    // directory and nested directories and sum them up in order to obtain the
    // directory's file size.
      printf("%lld\t%s\n", size, entry->d_name);
    } else {
      printf("%s\n", entry->d_name);
    }

    memset(rel_path, 0, sizeof rel_path);
  }

  return 0;
}
