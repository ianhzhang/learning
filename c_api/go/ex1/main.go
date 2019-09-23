package main

import "fmt"

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lgb 
// #include <test.h>
import "C"
// ihz: cannot be empty line between #include and import C

func main() {
  fmt.Printf("Invoking c library...\n")
  a := C.test()
  fmt.Println(a)
}

