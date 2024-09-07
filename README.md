# Exemplo API

Este é um exemplo de API escrita em Go utilizando o framework [Go-Chi](https://github.com/go-chi/chi) e GORM como ORM. A API também inclui autenticação JWT e documentação gerada com Swagger.

## Descrição

A API fornece funcionalidades para gerenciar usuários e produtos, com suporte a autenticação JWT e proteção de rotas. Além disso, a documentação interativa da API pode ser acessada via Swagger.

## Tecnologias Utilizadas

- Go (Golang)
- Go-Chi (Router HTTP)
- JWT Auth
- GORM (ORM)
- SQLite (Banco de Dados)
- Swagger (Documentação)

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio

2. Instale as dependências && inicializacao do projeto

```
go mod tidy
go run main.go

```

```
Endpoints

Produtos
GET /products: Retorna a lista de produtos (rota protegida)
POST /products: Cria um novo produto (rota protegida)
GET /products/{id}: Retorna um produto específico (rota protegida)
PUT /products: Atualiza um produto existente (rota protegida)
DELETE /products/{id}: Deleta um produto (rota protegida)
Usuários
GET /users: Retorna a lista de usuários
POST /users: Cria um novo usuário
GET /users/{id}: Retorna um usuário específico
PUT /users: Atualiza um usuário existente
DELETE /users/{id}: Deleta um usuário
Autenticação
POST /login: Retorna um token JWT ao realizar o login.
Documentação Swagger
GET /docs: Acesse a documentação interativa da API em http://localhost:8080/docs.
Configuração

No arquivo main.go, o banco de dados SQLite é utilizado como padrão, criando um arquivo test.db para armazenar os dados. Para trocar o banco de dados ou modificar outras configurações, edite as configurações no arquivo de configuração.

Licença

Este projeto está licenciado sob a Licença Apache 2.0.

Contato

Caso tenha dúvidas ou sugestões, entre em contato:

Email: support@swagger.io

```