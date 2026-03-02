package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	env := getEnv("ENVIRONMENT", "dev")

	host := os.Getenv(fmt.Sprintf("DB_HOST_%s", strings.ToUpper(env)))
	port := os.Getenv(fmt.Sprintf("DB_PORT_%s", strings.ToUpper(env)))
	user := os.Getenv(fmt.Sprintf("DB_USER_%s", strings.ToUpper(env)))
	password := os.Getenv(fmt.Sprintf("DB_PASSWORD_%s", strings.ToUpper(env)))
	dbname := os.Getenv(fmt.Sprintf("DB_NAME_%s", strings.ToUpper(env)))

	log.Printf("Conectando ao banco: host=%s port=%s user=%s dbname=%s", host, port, user, dbname)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Erro ao abrir conexão: %v", err)
		panic(fmt.Sprintf("Erro ao conectar: %v", err))
	}

	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(2)

	log.Println("Pool de conexões configurado")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
