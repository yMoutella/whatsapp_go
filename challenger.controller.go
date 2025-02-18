package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func challengeController(c *gin.Context) {

	err := godotenv.Load()

	if err != nil {

		log.Print(err)

	}

	token := os.Getenv("WEBHOOK_VERIFY_TOKEN")
	mode := c.Query("hub.mode")
	challenge := c.Query("hub.challenge")
	verify_token := c.Query("hub.verify_token")

	if mode == "subscribe" && verify_token == token {
		c.JSON(http.StatusOK, challenge)
		return
	}

	c.JSON(http.StatusBadRequest, challenge)

}
