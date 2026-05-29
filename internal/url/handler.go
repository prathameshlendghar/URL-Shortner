package url

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	customValidator "github.com/prathameshlendghar/URL-Shortner/pkg/validator"
)

type Handler struct {
	service   *Service
	validator *validator.Validate
}

func NewHandler(service *Service, v *validator.Validate) *Handler {
	return &Handler{
		validator: v,
		service:   service,
	}
}

type contextKey string

const UserIDKey contextKey = "userID"

func GetDummyUserThing(r **http.Request) {
	dummyUserID := 1 // Ensure user ID 1 exists in your Postgres DB!
	ctx := context.WithValue((*r).Context(), UserIDKey, dummyUserID)
	*r = (*r).WithContext(ctx)
}

func (h *Handler) CreateNewShortUrl(w http.ResponseWriter, r *http.Request) {
	GetDummyUserThing(&r)

	var newShortUrlReq CreateUrl

	err := json.NewDecoder(r.Body).Decode(&newShortUrlReq)
	if err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if err = h.validator.Struct(newShortUrlReq); err != nil {
		http.Error(w, "Payload validation failed", http.StatusBadGateway)
		return
	}

	isValid := customValidator.IsValidURL(newShortUrlReq.ActualUrl)
	if !isValid {
		http.Error(w, "Invalid long url format", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		// If there's no ID in the context, they aren't logged in.
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	resp, err := h.service.CreateNewShortUrl(r.Context(), newShortUrlReq, userID)
	if err != nil {
		slog.ErrorContext(r.Context(), "Failed creating new short url", slog.Any("error", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.ErrorContext(r.Context(), "Failed to encode created URL response", slog.Any("error", err))
	}
}

func (h *Handler) RedirectShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("alias")

	resp, err := h.service.GetActualUrl(r.Context(), shortUrl)
	if err != nil {
		if err == ErrUrlNotFound {
			http.Error(w, "Invalid short url", http.StatusNotFound)
			return
		}
		if err == ErrUrlNotActive {
			http.Error(w, "Url is not accessible", http.StatusGone)
			return
		}
		slog.ErrorContext(r.Context(), "Failed to fetch the mapped URL", slog.Any("error", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	//FIX: Prevent all browser caching to guarantee 100% accurate click tracking
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache") // For HTTP/1.0 legacy clients
	w.Header().Set("Expires", "0")       // Proxies

	//FIX: Use 302 to force a safe GET request on the destination
	http.Redirect(w, r, resp, http.StatusFound)
}