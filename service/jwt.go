package service

import (
	"time"

	"example.com/tech-challange-auth-issuer/config"
	"github.com/golang-jwt/jwt/v5"
)

var cfg *config.Config

func Init() error {
	var err error
	cfg, err = config.Load()
	return err
}

func GenerateToken(customerID string) (string, error) {
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
