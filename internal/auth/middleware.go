package auth

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "userID"

func getJWTSecret(token *jwt.Token) (interface{}, error) {
	// Here we can write the logic of getting the secret from wherever stored
	// As of now it is hardcoded in program so no issue
	return jwtSecret, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}
		tokenString := cookie.Value
		claims := &CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, getJWTSecret)

		if err != nil || !token.Valid {
			// Token is expired, tampered with, or signed with the wrong key. Kick them out.
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Assumes claims.UserID is a string. If it's an int, check: claims.UserID == 0
		if claims.UserID == 0 {
			http.Error(w, "Unauthorized: Malformed token payload", http.StatusUnauthorized)
			return // Short-circuit right here! Do NOT call next.ServeHTTP
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)

		requestWithUserInfo := r.WithContext(ctx)

		next.ServeHTTP(w, requestWithUserInfo)
	})
}
