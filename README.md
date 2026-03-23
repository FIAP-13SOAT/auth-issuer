# Tech Challenge - Auth Issuer

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
├── database/
│   └── db.go           # Configuração do banco de dados
├── handler/
│   └── auth.go         # Lógica de autenticação
├── repository/
│   └── user.go         # Queries do banco
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
```


## 📖 Documentação da API

A documentação completa da API está disponível como coleção Postman em [`docs/postman_collection.json`](docs/postman_collection.json). Importe o arquivo no [Postman](https://www.postman.com/) para testar os endpoints interativamente.

### `POST /login`

Autentica um cliente pelo CPF ou CNPJ e retorna um token JWT.

**Request Body:**

```json
{
  "document": "12345678900"
}
```

| Campo      | Tipo   | Obrigatório | Descrição                          |
|------------|--------|-------------|------------------------------------|
| `document` | string | Sim         | CPF (11 dígitos) ou CNPJ (14 dígitos) do cliente |

**Respostas:**

| Código | Descrição                  | Body                                                        |
|--------|----------------------------|-------------------------------------------------------------|
| `200`  | Autenticação bem-sucedida  | `{ "token": "<JWT>" }`                                      |
| `400`  | Documento inválido ou vazio | `{ "message": "O campo 'document' é obrigatório" }` ou `{ "message": "CPF/CNPJ inválido" }` |
| `404`  | Usuário não encontrado     | `{ "message": "Usuário não encontrado" }`                   |

**Formato do Token JWT:**

O token retornado é um JWT assinado contendo o `customer_id` do cliente autenticado. Utilize-o no header `Authorization: Bearer <token>` para acessar os endpoints protegidos da aplicação principal.

**Exemplo com cURL:**

```bash
curl -X POST https://<api-gateway-url>/login \
  -H "Content-Type: application/json" \
  -d '{"document": "12345678900"}'
```

**Resposta de sucesso:**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

## Como Testar localmente

Instruções para teste localmente [aqui](test_instructions.md)
