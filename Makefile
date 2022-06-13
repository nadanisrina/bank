DB_URL=postgresql://postgres:postgres@localhost:5432/bank?sslmode=disable
postgres: 
	docker run --name image-postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:14.3-alpine

createdb: 
	docker exec -it image-postgres createdb --username=postgres --owner=postgres bank

dropdb:
	docker exec -it image-postgres dropdb --username=postgres bank

migrateup: 
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown: 
	migrate -path db/migration -database "${DB_URL}" -verbose down

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test