package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"com.fiapchallenge/tech-challange-auth-issuer/config"
	"com.fiapchallenge/tech-challange-auth-issuer/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
)

func ValidatorHandler(ctx context.Context, request events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	token := extractToken(request.Headers)
	if token == "" {
		log.Println("Token não encontrado no header")
		return denyResponse(), nil
	}

	cfg, err := config.Load()
	if err != nil {
		log.Printf("Erro ao carregar config: %v", err)
		return denyResponse(), nil
	}

	claims, err := validateJWT(token, cfg.JWT.Secret)
	if err != nil {
		log.Printf("Token inválido: %v", err)
		return denyResponse(), nil
	}

	sub, _ := claims.GetSubject()
	role := claims.Role

	log.Printf("Token válido: sub=%s, role=%s", sub, role)
	return allowResponse(sub, role), nil
}

func extractToken(headers map[string]string) string {
	auth := headers["authorization"]
	if auth == "" {
		auth = headers["Authorization"]
	}
	if strings.HasPrefix(auth, "Bearer ") {
		return strings.TrimPrefix(auth, "Bearer ")
	}
	return ""
}

func validateJWT(tokenStr string, secret string) (*service.CustomClaims, error) {
	claims := &service.CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}
	return claims, nil
}

func allowResponse(sub string, role string) events.APIGatewayV2CustomAuthorizerSimpleResponse {
	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: true,
		Context: map[string]interface{}{
			"sub":  sub,
			"role": role,
		},
	}
}

func denyResponse() events.APIGatewayV2CustomAuthorizerSimpleResponse {
	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: false,
	}
}

func IsValidator() bool {
	return os.Getenv("LAMBDA_MODE") == "validator"
}
