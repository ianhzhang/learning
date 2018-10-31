package main

import (
	"fmt"
	"errors"
)


func main() {
	// This is for sub-function which has no recover function
	defer func() {
		if r:= recover(); r!=nil {
			fmt.Println("Recovered ",r)
		}
	}()
	fmt.Println("Start")
	a, err := myfun()
	fmt.Println("------------------")
	fmt.Println("err", err)
	fmt.Println("a=", a)
	if err==nil {
		fmt.Println("return ", a)
	}
}


func myfun() (x int, err error){
	err = nil

	// This is recover function for this fuction
	// We set x and err so the calling function can check it.
	defer func() {
		if r:= recover(); r!=nil {
			fmt.Println("Recovered2 ",r)
			err = errors.New("Ian Zhang Error")
			fmt.Println(err)	
			x = -1
		}
	}()
	//panic("I Panic")
	fmt.Println("After Panic")
	x = 1
	err = nil
	return			// same as: return x, err
}