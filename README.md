# Gin-TODO

Gin-Todo is a simple note-taking app developed using the Gin framework in Go, featuring a layered architecture. It helps you manage your daily tasks efficiently and effectively.

## How to use:

### Clone it to your device:
```bash
git clone https://github.com/capigiba/Note-Go.git
```

### Frontend source code:
```bash
https://github.com/capigiba/Gin-TODO-webapp.git
```

### Setup package:
```bash
go mod tidy
```

### MySQL setup:
Run MySQL on your command line and paste these commands:

```sql
CREATE USER 'notes'@'localhost' IDENTIFIED BY 'Str0ngP@ssw0rd!';

CREATE DATABASE notes;

GRANT ALL PRIVILEGES ON notes.* TO 'notes'@'localhost';

CREATE TABLE todos (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(255) NOT NULL,
    detail TEXT NOT NULL,
    complete BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

### Run server on port 8080:
```bash
go run cmd/app/main.go
```

### Register account:
```bash
curl -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword"
}'
```

### Login to get token:
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "email": "john.doe@example.com",
    "password": "securepassword"
}'
```

### Create your first todo:
```bash
curl -X POST http://localhost:8080/todos \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your_jwt_token_here" \
-d '{
    "title": "Buy groceries",
    "detail": "Milk, Bread, Eggs, Cheese"
}'
```
