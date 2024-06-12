package controller

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/vartzy/order-api-microservice/model"
	"github.com/vartzy/order-api-microservice/repository/order"
)

type Order struct {
    Repo *order.RedisRepo
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
    var body struct {
        CustomerID uuid.UUID`json:"customer_id"`
        LineItems []model.LineItem`json:"line_items"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    now := time.Now().UTC()

    order := model.Order{
        OrderID: rand.Uint64(),
        CustomerID: body.CustomerID,
        LineItems: body.LineItems,
        CreatedAt: &now,
    }

    err := o.Repo.Insert(r.Context(), order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
    cursorStr := r.URL.Query().Get("cursor")
    if cursorStr == "" {
        cursorStr = "0"
    }

    const decimal = 10
    const bitSize = 64
    cursor, err := strconv.ParseUint(cursorStr, decimal, bitSize)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    const size = 10
    res, err := o.Repo.FindAll(r.Context(), order.FindAllPage{
        Offset: cursor,
        Size: size,
    })
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var response struct {
        Items []model.Order`json:"items"`
        Next uint64`json:"next,omitempty"`
    }

    response.Items = res.Orders
    response.Next = res.Cursor

    resBytes, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(resBytes)
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
    idParam := chi.URLParam(r, "id")

    const base = 10
    const bitSize = 64

    orderID, err := strconv.ParseUint(idParam, base, bitSize)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ord, err := o.Repo.GetByID(r.Context(), orderID)
    if errors.Is(err, order.ErrNotExist) {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(ord); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
    var body struct {
        Status string `json:"status"`
    }

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    idParam := chi.URLParam(r, "id")

    const base = 10
    const bitSize = 64

    orderID, err := strconv.ParseUint(idParam, base, bitSize)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ord, err := o.Repo.GetByID(r.Context(), orderID)
    if errors.Is(err, order.ErrNotExist) {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    const completedStatus = "completed"
    const shippedStatus = "shipped"
    now := time.Now().UTC()

    switch body.Status {
    case shippedStatus:
        if ord.ShippedAt != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        ord.ShippedAt = &now
    case completedStatus:
        if ord.CompletedAt != nil || ord.ShippedAt == nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        ord.CompletedAt = &now
    default:
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = o.Repo.UpdateByID(r.Context(), ord)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(ord); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
    idParam := chi.URLParam(r, "id")

    const base = 10
    const bitSize = 64

    orderID, err := strconv.ParseUint(idParam, base, bitSize)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err = o.Repo.DeleteByID(r.Context(), orderID)
    if errors.Is(err, order.ErrNotExist) {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    } else if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "order deleted"}`))
}
