package products

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/cristianemek/go-ecommerce/internal/json"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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
	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("error getting product by id:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Write(w, http.StatusOK, product)
}

func (h *handle) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var tempProduct CreateProductParams
	if err := json.Read(r, &tempProduct); err != nil {
		log.Println("error reading request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest) //400 error en los datos enviados, mala estructura, etc
		return
	}

	product, err := h.service.CreateProduct(r.Context(), tempProduct)
	if err != nil {
		log.Println("error creating product:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError) //500 error en el servidor al intentar crear el producto, valores de los campos no validos
		return
	}
	json.Write(w, http.StatusCreated, product) //201 producto creado exitosamente, se devuelve el producto creado en la response
}
