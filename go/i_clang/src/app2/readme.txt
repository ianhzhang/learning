For dynamic library:
export LD_LIBRARY_PATH=/home/user/vmShare/git/hub/learningProj/go/i_clang/src/app1/myClib
export LD_LIBRARY_PATH=/home/user/vmShare/git/hub/learningProj/go/i_clang/src/app1/myClib

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -LmyClib -lgreet
// #include <stdlib.h>
// #include "myClib/greeter.h"

export LD_LIBRARY_PATH=/home/user/vmShare/git/hub/learningProj/go/i_clang/src/app1/myClib
./app1


For static library:
gcc -c greeter.c -o greeter.o
gcc -c greeter.c -o greeter.o

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -LmyClib -lgreet
// #include <stdlib.h>
// #include "myClib/greeter.h"


go build app1
./app1


app1:
main.go call C files.

app2:
main.go call go library.  go library call C files.  The library is the driver layer for isolation.