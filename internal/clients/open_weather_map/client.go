package open_weather_map

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/robbiekes/tg_bot/internal/service/weather"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	url = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=en"
)

type WeatherClient struct {
	client *http.Client
	apikey string
}

func NewWeatherAPI(client *http.Client, apikey string) *WeatherClient {
	return &WeatherClient{client: client, apikey: apikey}
}

func (w *WeatherClient) currentWeatherRequest() (*weather.CurrentWeatherResponse, error) {
	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}

	wURL := fmt.Sprintf(url, "Moscow", w.apikey)

	response, err := httpClient.Get(wURL)
	if err != nil {
		return nil, fmt.Errorf("making get request to api - %s\nstatus code - %d\n", err.Error(), response.StatusCode)
	}

	var output weather.CurrentWeatherResponse

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("reading response body - %s\nstatus code - %d\n", err.Error(), response.StatusCode)
	}

	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshall text - %s\nstatus code - %d\n", err.Error(), response.StatusCode)
	}

	return &output, nil
}

func (w *WeatherClient) GetCurrentWeather() (string, error) {
	response, err := w.currentWeatherRequest()
	if err != nil {
		return "", errors.Wrap(err, "getting current weather request")
	}

	if len(response.Weather) == 0 {
		return "", errors.New("no info about weather")
	}

	temp := fmt.Sprintf("%v °C", response.Main.Temp)
	feelsLike := fmt.Sprintf("%v °C", response.Main.FeelsLike)

	if response.Main.Temp > 0 {
		temp = "+" + temp
		feelsLike = "+" + feelsLike
	} else {
		temp = "-" + temp
		feelsLike = "-" + feelsLike
	}

	description := response.Weather[0].Description
	wind := response.Wind.Speed

	out := fmt.Sprintf("it's %s outside\nfeels like %s\n\nit's %s\nand wind velocity is %v m/s\n", temp, feelsLike, description, wind)

	return out, nil
}
