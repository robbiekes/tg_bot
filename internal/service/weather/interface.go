package weather

type WeatherClient interface {
	GetCurrentWeather() (string, error)
}
