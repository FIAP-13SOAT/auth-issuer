package repository

import (
	"context"
	_ "database/sql"
	"example.com/tech-challange-auth-issuer/database"
)

func GetCustomerID(ctx context.Context, document string) (string, error) {
	var customerID string
	err := database.DB.QueryRowContext(ctx,
		"SELECT id FROM customer WHERE cpf_cnpj = $1",
		document,
	).Scan(&customerID)

	return customerID, err
}
