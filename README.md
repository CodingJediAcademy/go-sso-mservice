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

## Tests
### Prepare
```bash
go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./tests/migrations --migrations-table=migrations_test
```
### Run App in test mode
```bash
go run ./cmd/sso --config=./config/local_tests.yaml
```
### Run tests
```bash
go test ./tests -count=1 -v
```
Параметр -count=1 — стандартный способ запустить тесты с игнорированием кэша, а -v добавить больше подробностей в вывод теста.