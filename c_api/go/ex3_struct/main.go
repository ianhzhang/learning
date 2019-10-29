package main

import (
  "fmt"
  "reflect"
)

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lgb 
// #include <test.h>
import "C"
// ihz: cannot be empty line between #include and import C

func main() {
  fmt.Printf("Invoking c library...\n")
  a := []float64{1.1,1.2}
  b := C.test((*C.double)(&a[0]), 2)
  fmt.Println(b)
  fmt.Println(reflect.TypeOf(b))
  fmt.Println(b.R2)
  fmt.Println(b.C2)

}
