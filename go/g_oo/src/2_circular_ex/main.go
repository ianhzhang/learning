package main

import (
	"fmt"
)

func main() {
	fmt.Println("Circluar Buffer Example\n===================")

	cbuff1 := New_CircularBuf(2)
	cbuff1.Append(2)
	cbuff1.Append(3)
	cbuff1.Append(4)
	cbuff1.Append(5)
	cbuff1.Append(6)
	cbuff1.Append(7)
	cbuff1.Append(8)
	cbuff1.Append(9)
	cbuff1.Append(10)
	cbuff1.Append(11)
	cbuff1.Append(12)
	cbuff1.Append(13)
	cbuff1.Append(14)
	cbuff1.Append(15)
	cbuff1.Append(16)
	max, min, avg := cbuff1.Stats()
	fmt.Printf("max %d, min %d, avg %f\n", max,min,avg)
}