package main

import (
	"Telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	// Создание клиента для бота Telegram
	bot, err := tgbotapi.NewBotAPI("TELEGRAM_BOT_API_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}