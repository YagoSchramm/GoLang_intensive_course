package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type userContextKey struct{}

type errorResponse struct {
	Error string `json:"error"`
}

func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				respondError(w, http.StatusUnauthorized, "Authorization header ausente")
				return
			}

			if !strings.HasPrefix(authorization, "Bearer ") {
				respondError(w, http.StatusUnauthorized, "Authorization inválido")
				return
			}

			tokenString := strings.TrimPrefix(authorization, "Bearer ")
			claims := &jwt.RegisteredClaims{}

			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("método de assinatura inválido")
				}
				return []byte(jwtSecret), nil
			})
			if err != nil || token == nil || !token.Valid {
				respondError(w, http.StatusUnauthorized, "Token inválido")
				return
			}

			if claims.Subject == "" {
				respondError(w, http.StatusUnauthorized, "Token sem usuário")
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), userContextKey{}, claims.Subject))
			r.Header.Set("user-id", claims.Subject)
			next.ServeHTTP(w, r)
		})
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorResponse{Error: message})
}
