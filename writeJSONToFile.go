package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func writeJSONToFile(resp *http.Response) error {
	// Writing the API Data to a File
	shaHash := sha256.New()
	// Adding the current timestamp to make the encoding unique
	shaHash.Write([]byte(time.Now().String()))
	encodedName := fmt.Sprintf("%x.json", shaHash.Sum(nil))

	// Creating the file, automaticaaly gives Read-Write permissions
	workingDirectory, err := os.Getwd()
	if err != nil {
		return errors.New(fmt.Sprintf("could'nt find working directory, %v", err))
	}
	filePath := workingDirectory + "/jsonFiles/" + encodedName
	newFile, err := os.Create(filePath)
	if err != nil {
		// Respond with err
		return errors.New(fmt.Sprintf("could'nt create file, %v", err))
	}
	// Writing onto file
	n, err := io.Copy(newFile, resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("error writing data to newFile, %v", err))
	}
	log.Println("Wrote these many characters:", n)
	return nil
}
