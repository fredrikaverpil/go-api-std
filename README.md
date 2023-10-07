# go-api-std

Go APIs using as much of the standard library as possible.

## Features

- REST server.
- Dummy store (with in-memory db), used by unit tests.
- Postgres store.
- Migrations via golang-migrate.
- Hashing of passwords.

## Architectural choices

- Layered API design:
  - domain, models and stores
    - services
      - rest
- Custom errors are defined in the domain layer and can be used throughout the layers (but not useful in the rest layer). Each custom domain error has an internal code which you map to a desired http status code.
- In the rest layer, the custom domain errors are translated into http errors (using an error-to-status mapper).
  - When using this mapper, you must define the http status codes you are willing to have your route return, so to prevent the endpoint to start returning new error codes unexpectedly, following e.g. domain logic refactorings. Could also potentially make it easier to keep endpoint docs in sync with the code.
- Ultimately, you should be able to tell from each rest server method which possible http status codes can be returned.

## Todo/WIP:

- Use SQLite for test db instead of Postgres.
- Simple CRUD.
- OpenAPI/Swagger.
- Error responses follows JSON:API spec.
- JWT tokens.
- Take some inspiration from
  - https://github.com/golang-standards/project-layot
  - https://github.com/araujo88/golang-rest-api-template
- gRPC server.
- gRPC API gateway.
- Protobuf (Google AIPs).
- GraphQL server.
