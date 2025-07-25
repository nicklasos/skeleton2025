
# 🛠️ Golang Web API Project Skeleton (Echo + sqlc + Swagger + Modular Structure)

## 🎯 Project Goal:
Generate a **starter Go web API project skeleton** that follows modern Go best practices with the following requirements:

---

## ✅ Project Tech Stack:
- **Go** (Latest stable version)
- **MySQL** (as the database)
- **Echo** (as the web framework)
- **sqlc** (for typed DB queries)
- **Swaggo/swag** (for Swagger / OpenAPI documentation)
- **Makefile** (for common dev tasks)

---

## ✅ Folder Structure:

```
myapp/
├── cmd/api/main.go
├── internal/
│   ├── users/
│   │   ├── queries.sql
│   │   ├── queries_gen.go
│   │   ├── handler.go
│   │   └── repository.go
│   └── db/
│       └── db.go
├── migrations/
│   └── (empty for now)
├── docs/
│   └── (generated by swag later)
├── sqlc.yaml
├── go.mod
├── go.sum
└── Makefile
```

---

## ✅ Features Needed:

### ✅ 1. **Basic Main Entry Point:**

- **Location:** `cmd/api/main.go`
- Echo server setup
- Routes setup for `/users`

---

### ✅ 2. **Users Module:**

Inside `/internal/users/`

- **SQLC Query Example:**  
Include at least:
  - Get user by ID
  - Create user
  - List users

```sql
-- name: GetUserByID :one
SELECT id, name, email FROM users WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (name, email) VALUES (?, ?);

-- name: ListUsers :many
SELECT id, name, email FROM users ORDER BY id;
```

- **sqlc-generated Go code** (`queries_gen.go`)
- **User handler** (`handler.go`) with at least:
  - GET `/users/:id`
  - POST `/users`
  - GET `/users`

- **Repository pattern** (`repository.go`) that calls sqlc-generated code.

---

### ✅ 3. **Database Connection Setup:**

Inside `/internal/db/db.go`  
Setup basic MySQL connection using `database/sql` with `sql.Open()`.

Use connection string from environment variables or hardcoded placeholder.

---

### ✅ 4. **Swagger Setup:**

- Annotate all handlers with Swaggo comments
- Generate placeholder Swagger docs under `/docs/`
- Setup `/swagger/*` route in Echo using swaggo's Echo middleware.

---

### ✅ 5. **Makefile with Common Commands:**

At minimum:

```makefile
run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

sqlc:
	sqlc generate

swagger:
	swag init --parseDependency --parseInternal

test:
	go test ./...

fmt:
	go fmt ./...
```

---

### ✅ 6. **go.mod Setup:**

- Use module path like: `module myapp`
- Import Echo, Swaggo, go-sql-driver/mysql, and sqlc dependencies.

---

### ✅ 7. **Simple README.md (Optional, if LLM has time):**

Explain how to:

- Run the server
- Generate sqlc
- Generate Swagger
- Build the project

---

## ✅ Additional Notes:

- For now, skip Docker.
- No unit tests needed for first version.
- No Git setup required.

---

## ✅ Stretch Goal (Optional):

If there’s time/space, optionally generate a second module:  
Example: `/internal/cities/` with 1 sample GET endpoint (`/cities`).

---

## ✅ Important:  

✅ Keep everything **simple and beginner-friendly**.  
✅ Focus on **clean structure and working compilation**.  
✅ Use idiomatic Go style.  
✅ For sqlc-generated code: you can generate it or leave it empty with a placeholder comment like `// generated by sqlc`.

---

✅ End of instruction.
