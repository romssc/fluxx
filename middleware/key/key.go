package key

import (
	"net/http"
)

func New(header string, key string, errs string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(header)
			if token != key {
				http.Error(w, errs, http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
