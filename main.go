package main

import (
	"log"

	"github.com/phk13/poc-tw/db"
	"github.com/phk13/poc-tw/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("No connection to DB")
		return
	}
	handlers.Handlers()
}
