all:
	gcc -Wall ctemplate.c configurator.c -o configurator

debug:
	gcc -g -O0 -Wall ctemplate.c configurator.c -o configurator
