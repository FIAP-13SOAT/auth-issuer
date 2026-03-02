package repository

import (
	"com.fiapchallenge/tech-challange-auth-issuer/database"
	"context"
	_ "database/sql"
)

func GetId(ctx context.Context, document string) (string, error) {
	var token string
	err := database.DB.QueryRowContext(ctx,
		"SELECT id FROM customer WHERE cpf_cnpj = $1",
		document,
	).Scan(&token)

	return token, err
}
