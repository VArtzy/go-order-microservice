package main

import (
    "context"
    "fmt"

    "github.com/vartzy/order-api-microservice/application"
)

func main() {
    app := application.New()

    err := app.Start(context.TODO())
    if err != nil {
        fmt.Println("failed to start app:", err)
    }
}
