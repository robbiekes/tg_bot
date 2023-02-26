package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) WeatherToday(inputMessage *tgbotapi.Message) error {
	c.weather.GetCurrentWeather()

	return nil
}
