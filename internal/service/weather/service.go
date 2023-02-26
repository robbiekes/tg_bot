package weather

type WeatherService struct {
	WeatherApiClient WeatherClient
}

func NewWeatherService(client WeatherClient) *WeatherService {
	return &WeatherService{WeatherApiClient: client}
}
