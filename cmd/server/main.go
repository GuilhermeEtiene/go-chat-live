package main

import (
	"log"
	"net/http"
	"os"

	"go-chat-live/internal/database"
	"go-chat-live/internal/user"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: erro ao carregar .env, usando variáveis do sistema")
	}

	database.ConnectDB()
	database.DB.AutoMigrate(&user.User{})

	r := mux.NewRouter()

	r.HandleFunc("/users", user.CriarUsuario).Methods("POST")
	r.HandleFunc("/users", user.ListUsuarios).Methods("GET")
	r.HandleFunc("/users/{id}", user.BuscarUsuarioPorId).Methods("GET")
	r.HandleFunc("/users/{id}", user.AtualizarUsuario).Methods("PUT")
	r.HandleFunc("/users/{id}", user.DeletarUsuario).Methods("DELETE")

	port := os.Getenv("REST_PORT")
	if port == "" {
		port = "8080" // padrão se variável não existir
	}

	log.Println("Servidor REST rodando em localhost:" + port)
	http.ListenAndServe(":"+port, r)
}
