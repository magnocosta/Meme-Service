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
