postgres:
	docker run --name postgres13 -dp 5432:5432 -e POSTGRES_USER=skygazer -e POSTGRES_PASSWORD=hamdalah postgres:12-alpine

createdb:
	docker exec -it postgres13 createdb --username=skygazer --owner=skygazer simple_bank

dropdb:
	docker exec -it postgres13 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://skygazer:hamdalah@localhost:5332/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://skygazer:hamdalah@localhost:5332/simple_bank?sslmode=disable" -verbose down

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

sqlcgen:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlcgen