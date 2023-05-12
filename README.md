# GoFinance

A simple Go based REST API server built using go-chi/chi HTTP router and PostgreSQL database.

```bash
❯ make help
environment                    Setup environment.
lint                           Running golangci-lint for code analysis.
migrate-all                    Rollback migrations, all migrations
migrate-create                 Create a DB migration files e.g `make migrate-create name=migration-name`
migrate-down                   Rollback migrations, latest migration (1)
migrate-up                     Run migrations UP
server                         Running application
```
