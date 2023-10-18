# go-api-std

Go APIs using as much of the standard library as possible.

Please note this is project is a playground for me while learning Go. 😄

## Quickstart

Start Docker Desktop, or similar. Then run...

```bash
brew install golang-migrate

make dbup
make createdb
make migrateup

make swagger

make server
```

## Features

- REST server.
- Dummy store (with in-memory db), used by unit tests.
- Postgres store.
- Migrations via golang-migrate.
- Hashing of passwords.
- Swagger
  - Generate JSON/YAML via github.com/swaggo/swag
  - Serve /docs via github.com/swaggo/http-swagger

## Architectural choices

- Layered API design:
  - domain, models and stores
    - services
      - rest
- [Custom errors](internal/domain/error.go) are defined in the domain layer and can be used throughout the layers (but not useful in the rest layer).
- In the rest layer, the custom domain errors are translated into http errors (using an [error-to-status mapper](internal/rest/error_resp_mapper.go)).
  - When using this mapper, you must define the http status codes you are willing to have your route return, so to prevent the endpoint to start returning new error codes unexpectedly, following e.g. domain logic refactorings. Could also potentially make it easier to keep endpoint docs in sync with the code.
- Ultimately, you should be able to tell from each rest server method which possible http status codes can be returned.

## Todo/WIP:

- Use SQLite for test db instead of Postgres.
- Simple CRUD.
- OpenAPI/Swagger.
- Error responses follows JSON:API spec.
- JWT tokens.
- Take some inspiration from
  - https://github.com/golang-standards/project-layout
  - https://github.com/araujo88/golang-rest-api-template
- gRPC server.
- gRPC API gateway.
- Protobuf (Google AIPs).
- GraphQL server.

## Error handling

- In API route functions, pass any error through the error-to-response mapper so to return a valid API error response.
- In service functions, wrap any errors with a suitable domain error and return them to the API route functions.
- In any other functions, just return regular, normal, errors.
