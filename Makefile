db_user = root
db_password = secret
db_host = localhost
db_port = 5432
db_name = test
ssl_mode = disable

.PHONY: install-tools dbup dbdown dropdb createdb psql migrateup migratedown test server swagger

install-tools:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@# TODO: use this instead? @go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

	@# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

dbup:
	@docker run --rm --name postgres15 -p 5432:5432 -e POSTGRES_USER=$(db_user) -e POSTGRES_PASSWORD=$(db_password) -e POSTGRES_DB=$(db_name) -d postgres:15-alpine

dbdown:
	@docker stop postgres15

dropdb:
	@docker exec -it postgres15 dropdb $(db_name)

createdb:
	@docker exec -it postgres15 createdb --username=$(db_user) --owner=$(db_user) $(db_name)

psql:
	@docker exec -it postgres15 psql -U $(db_user) -d $(db_name)

migrateup:
	@migrate -path db/migration -database "postgresql://$(db_user):$(db_password)@$(db_host):$(db_port)/$(db_name)?sslmode=$(ssl_mode)" -verbose up

migratedown:
	@migrate -path db/migration -database "postgresql://$(db_user):$(db_password)@$(db_host):$(db_port)/$(db_name)?sslmode=$(ssl_mode)" -verbose down

test:
	@go test -v -cover ./...

server:
	@go run cmd/server-rest/server-rest.go

swagger:
	@swag init -g cmd/server-rest/server-rest.go --output static --packageName docs
	  mv static/docs.go internal/docs/docs.go

