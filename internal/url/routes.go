package url

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/prathameshlendghar/URL-Shortner/internal/auth"
	"github.com/prathameshlendghar/URL-Shortner/internal/server"
)

func SetupRoutes(mux *http.ServeMux, db *sql.DB, v *validator.Validate) {

	repo := NewRepository(db)
	Service := NewService(repo)
	Handler := NewHandler(Service, v)

	// ==================== Core functionality routes ======================//
	mux.Handle("POST /api/v1/urls", server.Chain(http.HandlerFunc(Handler.CreateNewShortUrl), auth.AuthMiddleware))
	mux.HandleFunc("GET /{alias}", Handler.RedirectShortUrl)
	// mux.HandleFunc("GET /api/v1/urls", nil)
	// mux.HandleFunc("DELETE /api/v1/urls/{alias}", nil)

}
