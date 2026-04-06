# Surfbook v1

Minimal Go backend for the Surfbook project. This repo contains a basic web server, Postgres setup, and initial domain models for notebooks, tags, and related content.

## Tech stack
- Go 1.25.6
- PostgreSQL (via Docker)
- Gorilla Mux router

## Project layout
- `app/cmd/api/main.go` server entrypoint
- `app/foundation` database connection setup
- `app/model` domain models (notebooks)
- `app/repository` database repositories
- `app/service` business services
- `migrations` SQL migration files
- `docs/models.md` data model notes

## Prerequisites
- Go 1.25.6
- Docker + Docker Compose

## Setup
1. Start Postgres

```bash
# from repo root
docker-compose up -d
```

2. Create the database

```bash
docker-compose exec postgres psql -U postgres -c "CREATE DATABASE surfbook_dev;"
```

3. Run migrations

```bash
docker-compose exec -T postgres psql -U postgres -d surfbook_dev ./migrations/00001-create-tables.up.ddl.sql
```

4. Run the API

```bash
cd app
go run ./cmd/api
```

The API listens on `:8000` and currently does not register routes yet.

## Tests

```bash
cd app
go test ./...
```

## Notes
- The connection string is currently hardcoded in `app/cmd/api/main.go`:
  `postgres://postgres:pass@localhost:5432/surfbook_dev?sslmode=disable`
- Database schema reference: `docs/models.md`

## Troubleshooting
- Verify database exists:

```bash
docker-compose exec postgres psql -U postgres -tAc "SELECT 1 FROM pg_database WHERE datname='surfbook_dev';"
```
