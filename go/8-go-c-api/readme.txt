gcc -O2 -c test.c

ar q libgb.a test.o

go build main.go