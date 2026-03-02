package products

import (
	"log"
	"net/http"

	"github.com/cristianemek/go-ecommerce/internal/products"
	"github.com/cristianemek/go-ecommerce/internal/products/json"
)

type handle struct {
	service Service
}

//constructor

func NewHandler(service Service) *handle {
	return &handle{service: service}
}

// w http.ResponseWriter, es la response que se va a enviar al cliente, y r *http.Request es la solicitud que se recibe del cliente (donde viene body, query params, etc)
func (h *handle) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println("error listing products:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	products := []string{"product1", "product2", "product3"}

	json.Write(w, http.StatusOK, products)
}
