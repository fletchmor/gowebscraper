package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	c := chi.NewRouter()
	godotenv.Load()
	port := os.Getenv("PORT")
	c.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1 := chi.NewRouter()
	v1.Get("/err", handlerError)
	c.Mount("/v1", v1)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: c,
	}

	server.ListenAndServe()
}
