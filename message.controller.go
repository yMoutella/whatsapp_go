package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var messageInterface MessageInterface

func messageController(c *gin.Context) {

	if err := c.ShouldBindJSON(&messageInterface); err != nil {
		panic(err)
	}

	fmt.Println(messageInterface)

	c.Status(200)

	sendMessage()

}

func sendMessage() {
	//GRAPH_API_TOKEN := os.Getenv("GRAPH_API_TOKEN")
	menu := getMenu(1)

	var dto dtoMessage
	dto.createDto(messageInterface)

	_, err := json.Marshal(dto)

	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://graph.facebook.com/v22.0/"+dto.PhoneNumberId+"/messages", "application/json", bytes.NewBuffer([]byte(menu)))

	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("resp.Body: %v\n", string(data))
}
