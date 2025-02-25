package main

import (
	//	"bytes"
	//	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	//	"net/http"
	"os"
)

var messageInterface MessageInterface

func messageController(c *gin.Context) {

	if err := c.ShouldBindJSON(&messageInterface); err != nil {
		log.Printf("Error binding json: %s", err)
	}

	fmt.Println(messageInterface)

	c.Status(200)

	sendMessage()

}

func getToken() string {

	err := godotenv.Load()

	GRAPH_API_TOKEN := os.Getenv("GRAPH_API_TOKEN")

	if err != nil {
		log.Fatalf("Error to get env variable: %s", err)
	}

	return GRAPH_API_TOKEN
}

func setDto() dtoMessage {

	var dto dtoMessage
	dto.createDto(messageInterface)

	return dto
}

type State struct {
	state      string
	lastUpdate string
	isManual   bool
}

type RequestPayload struct {
	jsonData     []byte
	pathVariable string
	token        string
}

func createPayload() RequestPayload {

	//	dto := setDto()
	//
	//	token := getToken()
	//	message, err := getMenu(opt)
	//
	//	message.To = dto.PhoneNumber
	//
	//	if err != nil {
	//		log.Fatalf("Error getting menu options: %s", err)
	//	}
	//
	//	jsonData, err := json.Marshal(message)
	//
	//	if err != nil {
	//		log.Printf("Error marshaling data: %s", err)
	//	}
	//
	//	return RequestPayload{
	//		jsonData:     jsonData,
	//		pathVariable: dto.PhoneNumberId,
	//		token:        token,
	//	}

	return RequestPayload{}
}

func sendMessage() {

	// req, err := http.NewRequest("POST", "https://graph.facebook.com/v22.0/"+payload.pathVariable+"/messages", bytes.NewBuffer(payload.jsonData))
	//
	//	if err != nil {
	//		log.Printf("Error in request: %s", err)
	//	}
	//
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer "+payload.token)
	//
	// client := &http.Client{}
	// resp, err := client.Do(req)
	//
	//	if err != nil {
	//		log.Printf("Error requesting the client: %s", err)
	//	}
	//
	// fmt.Printf("resp.Body: %v\n", resp)
}
