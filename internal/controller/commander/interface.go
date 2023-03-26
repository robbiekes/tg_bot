package commander

type WeatherService interface {
	GetCurrentWeather() (string, error)
}
