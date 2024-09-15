package main

import (
	"encoding/json"
	"log"
	"os"
)

func extractData[T Wind | Basic](filePath string, params T) ([]byte, error) {
	// Reading data frm the file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Couldn't read from file")
		return []byte{}, err
	}

	// Extracting what I wanna send to the user
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		log.Println("Couldn't unMarshall")
		return []byte{}, err
	}

	// converting the extracted data to json again
	dat, err := json.Marshal(params)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", params)
		return []byte{}, err
	}
	return dat, nil
}
