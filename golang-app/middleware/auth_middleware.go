package middleware

import (
	"fmt"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if userName, pass, ok := r.BasicAuth(); ok {
			if userName == "user" && pass == "pass" {
				fmt.Println("Authenticated successfully")
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)

	})
}
