package commander

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) error {
	text := "you wrote: " + inputMessage.Text

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	_, err := c.bot.Send(msg)
	if err != nil {
		return errors.New("error sending message - " + err.Error())
	}

	return nil
}
