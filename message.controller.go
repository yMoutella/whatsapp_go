package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var messageInterface MessageInterface

func messageController(c *gin.Context) {

	if err := c.ShouldBindJSON(&messageInterface); err != nil {
		log.Printf("Error binding json: %s", err)
	}

	fmt.Println(messageInterface)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status":  200,
	})

	go sendMessage()
	return

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

func sendMessage() {

	dto := setDto()

	if isManual(dto.PhoneNumber) {
		return
	}

	payload := buildPayload(dto)

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v22.0/"+payload.pathVariable+"/messages", bytes.NewBuffer(payload.jsonData))

	if err != nil {
		log.Printf("Error in request: %s", err)
	}

	fmt.Printf("\n\n%s\n\n", payload)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+payload.token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error requesting the client: %s", err)
	}

	if resp.Status != "200" {
		log.Printf("Status not good received from facebook: %s", resp.Status)
	}

	fmt.Printf("resp.Body: %v\n", resp)
}
