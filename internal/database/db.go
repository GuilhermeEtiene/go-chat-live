package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é a instância global do banco de dados PostgreSQL
var DB *gorm.DB

// ConnectDB estabelece conexão com o banco PostgreSQL usando variáveis de ambiente.
// Configura valores padrão caso as variáveis não estejam definidas.
func ConnectDB() {
	var err error

	// Configurações do PostgreSQL via variáveis de ambiente
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "chatuser"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "chatpass"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "chatdb"
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	log.Printf("Connecting to PostgreSQL: %s:%s/%s", host, port, dbname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection error:", err)
	}

	log.Println("PostgreSQL connection established successfully!")
}
