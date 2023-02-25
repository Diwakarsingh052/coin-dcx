package main

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func ConnectToRedis() (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count := 0
	//connecting to redis and retrying 5 times in case redis is not ready to accept connections
	for {
		if count == 5 {
			return nil, errors.New("cannot connect to redis")
		}

		rdb := redis.NewClient(&redis.Options{
			Addr: "localhost:6379", // this host name is from the docker file
		})

		//ping redis to make sure connection is successful
		err := rdb.Ping(ctx).Err()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second) // waiting for one second before making another connection request
			log.Println("retrying connection to redis")
			count++
			continue
		}

		//if success return the connection
		return rdb, nil
	}

}
