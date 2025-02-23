GO_BIN := $(shell go env GOPATH)/bin
GOOSE := $(GO_BIN)/goose

app.start:
	go run ./cmd/api/main.go

docker.start:
	docker-compose -f build/container/docker-compose.yml up -d

docker.stop:
	docker-compose -f build/container/docker-compose.yml down

docker.build:
	docker-compose -f build/container/docker-compose.yml build

docker.ps:
	docker-compose -f build/container/docker-compose.yml ps

docker.db.start:
	docker-compose -f build/container/docker-compose.yml up -d db

docker.db.stop:
	docker-compose -f build/container/docker-compose.yml rm -fs db

docker.db.sql:
	docker-compose -f build/container/docker-compose.yml run db psql -h db -U meme -d meme_db

tools.goose.up:
	@$(GOOSE) up

tools.goose.down:
	@$(GOOSE) down

tools.goose.version:
	@$(GOOSE) -version

tools.goose.install:
	@echo "🔧 Checking for goose..."
	@if ! [ -x "$(GOOSE)" ]; then \
		echo "🚀 Installing goose..."; \
		go install github.com/pressly/goose/v3/cmd/goose@latest; \
	else \
		echo "✅ Goose is already installed"; \
	fi

tools.goose.create:
	@if [ -z "$(name)" ]; then \
		echo "❌ Error: Please provide a migration name. Usage: make tools.goose.migration name=your_migration"; \
		exit 1; \
	fi
	@echo "🚀 Creating new migration: $(name)"
	@$(GOOSE) create $(name) sql
