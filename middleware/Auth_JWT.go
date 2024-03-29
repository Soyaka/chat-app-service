package middleware

import (
	utils "main/Utils"
	"net/http"
	"time"
)

func AuthJWT(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		cliams, err := utils.VerifyJWToken(c.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if time.Until(cliams.ExpiresAt.Time) > 30*time.Second {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
