package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func runCacheRedis() {
	client := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
		})

	ctx := context.Background()

	err := client.Set(ctx, "JamsMendez", "something json", time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get(ctx, "JamsMendez").Result()
	if err == redis.Nil {
		fmt.Println("key not found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("JamsMendez", val)
	}

	time.Sleep(2 * time.Minute)

	val, err = client.Get(ctx, "JamsMendez").Result()
	if err == redis.Nil {
		fmt.Println("key not found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("JamsMendez", val)
	}
}
