package open_weather_map

import (
	"fmt"
	"io"
	"net/http"
)

const (
	url  = "https://api.openweathermap.org/data/2.5/weather"
	host = "open-weather-map27.p.rapidapi.com"
)

type WeatherAPI struct {
	client *http.Client
	apikey string
}

func NewWeatherAPI(client *http.Client, apikey string) *WeatherAPI {
	return &WeatherAPI{client: client, apikey: apikey}
}

func (w *WeatherAPI) GetCurrectWeather() {

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", w.apikey)
	req.Header.Add("X-RapidAPI-Host", host)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
