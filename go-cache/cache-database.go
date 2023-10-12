package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

func runCacheRedisDatabase() {
	var sql sql.DB
	getUsers(context.Background(), &sql)
}

func NewRedisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return client
}

func getUsers(ctx context.Context, db *sql.DB) ([]User, error) {
	// Get Redis client
	redisClient := NewRedisClient(ctx)

	// Check if the data is available in the Redis cache
	val, err := redisClient.Get(ctx, "users").Result()
	if err == redis.Nil {
		// If the data is not available in the Redis cache, query the database
		rows, err := db.QueryContext(ctx, "SELECT * FROM users")
		if err != nil {
		}
		defer rows.Close()

		// Loop through the rows and store the user data in a slice
		users := make([]User, 0)
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email)
			if err != nil {
				return nil, err
			}
			users = append(users, user)
		}

		// Cache the user data in Redis
		data, err := json.Marshal(users)
		if err != nil {
			return nil, err
		}
		err = redisClient.Set(ctx, "users", data, time.Minute*5).Err()
		if err != nil {
			return users, err
		}

		// Return the user data
		return users, nil

	} else if err != nil {
		return nil, err
	}

	// If the data is available in the Redis cache, unmarshal the data and return it
	var users []User
	err = json.Unmarshal([]byte(val), &users)
	if err != nil {

	}

	return users, nil
}
