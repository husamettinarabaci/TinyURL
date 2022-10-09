tinyurldb: 
	docker run --name tinyurldb -p 5432:5432 -e POSTGRES_USER=tinyurl -e POSTGRES_PASSWORD=tinyurl -d postgres:latest

createdb:
	docker exec -it tinyurl_db createdb --username=root --owner=root tinyurl

dropdb:
	docker exec -it tinyurl_db dropdb tinyurl

migrateup:
	migrate -path db/migration -database "postgresql://tinyurl:tinyurl@localhost:5432/tinyurl?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://tinyurl:tinyurl@localhost:5432/tinyurl?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc test

