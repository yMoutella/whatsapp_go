package main

import (
	"encoding/json"
	"log"
)

var lastOption bool

func isManual(dto dtoMessage) bool {

	if lastOption {
		return true
	}

	dbState, err := getState(dto.PhoneNumber)

	if err != nil {
		log.Printf("Error getting state: %s", err)
	}

	var state State
	json.Unmarshal([]byte(dbState), &state)

	if state.isManual {
		return true
	}

	return false

}

func isStartingChat(dto string) bool {

	dbState, err := getState(dto)

	if err != nil {
		log.Printf("Error getting state: %s", err)
	}

	if dbState == "" {
		return true
	}

	return false

}
