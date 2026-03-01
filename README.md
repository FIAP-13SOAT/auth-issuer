# Tech Challenge - Auth Issuer

[![CI/CD](https://github.com/seu-usuario/tech-challange-auth-issuer/workflows/CI%2FCD/badge.svg)](https://github.com/seu-usuario/tech-challange-auth-issuer/actions)
[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Lambda function para autenticação de clientes via CPF/CNPJ, retornando um token JWT.

## 📋 Descrição

Serviço serverless que autentica clientes consultando o banco de dados PostgreSQL e retorna um token JWT para acesso aos demais serviços.

## 🏗️ Arquitetura

```
┌─────────────┐      ┌──────────────┐      ┌────────────┐
│   Cliente   │─────▶│    Lambda    │─────▶│ PostgreSQL │
│             │◀─────│ Auth Issuer  │◀─────│    (RDS)   │
└─────────────┘      └──────────────┘      └────────────┘
   CPF/CNPJ              JWT Token           Validação
```

## 🚀 Tecnologias

- **Go 1.25** - Linguagem de programação
- **AWS Lambda** - Execução serverless
- **PostgreSQL** - Banco de dados
- **LocalStack** - Ambiente de desenvolvimento local

## 📁 Estrutura do Projeto

```
tech-challange-auth-issuer/
├── main.go              # Entry point da Lambda
├── config/
│   └── config.go       # Configuração centralizada
├── database/
│   └── db.go           # Conexão com PostgreSQL
├── handler/
│   ├── auth.go         # Handler de autenticação
│   └── auth_test.go    # Testes do handler
├── repository/
│   └── customer.go     # Queries de customer
├── service/
│   └── jwt.go          # Geração de tokens JWT
├── validator/
│   ├── document.go     # Validação CPF/CNPJ
│   └── document_test.go
├── models/
│   └── types.go        # Structs de entrada/saída
└── README.md
```

## 🔧 Configuração

### Variáveis de Ambiente

**Desenvolvimento:**
```bash
ENVIRONMENT=dev
DB_HOST_DEV=garage-postgres
DB_PORT_DEV=5432
DB_USER_DEV=secret
DB_PASSWORD_DEV=secret
DB_NAME_DEV=garage
```

**Produção:**
```bash
ENVIRONMENT=prod
DB_HOST_PROD=<rds-endpoint>
DB_PORT_PROD=5432
DB_USER_PROD=<user>
DB_PASSWORD_PROD=<password>
DB_NAME_PROD=<database>
JWT_SECRET=<secret-key>
```

## 🚀 Quick Start

```bash
# Subir ambiente
docker-compose up -d

# Build e deploy local
make deploy-local

# Testar
make invoke-local
```

## 📚 Documentação

- [Instruções de Teste](test_instructions.md)
- [Documentação da API](API.md)
- [Guia de Contribuição](CONTRIBUTING.md)

## ✅ Funcionalidades

- ✅ Validação de CPF/CNPJ
- ✅ Autenticação via banco de dados
- ✅ Geração de JWT com expiração
- ✅ Testes unitários
- ✅ CI/CD automatizado
- ✅ Linting e análise estática

## 🧪 Testes

```bash
make test
```

## 📝 Licença

MIT
