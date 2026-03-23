package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strings"

	"com.fiapchallenge/tech-challange-auth-issuer/models"
	"com.fiapchallenge/tech-challange-auth-issuer/repository"
	"com.fiapchallenge/tech-challange-auth-issuer/service"
	"com.fiapchallenge/tech-challange-auth-issuer/validator"
	"github.com/aws/aws-lambda-go/events"
)

func AuthHandler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var input models.Input
	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"error":"JSON inválido"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	if len(strings.TrimSpace(input.Document)) == 0 {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"error":"O campo 'document' é obrigatório"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	if !validator.IsValidDocument(input.Document) {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 400,
			Body:       `{"error":"CPF/CNPJ inválido"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	customerID, err := repository.GetId(ctx, input.Document)
	if err == sql.ErrNoRows {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 404,
			Body:       `{"error":"Usuário não encontrado"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"error":"Erro interno"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	token, err := service.GenerateToken(customerID)
	if err != nil {
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 500,
			Body:       `{"error":"Erro ao gerar token"}`,
			Headers:    map[string]string{"Content-Type": "application/json"},
		}, nil
	}

	responseBody, _ := json.Marshal(models.Output{Token: token})
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(responseBody),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}
