package middlew

import (
	"net/http"

	"github.com/phk13/poc-tw/db"
)

/* CheckDB is a middleware to check current DB status and send 500 if it's down.*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "DB connection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
