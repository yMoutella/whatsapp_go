package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", challengeController)
	r.POST("/", messageController)
	r.Run(":8080")

}
