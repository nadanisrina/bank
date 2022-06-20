DB_URL=postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable
postgres: 
	docker run --name image-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14.3-alpine

createdb: 
	docker exec -it image-postgres createdb --username=postgres --owner=postgres bank

dropdb:
	docker exec -it image-postgres dropdb --username=postgres bank

migrateup: 
	migrate -path db/migrations -database "${DB_URL}" -verbose up

migrateversion: 
	migrate -source=file://db/migrations -database "${DB_URL}" up 2

migratedown: 
	migrate -path db/migrations -database "${DB_URL}" -verbose down

newmigrate: 
	migrate create -ext sql -dir db/migrations -seq alter_users_column_token_up

test: 
	go test -v -cover ./...

run: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateversion test