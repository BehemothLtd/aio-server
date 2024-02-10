# Intall Go

```
brew install go
```

# Install go-migrate

```bash
brew install golang-migrate
```

# DB Migration

## Create new migration

```bash
migrate create -ext sql -dir database/migrations -seq create_users_table
```

## Migrate

```bash
make db.migrate
```

# Graphql generate command

```bash
make graphql.generate
```

# Install dependencies

```
go mod tidy
```

# Kickoff

```bash
go run main.go
```
