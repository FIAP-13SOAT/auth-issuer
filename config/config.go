package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Environment string
	Database    DatabaseConfig
	JWT         JWTConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret     string
	Issuer     string
	Expiration int
}

func Load() (*Config, error) {
	env := getEnv("ENVIRONMENT", "dev")
	envUpper := strings.ToUpper(env)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET não configurado")
	}

	cfg := &Config{
		Environment: env,
		JWT: JWTConfig{
			Secret:     jwtSecret,
			Issuer:     "tech-challange-auth-issuer",
			Expiration: 15,
		},
	}

	// DB config é opcional (validator não precisa)
	dbHost := os.Getenv(fmt.Sprintf("DB_HOST_%s", envUpper))
	dbPort := os.Getenv(fmt.Sprintf("DB_PORT_%s", envUpper))
	dbUser := os.Getenv(fmt.Sprintf("DB_USER_%s", envUpper))
	dbPassword := os.Getenv(fmt.Sprintf("DB_PASSWORD_%s", envUpper))
	dbName := os.Getenv(fmt.Sprintf("DB_NAME_%s", envUpper))

	if dbHost != "" && dbPort != "" && dbUser != "" && dbPassword != "" && dbName != "" {
		cfg.Database = DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
		}
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
