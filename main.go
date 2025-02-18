package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", challengeController)
	r.POST("/", messageController)
	r.Run(":8080")

	fmt.Printf("Running in port 8080")

}
