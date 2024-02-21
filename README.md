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

# Install dependencies

Install `go-enum`

```bash
curl -fsSL "https://github.com/abice/go-enum/releases/download/v0.6.0/go-enum_$(uname -s)_$(uname -m)" -o go-enum

chmod +x go-enum (Not required maybe)
```

```bash
go mod tidy
```

# Kickoff

```bash
go run main.go
```

## Enum generator

```bash
./go-enum --sqlint --marshal -f ./enums/{path_to_enum_file}
```
