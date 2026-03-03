package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cristianemek/go-ecommerce/internal/products/json"
	"github.com/go-chi/chi/v5"
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
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println("error listing products:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, products)
}

func (h *handle) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(r.Context(), int64(id))
	if err != nil {
		log.Println("error getting product by id:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, product)
}
