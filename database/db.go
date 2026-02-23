package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	env := getEnv("ENVIRONMENT", "dev")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv(fmt.Sprintf("DB_HOST_%s", strings.ToUpper(env))),
		os.Getenv(fmt.Sprintf("DB_PORT_%s", strings.ToUpper(env))),
		os.Getenv(fmt.Sprintf("DB_USER_%s", strings.ToUpper(env))),
		os.Getenv(fmt.Sprintf("DB_PASSWORD_%s", strings.ToUpper(env))),
		os.Getenv(fmt.Sprintf("DB_NAME_%s", strings.ToUpper(env))),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("Erro ao conectar: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	if err = DB.Ping(); err != nil {
		panic(fmt.Sprintf("Erro ao pingar: %v", err))
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
