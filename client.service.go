package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func createClient() *redis.Client {

	opt, err := redis.ParseURL("redis://default:jQRzNvLWa5ysyomgZf1rbHwDgrwMJ0cG@redis-19131.c336.samerica-east1-1.gce.redns.redis-cloud.com:19131")

	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(opt)

	return redisClient

}

func getRedisClient() *redis.Client {

	if redisClient == nil {
		fmt.Print(redisClient)
		return createClient()
	}

	fmt.Print(redisClient)
	return redisClient

}

func setValue() error {
	client := getRedisClient()
	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()

	return err

}

func getValue(state string) (string, error) {

	client := getRedisClient()
	ctx := context.Background()

	val, err := client.Get(ctx, state).Result()

	return val, err
}
