# 📝 Go Todo App

A simple **REST API** built with **Go + Gin + GORM (MySQL)**. **Clean / Hexagonal Architecture(golang-clean-context)**
<br />
Project follows a modular structure (`internal/todo`, `internal/users`) and uses **Clean Architecture** style (index → service → dao).

---

## 📂 Project Structure
```
.
├── db/                # Database connection (mysql.go)
├── internal/
│   ├── todo/          # Todo module (models, services)
│   │   ├── models/
│   │   ├── daos/      # Data access object (later)
│   │   ├── services/
│   │   └── index.go   # router
│   └── users/         # Users module
│       └── index.go   # router
├── go.mod
├── go.sum
├── main.go            # Application entrypoint
```

---

## 🚀 Getting Started

### 1. Prerequisites
- Go >= 1.20
- MySQL >= 8.0
- (Optional) [Air](https://github.com/cosmtrek/air) for hot-reload

### 2. Clone repository
```bash
git clone https://github.com/delpikye-v/todo-app.git
cd todo-app
```

### 3. Setup environment
Create `.env` file:
```env
DB_USER=root
DB_PASS=secret
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=todo_db
```

### 4. Run database (Docker optional)
```bash
docker run --name mysql \
  -e MYSQL_ROOT_PASSWORD=secret \
  -e MYSQL_DATABASE=todo_db \
  -p 3306:3306 \
  -d mysql:8
```

### 5. Install dependencies
```bash
go mod tidy
```

### 6. Run the app
```bash
go run main.go
```

Or with hot-reload:
```bash
air
```

Server will start at:
```
http://localhost:8080
```

---

## 📡 API Endpoints

### ✅ Todos
- `GET /todos` → list all todos
- `GET /todos/:id` → get a todo by id
- `POST /todos` → create a new todo
- `PUT /todos/:id` → update todo
- `DELETE /todos/:id` → delete todo

### 👤 Users
- `GET /users` → list all users
  *(Extend CRUD similar to todos)*

---

## 📌 CRUD Examples (with `curl`)

### Create Todo
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"task": "Learn Go", "done": false}'
```

### Get All Todos
```bash
curl http://localhost:8080/todos
```

🔹 Example Response:
```json
[
  {
    "id": 1,
    "task": "Learn Go",
    "done": false
  },
  {
    "id": 2,
    "task": "Write Gin API",
    "done": true
  }
]
```

### Get Todo by ID
```bash
curl http://localhost:8080/todos/1
```

🔹 Example Response:
```json
{
  "id": 1,
  "task": "Learn Go",
  "done": false
}
```

### Update Todo
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"task": "Learn Go + GORM", "done": true}'
```

### Delete Todo
```bash
curl -X DELETE http://localhost:8080/todos/1
```

---

## 🛠 Tech Stack
- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/) (MySQL ORM)  => This source code is for teaching purposes and demonstrates clean code practices.
- [Air](https://github.com/cosmtrek/air) (Hot reload)
- MySQL

---

## 📌 Development Notes
- Follows **Clean Architecture**:
  - `services` → business logic
  - `dao` → DB access
- Easily extendable: add new module under `internal/`.

---

## ✅ Next Steps
- Add **JWT Authentication** (users module).
- Dockerize app + db for easy deployment.
- etc...

---

## License

This project is licensed under the MIT License.

---

Made with ❤️ by **DelpiKye**
