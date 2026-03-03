package orders

import (
	"log"
	"net/http"

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
func (h *handle) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var tempOrder createOrderParams //struct temporal para recibir el body de la request, con los campos necesarios para crear una orden, sin exponer el struct de la base de datos
	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.PlaceOrder(r.Context(), tempOrder)

	json.Write(w, http.StatusOK, nil)
}
