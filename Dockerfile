FROM golang:1.21.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main /app/.env .

EXPOSE 3000

CMD ["./main"]
