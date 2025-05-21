FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /app

RUN go install mvdan.cc/gofumpt@latest

# gets all project files by docker compose volumes
ENTRYPOINT ["sh", "-c", "\
    DIFF=$(gofumpt -d .); \
    if [ -n \"$DIFF\" ]; then \
        echo '❌ Code is not formatted:'; \
        echo \"$DIFF\"; \
        exit 1; \
    else \
        echo '✅ Code is formatted'; \
    fi"]