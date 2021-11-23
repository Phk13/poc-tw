package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/phk13/poc-tw/middlew"
	"github.com/phk13/poc-tw/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers sets ports, handler and begins listening */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
