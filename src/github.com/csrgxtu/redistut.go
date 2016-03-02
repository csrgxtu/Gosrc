package main

import (
	"fmt"
	"gopkg.in/redis.v3"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.200.2:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	count, _ := client.Keys("statistics:online:user:*").Result()
	fmt.Println(len(count))
	//fmt.Println(len(client.Keys("statistics:online:user:*").Result()))
}
