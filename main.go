package main

import (
	"com.fiapchallenge/tech-challange-auth-issuer/database"
	"com.fiapchallenge/tech-challange-auth-issuer/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	database.Init()
	lambda.Start(handler.AuthHandler)
}
