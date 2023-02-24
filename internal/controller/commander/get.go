package commander

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) error {
	args := inputMessage.CommandArguments()

	//  TODO: вынести отправку сообщения в отдельную ф-ю
	id, err := strconv.Atoi(args)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Invalid command argument :(")

		_, err = c.bot.Send(msg)
		if err != nil {
			return errors.New("error sending message - " + err.Error())
		}

		return nil
	}

	item, err := c.productService.ItemByID(id)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "No product with that ID :(")

		_, err = c.bot.Send(msg)
		if err != nil {
			return errors.New("error sending message - " + err.Error())
		}

		return nil
	}

	text := fmt.Sprintf("Here the product: %s", item.Title)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	_, err = c.bot.Send(msg)
	if err != nil {
		return errors.New("error sending message - " + err.Error())
	}

	return nil
}
