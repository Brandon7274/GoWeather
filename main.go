package main

import (
	"encoding/json"
	"fmt"
	"log"

	httphandler "github.com/Brandon7274/GoWeather/Handler"
)

const userAgent = "terminal-weather-app (brandon@example.com)"

func main() {
	fmt.Println("ITS ALIVE!")

	url := "https://api.weather.gov/gridpoints/LWX/33,35/forecast"

	data, _ := httphandler.GetWeather(url, 3)

	fmt.Println("--------------------------------------")
	fmt.Println("          NYC WEATHER REPORT          ")
	fmt.Println("--------------------------------------")

	marshaled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("marshaling error: %s", err)
	}

	fmt.Println(string(marshaled))

}

//"https://api.weather.gov/gridpoints/LWX/40,74/forecast"
