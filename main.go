package main

import (
	"example.com/tech-challange-auth-issuer/database"
	"example.com/tech-challange-auth-issuer/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	database.Init()
	lambda.Start(handler.AuthHandler)
}
