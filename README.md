# Sharing Vision Backend

Simple REST API for articles, built with Go (Gin + MySQL).

## Prerequisites

- Go 1.26+
- MySQL (or TiDB / PlanetScale serverless)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate) (for step 3)

## How to run

1. Set up a MySQL database, or use TiDB / PlanetScale serverless.
2. Create the database:

```sql
create database articles;
```

3. Run migrations (requires the `migrate` CLI):

```bash
make deploy-migrate url='mysql://<username>:<password>@tcp(<hostname>:<port>)/<dbname>?<queries>'
```

4. Copy `.env.template` to `.env` and fill in the values:

```bash
cp .env.template .env
```

5. Start the service:

```bash
make
```

## Endpoints

| Method | Path         |
| ------ | ------------ |
| POST   | `/article`   |
| GET    | `/article`   |
| GET    | `/article/:id` |
| PUT    | `/article/:id` |
| DELETE | `/article/:id` |

