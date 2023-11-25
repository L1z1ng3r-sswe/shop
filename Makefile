run:
	go run cmd/main.go

migrate:
	migrate create -ext sql -dir migrations -seq schema

migrate-up:
	migrate -path migrations -database "postgres://postgres:asyl12345.@localhost:5432/postgres?sslmode=disable" up

migrate-down: 
	migrate -path migrations -database "postgres://postgres:asyl12345.@localhost:5432/postgres?sslmode=disable" down

migrate-dirty:
	migrate -path migrations -database "postgres://postgres:asyl12345.@localhost:5432/postgres?sslmode=disable" force 1
