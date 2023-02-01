package routers

import (
	"encoding/json"
	"net/http"

	"github.com/satico0518/twitterGo/bd"
	"github.com/satico0518/twitterGo/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Body error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email not provided: ", http.StatusBadRequest)
		return
	}

	_, userExists, _ := bd.UserExists(t.Email)
	if userExists {
		http.Error(w, "User already registered", http.StatusBadRequest)
		return
	}

	_, status, err := bd.RegisterUser(t)
	if err != nil || !status {
		http.Error(w, "Error trying to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
