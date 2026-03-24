package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"com.fiapchallenge/tech-challange-auth-issuer/models"
	"com.fiapchallenge/tech-challange-auth-issuer/repository"
	"com.fiapchallenge/tech-challange-auth-issuer/service"
	"com.fiapchallenge/tech-challange-auth-issuer/validator"
	"github.com/aws/aws-lambda-go/events"
	"golang.org/x/crypto/bcrypt"
)

func AuthHandler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	path := request.RawPath
	if path == "/admin/login" {
		return handleAdminLogin(ctx, request)
	}
	return handleCustomerLogin(ctx, request)
}

func handleCustomerLogin(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var input models.Input
	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		return jsonResponse(400, `{"error":"JSON inválido"}`)
	}

	if len(strings.TrimSpace(input.Document)) == 0 {
		return jsonResponse(400, `{"error":"O campo 'document' é obrigatório"}`)
	}

	if !validator.IsValidDocument(input.Document) {
		return jsonResponse(400, `{"error":"CPF/CNPJ inválido"}`)
	}

	customerID, err := repository.GetCustomerIdByDocument(ctx, input.Document)
	if err == sql.ErrNoRows {
		return jsonResponse(404, `{"error":"Usuário não encontrado"}`)
	}
	if err != nil {
		return jsonResponse(500, `{"error":"Erro interno"}`)
	}

	token, err := service.GenerateToken(customerID, "CUSTOMER")
	if err != nil {
		return jsonResponse(500, `{"error":"Erro ao gerar token"}`)
	}

	responseBody, _ := json.Marshal(models.Output{Token: token})
	return jsonResponse(200, string(responseBody))
}

func handleAdminLogin(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	var input models.AdminInput
	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		return jsonResponse(400, `{"error":"JSON inválido"}`)
	}

	if len(strings.TrimSpace(input.Email)) == 0 || len(strings.TrimSpace(input.Password)) == 0 {
		return jsonResponse(400, `{"error":"Os campos 'email' e 'password' são obrigatórios"}`)
	}

	user, err := repository.GetUserByEmail(ctx, input.Email)
	if err == sql.ErrNoRows {
		return jsonResponse(401, `{"error":"Credenciais inválidas"}`)
	}
	if err != nil {
		return jsonResponse(500, `{"error":"Erro interno"}`)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return jsonResponse(401, `{"error":"Credenciais inválidas"}`)
	}

	token, err := service.GenerateToken(user.ID, user.Role)
	if err != nil {
		return jsonResponse(500, `{"error":"Erro ao gerar token"}`)
	}

	responseBody, _ := json.Marshal(models.Output{Token: token})
	return jsonResponse(200, string(responseBody))
}

func jsonResponse(status int, body string) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       body,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}
