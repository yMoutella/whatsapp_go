package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func messageController(c *gin.Context) {

	err := setValue()

	if err != nil {
		panic(err)
	}

	value, err := getValue()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, value)

	getRedisClient().Close()
}
