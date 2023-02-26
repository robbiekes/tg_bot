package commander

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot      *tgbotapi.BotAPI
	weather  WeatherService
	commands map[string]func(*tgbotapi.Message) error
}
