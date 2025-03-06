package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJSONUtils(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
