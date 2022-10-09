tinyurldb: 
	docker run --name tinyurldb -p 5432:5432 -e POSTGRES_USER=tinyurl -e POSTGRES_PASSWORD=tinyurltinyurl -d postgres:latest

createdb:
	docker exec -it tinyurl_db createdb --username=tinyurl --owner=tinyurl tinyurl

dropdb:
	docker exec -it tinyurl_db dropdb tinyurl

migrateup:
	migrate -path db/migration -database "postgresql://tinyurl:tinyurltinyurl@tinyurl.cruwjtugxx5v.us-east-1.rds.amazonaws.com:5432/tinyurl" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://tinyurl:tinyurltinyurl@tinyurl.cruwjtugxx5v.us-east-1.rds.amazonaws.com:5432/tinyurl" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc test server migrateuplocal

