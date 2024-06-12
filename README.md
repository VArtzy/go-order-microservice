# go-order-microservice
Golang orders REST API microservices.

## How to use

### You can setting up .env file before or using default env value provided.

## Via Docker

- Build the container

```docker-compose build```

- Start the container

```docker-compose up```

## Via Command Line

- Build microservice

Prerequisite:
- go 1.21.5
- redis

```go build main.go```

- Running your redis

make sure you have redis server listened at given port (6379 default). Every OS have different setup to run redis server.

- Running executeable from go build
