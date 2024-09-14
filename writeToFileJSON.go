package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func writeToFileJSON(resp *http.Response) (string, error) {
	// Writing the API Data to a File
	// Creating unique file nm using uuid
	encodedName := fmt.Sprintf("%v.json", uuid.New())

	// Creating the file, automaticaaly gives Read-Write permissions
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", errors.New(fmt.Sprintf("could'nt find working directory, %v", err))
	}
	filePath := workingDirectory + "/jsonFiles/" + encodedName
	newFile, err := os.Create(filePath)
	if err != nil {
		// Respond with err
		return "", errors.New(fmt.Sprintf("could'nt create file, %v", err))
	}
	// Writing onto file
	n, err := io.Copy(newFile, resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error writing data to newFile, %v", err))
	}
	log.Println("Wrote these many characters:", n)
	return filePath, nil
}
