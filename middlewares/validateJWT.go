package middlewares

import (
	"net/http"

	"github.com/satico0518/twitterGo/routers"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token Error!", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
