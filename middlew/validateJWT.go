package middlew

import (
	"net/http"

	"github.com/phk13/poc-tw/routers"
)

/* ValidateJWT allows validation of the JWT retrieved in HTTP request.*/
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error in token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
