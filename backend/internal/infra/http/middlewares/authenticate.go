package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"sub-watch-backend/internal/application/auth"
	app_errors "sub-watch-backend/internal/application/errors"
)

func Authenticate(jwtService *auth.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				writeAuthError(w, "Missing authorization header")
				return
			}

			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				writeAuthError(w, "Invalid authorization header format")
				return
			}

			token := headerParts[1]

			userID, err := jwtService.ValidateToken(token)
			if err != nil {
				writeAuthError(w, "Invalid authentication token")
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func writeAuthError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	errorResponse := app_errors.ErrorsResponseDTO{
		Errors: []app_errors.Error{
			{
				Code:    http.StatusUnauthorized,
				Type:    app_errors.ERROR_INVALID_AUTHENTICATION_TOKEN,
				Message: message,
			},
		},
	}

	json.NewEncoder(w).Encode(errorResponse)
}