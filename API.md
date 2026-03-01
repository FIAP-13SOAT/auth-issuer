# API Documentation

## Endpoint: Auth Issuer

### Request

**Method:** Lambda Invoke  
**Function:** IssuerLambda

**Payload:**
```json
{
  "document": "033.326.420-73"
}
```

**Campos:**
- `document` (string, obrigatório): CPF (11 dígitos) ou CNPJ (14 dígitos), com ou sem formatação

### Response

#### Sucesso (200)
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZWNoLWNoYWxsYW5nZS1hdXRoLWlzc3VlciIsInN1YiI6IjEyMzQ1Njc4LTkwYWItY2RlZi0xMjM0LTU2Nzg5MGFiY2RlZiIsImV4cCI6MTcwMDAwMDAwMCwiaWF0IjoxNzAwMDAwMDAwfQ.signature"
}
```

**Token Claims:**
- `iss`: Issuer (tech-challange-auth-issuer)
- `sub`: Customer ID (UUID)
- `iat`: Issued At (timestamp)
- `exp`: Expiration (15 minutos)

#### Erros

**Documento vazio:**
```json
{
  "errorMessage": "O campo 'document' é obrigatório",
  "errorType": "Error"
}
```

**CPF/CNPJ inválido:**
```json
{
  "errorMessage": "CPF/CNPJ inválido",
  "errorType": "Error"
}
```

**Usuário não encontrado:**
```json
{
  "errorMessage": "Usuário não encontrado",
  "errorType": "Error"
}
```

### Exemplos

#### CPF com formatação
```bash
aws lambda invoke \
  --function-name IssuerLambda \
  --payload '{"document":"033.326.420-73"}' \
  response.json
```

#### CPF sem formatação
```bash
aws lambda invoke \
  --function-name IssuerLambda \
  --payload '{"document":"03332642073"}' \
  response.json
```

#### CNPJ
```bash
aws lambda invoke \
  --function-name IssuerLambda \
  --payload '{"document":"11.222.333/0001-81"}' \
  response.json
```

### Validações

1. **Campo obrigatório**: document não pode ser vazio ou apenas espaços
2. **Formato válido**: CPF/CNPJ deve ter dígitos verificadores corretos
3. **Usuário existente**: CPF/CNPJ deve estar cadastrado no banco

### Segurança

- Token JWT assinado com HS256
- Secret configurável via variável de ambiente `JWT_SECRET`
- Expiração de 15 minutos
- Conexão com banco via SSL em produção
