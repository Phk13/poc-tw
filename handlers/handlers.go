package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/phk13/poc-tw/middlew"
	"github.com/phk13/poc-tw/routers"
	"github.com/rs/cors"
)

/* Handlers sets ports, handler and begins listening */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readtweets", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
