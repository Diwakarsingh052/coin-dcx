package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	// create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// create stream event data
	data := map[string]interface{}{
		"email": "johndoe@example.com",
		"title": "My First Post",
	}

	// add event to Redis stream
	err := rdb.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "article:add:events",
		Values: data,
	}).Err()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Event added to stream successfully!")
}
