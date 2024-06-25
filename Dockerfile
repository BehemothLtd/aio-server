FROM golang:1.21.6-bullseye as builder-dev

WORKDIR /app

RUN apt-get update && \
  apt-get install -y redis redis-tools postgresql postgresql-client && \
  apt-get install -y make tar gzip nano vim netcat net-tools telnet iputils-ping && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
  mv migrate /bin/
COPY . /app
RUN go mod download
RUN export CGO_ENABLED=0 && export GOOS=linux && export GO111MODULE=on && go build -ldflags "-s -w" -o ./rake rake/execute.go

FROM golang:1.21.6-alpine as builder

WORKDIR /app

RUN apk add upx

COPY go.mod go.sum /app/
RUN go mod download

COPY . /app
RUN export CGO_ENABLED=0 && export GOOS=linux && export GO111MODULE=on && \
  go build -ldflags "-s -w" -o .build/main . &&\
  go build -ldflags "-s -w" -o .build/workers workers/server.go
RUN upx -9 .build/main && \
    upx -9 .build/workers

FROM gcr.io/distroless/static:latest@sha256:41972110a1c1a5c0b6adb283e8aa092c43c31f7c5d79b8656fbffff2c3e61f05
WORKDIR /app
COPY --from=builder /app/.build .
COPY ./gql/schemas /app/gql/schemas

CMD ["/app/main"]
