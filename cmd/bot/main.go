package main

import (
	"log"
	"tgbot3/pkg/admin"
	"tgbot3/pkg/config"
	"tgbot3/pkg/handler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.GetConfig("config/config.yml")
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Println(err)
	}

	// Debug mode
	bot.Debug = true

	// Создаем администратора
	admin := admin.NewAdmin(cfg.AdminID)

	// Создаем обработчик
	handler := handler.NewHandler(admin)

	// Включаем обработчик входящих сообщений
	err = handler.Start(bot)
	if err != nil {
		panic(err)
	}
}
