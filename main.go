package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", challengeController)
	r.POST("/", messageController)
	err := r.Run(":8080")

	if err != nil {
		log.Printf("Error on run server: %s", err)
	}

	fmt.Printf("Running in port 8080")

}
