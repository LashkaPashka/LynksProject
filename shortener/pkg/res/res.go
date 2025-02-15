package res

import (
	"encoding/json"
	"net/http"
)

func Encode(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&data)
	w.WriteHeader(statusCode)
}