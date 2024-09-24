// The WGS84 coordinates for Amrita Vishwa Vidyapeetham in Coimbatore, India, are approximately:
// lat: 11.0587° N
// lon: 76.9434° E
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Will make all the handler functions as methods on this,
// so i don't manually have to pass the client to every function if they update or delete
type apiConifg struct {
	baseURL  string
	dbClient *mongo.Client
}

func main() {
	fmt.Println("Helo world")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err) // checking wheather we are able to access the env file or not
	}
	apiKey := os.Getenv("API_KEY") // getting the api key if not there logging the message
	if apiKey == "" {
		log.Fatalf("API_KEY environment variable not set")
	}

	urlString := os.Getenv("BaseURL")
	if urlString == "" {
		log.Fatalf("BaseURL environment variable not set")
	}

	dbURI := os.Getenv("dbURI")
	if dbURI == "" {
		log.Fatalf("dbURI environment variable not set")
	}

	// Copy pasted from mongoDB website- setting up the db connection
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Counld'nt create a client to connect to the db")
	}
	// To disconnect the client from the db when the main ends
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Couldn't disconnect the db connection ")
		}
	}()
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		log.Fatalf("Flopped the connection to the db loser")
	}
	// log.Println(result)
	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	mux := http.NewServeMux() // setting up the routes and starting
	mux.HandleFunc("/air", func(w http.ResponseWriter, r *http.Request) {
		handlerGetAir(w, r, apiKey)
	})
	// [METHOD] /path, the space b/w method and / is important
	// Creating Path for Basic Package and Wind Package
	apiCfg := apiConifg{
		baseURL:  urlString,
		dbClient: client,
	}
	mux.HandleFunc("GET /basic", apiCfg.handlerGetBasic)
	mux.HandleFunc("GET /wind", apiCfg.handlerGetWind)

	// If u don't pass in your ServeMux, it will use the default ServeMux
	log.Println("Server starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
