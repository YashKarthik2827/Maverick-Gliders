package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Helo world")


	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)// checking wheather we are able to access the env file or not 
	}
	apiKey := os.Getenv("API_KEY")// getting the api key if not there logging the message
	if apiKey == "" {
		log.Fatalf("API_KEY environment variable not set")
	}

	
	mux := http.NewServeMux() // setting up the routes and starting 
	mux.HandleFunc("/air", func(w http.ResponseWriter, r *http.Request) {
		handlerGetAir(w, r, apiKey)
	})
	log.Println("Server starting at port 8080")
	http.ListenAndServe(":8080", mux)
}
