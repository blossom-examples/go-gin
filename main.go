package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Create router
	r := gin.Default()

	// API routes
	api := r.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			name := c.DefaultQuery("name", "World")
			c.JSON(http.StatusOK, gin.H{
				"message":   fmt.Sprintf("Hello, %s!", name),
				"timestamp": time.Now().UTC().Format(time.RFC3339),
			})
		})

		api.POST("/echo", func(c *gin.Context) {
			var data map[string]interface{}
			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":   "Echo response",
				"received":  data,
				"timestamp": time.Now().UTC().Format(time.RFC3339),
			})
		})
	}

	// Serve index.html for root path
	r.GET("/", func(c *gin.Context) {
		c.File("public/index.html")
	})

	// Serve other static files
	r.Static("/static", "./public")

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	fmt.Printf("Starting server...\n")
	fmt.Printf("- Environment: %s\n", gin.Mode())
	fmt.Printf("- Port: %s\n", port)
	fmt.Printf("- URL: http://localhost:%s\n", port)
	fmt.Printf("\nReady! Visit http://localhost:%s to see the demo\n", port)

	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}