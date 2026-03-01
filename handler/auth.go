package handler

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"example.com/tech-challange-auth-issuer/models"
	"example.com/tech-challange-auth-issuer/repository"
	"example.com/tech-challange-auth-issuer/service"
)

func AuthHandler(ctx context.Context, input models.Input) (models.Output, error) {
	if len(strings.TrimSpace(input.Document)) == 0 {
		return models.Output{}, errors.New("O campo 'document' é obrigatório")
	}

	customerID, err := repository.GetId(ctx, input.Document)
	if err == sql.ErrNoRows {
		return models.Output{}, errors.New("Usuário não encontrado")
	}
	if err != nil {
		return models.Output{}, err
	}

	token, err := service.Generate(customerID)

	if err == sql.ErrNoRows {
		return models.Output{}, errors.New("Ocorreu um erro desconhecido, por favor tente novamente.")
	}
	if err != nil {
		return models.Output{}, err
	}

	return models.Output{Token: token}, nil
}
