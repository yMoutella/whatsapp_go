package main

import (
	"context"
	"encoding/json"
	"log"
)

func getMenu[T Message_List_Request](opt string) (T, error) {

	ctx := context.Background()
	client := getRedisClient()

	menu, err := client.HGet(ctx, "menu", opt).Result()

	if err != nil {
		log.Fatalf("Error getting hash: %s", err)
	}

	var menuObject T

	json.Unmarshal([]byte(menu), &menuObject)

	return menuObject, nil
}
