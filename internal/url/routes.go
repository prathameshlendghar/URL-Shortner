package url

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB, v *validator.Validate) {

	repo := NewRepository(db)
	Service := NewService(repo)
	Handler := NewHandler(Service, v)

	// ==================== Core functionality routes ======================//
	mux.HandleFunc("POST /api/v1/urls", Handler.CreateNewShortUrl)
	mux.HandleFunc("GET /{alias}", Handler.RedirectShortUrl)
	// mux.HandleFunc("GET /api/v1/urls", nil)
	// mux.HandleFunc("DELETE /api/v1/urls/{alias}", nil)

}
