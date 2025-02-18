package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func messageController(c *gin.Context) {
	setValue()
	c.JSON(http.StatusOK, getValue())
}
