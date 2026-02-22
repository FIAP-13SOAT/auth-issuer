package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// Estrutura de entrada
type Input struct {
	Document string `json:"document"`
}

// Estrutura de saída
type Output struct {
	Token string `json:"token"` // Token JWT
}

// handler para autenticar
func handler(ctx context.Context, input Input) (Output, error) {
	token := "teste"

	// Criar o output
	output := Output{
		Token: token,
	}

	return output, nil
}

func main() {
	lambda.Start(handler)
}
