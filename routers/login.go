package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/satico0518/twitterGo/bd"
	"github.com/satico0518/twitterGo/jwt"
	"github.com/satico0518/twitterGo/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil || len(t.Email) == 0 {
		http.Error(w, "Wrong Credentials", http.StatusBadRequest)
		return
	}
	user, exists := bd.TryLogin(t.Email, t.Password)
	if !exists {
		http.Error(w, "Wrong Credentials", http.StatusUnauthorized)
		return
	}

	jwtKey, err := jwt.GetToken(user)
	if err != nil {
		http.Error(w, "Error gettting token"+err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
