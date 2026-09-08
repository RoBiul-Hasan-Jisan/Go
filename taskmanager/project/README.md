# Task Manager

A simple full-stack task manager: a Go (Gin + GORM + PostgreSQL) REST API
backend, and a Next.js (App Router + TypeScript + Tailwind CSS) frontend.

```
Frontend: http://localhost:3000
Backend:  http://localhost:8080
```

## Project structure

```
project/
├── backend/            Go REST API (Gin, GORM, PostgreSQL, JWT)
├── frontend/            Next.js app (TypeScript, Tailwind CSS)
├── README.md
└── .gitignore
```

## Prerequisites

- Go 1.22+
- Node.js 18.18+ and npm
- PostgreSQL 13+ (running locally or accessible remotely)

## 1. Database setup

Create a database for the app. From the `psql` shell or any PostgreSQL client:

```sql
CREATE DATABASE taskdb;
```

That's it — the backend automatically creates the `users` and `tasks` tables
(and keeps them in sync) on startup via GORM's `AutoMigrate`. No manual
migration scripts are needed.

## 2. Backend setup

```bash
cd backend
cp .env.example .env
```

Edit `backend/.env` with your own values:

```env
PORT=8080
DATABASE_URL=postgresql://user:password@localhost:5432/taskdb
JWT_SECRET=your_secret
JWT_EXPIRY_HOURS=72
FRONTEND_URL=http://localhost:3000
```

| Variable           | Description                                              |
|--------------------|------------------------------------------------------------|
| `PORT`             | Port the API server listens on                            |
| `DATABASE_URL`     | PostgreSQL connection string                               |
| `JWT_SECRET`       | Secret used to sign/verify JWTs — use a long random value  |
| `JWT_EXPIRY_HOURS` | How long issued tokens stay valid (default `72`)           |
| `FRONTEND_URL`     | Origin allowed by CORS (your Next.js dev/production URL)   |

Install dependencies and run the server:

```bash
go mod tidy
go run ./cmd/server
```

The API will start on `http://localhost:8080`. Check it's alive:

```bash
curl http://localhost:8080/api/health
```

## 3. Frontend setup

In a separate terminal:

```bash
cd frontend
cp .env.local.example .env.local
```

`frontend/.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

Install dependencies and run the dev server:

```bash
npm install
npm run dev
```

The app will start on `http://localhost:3000`.

## Running both together

```bash
# Terminal 1
cd backend
go run ./cmd/server

# Terminal 2
cd frontend
npm run dev
```

Then open `http://localhost:3000` in your browser, register an account, and
start creating tasks.

## Database schema

**users**

| Column          | Type      | Notes                    |
|-----------------|-----------|---------------------------|
| id              | uint      | primary key                |
| name            | string    |                            |
| email           | string    | unique                     |
| password_hash   | string    | bcrypt hash, never exposed |
| created_at      | timestamp |                            |
| updated_at      | timestamp |                            |

**tasks**

| Column      | Type      | Notes                              |
|-------------|-----------|--------------------------------------|
| id          | uint      | primary key                          |
| user_id     | uint      | foreign key → users.id (cascade delete) |
| title       | string    |                                       |
| description | string    |                                       |
| completed   | bool      | default `false`                      |
| created_at  | timestamp |                                       |
| updated_at  | timestamp |                                       |

A user can have many tasks; deleting a user deletes their tasks.

## API documentation

All responses use a consistent envelope:

```json
{
  "success": true,
  "message": "Task created successfully",
  "data": {}
}
```

On errors, `success` is `false`, `data` is omitted, and `message` describes
the problem. Protected routes require an `Authorization: Bearer <token>`
header.

### `GET /api/health`

Health check. No auth required.

**200 OK**
```json
{ "success": true, "message": "Service is healthy", "data": { "status": "ok" } }
```

---

### `POST /api/auth/register`

Create a new account.

**Body**
```json
{ "name": "Jane Doe", "email": "jane@example.com", "password": "secret123" }
```

**201 Created**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "token": "eyJ...",
    "user": { "id": 1, "name": "Jane Doe", "email": "jane@example.com", "created_at": "...", "updated_at": "..." }
  }
}
```

Errors: `400` invalid input, `409` email already registered.

---

### `POST /api/auth/login`

**Body**
```json
{ "email": "jane@example.com", "password": "secret123" }
```

**200 OK** — same shape as register.

Errors: `400` invalid input, `401` invalid credentials.

---

### `GET /api/tasks`  🔒

Returns the authenticated user's tasks, newest first.

**200 OK**
```json
{ "success": true, "message": "Tasks retrieved successfully", "data": [ { "id": 1, "title": "...", "...": "..." } ] }
```

---

### `POST /api/tasks`  🔒

**Body**
```json
{ "title": "Buy groceries", "description": "Milk, eggs, bread" }
```

**201 Created** — the created task.

---

### `GET /api/tasks/:id`  🔒

**200 OK** — the task, if it belongs to the authenticated user.
**404** if not found or not owned by the user.

---

### `PUT /api/tasks/:id`  🔒

Partial update — send only the fields you want to change.

**Body**
```json
{ "title": "Buy groceries", "completed": true }
```

**200 OK** — the updated task.

---

### `DELETE /api/tasks/:id`  🔒

**200 OK**
```json
{ "success": true, "message": "Task deleted successfully", "data": null }
```

🔒 = requires `Authorization: Bearer <token>` header.

### Logout

Logout is handled entirely on the frontend: JWTs are stateless, so "logging
out" simply means the client deletes the stored token (`localStorage`) and
redirects to the login page. There's no server-side session to invalidate.

## How authentication works

1. On register/login, the backend returns a signed JWT alongside the user.
2. The frontend stores the token in `localStorage` (see `frontend/lib/auth.ts`).
3. Every request to a protected endpoint attaches `Authorization: Bearer <token>`
   via an axios interceptor (see `frontend/lib/api.ts`).
4. The backend's `AuthRequired` middleware validates the token and injects
   the authenticated user's ID into the request context, so every task
   endpoint automatically scopes reads/writes to that user only.

## Tech stack

**Backend:** Go, Gin, GORM, PostgreSQL, JWT (`golang-jwt/jwt`), bcrypt, CORS.

**Frontend:** Next.js 15 (App Router), TypeScript, Tailwind CSS, Axios.
