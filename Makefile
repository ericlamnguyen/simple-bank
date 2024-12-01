start_local_postgres_container:
	docker run -d --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:14-alpine
.PHONY: start_postgres_container

start_local_simple_bank_container:
	docker run --name simple_bank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgres://root:secret@postgres:5432/simple_bank?sslmode=disable" simple_bank:latest
.PHONY: start_simple_bank_container

connect_local_postgres_container:
	docker exec -it postgres /bin/sh
.PHONY: connect_postgres_container

create_db_simple_bank:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
.PHONY: create_db_simple_bank

drop_db_simple_bank:
	docker exec -it postgres dropdb simple_bank
.PHONY: drop_db_simple_bank

migrate_up:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
.PHONY: migrate_up

migrate_down:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
.PHONY: migrate_down

migrate_up_1:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
.PHONY: migrate_up_1

migrate_down_1:
	migrate -path ./db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
.PHONY: migrate_down_1

test:
	go test -count=1 ./...
.PHONY: test

test_with_coverage:
	go test -count=1 -v -cover ./...
.PHONY: test_with_coverage

server_start:
	go run main.go
.PHONY: server_start

sqlc:
	sqlc generate
.PHONY: sqlc

mockgen:
	mockgen -package mockdb -destination db/mock/store.go github.com/ericlamnguyen/simple-bank/db/sqlc Store
.PHONY: mockgen

generate_32_byte_symmetric_key:
	openssl rand -hex 16
.PHONY: generate_32_byte_symmetric_key

retrieve_aws_secrets:
	aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq --raw-output 'to_entries|map("\(.key)=\(.value)")|.[]'
.PHONY: retrieve_aws_secrets

docker_login_to_ECR:
	aws ecr get-login-password | docker login --username AWS --password-stdin 273354660396.dkr.ecr.us-east-1.amazonaws.com
.PHONY: docker_login_to_ECR
