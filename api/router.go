package api

import (
	"net/http"

	"github.com/prathameshlendghar/URL-Shortner/handlers"
)

func RoutesSetup() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("/geturl", handlers.GetMainURL)
	mux.HandleFunc("/shorten", handlers.ShortenURL)
	mux.HandleFunc("/editurl", handlers.EditURL)
	mux.HandleFunc("/deleteurl", handlers.DeleteURL)
	
	return mux
}
