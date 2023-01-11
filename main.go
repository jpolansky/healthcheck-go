package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	router := gin.Default()
	router.GET("/health-check",healthCheck)
	router.GET("/echo/:echo", func(c *gin.Context) {
		echo := c.Param("echo")
		c.JSON(http.StatusOK, gin.H{
			"echo": echo,
		})
	})
	fmt.Println("Running on /echo:echo and /health-check")
	router.Run()
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message":"success"})
}