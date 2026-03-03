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

func Read(r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() //evitar que se envien campos que no existen en el struct, para evitar errores de tipeo o campos innecesarios
	return decoder.Decode(data)     //decodear el body de la request al struct data, y devolver error si no se puede decodear
}
