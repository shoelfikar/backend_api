package middleware

import (
	"net/http"

	"github.com/rs/cors"
	// "github.com/shoelfikar/golang_backend_api/helper"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedHeaders: []string{"Content-Type", "Authorization", "X-Requested-With"},
			AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
			MaxAge:         0,
		})
		
		// You can add custom logic here if needed before applying CORS.
		// helper.DefaultLoggingInfo("CORS middleware is enabled")
		// Apply CORS by calling the CORS handler's ServeHTTP method.
		c.ServeHTTP(w, r, func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	})

}
