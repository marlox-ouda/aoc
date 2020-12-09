/*
 * Solve of 2020.08 Advent of code
 *
 * Author: Marlox (git@mx.ouda.fr)
 *
 * Performance:
 *   0.16 ms with write_in_file
 *   0.19 ms with print
 */

//******************************************************************************
// Macros definitions
//******************************************************************************

//******************************************************************************
// Macros related to system call and I/O management

#define ARCH X86_64
// Syscall definition (x86_64 only)
#if ARCH == X86_64
  #define SYS_OPEN 2
  #define SYS_WRITE 1
  #define SYS_CLOSE 3
  #define SYS_STAT 5
  #define SYS_MMAP 9
  #define SYS_MUNMAP 11
  #define SYS_EXIT 60
  #define SYS_FSYNC 74
#endif

// Entrypoint definition (x86_64 only)
#if ARCH == X86_64
  #define ENTRYPOINT _start
#endif

// Assembly instructions (x86_64 only)
#if ARCH == X86_64
  #define SYSCALL_INTR "syscall"
  #define SYSCALL_MODIFIED_REGS "rcx", "r11"
  #define SYSCALL_OUTPUT_REG "eax"
  #define SYSCALL_ARG1_REG "0"
  #define SYSCALL_ARG2_REG "D"
  #define SYSCALL_ARG3_REG "S"
  #define SYSCALL_ARG4_REG "d"
  #define SYSCALL_ARG5_REG "r10"
  #define SYSCALL_ARG6_REG "r8"
  #define SYSCALL_ARG7_REG "r9"
  // define syscall register
#endif

// File mode macros
#define O_RDONLY 0
#define O_WRONLY 1
#define O_CREAT 64
#define O_TRUNC 512
#define O_NOATIME 262144

// Mmap parameter macros
#define MAP_PRIVATE 2
#define PROT_READ 1
#define MAP_FAILED ((void *) -1)

// Pre-allocated file descriptor
#define FD_STDIN 0
#define FD_STDOUT 1
#define FD_STDERR 1

//******************************************************************************
// Recurrent macros among programs

// Preallocated buffers
#define OUTPUT_BUFFER_LEN 50

// ASCII char value
#define CHAR_NEWLINE 10
#define CHAR_ZERO 48
#define CHAR_A 97

// Return codes
#define SUCCESS_CODE 0
#define EIOINPUTERR -1
#define EIOOUTPUTERR -2
#define EMEMERR -3
#define EFMTINPUTERR -4

//******************************************************************************
// Macro related to program

// Preallocated buffers
#define NUMBER_BUFFER_LEN 1000

//******************************************************************************
// Types definitions
//******************************************************************************

//******************************************************************************
// Types related to system call and I/O management

typedef unsigned long size_t;

struct stat {
  unsigned long st_dev;
  unsigned long st_ino;
  unsigned int st_mode;
  unsigned long st_nlink;
  unsigned int st_uid;
  unsigned int st_gid;
  unsigned long st_rdev;
  size_t st_size;
  unsigned long st_blksize;
  unsigned long st_blocks;
  unsigned long st_atime;
  unsigned long st_mtime;
  unsigned long st_ctime;
};

struct mapped_ptr {
  const char * first_addr;
  size_t size;
};

//******************************************************************************
// Types specific to program

typedef unsigned long long number_t;

//******************************************************************************
// Functions definitions
//******************************************************************************

//******************************************************************************
// Functions related to system calls

// Exit the program with given return code
//
// Args:
//   retcode: the return code of the program
static inline void sys_exit(const int retcode) {
  register long eax asm("eax");
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_EXIT), "D" (retcode)
      : "rcx", "r11"
  );
}

// Write data buffer into the given file descriptor
//
// Args:
//   fd: the file descriptor where data is written to
//   data_buf: the buffer containing data to write
//   data_len: the size of the buffer to write
static inline void sys_write(const int fd, const char* const data_buf, const size_t data_len) {
  register long eax asm("eax");
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_WRITE), "D" (fd), "S" (data_buf), "d" (data_len)
      : "rcx", "r11" , "memory"
  );
}

// Close the given file descriptor
//
// Args:
//   fd: the file descriptor where data is written to
static inline void sys_close(const int fd) {
  register long eax asm("eax");
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_CLOSE), "D" (fd)
      : "rcx", "r11"
  );
}

