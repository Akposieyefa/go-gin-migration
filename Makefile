DB_MIGRATIONS_DIR :=./database/migrations
DB_URL :="postgresql://dbuser:dbpassword!@localhost:5432/dbname?sslmode=disable"
run:
	@go run cmd/main.go
create-migration:
ifdef table_name
	migrate create -ext sql -dir $(DB_MIGRATIONS_DIR) create_$(table_name)_table
else 
	@echo "Please provide a table_name argument", e.g make create-migration table_name=love
endif
	
migrate-up:
	migrate -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR) -verbose  up
migrate-down:
	migrate -database=$(DB_URL) -path=$(DB_MIGRATIONS_DIR) down

