package chat

import (
	"sync"
)

type Hub struct {
	clients    map[string][]*Client // Sala → Clientes conectados
	register   chan *Client
	unregister chan *Client
	broadcast  chan Message
	mu         sync.Mutex
}

type Message struct {
	RoomID  string
	Payload []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string][]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan Message),
	}
}

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
				c.Send <- msg.Payload
			}
			h.mu.Unlock()
		}
	}
}
