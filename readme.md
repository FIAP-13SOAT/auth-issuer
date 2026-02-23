## Como testar usando Localstack

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap
chmod +x bootstrap
zip function.zip bootstrap
```

Publicar no Localstack
```bash
aws --endpoint-url=http://localhost:4566 lambda delete-function \
  --function-name IssuerLambda \
  --region sa-east-1

aws --endpoint-url=http://localhost:4566 lambda create-function \
  --function-name IssuerLambda \
  --runtime provided.al2023 \
  --role arn:aws:iam::000000000000:role/lambda-role \
  --handler bootstrap \
  --zip-file fileb://function.zip \
  --environment '{"Variables":{"ENVIRONMENT":"dev","DB_HOST_DEV":"garage-postgres","DB_PORT_DEV":"5432","DB_USER_DEV":"secret","DB_PASSWORD_DEV":"secret","DB_NAME_DEV":"garage"}}' \
  --region sa-east-1
```

Para testar:
```bash
aws --endpoint-url=http://localhost:4566 lambda invoke   --function-name IssuerLambda   --cli-binary-format raw-in-base64-out   --payload '{"document":"12345678900"}'  --region sa-east-1 response.json
```
