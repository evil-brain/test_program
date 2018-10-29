CC = gcc
LD = gcc

CFLAGS = -Wall -pipe 
OFLAGS = -c -I/usr/include
LFLAGS = $(CFLAGS) -L/usr/lib/

SOURCES = $(wildcard *.c)
OBJECTS = $(SOURCES:.c=.o)

DEBUG = no
PROFILE = no
PEDANTIC = no
OPTIMIZATION = -O3

ifeq ($(DEBUG), yes)
	CFLAGS += -g
	OPTIMIZATION = -O0
endif

ifeq ($(PROFILE), yes)
	CFLAGS += -pg
endif

CFLAGS += $(OPTIMIZATION)

all: project

project: $(OBJECTS)
	$(CC) $(OBJECTS) $(CFLAGS) -o test_program

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

clean:
	rm -rf *.o test_program

rebuild: clean all

.PHONY : clean
.SILENT : clean
