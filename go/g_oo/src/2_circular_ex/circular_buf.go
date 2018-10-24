package main

import (
	"fmt"
	"math"
)

const DEFAULT_BUF_LEN = 10

type CircularBuf struct{
	buff []int
	index int
	total int
	maxlen int
}

func New_CircularBuf(param ...int) *CircularBuf{
	fmt.Println("ctor")

	p := new(CircularBuf)
	p.index = 0
	p.total = 0
	p.maxlen = DEFAULT_BUF_LEN
	
	if len(param) > 0 {
		p.maxlen = param[0]
	}
	fmt.Println("maxlen:", p.maxlen)
	p.buff = make([]int, p.maxlen)

	return p
}

func (c *CircularBuf) Append(val int) {
	c.buff[c.index] = val
	c.index ++ 
	c.total ++
	if (c.index >= c.maxlen) {
		c.index = 0
	}
	if (c.total >= c.maxlen) {
		c.total = c.maxlen
	}
	
}

func (c *CircularBuf) Stats() (int, int, float32) {
	var len int = c.maxlen
	
	if (c.total < c.maxlen) {
		len = c.total
	}
	fmt.Println(math.MaxInt32)
	var min = math.MaxInt32
	var max = math.MinInt16
	var total int32 = 0

	for i:=0 ; i < len; i++ {
		fmt.Printf("%d ",c.buff[i])
		if (c.buff[i] > max) {
			max = c.buff[i]
		} else if (c.buff[i] < min) {
			min = c.buff[i]
		}
		total = total + int32(c.buff[i])
	}
	fmt.Println("\ntotal:", total)

	var avg float32 = float32(total)/float32(len)
	return max, min, avg
}


func (c *CircularBuf) Display() {
	var len int = c.maxlen
	
	if (c.total < c.maxlen) {
		len = c.total
	}
	fmt.Println(c.total)
	
	for i:=0 ; i < len; i++ {
		fmt.Printf("%d ",c.buff[i])
	}
	fmt.Println()
}