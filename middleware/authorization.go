package middleware

import (
	"context"
	"net/http"

	"github.com/luquxSentinel/housebid/tokens"
)

func AuthorizationMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwttoken := r.Header.Get("authorization")
		if jwttoken == "" {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
		}

		uid, err := tokens.ValidateJWT(jwttoken)
		if err != nil {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
		}

		ctx := context.WithValue(r.Context(), "uid", uid)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
