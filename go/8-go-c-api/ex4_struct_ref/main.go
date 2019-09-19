package main

import (
  "fmt"
  "reflect"
  "unsafe"
)

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lgb 
// #include <test.h>
import "C"
// ihz: cannot be empty line between #include and import C

type RCResult struct {
  R1 float64
  C1 float64
  R2 float64
  C2 float64
};


func main() {
  fmt.Printf("Invoking c library...\n")
  a := []float64{1.1,1.2}
  rslt := RCResult{3.1,3.2,3.3,3.5}

  fmt.Println(rslt)
  println("====================================")
  b := C.test((*C.double)(&a[0]), 2, (*C.RCResult)(unsafe.Pointer(&rslt)) )
  fmt.Println(b)
  fmt.Println(reflect.TypeOf(b))
  fmt.Println(rslt);
}
