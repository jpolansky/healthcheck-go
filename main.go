package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/health-check",healthCheck)
	router.Run("localhost:8080")
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"success"})
}