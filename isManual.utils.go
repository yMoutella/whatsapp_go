package main

import (
	"encoding/json"
	"log"
)

func isManual(phone string) bool {

	if lastOption {
		return true
	}

	dbState, err := getState(phone)

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
