package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
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

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerUserReq RegisterUserReq

	if err := json.NewDecoder(r.Body).Decode(&registerUserReq); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(registerUserReq); err != nil {
		http.Error(w, "Payload validation failed", http.StatusBadRequest)
		return
	}

	err := h.service.RegisterUser(r.Context(), registerUserReq)
	if err != nil {
		if err == ErrDuplicateKey {
			http.Error(w, "Username/Email already registered", http.StatusConflict)
		}
		slog.ErrorContext(r.Context(), "Unable to create new user", slog.Any("error", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
	})
}
