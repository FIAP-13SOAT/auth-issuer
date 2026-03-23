package database

import (
	"database/sql"
	"fmt"
	"log"

	"com.fiapchallenge/tech-challange-auth-issuer/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar configuração: %v", err))
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Erro ao abrir conexão: %v", err)
		panic(fmt.Sprintf("Erro ao conectar: %v", err))
	}

	if err = DB.Ping(); err != nil {
		log.Printf("Erro ao verificar conexão: %v", err)
		panic(fmt.Sprintf("Banco inacessível: %v", err))
	}

	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(2)
}
