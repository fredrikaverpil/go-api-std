# variables
db_user = root
db_password = secret
db_host = localhost
db_port = 5432
db_name = test
ssl_mode = disable

# commands
swag = go run github.com/swaggo/swag/cmd/swag@latest

# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain
migrate = go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest


.PHONY: dbup dbdown dropdb createdb psql migrateup migratedown test server swagger

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
	@$(migrate) -path db/migration -database "postgresql://$(db_user):$(db_password)@$(db_host):$(db_port)/$(db_name)?sslmode=$(ssl_mode)" -verbose up

migratedown:
	@$(migrate) -path db/migration -database "postgresql://$(db_user):$(db_password)@$(db_host):$(db_port)/$(db_name)?sslmode=$(ssl_mode)" -verbose down

test:
	@go test -v -cover ./...

vet:
	@go vet ./...

build:
	@go build -o bin/server-rest cmd/server-rest/server-rest.go

server:
	@go run cmd/server-rest/server-rest.go

swagger:
	@$(swag) init -g cmd/server-rest/server-rest.go --output static --packageName docs
	@mv static/docs.go internal/docs/docs.go

