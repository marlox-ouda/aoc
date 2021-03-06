#define CHAR_NEWLINE 10
#define CHAR_A 97
#define CHAR_ZERO 48
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

// calcule le nombre de bit à 1 dans un entier
// (sur les 26 premiers bit uniquement)
static inline unsigned short bit_count(unsigned int value) {
  unsigned short idx;
  unsigned short count = 0;
  for (idx = 0; idx < 26; idx++)
    if (((1 << idx) & value) > 0)
      count += 1;
  return count;
}

static inline void sys_exit(int retcode) {
  register long eax asm("eax");
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_EXIT), "D" (retcode)
      : "rcx", "r11"
  );
}

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

void _start() {
  int fd;
  const char * first_addr;
  const char * addr;
  const char * last_addr;
  // ensemble des déclarations associées à une personne
  unsigned int current_declaration = 0;
  // ensemble des déclarations communes à chaque membre du groupe
  unsigned int current_shared_group_declaration = 0;
  // ensemble des déclarations présentes une fois au moins dans un groupe
  unsigned int current_added_group_declaration = 0;
  // somme du nombre de déclarations présente une fois au moins dans les groupes
  unsigned short declaration_run_one = 0;
  // somme du nombre de déclarations communes à chaque groupe
  unsigned short declaration_run_two = 0;
  struct stat st;
  char * output_char;
  char output_buffer[OUTPUT_LEN];
  // register
  register long eax asm("eax");
  register long r10 asm("r10");
  register long r8 asm("r8");
  register long r9 asm("r9");
  // hardcoded path
  const char* const path = "/home/user/aoc/2020/06/input6.txt";
  const char* const output_path = "/tmp/output6.txt";

  //if ((fd = open("/home/user/aoc/2020/06/input6.txt", O_RDONLY | O_NOATIME)) < 0)
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
    if (*addr == CHAR_NEWLINE) {
      if (current_declaration == 0) {
        // cas d’un nouveau groupe
        // note: current_declaration == 0 implique que la branche else
        //       d'être exécuté ; pas besoin de refaire les opérations
        declaration_run_one += bit_count(current_added_group_declaration);
        declaration_run_two += bit_count(current_shared_group_declaration);
        current_shared_group_declaration = 0;
        current_added_group_declaration = 0;
      } else {
        // cas d’une nouvelle déclaration au sein d'un groupe
        if (current_shared_group_declaration == 0)
          current_shared_group_declaration = current_declaration;
        else
          current_shared_group_declaration &= current_declaration;
        current_added_group_declaration |= current_declaration;
        current_declaration = 0;
      }
    } else if (CHAR_A <= *addr && *addr < CHAR_A+26) {
      // converti les lettres de 'a' à 'z' en bit dans un entier
      current_declaration |= (1 << (*addr - CHAR_A));
    }
    addr += 1;
  }
  //munmap(first_addr, st.st_size);
  asm volatile (
      "syscall"
      : "=r" (eax)
      : "0" (SYS_MUNMAP), "D" (first_addr), "S" (st.st_size)
      : "rcx", "r11"
  );
  declaration_run_one += bit_count(current_added_group_declaration);
  declaration_run_two += bit_count(current_shared_group_declaration);
  //printf("%hu\t%hu\n", declaration_run_one, declaration_run_two);
  output_char = output_buffer + OUTPUT_LEN;
  *(--output_char) = '\0';
  *(--output_char) = '\n';
  while (declaration_run_two != 0) {
    *(--output_char) = CHAR_ZERO + (declaration_run_two % 10);
    declaration_run_two = declaration_run_two / 10;
  }
  *(--output_char) = '\t';
  while (declaration_run_one != 0) {
    *(--output_char) = CHAR_ZERO + (declaration_run_one % 10);
    declaration_run_one = declaration_run_one / 10;
  }
  // output in a file is 66% faster (0.09s instead of 0.13s in STDOUT)
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
