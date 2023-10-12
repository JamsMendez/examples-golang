package producer

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
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
		time.Sleep(2 * time.Second)

		uuID := uuid.New()

		cmd := client.RPush(ctx, redisList, uuID.String())

		_, err := cmd.Result()
		if err != nil {
			log.Println("redis.RPush.ERROR: ", err)

			continue
		}

		fmt.Print("\033[H\033[2J")

		fmt.Println("Redis Productor: ")
		fmt.Println("=================")

		cmd = client.LLen(ctx, redisList)
		total, err := cmd.Result()
		if err != nil {
			log.Println("redis.LLen.ERROR: ", err)

			continue
		}

		fmt.Printf("Total elements: %d\n", total)

		sliceCmd := client.LRange(ctx, redisList, 0, -1)
		values, err := sliceCmd.Result()
		if err != nil {
			log.Println("redis.LRange.ERROR: ", err)

			continue
		}

		fmt.Printf("LIST:\n%s\n", strings.Join(values, ", "))
	}
}
