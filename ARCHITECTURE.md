# Arquitetura

## Visão Geral

Sistema serverless de autenticação baseado em AWS Lambda que valida CPF/CNPJ e emite tokens JWT.

## Componentes

### 1. Lambda Function (Go)
- **Runtime**: Custom Runtime (provided.al2023)
- **Handler**: AuthHandler
- **Timeout**: 30s
- **Memory**: 128MB

### 2. PostgreSQL (RDS)
- **Tabela**: customer
- **Campos**: id (UUID), cpf_cnpj (VARCHAR)
- **Índice**: cpf_cnpj (único)

### 3. JWT Token
- **Algoritmo**: HS256
- **Claims**: sub (customer_id), iss, iat, exp
- **Expiração**: 15 minutos

## Fluxo de Autenticação

```
┌─────────┐         ┌─────────┐         ┌──────────┐         ┌──────────┐
│ Cliente │         │  Lambda │         │Validator │         │PostgreSQL│
└────┬────┘         └────┬────┘         └────┬─────┘         └────┬─────┘
     │                   │                   │                     │
     │ POST {document}   │                   │                     │
     ├──────────────────>│                   │                     │
     │                   │                   │                     │
     │                   │ Validar formato   │                     │
     │                   ├──────────────────>│                     │
     │                   │                   │                     │
     │                   │ CPF/CNPJ válido   │                     │
     │                   │<──────────────────┤                     │
     │                   │                   │                     │
     │                   │ SELECT id WHERE cpf_cnpj = ?            │
     │                   ├─────────────────────────────────────────>│
     │                   │                                          │
     │                   │              customer_id                 │
     │                   │<─────────────────────────────────────────┤
     │                   │                   │                     │
     │                   │ Gerar JWT         │                     │
     │                   ├───────────┐       │                     │
     │                   │           │       │                     │
     │                   │<──────────┘       │                     │
     │                   │                   │                     │
     │  {token: "..."}   │                   │                     │
     │<──────────────────┤                   │                     │
     │                   │                   │                     │
```

## Camadas

### Handler Layer
- Recebe requisição
- Valida entrada
- Orquestra fluxo
- Retorna resposta

### Validator Layer
- Valida formato CPF/CNPJ
- Calcula dígitos verificadores
- Remove formatação

### Repository Layer
- Consulta banco de dados
- Gerencia conexões
- Executa queries

### Service Layer
- Gera tokens JWT
- Configura claims
- Assina tokens

## Segurança

### Validações
1. Campo obrigatório
2. Formato CPF/CNPJ válido
3. Usuário existe no banco

### Proteções
- JWT secret via variável de ambiente
- Conexão SSL com banco (produção)
- Pool de conexões limitado
- Timeout de requisição

## Escalabilidade

- Lambda escala automaticamente
- Pool de conexões configurável
- Stateless (sem sessão)
- Cache de conexões

## Monitoramento

### Logs
- CloudWatch Logs
- Structured logging
- Níveis: INFO, ERROR

### Métricas
- Invocações
- Duração
- Erros
- Throttles

## Ambientes

### Desenvolvimento (LocalStack)
- PostgreSQL local
- Lambda local
- Sem SSL

### Produção (AWS)
- RDS PostgreSQL
- Lambda em VPC
- SSL obrigatório
