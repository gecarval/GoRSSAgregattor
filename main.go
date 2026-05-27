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

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("ERROR: PORT environment value not exported.")
	} else {
		log.Println("Listening on PORT:", port+".")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<head></head><body><h1>Welcome</h1></body>"))
	})
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
