package main

import (
	"example.com/tech-challange-auth-issuer/database"
	"example.com/tech-challange-auth-issuer/handler"
	"example.com/tech-challange-auth-issuer/service"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	if err := service.Init(); err != nil {
		log.Fatalf("Erro ao inicializar service: %v", err)
	}

	database.Init()
	lambda.Start(handler.AuthHandler)
}
