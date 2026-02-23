package handler

import (
	"context"
	"database/sql"
	"errors"
	"example.com/tech-challange-auth-issuer/models"
	"example.com/tech-challange-auth-issuer/repository"
	"strings"
)

func AuthHandler(ctx context.Context, input models.Input) (models.Output, error) {
	if len(strings.TrimSpace(input.Document)) == 0 {
		return models.Output{}, errors.New("O campo 'document' é obrigatório")
	}

	token, err := repository.GetTokenByDocument(ctx, input.Document)
	if err == sql.ErrNoRows {
		return models.Output{}, errors.New("Usuário não encontrado")
	}
	if err != nil {
		return models.Output{}, err
	}

	return models.Output{Token: token}, nil
}
