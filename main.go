package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Template struct {
	Name      string            `json:"template"`
	Variables map[string]string `json:"variables"`
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the root endpoint
	router.GET("/generateDocument", func(c *gin.Context) {

		// Read the request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}

		// Parse the JSON data into a struct
		var template Template
		err = json.Unmarshal(body, &template)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		// send response
		c.JSON(http.StatusOK, gin.H{"document": generateDocument(template)})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
