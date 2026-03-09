## 🛠️ Build

```bash
# Compilar para AWS Lambda (Linux)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap

# Criar pacote de deploy
chmod +x bootstrap
zip function.zip bootstrap
```

## 🐳 Desenvolvimento Local (LocalStack)

### Pré-requisitos

- Docker e Docker Compose
- AWS CLI
- Go 1.25+

### Subir ambiente

```bash
# Iniciar PostgreSQL e LocalStack
docker-compose up -d

# Verificar se estão na mesma rede
docker network inspect tech-challange_garage-network
```

### Deploy no LocalStack

```bash
# Build
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap
zip function.zip bootstrap

# Deletar função anterior (se existir)
aws --endpoint-url=http://localhost:4566 lambda delete-function \
  --function-name IssuerLambda \
  --region sa-east-1

# Criar função
aws --endpoint-url=http://localhost:4566 lambda create-function \
  --function-name IssuerLambda \
  --runtime provided.al2023 \
  --role arn:aws:iam::000000000000:role/lambda-role \
  --handler bootstrap \
  --zip-file fileb://function.zip \
  --environment '{"Variables":{"ENVIRONMENT":"dev","DB_HOST_DEV":"garage-postgres","DB_PORT_DEV":"5432","DB_USER_DEV":"secret","DB_PASSWORD_DEV":"secret","DB_NAME_DEV":"garage", "JWT_SECRET": "123"}}' \
  --region sa-east-1
```

### Testar

```bash
# Invocar Lambda
aws --endpoint-url=http://localhost:4566 lambda invoke \
  --function-name IssuerLambda \
  --cli-binary-format raw-in-base64-out \
  --payload '{"document":"033.326.420-73"}' \
  --region sa-east-1 \
  response.json

# Ver resposta
cat response.json
```

## 📥 Request/Response

### Request
```json
{
  "document": "12345678900"
}
```

### Response (Sucesso)
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Response (Erro)
```json
{
  "errorMessage": "Usuário não encontrado",
  "errorType": "Error"
}
```
