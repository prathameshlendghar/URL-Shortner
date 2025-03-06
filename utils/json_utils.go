package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func WriteJSONUtils(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func JSONRequestDecode(body io.Reader, decodedStruct *interface{}) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(decodedStruct)

	return err
}
