#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <ifaddrs.h>

#include "ctemplate.h"

char **interfaces;
char ext_if[10];
char int_if[10];
char dhcp_if[10];

int length(char *arr[]) {
  int count = 0;
  while(arr[count] != NULL) count++;
  return count;
}

void get_interfaces()
{
  int i = 0;
  struct ifaddrs *ifap, *ifa;
  getifaddrs(&ifap);

  for (ifa = ifap; ifa; ifa = ifa->ifa_next) {
    if (i > 0 && ifa->ifa_name == interfaces[i - 1]) { continue; }
    if ((int)ifa->ifa_flags >= 0) { continue; }
    if (memcmp(ifa->ifa_name, "lo0", 3) == 0) { continue; }
    i++;
    interfaces = realloc(interfaces, sizeof(interfaces) * i);
    interfaces[i - 1] = ifa->ifa_name;
  }

  interfaces = realloc(interfaces, sizeof(interfaces) * (i + 1));
  interfaces[i] = NULL;

  freeifaddrs(ifap);
}

void print_interfaces()
{
  int i, l = length(interfaces);
  printf("[");
  for (i = 0; i < l; i++) {
    printf("%s", interfaces[i]);
    if (i < l - 1) {
      printf(", ");
    }
  }
  printf("]");
}

void prompt(char *message)
{
  printf("%s Choices are ", message);
  print_interfaces();
  printf(": ");
}

void write_pf_conf()
{
  TMPL_varlist *mylist;
  FILE *pfconf;
  pfconf = fopen("pf.conf", "w+");

  mylist = TMPL_add_var(0, "ext_if", ext_if, "int_if", int_if, 0);
  TMPL_write("pf.conf.tmpl", 0, 0, mylist, pfconf, stderr);
  TMPL_free_varlist(mylist);
}

int main(int argc, char **argv)
{
  get_interfaces();
  prompt("Select external interface.");
  fgets(ext_if, 10, stdin);
  strtok(ext_if, "\n");
  prompt("Select internal interface.");
  fgets(int_if, 10, stdin);
  strtok(int_if, "\n");
  prompt("Select DHCP interface.");
  fgets(dhcp_if, 10, stdin);
  strtok(dhcp_if, "\n");

  write_pf_conf();

  return 0;
}
