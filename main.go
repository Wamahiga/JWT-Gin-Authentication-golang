package main

import (
	"github.com/Wamahiga/JWT-Gin-Authentication-golang/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT") // Retrieve the PORT environment variable

	if port == "" { // Check if the PORT environment variable is not set
		port = "8000" // Set default port to 8000 if the environment variable is not set
	}

	router := gin.New()      // Create a new Gin router
	router.Use(gin.Logger()) // Apply middleware to log HTTP requests

	routes.AuthRoutes(router) // Set up authentication-related routes
	routes.UserRoutes(router) // Set up user-related routes

	router.GET("/api-1", func(c *gin.Context) { // Define a GET route for /api-1
		c.JSON(200, gin.H{"Success": "Access granted for api-1!"}) // Respond with a JSON message
	})

	router.GET("/api-2", func(c *gin.Context) { // Define a GET route for /api-2
		c.JSON(200, gin.H{"Success": "Access granted for api-2!"}) // Respond with a JSON message
	})

	router.Run(":" + port) // Start the server on the specified port
}
