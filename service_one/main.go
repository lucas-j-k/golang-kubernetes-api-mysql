package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define a route
	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(200, gin.H{"message": "Pong from service_one"})
	})

	router.GET("/hello", func(c *gin.Context) {
		// Return JSON response
		c.JSON(200, gin.H{"message": "Goodbye"})
	})

	// Start the server
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
