package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (apiCfg apiConifg) handlerGetBasic(w http.ResponseWriter, r *http.Request) {
	// Extracting the latitude and longitude from the query string
	log.Println(r.URL.Query())
	queryStr := r.URL.Query()
	if queryStr["lat"] == nil || queryStr["lon"] == nil {
		// Error from user do sending 4XX code
		log.Println("Missing Coordinates")
		w.WriteHeader(400)
		return
	}
	client := http.Client{
		Timeout: time.Duration(2) * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("%s&lat=%v&lon=%v&format=json", apiCfg.baseURL, queryStr["lat"][0], queryStr["lon"][0]))
	if err != nil {
		// Respond with Error
		log.Println("Oh NORR Error, could't get response", err)
		w.WriteHeader(500)
		return
	}
	defer resp.Body.Close()
	log.Println("API Call made successfully")
	// log.Println(resp.Body)

	// Writing in file
	// Any error from these 2 functions is server side so sending 5XX status code
	filePath, err := WriteJSONToFile(resp)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	var params Basic
	params, err = parseJSON(filePath, params)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}

	// CRUD
	insertId, err := apiCfg.insertRecord(params)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error storing data in db, %v", err)
		return
	}

	resultCursor := apiCfg.readRecord(insertId)
	var result Basic
	err = resultCursor.Decode(&result)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Error marshalling result record %v", err)
	}
	log.Println(result)

	dat, err := createJSON(params)
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	// Sending the Response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write(dat)
	if err != nil {
		log.Println("Couldn't write to response body", err)
	}

}
