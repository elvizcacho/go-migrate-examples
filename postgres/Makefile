

make migrate-up:
	go run migrate.go -up
make migrate-down:
	go run migrate.go -down

make migrate-create:
	go run migrate.go -create="$(name)" -migrationsPath="./db/migrations"

make up:
	docker network create migrations_postgres_network || true
	docker-compose up -d