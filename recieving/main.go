package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/soulchyld/wman/database"
)

func main() {
	// Set up database connection and migrate tables
	db := database.Connect()
	defer db.Close()

	// Set up Gin router
	router := gin.Default()

	// Set up routes
	// router.POST("/receiving", createReceivingHandler)

	// Start server
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server")
	}
}
