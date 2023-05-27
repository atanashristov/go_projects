DB_URL=postgresql://postgres:post6res@127.0.0.1:5433/simple_bank?sslmode=disable

network:
	docker network create simplebank-network

postgres:
	docker run --name simplebank-postgres15 --network simplebank-network -p 5433:5432 -e POSTGRES_PASSWORD=post6res -d postgres:15.3-alpine

createdb:
	docker exec -it simplebank-postgres15 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it simplebank-postgres15 dropdb --username=postgres simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

psql:
	docker exec -it simplebank-postgres15 psql -U postgres simple_bank

dbshell:
	docker exec -it simplebank-postgres15 /bin/sh

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration psql dbshell sqlc
