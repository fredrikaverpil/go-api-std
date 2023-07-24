dbup:
	@docker run --rm --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=test -d postgres:15-alpine

dbdown:
	@docker stop postgres15

psql:
	@docker exec -it postgres15 psql -U root -d test

test:
	@go test -v -cover ./...

server:
	@go run main.go

.PHONY: dbup dbdown psql test server
