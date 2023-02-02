package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/satico0518/twitterGo/middlewares"
	"github.com/satico0518/twitterGo/routers"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods(http.MethodPost)
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods(http.MethodPost)
	router.HandleFunc("/profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.Profile))).Methods(http.MethodGet)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
