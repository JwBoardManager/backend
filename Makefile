# Carrega variáveis do .env
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: migrate-up migrate-down sqlc-generate run

# Criar nova migração
migrate-new:
	migrate create -ext sql -dir db/migrations -seq $(name)

# Aplicar todas as migrações
migrate-up:
	migrate -database "$(DATABASE_URL)" -path db/migrations up

# Reverter última migração
migrate-down:
	migrate -database "$(DATABASE_URL)" -path db/migrations down 1

# Gerar código SQLC
sqlc-generate:
	sqlc generate -f db/sqlc/sqlc.yaml

# Obter a última versão aplicada das migrações
LAST_VERSION := $(shell psql "$(DATABASE_URL)" -At -c "SELECT version FROM schema_migrations ORDER BY version DESC LIMIT 1;")

# Forçar a correção do estado de migração "dirty" para a última versão
migrate-force:
	@echo "⚠️ Banco de dados está 'dirty'. Corrigindo para versão $(LAST_VERSION)..."
	migrate -database "$(DATABASE_URL)" -path db/migrations force $(LAST_VERSION)


# Resetar todas as migrações (⚠️ Isso deleta todos os dados!)
migrate-reset:
	@echo "⚠️ Resetando todas as migrações e recriando o banco..."
	migrate -database "$(DATABASE_URL)" -path db/migrations down
	migrate -database "$(DATABASE_URL)" -path db/migrations up

# Rodar aplicação
run:
	go run cmd/server/main.go
