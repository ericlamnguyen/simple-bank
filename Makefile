start_postgres_instance:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
.PHONY: start_postgres_instance

connect_postgres_instance:
	docker exec -it postgres14 /bin/sh
.PHONY: connect_postgres_instance

create_db:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
.PHONY: create_db

drop_db:
	docker exec -it postgres14 dropdb simple_bank
.PHONY: drop_db

migrate_up:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
.PHONY: migrate_up

migrate_down:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
.PHONY: migrate_down

sqlc:
	sqlc generate
.PHONY: sqlc

test:
	go test -v -cover ./...
.PHONY: test

server_start:
	go run main.go
.PHONY: server_start