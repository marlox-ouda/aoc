#define _GNU_SOURCE 1
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <fcntl.h>

#define CHAR_NEWLINE 10
#define CHAR_0 48
#define CHAR_A 97
#define CHAR_J 106
#define CHAR_N 110
#define CHAR_PLUS 43
#define CHAR_MOINS 45

typedef char keyword_t;
typedef char bool_t;

static const keyword_t ACCUMULATOR = 1;
static const keyword_t NOP = 2;
static const keyword_t JUMP = 3;
static const bool_t TRUE = 1;
static const bool_t FALSE = 0;

struct instruction_t {
  keyword_t keyword;
  short value;
  bool_t matched;
};

int main() {
  int fd;
  struct stat st;

  const char * first_addr;
  const char * addr;
  const char * last_addr;

  unsigned short instructions_number = 0;
  struct instruction_t * instructions;
  struct instruction_t * current_instruction;
  struct instruction_t * modified_instruction;
  struct instruction_t * last_instruction;
  short value;
  short position = 0;
  short accumulator = 0;
  bool_t success = FALSE;

  if ((fd = open("/home/user/aoc/2020/08/input8.txt", O_RDONLY | O_NOATIME)) < 0)
    return -1;
  if (fstat(fd, &st) < 0)
    return -2;
  // l’ensemble du fichier est mappé en mémoire en un bloc
  if ((first_addr = mmap(NULL, st.st_size, PROT_READ, MAP_PRIVATE, fd, 0)) == MAP_FAILED)
    return -3;
  close(fd);
  last_addr = first_addr + st.st_size;
  addr = first_addr;
  while (addr < last_addr) {
    if (*addr == CHAR_NEWLINE)
      ++instructions_number;
    addr += 1;
  }
  //instructions_number -= 1;
  instructions = malloc((instructions_number) * sizeof(struct instruction_t));
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
          value += (*addr - CHAR_0);
        }
        current_instruction->value = -value;
        break;
      case CHAR_PLUS:
        value = 0;
        while (*++addr != CHAR_NEWLINE) {
          value *= 10;
          value += (*addr - CHAR_0);
        }
        current_instruction->value = value;
        break;
    }
  }
  munmap((char*)first_addr, st.st_size);
  while (1) {
    current_instruction = instructions + position;
    //printf("%hd: %hd\n", position, current_instruction->value);
    if (current_instruction->keyword == JUMP) {
      position += current_instruction->value;
    } else {
      ++position;
      if (current_instruction->keyword == ACCUMULATOR) {
        accumulator += current_instruction->value;
      }
    }
    if (current_instruction->matched == TRUE) {
      break;
    }
    current_instruction->matched = TRUE;
  }
  printf("run1:\t%hd\n", accumulator);

  last_instruction = instructions + instructions_number - 1;
  for (modified_instruction = instructions; modified_instruction <= last_instruction; ++modified_instruction) {
    if (modified_instruction->keyword == ACCUMULATOR)
      continue;
    modified_instruction->keyword ^= (JUMP | NOP);
    for (current_instruction = instructions; current_instruction <= last_instruction; ++current_instruction)
      current_instruction->matched = FALSE;
    position = 0;
    accumulator = 0;
    while (1) {
      current_instruction = instructions + position;
      //printf("%hd: %hd\n", position, current_instruction->value);
      if (current_instruction->keyword == JUMP) {
        position += current_instruction->value;
      } else {
        ++position;
        if (current_instruction->keyword == ACCUMULATOR) {
          accumulator += current_instruction->value;
        }
      }
      if (current_instruction->matched == TRUE)
        break;
      if (position == instructions_number) {
        success = TRUE;
        break;
      }
      if (position < 0 || position > instructions_number)
        break;
      current_instruction->matched = TRUE;
    }
    if (success == TRUE)
      break;
    modified_instruction->keyword ^= (JUMP | NOP);
  }
  free(instructions);
  printf("run2:\t%hd\n", accumulator);
}
