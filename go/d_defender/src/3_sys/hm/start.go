package hm

import (
	"fmt"
	"time"
	"strconv"
)

func Say(s string) {
	i := 0
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s + " = " + strconv.Itoa(i))
		i = i + 1
	}
}