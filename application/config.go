package application

import (
    "os"
    "strconv"
    "fmt"

    "github.com/joho/godotenv"
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

    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    redisAddr := os.Getenv("REDIS_ADDR")
    if redisAddr != "" {
        cfg.RedisAddr = redisAddr
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
