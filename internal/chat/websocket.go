package chat

import (
	"log"
	"net/http"

	"go-chat-live/internal/user"

	"github.com/gorilla/websocket"
)

// upgrader configura o upgrade de conexões HTTP para WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permite conexões de qualquer origem
	},
}

// ServeWs handles WebSocket requests and authenticates users via JWT.
// Creates a new client and registers it in the Hub for real-time communication.
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room")
	if roomID == "" {
		http.Error(w, "Room ID is required", http.StatusBadRequest)
		return
	}

	// Validate JWT token
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "JWT token is required", http.StatusUnauthorized)
		return
	}

	claims, err := user.ValidateJWT(token)
	if err != nil {
		http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
		return
	}

	userID, _ := claims["user_id"].(float64)
	email, _ := claims["email"].(string)

	log.Printf("Attempting to find user ID: %v", userID)

	// Fetch complete user data
	userData, err := user.FindById(int(userID))
	if err != nil {
		log.Printf("User not found for ID %v: %v", userID, err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade error:", err)
		return
	}

	client := &Client{
		ID:        r.RemoteAddr,
		Conn:      conn,
		RoomID:    roomID,
		Send:      make(chan []byte),
		UserID:    userData.ID,
		UserName:  userData.Name,
		UserEmail: email,
	}

	hub.register <- client

	go client.readPump(hub)
	go client.writePump()
}

// readPump reads messages from WebSocket connection and sends them to the Hub.
// Runs in a separate goroutine for each client.
func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		hub.broadcast <- Message{
			RoomID:   c.RoomID,
			Content:  string(msg),
			UserName: c.UserName,
			Sender:   c,
		}
	}
}

// writePump sends messages from the Send channel to the WebSocket connection.
// Runs in a separate goroutine for each client.
func (c *Client) writePump() {
	defer c.Conn.Close()

	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
