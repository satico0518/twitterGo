package routers

import (
	"encoding/json"
	"net/http"

	"github.com/satico0518/twitterGo/bd"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 2 {
		http.Error(w, "Invalid Id param", http.StatusBadRequest)
		return
	}

	user, err := bd.Profile(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
