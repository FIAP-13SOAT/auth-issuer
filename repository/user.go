package repository

import (
	"com.fiapchallenge/tech-challange-auth-issuer/database"
	"context"
	_ "database/sql"
)

func GetCustomerIdByDocument(ctx context.Context, document string) (string, error) {
	// Normaliza o documento removendo pontuação
	cleaned := cleanDocument(document)

	var id string
	err := database.DB.QueryRowContext(ctx,
		"SELECT id FROM customer WHERE REPLACE(REPLACE(REPLACE(cpf_cnpj, '.', ''), '-', ''), '/', '') = $1",
		cleaned,
	).Scan(&id)
	return id, err
}

// cleanDocument remove pontuação do CPF/CNPJ
func cleanDocument(doc string) string {
	result := ""
	for _, c := range doc {
		if c >= '0' && c <= '9' {
			result += string(c)
		}
	}
	return result
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
