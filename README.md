Run

```
newgrp docker
```

For running the access to docker daemon

Start the Docker DB
If you created docker-compose.yml:
Bash

```
docker compose up -d
```

Or if starting the existing container directly:
Bash

```
docker start gator-db
```

Bash

```
psql "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
```

Run Migration
Step 1:

```
cd sql/schema/
```

Step 2:

```
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up
```

To Undo Migrations

```
goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" down
```
