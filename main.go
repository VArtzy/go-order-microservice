package main

import (
    "context"
    "fmt"
    "os"
    "os/signal"

    "github.com/vartzy/order-api-microservice/application"
)

func main() {
    app := application.New()

    ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
    defer cancel()

    err := app.Start(ctx)
    if err != nil {
        fmt.Println("failed to start app:", err)
    }
}
