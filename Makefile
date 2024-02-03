include .env
export

MIGRATIONS_FOLDER =database/migrations
DATABASE_URL = postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST)/$(DB_NAME)?sslmode=disable

run.serve: 
	go run main.go

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

database.import:
	go run database/import.go

migrate.clean:
	make migrate.down
	make migrate.up
	make database.import