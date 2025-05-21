FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install mvdan.cc/gofumpt@latest

COPY go.mod go.sum ./
RUN go mod download

# gets other files by docker compose volumes
CMD ["air", "-c", ".air.toml"]