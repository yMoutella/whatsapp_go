package main

import (
	"context"
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
		return createClient()
	}

	return redisClient

}

func setValue() {
	client := getRedisClient()
	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()

	if err != nil {
		panic(err)
	}

	return

}

func getValue() string {

	client := getRedisClient()
	ctx := context.Background()
	val, err := client.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	client.Close()

	return val
}
