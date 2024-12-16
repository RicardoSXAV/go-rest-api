.PHONY: migrate-up migrate-down migrate-create

migrate-up:
	migrate -path migrations -database "mysql://api_user:Password123@tcp(localhost:3306)/go_api_db" up

migrate-down:
	migrate -path migrations -database "mysql://api_user:Password123@tcp(localhost:3306)/go_api_db" down

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name