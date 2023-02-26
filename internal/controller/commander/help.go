package commander

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) error {
	text := "/today - open_weather_map for today\n" +
		"/week - open_weather_map for a week\n" +
		"/clothes - open_weather_map for today with clothes recommendations\n" +
		"/list - list all products\n" +
		"/get - get a product by id (example: /get 1)"

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Ping", "Pong"),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		return errors.New("error sending message - " + err.Error())
	}

	return nil
}
