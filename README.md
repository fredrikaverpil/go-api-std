# go-api-std

Go APIs using as much of the standard library as possible.

## Features

- REST server.
- Dummy store (with in-memory db), used by unit tests.
- Postgres store.
- Migrations via golang-migrate.
- Hashing of passwords.

## Todo/WIP:

- Use SQLite for test db instead of Postgres.
- Simple CRUD.
- OpenAPI/Swagger.
- Error responses follows JSON:API spec.
- JWT tokens.
- Take some inspiration from
  - https://github.com/golang-standards/project-layot
  - https://github.com/araujo88/golang-rest-api-template
- gRPC server
- gRPC API gateway
- GraphQL server
