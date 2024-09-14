package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func extractData[T Wind | Basic](w http.ResponseWriter, filenm string, params T) error {
	// Reading data frm the file
	jsonData, err := os.ReadFile(filenm)
	if err != nil {
		log.Println("Couldn't read from file")
		return err
	}

	// Extracting what I wanna send to the user
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		log.Println("Couldn't unMarshall")
		return err
	}

	// converting the extracted data to json again
	dat, err := json.Marshal(params)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", params)
		return err
	}

	// Sending the Response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
	return nil
}
