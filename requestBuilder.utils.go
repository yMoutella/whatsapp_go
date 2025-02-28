package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var lastOption bool

func isStartingChat(phone string) bool {

	dbState, err := getState(phone)

	if err != nil {
		log.Printf("Error getting state: %s", err)
	}

	if dbState == "" {
		return true
	}
	return false

}

type RequestPayload struct {
	jsonData     []byte
	pathVariable string
	token        string
}

func getToken() string {

	err := godotenv.Load()

	GRAPH_API_TOKEN := os.Getenv("GRAPH_API_TOKEN")

	if err != nil {
		log.Fatalf("Error to get env variable: %s", err)
	}

	return GRAPH_API_TOKEN
}

func buildPayload(dto dtoMessage) RequestPayload {

	menuOption := dto.List_Reply
	var message Message_List_Request

	if isStartingChat(dto.PhoneNumber) {
		_message, err := getMenu("1")

		if err != nil {
			log.Fatalf("Error in getting menu options: %s", err)
		}
		message = _message

	} else {
		_message, err := getMenu(menuOption)
		message = _message

		if err != nil {
			log.Fatalf("Error in getting menu options: %s", err)
		}

		isManual := isManual(dto.PhoneNumber)

		setState(dto.List_Reply, dto.PhoneNumber, isManual)
	}

	message.To = dto.PhoneNumber

	jsonData, err := json.Marshal(message)

	if err != nil {
		log.Fatalf("Error marshaling  jsonData: %s", err)
	}

	token := getToken()

	return RequestPayload{
		jsonData:     jsonData,
		token:        token,
		pathVariable: dto.PhoneNumber,
	}
}