// Stat the given file descriptor
//
// Args:
//   fd: the file descriptor to stat
//   st: the reference stat structure to fill with stat data
static inline void sys_fstat(const int fd, const struct stat* st) {
  register long eax asm("eax");
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_STAT), "D" (fd), "S" (st)
      : "rcx", "r11"
  );
  if (eax < 0)
    sys_exit(EIOINPUTERR);
}

static inline void sys_munmap(const char* const addr, const size_t size) {
  register long eax asm("eax");
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_MUNMAP), "D" (addr), "S" (size)
      : "rcx", "r11"
  );
}

//******************************************************************************
// Functions related to I/O management

// Map the whole file in memory
//
// Args:
//   path: the path to the file to map
// Return:
//   struct mapped_ptr: structure with pointer to the beginning of the file
//                      and file size
static inline struct mapped_ptr map_file(const char * const path) {
  struct mapped_ptr input;
  int fd;
  struct stat st;

  register long r10 asm("r10");
  register long r8 asm("r8");
  register long r9 asm("r9");

  //if ((fd = open(path, O_RDONLY | O_NOATIME)) < 0)
  asm volatile (
      "syscall"
      : "=a" (fd)
      : "0" (SYS_OPEN), "D" (path), "S" (O_RDONLY | O_NOATIME)
      : "rcx", "r11"
  );
  if (fd < 0)
    sys_exit(EIOINPUTERR);
  sys_fstat(fd, &st);
  input.size = st.st_size;
  // l’ensemble du fichier est mappé en mémoire en un bloc
  //if ((addr = mmap(NULL, input.size, PROT_READ, MAP_PRIVATE, fd, 0)) < 0)
    //return -3;
  r10 = MAP_PRIVATE;
  r8 = fd;
  r9 = 0;
  asm volatile (
      "syscall"
      : "=a" (input.first_addr)
      : "0" (SYS_MMAP), "D" (0), "S" (input.size), "d" (PROT_READ), "r" (r10), "r" (r8), "r" (r9)
      : "rcx", "r11", "memory"
  );
  if (input.first_addr == MAP_FAILED)
    sys_exit(EMEMERR);
  //close(fd);
  sys_close(fd);
  return input;
}

// Unmap the given mapping
//
// Args:
//   mapping: the mapping to unmap
static inline void unmap(struct mapped_ptr mapping) {
  sys_munmap(mapping.first_addr, mapping.size);
}

// Print the given data to provided file
//
// Note: writing in /tmp/ file increase performance of 0.02 to 0.04ms
//
// Args:
//   path: the path in witch data has to be written
//   data_buf: the buffer containing data to write
//   data_len: the size of the buffer to write
static inline void write_in_file(const char* const path, const char* const data_buf, const size_t data_len) {
  int fd;
  asm volatile (
      "syscall"
      : "=a" (fd)
      : "0" (SYS_OPEN), "D" (path), "S" (O_WRONLY | O_CREAT | O_TRUNC), "d" (0777)
      : "rcx", "r11", "memory"
  );
  if (fd < 0)
    sys_exit(EIOOUTPUTERR);
  sys_write(fd, data_buf, data_len);
  sys_close(fd);
}

// Print the given data to stdout
//
// Args:
//   data_buf: the buffer containing data to write
//   data_len: the size of the buffer to write
static inline void print(const char* const data_buf, const size_t data_len) {
  sys_write(FD_STDOUT, data_buf, data_len);
}

//******************************************************************************
// Recurrent functions among programs

// Prepend short value in the buffer (in decimal base)
//
// WARNING: required space must be reserved (6 bytes max per call)
// 
// Args:
//   buf: the buffer to prepend with the short value
//   value: the value to prepend
// Return :
//   char*: the position of the buffer after the prepend
static inline char * prepend_short_in_buf(char *buf, short value) {
  if (value < 0) {
    *(--buf) = '-';
    value = -value;
  }
  while (value != 0) {
    *(--buf) = CHAR_ZERO + (value % 10);
    value = value / 10;
  }
  return buf;
}

//******************************************************************************
// Recurrent functions among programs

