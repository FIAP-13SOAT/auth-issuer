package service

import (
	"fmt"
	"time"

	"example.com/tech-challange-auth-issuer/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(customerID string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar configuração: %v", err))
	}

	claims := jwt.RegisteredClaims{
		Subject:   customerID,
		Issuer:    cfg.JWT.Issuer,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWT.Expiration) * time.Minute)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
