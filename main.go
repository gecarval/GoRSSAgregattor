package main

import (
	"fmt"
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
		log.Fatal("ERROR: no PORT environment exported.")
	} else {
		fmt.Println("Listening from PORT:", port+".")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<body><h1>Welcome</h1></body>"))
	})
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
