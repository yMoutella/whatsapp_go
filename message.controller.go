package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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

func getToken() (string, error) {

	err := godotenv.Load()

	GRAPH_API_TOKEN := os.Getenv("GRAPH_API_TOKEN")

	return GRAPH_API_TOKEN, err
}

func setDto() dtoMessage {

	var dto dtoMessage
	dto.createDto(messageInterface)

	return dto
}

func sendMessage() {

	GRAPH_API_TOKEN, err := getToken()

	if err != nil {
		log.Printf("Error getting token: %s", err)
	}

	dto := setDto()
	message, err := getMenu(dto.List_Reply)
	message.To = dto.PhoneNumber

	jsonData, err := json.Marshal(message)

	if err != nil {
		log.Printf("Error marshaling data: %s", err)
	}

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v22.0/"+dto.PhoneNumberId+"/messages", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Printf("Error in request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+GRAPH_API_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error requesting the client: %s", err)
	}

	fmt.Printf("resp.Body: %v\n", resp)
}
