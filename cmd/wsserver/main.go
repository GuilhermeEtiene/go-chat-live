package main

import (
	"log"
	"net/http"
	"os"

	"go-chat-live/internal/chat"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: erro ao carregar .env, usando variáveis do sistema")
	}

	hub := chat.NewHub()
	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	port := os.Getenv("WS_PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("Servidor WebSocket rodando em localhost:" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Erro no servidor:", err)
	}
}
