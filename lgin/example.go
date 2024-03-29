package main

import "github.com/gin-gonic/gin"

func main() {
	// Creates a gin router with default middleware
	// logger and recovery (crash-free) middleware
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
