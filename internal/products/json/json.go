package json

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")

	//write header al final despues de header set, si no no se asignaria el header content type, y se devolveria el default que es text/plain
	w.WriteHeader(status)           //devovler estado segun el tipo de respuesta, http standars status code
	json.NewEncoder(w).Encode(data) //encodear el data a json y escribirlo en la respuesta
}
