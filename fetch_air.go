package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func FetchAirData(lat, lon, apiKey string) (*http.Response, error) {
	url := fmt.Sprintf("http://my.meteoblue.com/packages/air-1h_air-day?lat=%s&lon=%s&apikey=%s", lat, lon, apiKey)
	//using the api key hitting the end point taken from the meteoblue docs with the required lat and log and with the api key
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode) // suppose facing any error logging that error
	}
	return resp, nil
}
func handlerGetAir(w http.ResponseWriter, r *http.Request, apiKey string) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")

	resp, err := FetchAirData(lat, lon, apiKey)
	if err != nil {
		http.Error(w, "Failed to fetch air data", http.StatusInternalServerError)
		// over here the error might happen coz of wrong api key lat and lon with random numbers not according to the spcified format
		return
	}
	defer resp.Body.Close() // closing the class using defer already expained how defer works in write json (lifo)
	filePath,err := WriteJSONToFile(resp)
	if err != nil {
		http.Error(w, "Failed to write JSON to file", http.StatusInternalServerError)
		return
	}
	

	// Extracting the data;
	var params Air
	dat, err := extractData(filePath, &params)
	if err != nil {
		http.Error(w, "Failed to extract the air data", http.StatusInternalServerError)
		log.Println("Error extracting air data:", err) //  logging the error
		return
	}
	w.Header().Add("Content-Type", "application/json") 
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(dat)
	if err != nil {
		log.Println("Couldn't write response body:", err) //  logging the error
	}
}
