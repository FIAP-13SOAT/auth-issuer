## Como testar usando Localstack

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap
chmod +x bootstrap
zip function.zip bootstrap
```

Publicar no Localstack
```bash
aws --endpoint-url=http://localhost:4566 lambda create-function \
  --function-name IssuerLambda2 \
  --runtime provided.al2023 \
  --role arn:aws:iam::000000000000:role/lambda-role \
  --handler bootstrap \
  --zip-file fileb://function.zip \
  --environment '{"Variables":{"ENVIRONMENT":"dev","DB_HOST_DEV":"host.docker.internal","DB_PORT_DEV":"5432","DB_USER_DEV":"secret","DB_PASSWORD_DEV":"secret","DB_NAME_DEV":"garage"}}' \
  --region sa-east-1
```
