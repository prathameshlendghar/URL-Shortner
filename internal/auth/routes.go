package auth

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB, v *validator.Validate) {

	repo := NewRepository(db)
	Service := NewService(repo)
	Handler := NewHandler(Service, v)

	// // ========================= Auth Routes ==============================//
	mux.HandleFunc("POST /api/v1/auth/register", Handler.RegisterUser)
	// mux.HandleFunc("POST /api/v1/auth/login", nil)

}
