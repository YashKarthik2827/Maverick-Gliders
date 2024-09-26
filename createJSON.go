package main

import (
	"encoding/json"
	"log"
)

func createJSON[T Wind | Basic](params T) ([]byte, error) {
	// converting the data to json
	dat, err := json.Marshal(params)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", params)
		return []byte{}, err
	}
	return dat, nil
}
