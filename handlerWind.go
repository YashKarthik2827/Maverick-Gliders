package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlerGetWind(w http.ResponseWriter, r *http.Request) {
	// Extracting the latitude and longitude from the Query
	log.Println(r.URL.Query())
	queryStr := r.URL.Query()
	if queryStr["lat"] == nil || queryStr["lon"] == nil {
		// Respond with error in JSON
		log.Fatal("Missing Coordinates")
	}
	log.Println(queryStr)
	client := http.Client{
		Timeout: time.Duration(2) * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("https://my.meteoblue.com/packages/basic-15min_basic-day_wind-15min_wind-day?apikey=dJK6DxymAIS4Oqcz&lat=%v&lon=%v&format=json", queryStr["lat"][0], queryStr["lon"][0]))
	defer resp.Body.Close()
	if err != nil {
		// Respond with Error
		log.Fatal("Oh NORR Error, could't get ")
	}
	log.Println("API Call made successfully")
	// log.Println(resp.Body)

	err = writeJSONToFile(resp)
	if err != nil {
		// Respond with JSON later
		log.Fatal(err)
	}
}
