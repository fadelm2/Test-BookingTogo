package config

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"net/http"
)

func NewGorillaRouter(config *viper.Viper) *mux.Router {
	r := mux.NewRouter()

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost:8080",
			"https://test.fadelweb.site",
			"https://testapi.fadelweb.site",
			"http://localhost:3000",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Accept-Encoding",
			"Authorization",
			"X-CSRF-Token",
			"X-Requested-With",
		},
		AllowCredentials: true,
	})

	// Error handling middleware

	// Wrap router with CORS
	return c.Handler(r).(*mux.Router)
}
