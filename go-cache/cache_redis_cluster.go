package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func runCacheRedisCluster() {
	cluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"node1:6379", "node2:6379", "node2:6379"},
	})

	ctx := context.Background()

	err := cluster.Ping(ctx).Err()
	if err != nil {
		log.Fatal(err)
	}

	err = cluster.Set(ctx, "JamsMendez", "something JSON", 0).Err()
	if err != nil {
		log.Fatal()
	}

	val, err := cluster.Get(ctx, "JamsMendez").Result()
	if err != nil {
		log.Fatal()
	}

	fmt.Println("JamsMendez: ", val)

	err = cluster.Del(ctx, "JamsMendez").Err()
	if err != nil {
		log.Fatal(err)
	}
}
