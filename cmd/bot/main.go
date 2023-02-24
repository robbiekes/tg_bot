package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/robbiekes/tg_bot/internal/controller/commander"
	"github.com/robbiekes/tg_bot/internal/service/product"
	"log"
	"os"
)

func main() {
	godotenv.Load()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	productService := product.NewService()
	commandHandler := commander.NewCommander(bot, productService)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		err = commandHandler.HandleUpdate(update)
		if err != nil {
			log.Fatal(err)
		}
	}
}
