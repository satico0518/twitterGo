package middlewares

import (
	"net/http"

	"github.com/satico0518/twitterGo/bd"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.CheckConnetion() {
			http.Error(w, "DB Connetion Lost", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
