// Package main implements the WebSocket server for real-time chat communication.
// This server handles WebSocket connections, user authentication via JWT,
// and message broadcasting between clients in chat rooms.
package main

import (
	"log"
	"net/http"
	"os"

	"go-chat-live/internal/chat"
	"go-chat-live/internal/database"
	"go-chat-live/internal/user"

	"github.com/joho/godotenv"
)

// main initializes and starts the WebSocket server with database connection,
// chat hub for managing connections, and JWT-authenticated WebSocket endpoint.
func main() {
	loadEnvironmentVariables()
	setupDatabase()

	hub := initializeChatHub()
	setupWebSocketEndpoint(hub)

	startWebSocketServer()
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

// initializeChatHub creates and starts the chat hub in a separate goroutine
func initializeChatHub() *chat.Hub {
	hub := chat.NewHub()
	go hub.Run()
	return hub
}

// setupWebSocketEndpoint configures the /ws endpoint for WebSocket connections
func setupWebSocketEndpoint(hub *chat.Hub) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})
}

// startWebSocketServer starts the HTTP server on configured port
func startWebSocketServer() {
	port := os.Getenv("WS_PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("WebSocket server starting on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("WebSocket server error:", err)
	}
}
