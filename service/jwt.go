package service

import (
	"fmt"
	"time"

	"com.fiapchallenge/tech-challange-auth-issuer/config"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func GenerateToken(subjectID string, role string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar configuração: %v", err))
	}

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subjectID,
			Issuer:    cfg.JWT.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWT.Expiration) * time.Minute)),
		},
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}
