package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Hello Redis")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set("ihz", "2", 0).Err()	// 0: expiration in nanosecond
	if err != nil {
		panic(err)
	}

	val, err := client.Get("ihz").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ihz:", val)
}