package commander

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewCommander(bot *tgbotapi.BotAPI, weather WeatherService) *Commander {
	c := &Commander{
		bot:      bot,
		weather:  weather,
		commands: make(map[string]func(*tgbotapi.Message) error),
	}

	c.initCommands()

	return c
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) error {
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackData())

		_, err := c.bot.Send(msg)
		if err != nil {
			return errors.New("error doing default behaviour - " + err.Error())
		}

		return nil
	}

	if update.Message != nil { // If we got a message

		commandName := update.Message.Command()

		command, ok := c.commands[commandName]
		if ok {
			err := command(update.Message)
			if err != nil {
				return errors.New("error handling command - " + err.Error())
			}
		} else {
			err := c.Default(update.Message)
			if err != nil {
				return errors.New("error doing default behaviour - " + err.Error())
			}
		}
	}

	return nil
}

func (c *Commander) initCommands() {
	c.commands["help"] = c.Help
	c.commands["today"] = c.WeatherToday
}
