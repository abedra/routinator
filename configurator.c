#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <ifaddrs.h>

char **interfaces;

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

int main(int argc, char **argv)
{
  get_interfaces();
  printf("Interfaces: ");
  print_interfaces();
  printf("\n");

  return 0;
}
