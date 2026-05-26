package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func SetupRoutes(db *sql.DB, v *validator.Validate) *http.ServeMux {
	// Mux ~ Multiplexer i.e Router. We register routes and their respective handler in it, and return the router back to the webserver serving form main
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Engine is running smoothly!"))
		if err != nil {
			log.Println("Error writing response: ", err)
		}
	})

	repo := NewRepository(db)
	Service := NewService(repo)
	Handler := NewHandler(Service, v)

	// ==================== Core functionality routes ======================//
	mux.HandleFunc("POST /api/v1/urls", nil)
	mux.HandleFunc("GET /{alias}", nil)
	mux.HandleFunc("GET /api/v1/urls", nil)
	mux.HandleFunc("DELETE /api/v1/urls/{alias}", nil)

	// ========================= Auth Routes ==============================//
	mux.HandleFunc("POST /api/v1/auth/register", nil)
	mux.HandleFunc("POST /api/v1/auth/login", nil)

	return mux
}
