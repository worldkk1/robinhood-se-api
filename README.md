# API for manage card and comment
API for managing cards (tasks) and comments.

---

## Features

- Create task
- Get task list
- Get task detail
- Edit task detail
- Archive task
- Get card history
- Add comment for any task
- Edit comment
- Delete comment
- Authenticate and authorize user

---

## Getting Started

### 1. Development (Local)

#### Prerequisites

- Go 1.25+
- PostgreSQL
- Redis
- goose (https://github.com/pressly/goose)

#### Steps

```sh
# 1. Copy environment file
cp .env.example .env

# 2. Start PostgreSQL and Redis (manually or via Docker)
docker-compose up task-db task-cache

# 3. Run database migrations
goose up -dir internal/database/migrations

# 4. Start the API server
go run ./cmd/main.go
```

API will run at `http://localhost:3000/v1/`

---

### 2. Docker Compose

```sh
docker-compose up
```

API will run at `http://localhost:3000/v1/`

---

## API Endpoints

### Auth

#### Register

```http
POST /v1/auth/register
Content-Type: application/json

{
  "name": "Alice",
  "email": "alice@example.com",
  "password": "password123",
  "roleId": "ae4c58a6-101a-4b0b-a63e-e187d1920c7e"
}

Role ID
admin: ae4c58a6-101a-4b0b-a63e-e187d1920c7e
user: 756b0c5c-e9ff-4e91-aa21-49c0dfdc653c
```

#### Login

```http
POST /v1/auth/login
Content-Type: application/json

{
  "email": "alice@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "accessToken": "...",
  "refreshToken": "..."
}
```

---

### Tasks

#### Create Task (Admin only)

```http
POST /v1/tasks
Authorization: Bearer <accessToken>
Content-Type: application/json

{
  "title": "New Task",
  "description": "Task details"
}
```

#### Get Task List

```http
GET /v1/tasks?offset=0&limit=10
Authorization: Bearer <accessToken>
```

Response:

```json
{
  "offset": 0,
  "limit": 10,
  "total": 1,
  "data": [
    {
      "id": "...",
      "title": "New Task",
      "description": "Task details",
      "status": "to_do",
      "userId": "...",
      "archivedAt": null,
      "createdAt": "...",
      "updatedAt": "...",
      "user": {
        "id": "...",
        "name": "Alice",
        "email": "alice@example.com"
      }
    }
  ]
}
```

#### Get Task Detail

```http
GET /v1/tasks/{id}
Authorization: Bearer <accessToken>
```

#### Edit Task

```http
PATCH /v1/tasks/{id}
Authorization: Bearer <accessToken>
Content-Type: application/json

{
  "title": "Updated Task",
  "description": "Updated details",
  "status": "in_progress"
}
```

#### Archive Task

```http
PATCH /v1/tasks/{id}/archive
Authorization: Bearer <accessToken>
```

---

### Comments

#### Add Comment

```http
POST /v1/tasks/{id}/comments
Authorization: Bearer <accessToken>
Content-Type: application/json

{
  "content": "This is a comment"
}
```

#### Get Comments

```http
GET /v1/tasks/{id}/comments
Authorization: Bearer <accessToken>
```

#### Edit Comment

```http
PATCH /v1/tasks/{id}/comments/{commentId}
Authorization: Bearer <accessToken>
Content-Type: application/json

{
  "content": "Updated comment"
}
```

#### Delete Comment

```http
DELETE /v1/tasks/{id}/comments/{commentId}
Authorization: Bearer <accessToken>
```

---

## Rate Limiting

All API endpoints are protected by rate limiting to prevent abuse.

- **Limit:** 1 request per second
- **Burst:** Up to 10 requests allowed in a short burst
- **Response:** If you exceed the limit, you will receive HTTP `429 Too Many Requests`


Rate limiting is enforced per client using Redis.

---

## License

MIT
