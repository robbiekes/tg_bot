package weather

import "github.com/pkg/errors"

type WeatherService struct {
	WeatherApiClient WeatherClient
}

func NewWeatherService(client WeatherClient) *WeatherService {
	return &WeatherService{WeatherApiClient: client}
}

func (s *WeatherService) GetCurrentWeather() (string, error) {
	weather, err := s.WeatherApiClient.GetCurrentWeather()
	if err != nil {
		return "", errors.Wrap(err, "getting current weather")
	}

	return weather, nil
}
