FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o binapp ./cmd/main.go

FROM scratch

COPY --from=build /app/binapp /

ENTRYPOINT [ "/binapp" ]