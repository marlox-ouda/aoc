#define _GNU_SOURCE 1
#include <stdio.h>
#include <unistd.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <fcntl.h>

static inline unsigned short bit_count(unsigned int value) {
  unsigned short idx;
  unsigned short count = 0;
  for (idx = 0; idx < 26; idx++)
    if (((1 << idx) & value) > 0)
      count += 1;
  return count;
}

#define CHAR_NEWLINE 10
#define CHAR_A 97

int main() {
  int fd;
  char * addr;
  char * last_addr;
  unsigned int current_declaration = 0;
  unsigned int current_shared_group_declaration = 0;
  unsigned int current_added_group_declaration = 0;
  unsigned short declaration_run_one = 0;
  unsigned short declaration_run_two = 0;
  struct stat st;

  if ((fd = open("/home/user/aoc/2020/06/input6.txt", O_RDONLY | O_NOATIME)) < 0)
    return -1;
  if (fstat(fd, &st) < 0)
    return -2;
  if ((addr = mmap(NULL, st.st_size, PROT_READ, MAP_PRIVATE, fd, 0)) < 0)
    return -3;
  last_addr = addr + st.st_size;
  while (addr < last_addr) {
    if (*addr == 10) {
      if (current_declaration == 0) {
        declaration_run_one += bit_count(current_added_group_declaration);
        declaration_run_two += bit_count(current_shared_group_declaration);
        current_shared_group_declaration = 0;
        current_added_group_declaration = 0;
      } else {
        if (current_shared_group_declaration == 0)
          current_shared_group_declaration = current_declaration;
        else
          current_shared_group_declaration &= current_declaration;
        current_added_group_declaration |= current_declaration;
        current_declaration = 0;
      }
    } else if (CHAR_A <= *addr && *addr < CHAR_A+26) {
      current_declaration |= (1 << (*addr - CHAR_A));
    }
    addr += 1;
  }
  declaration_run_one += bit_count(current_added_group_declaration);
  declaration_run_two += bit_count(current_shared_group_declaration);
  printf("%hu\t%hu\n", declaration_run_one, declaration_run_two);
  close(fd);

}
