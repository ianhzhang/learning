package main

import (
	"fmt"
	"net"
	//"reflect"
	"github.com/go-redis/redis"
	"encoding/json"
	"3_test/lib"
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


	// ==============================================================================
	// ==============================================================================
	fmt.Println("UDP Server!")
	addr := net.UDPAddr{
		Port: 30002,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr) // code does not block here
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	
	ecgBuf := lib.New_CircularBuf(3000)
	icgBuf := lib.New_CircularBuf(3000)

	var buf [16384]byte
	for {
		// rLen, remoteAddr, err := conn.ReadFromUDP(buf[:])
		rLen, _, err := conn.ReadFromUDP(buf[:])
		// Do stuff with the read bytes
		if err!= nil {
			fmt.Println(err)
			panic(err)
		} else {
			bufStr := string(buf[:rLen])
			fmt.Println(bufStr)
			// Append it to Redis
			err := client.Set("ihz", bufStr, 0).Err()	// 0: expiration in nanosecond
			if err != nil {
				panic(err)
			}

			// Append value to circular buffer
			var result map[string]interface{}
			_ = json.Unmarshal([]byte(bufStr), &result)

			ecgVal := int(result["E"].(float64))
			icgVal := int(result["a"].(float64))

			ecgBuf.Append(ecgVal)
			icgBuf.Append(icgVal)

			// Calculate stats
		}
	}

}