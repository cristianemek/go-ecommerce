package orders

import (
	"log"
	"net/http"

	"github.com/cristianemek/go-ecommerce/internal/json"
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

	//1. leer payload y convertirlo a un struct temporal
	var tempOrder createOrderParams //struct temporal para recibir el body de la request, con los campos necesarios para crear una orden, sin exponer el struct de la base de datos
	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest) //error en los datos enviados por el cliente, no en el servidor
		return
	}

	//2. llamar al servicio donde se va a crear la orden en la base de datos, pasando el contexto de la request y el struct temporal con los datos necesarios para crear la orden
	createOrder, err := h.service.PlaceOrder(r.Context(), tempOrder) //aquí se llama al servicio para crear la orden, pasando el contexto de la request y el struct temporal con los datos necesarios para crear la orden
	if err != nil {
		log.Println(err)

		if err == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err == ErrProductNotStock {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError) //error en el servidor al intentar crear la orden, no en los datos enviados por el cliente
		return
	}

	//3. enviar la respuesta al cliente con el struct de la orden creada, usando el json.Write para convertir el struct a json y enviarlo en la response
	json.Write(w, http.StatusCreated, createOrder)
}
