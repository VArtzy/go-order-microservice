package types

import (
    "net/http"
    "github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
    rdb *redis.Client
}
