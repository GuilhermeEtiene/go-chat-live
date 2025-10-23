// Package main implements the REST API server for the chat application.
// This server handles user management, authentication and CRUD operations.
package main

import (
	"log"
	"os"

	"go-chat-live/internal/database"
	"go-chat-live/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// main initializes and starts the REST API server with database connection,
// CORS middleware, and user management routes.
func main() {
	loadEnvironmentVariables()
	setupDatabase()

	router := setupRouter()
	setupRoutes(router)

	startServer(router)
}

// loadEnvironmentVariables loads configuration from .env file
func loadEnvironmentVariables() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: could not load .env file, using system variables")
	}
}

// setupDatabase initializes database connection and runs migrations
func setupDatabase() {
	database.ConnectDB()
	database.DB.AutoMigrate(&user.User{})
}

// setupRouter creates Gin router with CORS middleware
func setupRouter() *gin.Engine {
	r := gin.Default()

	// CORS middleware for cross-origin requests
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	return r
}

// setupRoutes defines all API endpoints for user management
func setupRoutes(r *gin.Engine) {
	r.POST("/users", user.CreateUser)
	r.POST("/login", user.LoginUser)
	r.GET("/users", user.ListUsers)
	r.GET("/users/:id", user.GetUserById)
	r.PUT("/users/:id", user.UpdateUser)
	r.DELETE("/users/:id", user.DeleteUser)
}

// startServer starts the HTTP server on configured port
func startServer(r *gin.Engine) {
	port := os.Getenv("REST_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("REST API server starting on port %s", port)
	r.Run(":" + port)
}
