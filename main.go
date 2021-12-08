package main

import (
	"log"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/handlers"
	"github.com/spf13/viper"
)

func main() {
	log.Println(viper.GetViper().GetString("JWTSECRET"))
	if !db.CheckConnection() {
		log.Fatal("No connection to DB")
		return
	}
	handlers.Handlers()
}
