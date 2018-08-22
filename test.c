#include <stdlib.h>
#include <stdio.h>
#include "lib.h"

int main(int argc, char const *argv[]) {
  char out[1024];
  char *p = out;
  GoCall((char *)argv[1],&p);
  PrintGoString(p);

  struct Demo d;
  d = GetStructDemo(1,2);

  PrintStrutDemo(d);
  PrintGoInt(d.a,d.b);
  return 0;
}
