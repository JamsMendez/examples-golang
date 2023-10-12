package consumer

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

const redisList = "list_uuid"

func Start() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	for {
		time.Sleep(3 * time.Second)

		fmt.Print("\033[H\033[2J")

		fmt.Println("Redis Consumer: ")
		fmt.Println("================")

		valuesCmd := client.BLMPop(ctx, time.Second, "LEFT", 1, redisList)
		key, values, err := valuesCmd.Result()
		if err != nil {
			log.Println("redis.BLMPop.ERROR: ", err)

			continue
		}

		/*
			sliceCmd := client.BLPop(ctx, time.Second, redisList)
			// values [key, values]
			values, err := sliceCmd.Result()
			if err != nil {
				log.Println("redis.BLPop.ERROR: ", err)

				continue
			}
		*/

		if len(values) == 0 {
			// fmt.Println("redis.BLPop.[]String: ", values)
			fmt.Println("redis.BLMPop.[]String: ", values)

			continue
		}

		value := strings.Join(values, ", ")
		fmt.Printf("Elements consumed %s: %s\n", key, value)

		cmd := client.LLen(ctx, redisList)
		n, err := cmd.Result()
		if err != nil {
			fmt.Println("redis.LLen.ERROR: ", err)
		}

		fmt.Printf("Elements pending %d\n", n)
	}
}
