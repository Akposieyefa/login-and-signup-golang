package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/akposieyefa/login-and-signup/models"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Missing auth token",
				"success": false,
			})
			return
		}
		tk := &models.Claims{}

		// _, err := jwt.ParseWithClaims(header, tk, func(token jwt.Claims) (interface{}, error) {
		// 	return []byte("secret"), nil
		// })

		// if err != nil {
		// 	w.WriteHeader(http.StatusForbidden)
		// 	json.NewEncoder(w).Encode(err)
		// 	return
		// }

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
