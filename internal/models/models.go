package models

type WeatherResponse struct {
	Location struct {
		Name      string  `json:"name"`
		Region    string  `json:"region"`
		Country   string  `json:"country"`
		Lat       float64 `json:"lat"`
		Lon       float64 `json:"lon"`
		TzID      string  `json:"tz_id"`
		Localtime string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph      float64 `json:"wind_mph"`
		WindKph      float64 `json:"wind_kph"`
		WindDegree   int     `json:"wind_degree"`
		WindDir      string  `json:"wind_dir"`
		PressureMb   float64 `json:"pressure_mb"`
		PressureIn   float64 `json:"pressure_in"`
		PrecipMm     float64 `json:"precip_mm"`
		PrecipIn     float64 `json:"precip_in"`
		Humidity     int     `json:"humidity"`
		Cloud        int     `json:"cloud"`
		FeelsLikeC   float64 `json:"feelslike_c"`
		FeelsLikeF   float64 `json:"feelslike_f"`
		WindChillC   float64 `json:"windchill_c"`
		WindChillF   float64 `json:"windchill_f"`
		HeatIndexC   float64 `json:"heatindex_c"`
		HeatIndexF   float64 `json:"heatindex_f"`
		DewPointC    float64 `json:"dewpoint_c"`
		DewPointF    float64 `json:"dewpoint_f"`
		VisibilityKm float64 `json:"vis_km"`
		VisibilityMi float64 `json:"vis_miles"`
		UV           float64 `json:"uv"`
		GustMph      float64 `json:"gust_mph"`
		GustKph      float64 `json:"gust_kph"`
	} `json:"current"`
}


type WeatherResponse7Days struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Forecast struct {
		Forecastday []struct {
			Date       string `json:"date"`
			Day        struct {
				TempMaxC float64 `json:"maxtemp_c"`
				TempMinC float64 `json:"mintemp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type WeatherResponse30Days struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Forecast struct {
		Forecastday []struct {
			Date       string `json:"date"`
			Day        struct {
				TempMaxC float64 `json:"maxtemp_c"`
				TempMinC float64 `json:"mintemp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}