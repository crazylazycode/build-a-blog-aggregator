# Gator - CLI Blog Aggregator

Gator is a multi-user CLI application written in Go that fetches, parses, and aggregates RSS feeds into a PostgreSQL database in real time. It allows users to register accounts, manage RSS feed subscriptions, run a background aggregator service, and read aggregated posts right from the terminal.

---

## Prerequisites

Ensure you have the following installed on your system:

- **Go** (v1.20 or later)
- **PostgreSQL 16** (or **Docker** / **Docker Compose**)
- **Goose** (database migration tool)

---

## Installation

Install the `gator` CLI binary using `go install`:

```bash
go install [github.com/crazytobeawake/gator@latest](https://github.com/crazytobeawake/gator@latest)
```

Local Development Setup

1. Docker Permissions & Container Setup

If running Docker without sudo, activate the group permissions in your current shell:
Bash

newgrp docker

Start the database using Docker Compose:
Bash

docker compose up -d

(Or start an existing container directly with docker start gator-db) 2. Global Configuration File

Create or edit your local configuration file at ~/.gatorconfig.json:
JSON

{
"db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
"current_user_name": ""
}

3. Database Migrations

Apply database schema migrations using Goose:
Bash

cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
cd ../..

To roll back a migration:
Bash

cd sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" down
cd ../..

4. Direct Database Connection

Inspect the database using psql:
Bash

psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
