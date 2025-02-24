package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func createClient() *redis.Client {

	opt, err := redis.ParseURL("redis://default:jQRzNvLWa5ysyomgZf1rbHwDgrwMJ0cG@redis-19131.c336.samerica-east1-1.gce.redns.redis-cloud.com:19131")

	if err != nil {
		log.Printf("Error auth redis: %s", err)
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

func setState(state string, phone string, isManual bool) error {
	client := getRedisClient()
	ctx := context.Background()

	currentTime := time.Now()

	err := client.JSONSet(ctx, phone, "$", State{
		state:      state,
		isManual:   isManual,
		lastUpdate: currentTime.Format("0000-00-00 00:00:00"),
	}).Err()

	return err

}

func getState(phone string) (string, error) {

	client := getRedisClient()
	ctx := context.Background()

	val, err := client.JSONGet(ctx, phone, "$").Result()

	return val, err
}
