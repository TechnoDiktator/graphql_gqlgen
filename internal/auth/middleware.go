package auth

import (
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		header := r.Header.Get("Authorization")

		if header == "" {
			next.ServeHTTP(w, r)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			next.ServeHTTP(w, r)
			return
		}

		token := strings.TrimPrefix(
			header,
			"Bearer ",
		)

		claims, err := ParseToken(token)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := WithClaims(
			r.Context(),
			claims,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
