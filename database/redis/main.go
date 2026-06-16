package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("A função demorou: %v\n", time.Since(start))
	}()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	defer rdb.Close()
	ctx := context.Background()

	err := rdb.Set(ctx, "chave", "valor?", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "chave").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("chave?", val)

}
