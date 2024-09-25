package main

import (
	"encoding/json"
	"log"
	"os"
)

func extractData[T Basic | Wind | Air | Cloud ](filePath string, params *T) ([]byte, error) {
	// Reading data from the file
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Couldn't read from file")
		return []byte{}, err
	}

	// Extracting what I want to send to the user
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		log.Println("Couldn't unmarshal")
		return []byte{}, err
	}

	// Converting the extracted data to JSON again
	dat, err := json.Marshal(params)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", params)
		return []byte{}, err
	}
	return dat, nil
}

