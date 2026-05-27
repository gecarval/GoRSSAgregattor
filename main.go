package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("ERROR: environment file does not exist in root directory.")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" || port == "" {
		log.Fatal("ERROR: HOST or PORT environment value not exported.")
	}
	log.Println("Listening from ", host+":"+port+".")

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: servive(),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func servive() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<head></head><body><h1>Welcome</h1></body>"))
	})

	return r
}
