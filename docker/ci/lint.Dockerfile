FROM golangci/golangci-lint:v2.1.6
WORKDIR /app

# gets all project files by docker compose volumes
ENTRYPOINT [ "golangci-lint", "run" ]