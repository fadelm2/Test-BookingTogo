package main

import (
	"fmt"
	"github.com/bookingtogo/internal/config"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator(viperConfig)
	app := config.NewGorillaRouter(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		App:      app,
		Log:      log,
		Validate: validate,
		Config:   viperConfig,
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
			"http://localhost:8080",
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
			"Authorization",
		},
		AllowCredentials: true,
	})

	// WRAP HERE (Benar)
	handler := c.Handler(app)

	webPort := viperConfig.GetString("web.port")
	fmt.Println("Listening on :", webPort)
	err := http.ListenAndServe(":"+webPort, handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
