package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "redis-service:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
