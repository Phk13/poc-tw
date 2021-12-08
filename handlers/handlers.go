package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phk13/poc-tw/config"
	"github.com/phk13/poc-tw/middlew"
	"github.com/phk13/poc-tw/routers"
	"github.com/rs/cors"
)

/* Handlers sets port, handlers and begins listening */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/viewProfile", middlew.CheckDB(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")

	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlew.CheckDB(middlew.ValidateJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.CheckDB(middlew.ValidateJWT(routers.RemoveTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/createRelationship", middlew.CheckDB(middlew.ValidateJWT(routers.CreateRelationship))).Methods("POST")
	router.HandleFunc("/deleteRelationship", middlew.CheckDB(middlew.ValidateJWT(routers.RemoveRelationship))).Methods("DELETE")
	router.HandleFunc("/checkRelationship", middlew.CheckDB(middlew.ValidateJWT(routers.CheckRelationship))).Methods("GET")

	router.HandleFunc("/listUsers", middlew.CheckDB(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/getFollowerTweets", middlew.CheckDB(middlew.ValidateJWT(routers.GetFollowerTweets))).Methods("GET")

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+config.AppCfg.ServerPort, handler))
}
