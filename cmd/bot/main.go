package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/robbiekes/tg_bot/internal/clients/open_weather_map"
	"github.com/robbiekes/tg_bot/internal/controller/commander"
	"github.com/robbiekes/tg_bot/internal/service/weather"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	// API init
	apikey := os.Getenv("OWM_API_KEY")
	wapi := open_weather_map.NewWeatherAPI(http.DefaultClient, apikey)

	// bot init
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	weatherService := weather.NewWeatherService(wapi)
	commandHandler := commander.NewCommander(bot, weatherService)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		err = commandHandler.HandleUpdate(update)
		if err != nil {
			log.Fatal(err)
		}
	}
}
