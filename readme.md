# Todo Task Backend

Backend de gerenciamento de tarefas com Golang, GraphQL e autenticação via JWT.

## 🚀 Funcionalidades

- Gerenciamento de tarefas (CRUD)
- Autenticação com JWT
- API GraphQL

## 🛠️ Tecnologias Utilizadas

- Golang
- GraphQL
- JWT
- PostgreSQL

## ⚙️ Configuração Rápida

1. Clone o repositório:

```bash
    git clone https://github.com/seu-usuario/todo-task-backend.git
    cd todo-task-backend
```

2. Crie o arquivo `.env`:

```bash
DB_USERNAME=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_DATABASE=
SECRET_JWT_TOKEN=
```

3. Inicie o banco com Docker Compose:

```bash
docker compose up -d
```

4. Execute o servidor:

```bash
go run main.go # ou air
```

## 🧪 Exemplos de Query

acesse `http://localhost:8080/query/playground`

---
**Desenvolvido com ❤️ e Golang!**