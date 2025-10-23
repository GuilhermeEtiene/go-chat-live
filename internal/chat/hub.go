package chat

import (
	"encoding/json"
	"sync"
)

// Hub manages all active WebSocket connections and distributes messages between clients.
// Uses channels for asynchronous communication and mutex for concurrency safety.
type Hub struct {
	clients    map[string][]*Client // Mapping of rooms to connected clients
	register   chan *Client         // Channel to register new clients
	unregister chan *Client         // Channel to unregister clients
	broadcast  chan Message         // Channel for message broadcasting
	mu         sync.Mutex           // Mutex for concurrency protection
}

// Message represents an internal system message to be distributed.
type Message struct {
	RoomID   string  // Target room ID
	Content  string  // Message content
	UserName string  // Sender user name
	Sender   *Client // Client who sent the message
}

// ChatMessage represents the message structure sent to the client via WebSocket.
type ChatMessage struct {
	Content  string `json:"content"`  // Message content
	UserName string `json:"userName"` // User's name
}

// NewHub creates and initializes a new Hub instance.
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string][]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),
	}
}

// Run executes the main Hub loop to process events asynchronously.
// Manages client registration/unregistration and message distribution.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.RoomID] = append(h.clients[client.RoomID], client)
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			clients := h.clients[client.RoomID]
			for i, c := range clients {
				if c == client {
					h.clients[client.RoomID] = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			h.mu.Unlock()
		case msg := <-h.broadcast:
			h.mu.Lock()
			for _, c := range h.clients[msg.RoomID] {
				if c != msg.Sender {
					chatMsg := ChatMessage{
						Content:  msg.Content,
						UserName: msg.UserName,
					}
					msgBytes, _ := json.Marshal(chatMsg)
					c.Send <- msgBytes
				}
			}
			h.mu.Unlock()
		}
	}
}
