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
char domain_name[255];
char nameservers[1024];
char subnet[16];
char netmask[16];
char router[16];
char dhcp_start[16];
char dhcp_end[16];

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

void assign_interfaces()
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
}

void assign_dhcp_options()
{
  printf("Enter domain name: ");
  fgets(domain_name, 255, stdin);
  strtok(domain_name, "\n");

  printf("Enter nameservers: ");
  fgets(nameservers, 1024, stdin);
  strtok(nameservers, "\n");

  printf("Enter subnet: ");
  fgets(subnet, 16, stdin);
  strtok(subnet, "\n");

  printf("Enter netmask: ");
  fgets(netmask, 16, stdin);
  strtok(netmask, "\n");

  printf("Enter router address: ");
  fgets(router, 16, stdin);
  strtok(router, "\n");

  printf("Enter DHCP start: ");
  fgets(dhcp_start, 16, stdin);
  strtok(dhcp_start, "\n");

  printf("Enter DHCP end: ");
  fgets(dhcp_end, 16, stdin);
  strtok(dhcp_end, "\n");
}

void write_dhcpd_conf()
{
  TMPL_varlist *mylist;
  FILE *dhcpdconf;
  dhcpdconf = fopen("dhcpd.conf", "w+");

  mylist = TMPL_add_var(0, "domain", domain_name, "nameservers", nameservers, 0);
  mylist = TMPL_add_var(mylist, "subnet", subnet, "netmask", netmask, 0);
  mylist = TMPL_add_var(mylist, "router", router, "dhcp_start", dhcp_start, 0);
  TMPL_add_var(mylist, "dhcp_end", dhcp_end, 0);

  TMPL_write("dhcpd.conf.tmpl", 0, 0, mylist, dhcpdconf, stderr);
  TMPL_free_varlist(mylist);
}

int main(int argc, char **argv)
{
  assign_interfaces();
  write_pf_conf();

  assign_dhcp_options();
  write_dhcpd_conf();

  return 0;
}
