package main

// I wanna pu the metadata and the units for every response
type Base struct {
	Metadata struct {
		TimezoneAbbrevation string  `json:"timezone_abbrevation"`
		Latitude            float32 `json:"latitude"`
		ModelrunUTC         string  `json:"modelrun_utc"`
		Longitude           float32 `json:"longitude"`
		GenerationTimeMs    float32 `json:"generation_time_ms"`
	} `json:"metadata"`
	Units struct {
		Density          string `json:"density"`
		Precipitation    string `json:"precipitation"`
		Windspeed        string `json:"windspeed"`
		RelativeHumidity string `json:"relativehumidity"`
		Time             string `json:"time"`
		Temperature      string `json:"temperature"`
		Winddirection    string `json:"winddirection"`
		Pressure         string `json:"pressure"`
	} `json:"units"`
}

// Endpoint specific response
type Basic struct {
	Base
	DataDay struct {
		Time                 []string  `json:"time"`
		SealevelPressureMean []int     `json:"sealevelpressure_mean"`
		RelativeHumidityMean []float32 `json:"relativehumidity_mean"`
	} `json:"data_day"`
}

type Wind struct {
	Base
	DataDay struct {
		Time          []string  `json:"time"`
		AirdensityMax []float32 `json:"airdensity_max"`
		WindspeedMean []float32 `json:"windspeed_mean"`
		Winddirection []int     `json:"winddirection"`
	} `json:"data_day"`
}
