package middleware

import (
	"context"
	"net/http"

	"github.com/YagoSchramm/base-auth-v1/foundation"
	"github.com/YagoSchramm/base-auth-v1/model"
	"github.com/gorilla/mux"
)

func Identify(
	userLoader func(ctx context.Context, username string) (*model.UserEntityDomain, error),
) mux.MiddlewareFunc {
	return mux.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			username := foundation.FromBase64(token)
			user, err := userLoader(r.Context(), username)
			if err != nil {
				http.Error(w, "invalid Credentials", http.StatusForbidden)
				return
			}
			ctx := context.WithValue(r.Context(), "user", user)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	})
}
