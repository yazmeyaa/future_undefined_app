include ../.env
export

sqlc_cfg_path = "./internal/repository/postgres/sqlc/sqlc.yml"
sqlc-gen:
	sqlc generate -f $(sqlc_cfg_path)

grpc-gen:
	protoc

DB_HOST="localhost"
DB_USER=$(PROJECTS_DB_USERNAME)
DB_PASSWORD=$(PROJECTS_DB_PASSWORD)
DB_PORT=$(PROJECTS_DB_HOST_PORT)
DB_NAME=$(PROJECTS_DB_DATABASE)
DB_URL="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

migrate-up:
	@echo "Running migrations with DB_URL: $(DB_URL)"
	migrate -path ./migrations -database $(DB_URL) up

migrate-down:
	@echo "Down migrations with DB_URL: $(DB_URL)"
	migrate -path ./migrations -database $(DB_URL) down
