# IcompHub Go API

API made with Go, Gin, Gorm and Swagger for IcompHub applications

## How to run

The api depends on a **database**, make sure you have one running

If you don't have a database right there with you, don't worry, see [icomphub-db repository](https://github.com/IcompHub/icomphub-db) and be happy

## Update Enviroments

You must update the `.env`

```
INTERNAL_API_PORT=8000 #port for the api to run inside docker container
SWAGGER_API_URL=localhost:8016 #url to access swagger on your local machine
DB_HOST=host.docker.internal #docker can't see your localhost
DB_PORT=5416
DB_USER=icomphub
DB_PASSWORD=development0nly
DB_NAME=icomphub
```

## Setup DB

Inside the `db-setup.sql` file there is a sql script to create all tables, enums and data required to run de app, make shure to run it at you db

## Run Docker

Now just run this on your terminal:

```
docker compose up -d
```

Your api should start running at port 8016, take a look at http://localhost:8016/swagger/index.html
