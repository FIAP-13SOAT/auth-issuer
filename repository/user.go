package repository

import (
	"context"
	_ "database/sql"
	"example.com/tech-challange-auth-issuer/database"
)

func GetTokenByDocument(ctx context.Context, document string) (string, error) {
	var token string
	// todo: montar token
	err := database.DB.QueryRowContext(ctx,
		"SELECT id FROM users WHERE cpf_cnpj = $1",
		document,
	).Scan(&token)

	return token, err
}
