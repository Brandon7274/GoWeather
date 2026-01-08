package httphandler

import (
	"encoding/json"
	"net/http"
)

type WeatherResponse struct {
	Properties struct {
		Periods []struct {
			Number      int8   `json:"number"`
			Name        string `json:"name"`
			Temperature int8   `json:"temperature"`
			Icon        string `json:"icon"`
		} `json:"periods"`
	} `json:"properties"`
}

func GetWeather(url string, limit int) (WeatherResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return WeatherResponse{}, err
	}
	defer resp.Body.Close()

	var a WeatherResponse

	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		return WeatherResponse{}, err
	}

	if len(a.Properties.Periods) > limit {
		a.Properties.Periods = a.Properties.Periods[:limit]
	}

	return a, nil

}
