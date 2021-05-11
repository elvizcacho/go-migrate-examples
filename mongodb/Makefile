migrate-up:
	go run migrate.go -up -migrationsPath="./db/migrations"
migrate-down:
	go run migrate.go -down -migrationsPath="./db/migrations"
migrate-up-all:
	go run migrate.go -upAll -migrationsPath="./db/migrations"
migrate-down-all:
	go run migrate.go -downAll -migrationsPath="./db/migrations"

migrate-create:
	go run migrate.go -create="$(name)" -migrationsPath="./db/migrations"

up:
	docker network create migrations_mongodb_network || true
	docker-compose up -d