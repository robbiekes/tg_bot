package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

// TODO: make weather for a whole day

func (c *Commander) WeatherNow(inputMessage *tgbotapi.Message) error {
	weather, err := c.weather.GetCurrentWeather()
	if err != nil {
		return errors.Wrap(err, "getting current weather")
	}

	text := "weather for now:\n\n" + weather

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	_, err = c.bot.Send(msg)
	if err != nil {
		return errors.New("error sending message - " + err.Error())
	}

	return nil
}
