# Changelog

Todas as mudanças notáveis neste projeto serão documentadas neste arquivo.

O formato é baseado em [Keep a Changelog](https://keepachangelog.com/pt-BR/1.0.0/),
e este projeto adere ao [Semantic Versioning](https://semver.org/lang/pt-BR/).

## [1.0.0] - 2024-01-XX

### Adicionado
- Autenticação via CPF/CNPJ
- Validação de documentos brasileiros
- Geração de tokens JWT
- Integração com PostgreSQL
- Testes unitários
- CI/CD com GitHub Actions
- Linting com golangci-lint
- Makefile para automação
- Documentação completa
- Suporte a LocalStack

### Segurança
- JWT secret via variável de ambiente
- Validação de entrada
- Pool de conexões limitado
