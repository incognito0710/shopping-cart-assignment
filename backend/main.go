package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Allow cross-origin requests from your React dev server
	r.Use(cors.Default())

	// Connect DB (Postgres)
	ConnectDatabase()

	// Register routes
	SetupRoutes(r)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
