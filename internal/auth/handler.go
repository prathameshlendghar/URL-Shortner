package auth

import (
	"encoding/json"
	"errors"
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
			return
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

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	tokenResp, err := h.service.LoginUser(r.Context(), "abcd@gmail.com", "prathamesh")
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			http.Error(w, "Failed Login: Invalid login credentials", http.StatusUnauthorized)
			return
		}
		slog.ErrorContext(r.Context(), "Failed Login attempt", slog.Any("error", err))
		http.Error(w, "Internal error occurred: please try again later", http.StatusInternalServerError)
		return
	}

	// 3. PUSH IT: Set the token inside an HttpOnly Cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenResp.Token,
		Expires:  tokenResp.ExpiresAt,
		HttpOnly: true,                    // Keeps JS from reading it (XSS protection)
		Secure:   false,                   // Ensures HTTPS usage (set to false ONLY in local dev)
		SameSite: http.SameSiteStrictMode, // CSRF protection
		Path:     "/",                     // Available across the whole domain
	})

	// 4. Send a clean success response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Changed to 200 OK
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}
