package repository

import (
	"com.fiapchallenge/tech-challange-auth-issuer/database"
	"context"
	_ "database/sql"
)

func GetCustomerIdByDocument(ctx context.Context, document string) (string, error) {
	var id string
	err := database.DB.QueryRowContext(ctx,
		"SELECT id FROM customer WHERE cpf_cnpj = $1",
		document,
	).Scan(&id)
	return id, err
}

type UserInfo struct {
	ID       string
	Role     string
	Password string
}

func GetUserByEmail(ctx context.Context, email string) (*UserInfo, error) {
	var info UserInfo
	err := database.DB.QueryRowContext(ctx,
		`SELECT id, password, role FROM "user" WHERE email = $1`,
		email,
	).Scan(&info.ID, &info.Password, &info.Role)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
