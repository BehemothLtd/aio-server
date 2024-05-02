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

# Kickoff async server

```bash
go run workers/main.go
```

# Monitor Async queues

`.zshrc`

```bash
export GOPATH=$HOME/go
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```

```
go install github.com/hibiken/asynq/tools/asynq@latest

cd $GOPATH/bin

asynq dash
```

## Enum generator

note that `--sqlint` is only for enums of DB model that need enum attribute to be present as string but stored as `int` in DB

```bash
./go-enum --sqlint --marshal -f ./enums/{path_to_enum_file}
```
