package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("CHANGE_ME_SUPER_SECRET")

func Generate(customerID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:  customerID,
		Issuer:   "tech-challange-auth-issuer",
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(
			time.Now().Add(15 * time.Minute),
		),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
