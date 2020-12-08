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
#define EINPUTERR -1
#define EOUTPUTERR -2
#define EMEMERR -3

//******************************************************************************
// Macro related to program

// Preallocated buffers
#define INSTRUCTIONS_BUFFER_LEN 1000

// ASCII char value
#define CHAR_PLUS 43
#define CHAR_MOINS 45
#define CHAR_J 106
#define CHAR_N 110

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
// Types related to program

typedef char keyword_t;
typedef unsigned char bool_t;

struct instruction_t {
  keyword_t keyword;
  short value;
  bool_t matched;
};

static const keyword_t ACCUMULATOR = 1;
static const keyword_t NOP = 2;
static const keyword_t JUMP = 3;
static const bool_t FALSE = 0;
static const bool_t TRUE = 1;

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
  if (eax < 0) sys_exit(EINPUTERR);
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
    sys_exit(EINPUTERR);
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
    sys_exit(EOUTPUTERR);
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
// Main function
//******************************************************************************

void ENTRYPOINT() {
  // hardcoded path
  const char* const input_path = "/home/user/aoc/2020/08/input8.txt";
  const char* const output_path = "/tmp/output8.txt";

  const char * addr;
  const char * last_addr;

  unsigned short instructions_number = 0;
  //struct instruction_t * instructions;
  struct instruction_t instructions[INSTRUCTIONS_BUFFER_LEN];
  struct instruction_t * current_instruction;
  struct instruction_t * modified_instruction;
  struct instruction_t * last_instruction;
  bool_t checked_bit = TRUE;
  short value;
  short position = 0;
  short run1_accumulator = 0;
  short run2_accumulator = 0;
  bool_t success = FALSE;

  char * output_char;
  char output_buffer[OUTPUT_BUFFER_LEN];

  const struct mapped_ptr input = map_file(input_path);

  last_addr = input.first_addr + input.size;
  current_instruction = instructions;
  addr = input.first_addr;
  while (addr < last_addr) {
    switch (*addr) {
      case CHAR_NEWLINE:
        ++current_instruction;
        ++instructions_number;
        ++addr;
        break;
      case CHAR_A:
        current_instruction->keyword = ACCUMULATOR;
        current_instruction->matched = FALSE;
        addr += 4;
        break;
      case CHAR_N:
        current_instruction->keyword = NOP;
        current_instruction->matched = FALSE;
        addr += 4;
        break;
      case CHAR_J:
        current_instruction->keyword = JUMP;
        current_instruction->matched = FALSE;
        addr += 4;
        break;
      case CHAR_MOINS:
        value = 0;
        while (*++addr != CHAR_NEWLINE) {
          value *= 10;
          value += (*addr - CHAR_ZERO);
        }
        current_instruction->value = -value;
        break;
      case CHAR_PLUS:
        value = 0;
        while (*++addr != CHAR_NEWLINE) {
          value *= 10;
          value += (*addr - CHAR_ZERO);
        }
        current_instruction->value = value;
        break;
    }
  }
  unmap(input);
  while (1) {
    current_instruction = instructions + position;
    //printf("%hd: %hd\n", position, current_instruction->value);
    if (current_instruction->keyword == JUMP) {
      position += current_instruction->value;
    } else {
      ++position;
      if (current_instruction->keyword == ACCUMULATOR) {
        run1_accumulator += current_instruction->value;
      }
    }
    if (current_instruction->matched == TRUE) {
      break;
    }
    current_instruction->matched = TRUE;
  }

  last_instruction = instructions + instructions_number - 1;
  for (modified_instruction = instructions; modified_instruction <= last_instruction; ++modified_instruction) {
    if (modified_instruction->keyword == ACCUMULATOR)
      continue;
    modified_instruction->keyword ^= (JUMP | NOP);
    if (checked_bit == 128) {
      checked_bit = TRUE;
      for (current_instruction = instructions; current_instruction <= last_instruction; ++current_instruction)
        current_instruction->matched = FALSE;
    } else {
      checked_bit <<= 1;
    }
    position = 0;
    run2_accumulator = 0;
    while (1) {
      current_instruction = instructions + position;
      //printf("%hd: %hd\n", position, current_instruction->value);
      if (current_instruction->keyword == JUMP) {
        position += current_instruction->value;
      } else {
        ++position;
        if (current_instruction->keyword == ACCUMULATOR) {
          run2_accumulator += current_instruction->value;
        }
      }
      if ((current_instruction->matched & checked_bit) != 0)
        break;
      if (position == instructions_number) {
        success = TRUE;
        break;
      }
      if (position < 0 || position > instructions_number)
        break;
      current_instruction->matched = checked_bit;
    }
    if (success == TRUE)
      break;
    modified_instruction->keyword ^= (JUMP | NOP);
  }
  //free(instructions);
  //printf("%hd\t%hd\n", run1_accumulator, run2_accumulator);
  output_char = output_buffer + OUTPUT_BUFFER_LEN;
  *(--output_char) = '\0';
  *(--output_char) = '\n';
  output_char = prepend_short_in_buf(output_char, run2_accumulator);
  *(--output_char) = '\t';
  output_char = prepend_short_in_buf(output_char, run1_accumulator);

  //print(output_char, (output_buffer + OUTPUT_BUFFER_LEN - output_char - 1));
  write_in_file(output_path, output_char, (output_buffer + OUTPUT_BUFFER_LEN - output_char - 1));
  sys_exit(SUCCESS_CODE);
}
