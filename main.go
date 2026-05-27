package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
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
}
