#define OUTPUT_LEN 50

#define O_RDONLY 0
#define O_WRONLY 1
#define O_CREAT 64
#define O_TRUNC 512
#define O_NOATIME 262144

#define MAP_PRIVATE 2
#define PROT_READ 1
#define MAP_FAILED ((void *) -1)

#define SYS_OPEN 2
#define SYS_WRITE 1
#define SYS_CLOSE 3
#define SYS_STAT 5
#define SYS_MMAP 9
#define SYS_MUNMAP 11
#define SYS_EXIT 60
#define SYS_FSYNC 74

#define STDOUT 1

#define CHAR_NEWLINE 10
#define CHAR_ZERO 48
#define CHAR_A 97
#define CHAR_J 106
#define CHAR_N 110
#define CHAR_PLUS 43
#define CHAR_MOINS 45

typedef char keyword_t;
typedef unsigned char bool_t;

struct stat {
  unsigned long st_dev;
  unsigned long st_ino;
  unsigned int st_mode;
  unsigned long st_nlink;
  unsigned int st_uid;
  unsigned int st_gid;
  unsigned long st_rdev;
  unsigned long st_size;
  unsigned long st_blksize;
  unsigned long st_blocks;
  unsigned long st_atime;
  unsigned long st_mtime;
  unsigned long st_ctime;
};

struct instruction_t {
  keyword_t keyword;
  short value;
  bool_t matched;
};

static const keyword_t ACCUMULATOR = 1;
static const keyword_t NOP = 2;
static const keyword_t JUMP = 3;
static const bool_t TRUE = 1;
static const bool_t FALSE = 0;

static inline void sys_exit(int retcode) {
  register long eax asm("eax");
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_EXIT), "D" (retcode)
      : "rcx", "r11"
  );
}

void _start() {
  int fd;
  struct stat st;

  const char * first_addr;
  const char * addr;
  const char * last_addr;

  unsigned short instructions_number = 0;
  //struct instruction_t * instructions;
  struct instruction_t instructions[1000];
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
  char output_buffer[OUTPUT_LEN];
  // register
  register long eax asm("eax");
  register long r10 asm("r10");
  register long r8 asm("r8");
  register long r9 asm("r9");
  // hardcoded path
  const char* const path = "/home/user/aoc/2020/08/input8.txt";
  const char* const output_path = "/tmp/output8.txt";

  //if ((fd = open("/home/user/aoc/2020/08/input8.txt", O_RDONLY | O_NOATIME)) < 0)
  asm volatile (
      "syscall"
      : "=a" (fd)
      : "0" (SYS_OPEN), "D" (path), "S" (O_RDONLY | O_NOATIME)
      : "rcx", "r11"
  );
  if (fd < 0)
    sys_exit(-1);
  //if (fstat(fd, &st) < 0)
  //  return -2;
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_STAT), "D" (fd), "S" (&st)
      : "rcx", "r11"
  );
  if (eax < 0)
    sys_exit(-2);
  // l’ensemble du fichier est mappé en mémoire en un bloc
  //if ((addr = mmap(NULL, st.st_size, PROT_READ, MAP_PRIVATE, fd, 0)) < 0)
    //return -3;
  r10 = MAP_PRIVATE;
  r8 = fd;
  r9 = 0;
  asm volatile (
      "syscall"
      : "=a" (first_addr)
      : "0" (SYS_MMAP), "D" (0), "S" (st.st_size), "d" (PROT_READ), "r" (r10), "r" (r8), "r" (r9)
      : "rcx", "r11", "memory"
  );
  if (first_addr == MAP_FAILED)
    sys_exit(-3);
  //close(fd);
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_CLOSE), "D" (fd)
      : "rcx", "r11"
  );
  last_addr = first_addr + st.st_size;
  addr = first_addr;
  while (addr < last_addr) {
    if (*addr == CHAR_NEWLINE)
      ++instructions_number;
    addr += 1;
  }
  //instructions = malloc((instructions_number) * sizeof(struct instruction_t));
  current_instruction = instructions;
  addr = first_addr;
  while (addr < last_addr) {
    switch (*addr) {
      case CHAR_NEWLINE:
        ++current_instruction;
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
  //munmap(first_addr, st.st_size);
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_MUNMAP), "D" (first_addr), "S" (st.st_size)
      : "rcx", "r11"
  );
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
  output_char = output_buffer + OUTPUT_LEN;
  *(--output_char) = '\0';
  *(--output_char) = '\n';
  if (run2_accumulator < 0) {
    *(--output_char) = '-';
    run2_accumulator = -run2_accumulator;
  }
  while (run2_accumulator != 0) {
    *(--output_char) = CHAR_ZERO + (run2_accumulator % 10);
    run2_accumulator = run2_accumulator / 10;
  }
  *(--output_char) = '\t';
  if (run1_accumulator < 0) {
    *(--output_char) = '-';
    run1_accumulator = -run1_accumulator;
  }
  while (run1_accumulator != 0) {
    *(--output_char) = CHAR_ZERO + (run1_accumulator % 10);
    run1_accumulator = run1_accumulator / 10;
  }
  asm volatile (
      "syscall"
      : "=a" (fd)
      : "0" (SYS_OPEN), "D" (output_path), "S" (O_WRONLY | O_CREAT | O_TRUNC), "d" (0777)
      : "rcx", "r11", "memory"
  );
  if (fd < 0)
    sys_exit(-4);
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_WRITE), "D" (fd), "S" (output_char), "d" (output_buffer + OUTPUT_LEN - output_char - 1)
      : "rcx", "r11" , "memory"
  );
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_CLOSE), "D" (fd)
      : "rcx", "r11"
  );
  //return 0;
  sys_exit(0);
}
