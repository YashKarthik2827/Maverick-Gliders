package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid" // this is the lib imported for the uuid 
)

func WriteJSONToFile(resp *http.Response) error {
    
    uniqueID := uuid.New().String()// Generating a uuid for the file name 
    fileName := fmt.Sprintf("%s.json", uniqueID)

    workingDirectory, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("couldn't find working directory: %v", err)
    }
    directoryPath := workingDirectory + "/jsonFiles" 
    filePath := directoryPath + "/" + fileName


    err = os.MkdirAll(directoryPath, os.ModePerm)// over here we are creating a dir if it does not exits
    if err != nil {
        return fmt.Errorf("couldn't create directory: %v", err)
    }

    newFile, err := os.Create(filePath)// creating the file over here 
    if err != nil {
        return fmt.Errorf("couldn't create file: %v", err) // if not able to create with the file path logging the error faced
    }
    defer newFile.Close() /*
		
		closing the file using defer 
		basically after the program reaches the end this will get excuted if many defer is there lifo will be followed 
		*/

    n, err := io.Copy(newFile, resp.Body)
    if err != nil {
        return fmt.Errorf("error writing data to newFile: %v", err) // if not able to write the data into the file loggint the error faced
    }
    fmt.Printf("Wrote %d characters to file: %s\n", n, filePath) // n over here is the number of char and file path is file path
    return nil
}
