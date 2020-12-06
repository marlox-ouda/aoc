#define _GNU_SOURCE 1
#include <stdio.h>
#include <unistd.h>
#include <sys/stat.h>
#include <sys/mman.h>
#include <fcntl.h>

#define CHAR_NEWLINE 10
#define CHAR_A 97

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

int main() {
  int fd;
  char * addr;
  char * last_addr;
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

  if ((fd = open("/home/user/aoc/2020/06/input6.txt", O_RDONLY | O_NOATIME)) < 0)
    return -1;
  if (fstat(fd, &st) < 0)
    return -2;
  // l’ensemble du fichier est mappé en mémoire en un bloc
  if ((addr = mmap(NULL, st.st_size, PROT_READ, MAP_PRIVATE, fd, 0)) < 0)
    return -3;
  last_addr = addr + st.st_size;
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
  declaration_run_one += bit_count(current_added_group_declaration);
  declaration_run_two += bit_count(current_shared_group_declaration);
  printf("%hu\t%hu\n", declaration_run_one, declaration_run_two);
  close(fd);

}
