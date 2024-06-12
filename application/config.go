package application

import (
    "os"
    "strconv"
)

type Config struct {
    RedisAddr string
    ServerPort uint16
}

func LoadConfig() Config {
    cfg := Config{
        RedisAddr: "localhost:6379",
        ServerPort: 3000,
    }

    redisHost := os.Getenv("REDIS_ADDR")
    if redisHost != "" {
        cfg.RedisAddr = redisHost
    }

    serverPort := os.Getenv("SERVER_PORT")
    if serverPort != "" {
        port, err := strconv.ParseUint(serverPort, 10, 16)
        if err == nil {
            cfg.ServerPort = uint16(port)
        }
    }

    return cfg
}
