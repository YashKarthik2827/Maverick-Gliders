package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func FetchCloudData(lat, lon, apiKey string) (*http.Response, error) {
	url := fmt.Sprintf("http://my.meteoblue.com/packages/clouds-1h_clouds-day?lat=%s&lon=%s&apikey=%s", lat, lon, apiKey)
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}
	return resp, nil
}
func handlerGetCloud(w http.ResponseWriter, r *http.Request, apiKey string) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	res, err := FetchCloudData(lat, lon, apiKey)

	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()
	filePath, err := WriteJSONToFile(res)
	if err != nil {
		http.Error(w, "Failed to write JSON", http.StatusInternalServerError)
		return
	}
	// extracting the cloud data

	var params Cloud
	dat, err := extractData(filePath, &params)
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
