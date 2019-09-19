gcc -O2 -c test.c

ar q libgb.a test.o

go build main.go


user@user-VirtualBox:~/learning/go/8-go-c-api/ex2$ rm *.o *.a main
user@user-VirtualBox:~/learning/go/8-go-c-api/ex2$ gcc -c test.c 
user@user-VirtualBox:~/learning/go/8-go-c-api/ex2$ ar rcs libgb.a test.o 
user@user-VirtualBox:~/learning/go/8-go-c-api/ex2$ go build main.go 
user@user-VirtualBox:~/learning/go/8-go-c-api/ex2$ ./main 
