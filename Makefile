start_postgres_container:
	docker run -d --name postgres14 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:14-alpine
.PHONY: start_postgres_container

start_simple_bank_container:
	docker run --name simple_bank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgres://root:secret@postgres14:5432/simple_bank?sslmode=disable" simple_bank:latest
.PHONY: start_simple_bank_container

connect_postgres_container:
	docker exec -it postgres14 /bin/sh
.PHONY: connect_postgres_container

create_db_simple_bank:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
.PHONY: create_db_simple_bank

drop_db_simple_bank:
	docker exec -it postgres14 dropdb simple_bank
.PHONY: drop_db_simple_bank

migrate_up:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
.PHONY: migrate_up

migrate_up_1:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
.PHONY: migrate_up_1

migrate_down:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
.PHONY: migrate_down

migrate_down_1:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
.PHONY: migrate_down_1

sqlc:
	sqlc generate
.PHONY: sqlc

test:
	go test ./...
.PHONY: test

test_ignore_cache:
	go test -count=1 ./...
.PHONY: test_ignore_cache

test_with_coverage:
	go test -count=1 -v -cover ./...
.PHONY: test_with_coverage

server_start:
	go run main.go
.PHONY: server_start

mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/ericlamnguyen/simple-bank/db/sqlc Store
.PHONY: mockgen