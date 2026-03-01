package handler

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"example.com/tech-challange-auth-issuer/models"
	"example.com/tech-challange-auth-issuer/repository"
	"example.com/tech-challange-auth-issuer/service"
	"example.com/tech-challange-auth-issuer/validator"
)

func AuthHandler(ctx context.Context, input models.Input) (models.Output, error) {
	if len(strings.TrimSpace(input.Document)) == 0 {
		return models.Output{}, errors.New("O campo 'document' é obrigatório")
	}

	if !validator.IsValidDocument(input.Document) {
		return models.Output{}, errors.New("CPF/CNPJ inválido")
	}

	customerID, err := repository.GetCustomerID(ctx, input.Document)
	if err == sql.ErrNoRows {
		return models.Output{}, errors.New("Usuário não encontrado")
	}
	if err != nil {
		return models.Output{}, err
	}

	token, err := service.GenerateToken(customerID)

	if err == sql.ErrNoRows {
		return models.Output{}, errors.New("Ocorreu um erro desconhecido, por favor tente novamente.")
	}
	if err != nil {
		return models.Output{}, err
	}

	return models.Output{Token: token}, nil
}
