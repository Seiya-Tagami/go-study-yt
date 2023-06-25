// @see https://youtu.be/zPYjfgxYO7k

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				ChanceOfRain float64 `json:"chance_of_rain"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := "Japan"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	
	key := loadEnv()

	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + key + "&q=" + q + "&aqi=no&days=1&api=no&alerts=no")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}
	

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}
		
		message := fmt.Sprintf("%s - %.0fC, %.0f, %s\n",
		date.Format("15:04"),
		hour.TempC,
		hour.ChanceOfRain,
		hour.Condition.Text,
	)

	if hour.ChanceOfRain < 40 {
		fmt.Print(message)
	} else {
		color.Red(message)
	}
	}
}

func loadEnv() string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Cannot read .env: %v", err)
	}

	key := os.Getenv("API_ENDPOINT_KEY")

	return key
}
