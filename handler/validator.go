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
	path := request.RawPath
	method := request.RequestContext.HTTP.Method

	log.Printf("Token válido: sub=%s, role=%s, path=%s, method=%s", sub, role, path, method)

	// CUSTOMER só pode acessar rotas específicas
	if role == "CUSTOMER" && !isCustomerAllowed(path, method) {
		log.Printf("Acesso negado: CUSTOMER não pode acessar %s %s", method, path)
		return denyResponse(), nil
	}

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

// isCustomerAllowed define quais rotas o CUSTOMER pode acessar.
// Rotas permitidas:
//   - GET /vehicles, POST /vehicles, GET/PUT/DELETE /vehicles/{id}
//   - GET /service-types
//   - GET /public/*
func isCustomerAllowed(path string, method string) bool {
	// Veículos — CRUD completo
	if strings.HasPrefix(path, "/vehicles") {
		return true
	}

	// Tipos de serviço — apenas leitura
	if path == "/service-types" && method == "GET" {
		return true
	}

	// Rotas públicas
	if strings.HasPrefix(path, "/public/") {
		return true
	}

	return false
}
