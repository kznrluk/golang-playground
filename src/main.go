package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"time"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	for count := 5000; count > 0; count-- {
		rdb.ZAdd(ctx, "ranking", &redis.Z{
			Score:  float64(rand.Intn(5000000000)),
			Member: "HELLO! " + string(count),
		})
	}

	fmt.Println(val)
}
