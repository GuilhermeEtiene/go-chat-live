package chat

import (
	"github.com/gorilla/websocket"
)

// Client represents an active WebSocket connection from a user in the chat.
// Contains user information and channels for asynchronous communication.
type Client struct {
	ID        string          // Unique connection identifier
	Conn      *websocket.Conn // Active WebSocket connection
	RoomID    string          // ID of the room the client is connected to
	Send      chan []byte     // Channel for sending messages
	UserID    uint            // Authenticated user ID
	UserName  string          // User's name
	UserEmail string          // User's email
}
