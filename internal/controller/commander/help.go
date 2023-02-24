package commander

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) error {
	text := "/today - weather for today\n" +
		"/week - weather for a week\n" +
		"/clothes - weather for today with clothes recommendations\n" +
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
