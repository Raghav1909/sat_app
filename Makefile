BINARY_NAME := sat-cli
DB_FILE := db.sqlite3
MIGRATIONS_DIR := ./db/migrations

SQLC_CMD := sqlc generate
GOOSE_CMD := goose -dir $(MIGRATIONS_DIR) sqlite3 $(DB_FILE)

.PHONY: all
all: build

.PHONY: sqlc
sqlc:
	$(SQLC_CMD)

.PHONY: migrate-up
migrate-up:
	$(GOOSE_CMD) up

.PHONY: create-db
create-db:
	@touch $(DB_FILE)
	@echo "Database created: $(DB_FILE)"

.PHONY: delete-db
delete-db:
	@rm -f $(DB_FILE)
	@echo "Database deleted: $(DB_FILE)"

.PHONY: reset-db
reset-db: delete-db create-db migrate-up
	@echo "Database reset complete"

.PHONY: build
build:
	@go build -o $(BINARY_NAME) ./cmd/cli
	@echo "CLI application built: $(BINARY_NAME)"
