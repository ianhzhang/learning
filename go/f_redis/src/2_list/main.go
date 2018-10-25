package main

import (
	"fmt"
	"reflect"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Hello Redis")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	list_len := client.LLen("mList1").Val()
	fmt.Println(list_len)

	if list_len > 0 {
		items := client.LRange("mList1", 0 ,5).Val()
		fmt.Println(reflect.TypeOf(items))			// [] string
		fmt.Println(items)							// [ {}, {}]
		fmt.Println(len(items))						// 6

		for _, item := range(items) {
			fmt.Println(item)						// {"data": {"a": 20002152, "E": 1372}, "seq": 16776965}

			fmt.Println(reflect.TypeOf(item))		// string
		}
		client.LPop("mList1")
	}

}