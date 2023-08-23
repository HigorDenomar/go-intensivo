package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/higordenomar/go-intensivo/internal/entity"
)

func main() {
	route := chi.NewRouter()
	route.Use(middleware.Logger)

	route.Get("/order", Order)

	http.ListenAndServe(":3333", route)
}

func Order(w http.ResponseWriter, r *http.Request) {
	order, err := entity.NewOrder("3", 10.0, 1.0)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	order.CalculateFinalPrice()

	json.NewEncoder(w).Encode(order)
}
