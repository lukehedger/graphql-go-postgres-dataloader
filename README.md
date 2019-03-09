# GraphQL x Go x Postgres x DataLoader

Simple skeleton for a Go GraphQL API connected to a Postgres database with DataLoader for query batching and caching.

## Contributing

### Requirements
- Go v1.11
```sh
brew install go
```

- Postgres v11
```sh
brew install postgresql
```

### Setup
Build Go module:
```sh
go build ./...
```

### Run
Start the GraphQL server:
```sh
go run server/server.go
```

Try some operations:
```graphql
query sayHello {
  hello(id:"1")
}

query people {
  peopleViaLoader(ids:["1"]) {
    id
    name
  }
}
```

### DataLoader
Generate a new DataLoader
```sh
go run github.com/vektah/dataloaden -keys string github.com/lukehedger/graphql-go-postgres-dataloader/<Type>
```

## Architecture
- `db.go` - Database setup and connection
- `model.go` - Application type definitions
- `resolver.go` - Resolve API operations, with access to database
- `schema.graphql` - GraphQL schema in SDL
- `tools.go` - Executable for `dataloaden`
- `dataloaders/dataloader.go` - DataLoader middleware
- `dataloaders/personloader_gen.go` - Generated DataLoader
- `server/server.go` - Serve GraphQL API
- `util/schema.go` - Transform SDL to parseable string
