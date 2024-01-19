# go-sso-mservice

## Migrations
Main migrations
```bash
go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
```

Test Migrations
```bash
go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test
```