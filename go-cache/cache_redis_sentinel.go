package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func runCacheRedisSentinel() {
	sentinelClient := redis.NewFailoverClient(&redis.FailoverOptions{
		SentinelAddrs: []string{"localhost:26379", "localhost:26380", "localhost:26381"},
		MasterName:    "name-master",
		Password:      "", // set your Redis password here
		DB:            0,
	})

	ctx := context.Background()

	err := sentinelClient.Set(ctx, "JamsMendez", "something JSON", time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := sentinelClient.Get(ctx, "JamsMendez").Result()
	if err == redis.Nil {
		fmt.Println("key not found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("JamsMendez:", val)
	}

	time.Sleep(2 * time.Second)

	val, err = sentinelClient.Get(ctx, "JamsMendez").Result()
	if err == redis.Nil {
		fmt.Println("key not found")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("JamsMendez:", val)
	}
}
