package main
//https://karthikkaranth.me/blog/calling-c-code-from-go/

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -LmyClib -lgreet
// #include <stdlib.h>
// #include "myClib/greeter.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello")

	name := C.CString("Gopher")			// take a go string and return a pointer to a Cchar
	defer C.free(unsafe.Pointer(name))
	

	year := C.int(2018)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}