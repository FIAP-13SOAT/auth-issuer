package repository

import (
	"context"
	_ "database/sql"
	"example.com/tech-challange-auth-issuer/database"
)

func GetTokenByDocument(ctx context.Context, document string) (string, error) {
	var token string
	// TODO: Esse código retorna o ID do cliente como Token. Isso é apenas um placeholder até a implementação real
	err := database.DB.QueryRowContext(ctx,
		"SELECT id FROM customer WHERE cpf_cnpj = $1",
		document,
	).Scan(&token)

	return token, err
}