// Prepend short value in the buffer (in decimal base)
//
// WARNING: required space must be reserved (6 bytes max per call)
// 
// Args:
//   buf: the buffer to prepend with the short value
//   value: the value to prepend
// Return :
//   char*: the position of the buffer after the prepend
static inline char * prepend_ulong_in_buf(char *buf, number_t value) {
  while (value != 0) {
    *(--buf) = CHAR_ZERO + (value % 10);
    value = value / 10;
  }
  return buf;
}

//******************************************************************************
// Function specific to program

static inline unsigned short load_numbers(const number_t* numbers, const struct mapped_ptr* const input) {
  const char * addr;
  const char * last_addr;
  unsigned short numbers_count = 0;
  number_t* current_number = (number_t*)numbers;

  last_addr = input->first_addr + input->size;
  addr = input->first_addr;
  *current_number = 0;
  while (addr < last_addr) {
    if (*addr == CHAR_NEWLINE) {
      ++numbers_count;
      ++addr;
      *++current_number = 0;
    } else {
      *current_number *= 10;
      *current_number += (*addr++ - CHAR_ZERO);
    }
  }
  return numbers_count;
}

static inline const number_t* find_invalid_number(const number_t* const first_number, const number_t* const last_number) {
  number_t current_sum_1, current_sum_2;
  const number_t* current_number = first_number + 25;
  char success;
  int pos_number1, pos_number2;
  while (current_number < last_number) {
    success = 0;
    for (pos_number1 = 1; pos_number1 < 25; pos_number1++) {
      current_sum_1 = *(current_number - pos_number1);
      for (pos_number2 = pos_number1 + 1; pos_number2 < 26; pos_number2++) {
        current_sum_2 = *(current_number - pos_number2);
        if ((current_sum_1 + current_sum_2) == *current_number) {
          success = 1;
          break;
        }
      }
      if (success == 1)
        break;
    }
    if (success == 0)
      break;
    current_number++;
  }
  return current_number;
}

static inline number_t find_summing_range(
    const number_t* const first_number,
    const number_t run1_solution) {
  const number_t* range_first_number;
  const number_t* range_last_number;
  number_t min, max;
  number_t range_sum;
  range_first_number = range_last_number = first_number;
  range_sum = *first_number;
  while (1) {
    while (range_sum < run1_solution) {
      range_sum += *++range_last_number;
    }
    if (range_sum == run1_solution)
      break;
    while (range_sum > run1_solution)
      range_sum -= *range_first_number++;
    if (range_sum == run1_solution)
      break;
  }
  min = max = *range_first_number++;
  while (range_first_number <= range_last_number) {
    if (*range_first_number < min)
      min = *range_first_number;
    else if (*range_first_number > max)
      max = *range_first_number;
    range_first_number++;
  }
  return min + max;
}

//******************************************************************************
// Main function
//******************************************************************************

void ENTRYPOINT() {
  // hardcoded path
  const char* const input_path = "/home/user/aoc/2020/09/input9.txt";
  const char* const output_path = "/tmp/output9.txt";

  number_t numbers[NUMBER_BUFFER_LEN];
  number_t run1_xmas_sum = 0;
  number_t run2_xmas_sum = 0;

  char * output_char;
  char output_buffer[OUTPUT_BUFFER_LEN];

  const struct mapped_ptr input = map_file(input_path);
  const unsigned short numbers_count = load_numbers(numbers, &input);
  number_t* const last_number = numbers + numbers_count;
  unmap(input);

  run1_xmas_sum = *(find_invalid_number(numbers, last_number));
  run2_xmas_sum = find_summing_range(numbers, run1_xmas_sum);

  output_char = output_buffer + OUTPUT_BUFFER_LEN;
  *(--output_char) = '\0';
  *(--output_char) = '\n';
  output_char = prepend_ulong_in_buf(output_char, run2_xmas_sum);
  *(--output_char) = '\t';
  output_char = prepend_ulong_in_buf(output_char, run1_xmas_sum);

  //print(output_char, (output_buffer + OUTPUT_BUFFER_LEN - output_char - 1));
  write_in_file(output_path, output_char, (output_buffer + OUTPUT_BUFFER_LEN - output_char - 1));
  sys_exit(SUCCESS_CODE);
}
