package main

import (
	"net/http"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func handle_hello(w http.ResponseWriter, r *http.Request) {
	defer wg.Done()
	fmt.Println("handle hello in goroutine")
	time.Sleep(2*time.Second)
	fmt.Println("timeout")
	fmt.Fprintf(w, "Hello World goroutin2")
}

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/hello1", func(w http.ResponseWriter, r *http.Request) {
		wg.Add(1)
		go handle_hello(w,r)
		wg.Wait()
	})
	http.HandleFunc("/hello2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle hello2 in goroutine")
		//time.Sleep(10*time.Second)
		var cnt uint64
		cnt = 0
		fmt.Println("cnt=",cnt)
		for (cnt<9999999) {
			cnt = cnt+1
			fmt.Print(cnt)
		}
		fmt.Println("\ntimeout 2")
		fmt.Fprintf(w, "Hello World goroutin2")
	})
	http.HandleFunc("/hello3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handle hello3 in goroutine")
		fmt.Fprintf(w, "Hello World goroutin3")
	})
	http.ListenAndServe(":5000", nil)
}

/* 
What I see is even if hello2 is blocked by time.Sleep.   hello3 still can coming and get processed.
Therefore, the goroutine in hello1 is not necessary.
hello1 has to be used together with WaitGroup.  Otherwise, it will not work.

If it is from same tab, hello2 first, blocked, hell3 go through, then hello2 reply does not get to browser
*/