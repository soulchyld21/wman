package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up Gin router
	router := gin.Default()

	// Route to receiving service
	receivingGroup := router.Group("/receiving")
	{
		receivingGroup.POST("/", receivingHandler)
	}

	// Route to movement service
	movementGroup := router.Group("/movement")
	{
		movementGroup.POST("/", movementHandler)
	}

	// Route to picking service
	pickingGroup := router.Group("/picking")
	{
		pickingGroup.POST("/", pickingHandler)
	}

	// Route to consolidation service
	consolidationGroup := router.Group("/consolidation")
	{
		consolidationGroup.POST("/", consolidationHandler)
	}

	// Start server
	port := os.Getenv("PORT")
	log.Printf("Gateway listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func receivingHandler(c *gin.Context) {
	// Forward request to receiving service
	// TODO: implement forwarding logic
}

func movementHandler(c *gin.Context) {
	// Forward request to movement service
	// TODO: implement forwarding logic
}

func pickingHandler(c *gin.Context) {
	// Forward request to picking service
	// TODO: implement forwarding logic
}

func consolidationHandler(c *gin.Context) {
	// Forward request to consolidation service
	// TODO: implement forwarding logic
}
