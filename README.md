# Kemitbelajar Blog API

This is the backend API for the Kemitbelajar Blog, built with [Go](https://golang.org/) and the [Fiber](https://gofiber.io/) web framework.

## Prerequisites

- Go 1.25 or higher
- PostgreSQL
- Docker (optional, for containerized deployment)

## Getting Started

### 1. Environment Variables

Create a `.env` file in the root of the `backend` directory. You can copy the provided example:

```bash
cp .env.example .env
```

Make sure to update the environment variables according to your local setup:
- `DB_HOST`: Your database host.
- `DB_PORT`: Your database port.
- `DB_USER`: Your database user.
- `DB_PASSWORD`: Your database password.
- `DB_NAME`: Your database name.
- `SERVER_PORT`: Port for the Fiber server (default: 3001).
- `ALLOWED_ORIGINS`: Comma-separated URLs of allowed origins for CORS (default: `http://localhost:3000`).

### 2. Database Migrations

Migrations will automatically run when the application starts up. Ensure your PostgreSQL instance is running and the database specified in `DB_NAME` exists before starting the server.

### 3. Running Locally

To run the application locally without Docker, execute:

```bash
go run ./cmd/server/main.go
```

The server should start on `http://localhost:3001` (or your configured port) and you will see a message indicating the server is healthy.

## Docker

You can also run the application using Docker.

### Building the Image

```bash
docker build -t kemitbelajar-backend .
```

### Running the Container

If your PostgreSQL database is running on your **local host machine**, you cannot use `localhost` or `127.0.0.1` as the `DB_HOST` inside the container (because it refers to the container itself). 

Instead, modify your `.env` file (or pass it via `-e`) to use `host.docker.internal`:

```env
DB_HOST=host.docker.internal
```

Then run the container:

```bash
docker run -p 3001:3001 --env-file .env kemitbelajar-backend
```

## API Endpoints

The API is structured under the `/api` group. Available routes include:
- `GET /health` : Health check endpoint.
- `/api/posts/*` : Endpoints for managing blog posts.
- `/api/auth/*` : Endpoints for user authentication.

## Project Structure

- `/cmd/server` - Application entry point.
- `/internal` - Core application code (handlers, repositories, config, database).
- `/migrations` - SQL migration files for the database.
