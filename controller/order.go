package controller

import (
    "fmt"
    "net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Order created")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "List of orders")
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Order retrieved")
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Order updated")
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Order deleted")
}
