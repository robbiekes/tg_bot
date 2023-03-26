package commander

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) error {
	text := "/now - weather for now\n" +
		"/week - weather for a week\n" +
		"/clothes - weather for now with clothes recommendations\n"

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
