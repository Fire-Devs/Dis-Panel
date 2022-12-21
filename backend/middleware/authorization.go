package middleware

import (
	"github.com/Dispanel/config"
	"github.com/Dispanel/utils"
	"net/http"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != config.LoadConfig().Authorization {
			utils.WarningHandling("MIDDLEWARE", "Unauthorized access to "+r.Host)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
