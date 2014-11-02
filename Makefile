all:
	gcc -Wall -Werror configurator.c -o configurator

debug:
	gcc -g -O0 -Wall -Werror configurator.c -o configurator
