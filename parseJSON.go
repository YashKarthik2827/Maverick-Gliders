package main

import (
	"encoding/json"
	"log"
	"os"
)

// This function will parse the JSON data into our required data model
func parseJSON[T Wind | Basic](filePath string, params T) (T, error) {
	// Reading data frm the file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Couldn't read from file")
		return params, err
	}

	// Extracting what I wanna send to the user
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		log.Println("Couldn't unMarshall")
		return params, err
	}
	return params, nil
}
