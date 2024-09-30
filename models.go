package main

// I wanna put the metadata and the units for every response
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

//adding wind and air

type Air struct {
	Base
	DataDay struct {
		Time              []string  `json:"time"`
		AirTemperatureMax []float64 `json:"wetbulbglobetemperature_max"`
		AirTemperatureMin []float64 `json:"wetbulbglobetemperature_min"`
		AirPressure       []float64 `json:"convectivecloudbase_pressure"`
	} `json:"data_day"`
}

// Cloud response structure
type Cloud struct {
	Base
	DataDay struct {
		Time            []string  `json:"time"`
		CloudCoverMean  []float64 `json:"totalcloudcover_mean"`
		CloudCoverMax   []float64 `json:"totalcloudcover_max"`
		CloudCoverMin   []float64 `json:"totalcloudcover_min"`
		CloudHeightMean []float64 `json:"highclouds_mean"`
	} `json:"data_day"`
}
