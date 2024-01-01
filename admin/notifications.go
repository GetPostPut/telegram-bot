package admin

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Уведомление о запуске бота пользователем
func DefaultMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if update.Message.From.ID != Admin_id {
		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Пользователь %s запустил бота.\n\n", userName(update)))
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Пользователь %s отправил сообщение:\n\n%s", userName(update), update.Message.Text))
			bot.Send(msg)
		}
	}
}

// Отправление видео от пользователя администратору
func VideoMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Пользователь %s отправил видео:", userName(update)))
	bot.Send(msg)
	// Определение полученного видео
	video := (*update.Message.Video)
	fileConfig := tgbotapi.FileConfig{FileID: video.FileID}
	// Создание объекта tgbotapi.videoConfig для отправление видео
	videoConfig := tgbotapi.NewVideoShare(Admin_id, fileConfig.FileID)
	// Отправление видео администратору
	_, err := bot.Send(videoConfig)
	if err != nil {
		log.Println("Error sending video:", err)
	}
}

// Отправление голосового сообщения от пользователя администратору
func VoiceMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Пользователь %s отправил голосовое сообщение:", userName(update)))
	bot.Send(msg)
	// Определение полученного голосового сообщения
	voice := (*update.Message.Voice)
	fileConfig := tgbotapi.FileConfig{FileID: voice.FileID}
	// Создание объекта tgbotapi.voiceConfig для отправки голосового сообщения
	voiceConfig := tgbotapi.NewVoiceShare(Admin_id, fileConfig.FileID)
	// Отправление голосового сообщения администратору
	_, err := bot.Send(voiceConfig)
	if err != nil {
		log.Println("Error sending video note:", err)
	}
}

// Уведомление об отправленном пользователем видео сообщении
func VideoNoteMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Пользователь %s отправил видео сообщение:", userName(update)))
	bot.Send(msg)
	// Определение полученного видео сообщения
	videoNote := (*update.Message.VideoNote)
	fileConfig := tgbotapi.FileConfig{FileID: videoNote.FileID}
	// Создание объекта tgbotapi.videoNoteConfig для отправки видео сообщения
	videoNoteConfig := tgbotapi.NewVideoNoteShare(Admin_id, 0, fileConfig.FileID)
	// Отправление видео сообщения администратору
	_, err := bot.Send(videoNoteConfig)
	if err != nil {
		log.Println("Error sending video note:", err)
	}
}

// Уведомление об отправленном пользователем стикера
func StickerMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Пользователь %s отправил стикер:", userName(update)))
	bot.Send(msg)
	// Определение полученного стикера
	sticker := (*update.Message.Sticker)
	fileConfig := tgbotapi.FileConfig{FileID: sticker.FileID}
	// Создание объекта tgbotapi.stickerConfig для отправление стикера
	stickerConfig := tgbotapi.NewStickerShare(Admin_id, fileConfig.FileID)
	// Отправление стикера администратору
	_, err := bot.Send(stickerConfig)
	if err != nil {
		log.Println("Error sending sticker:", err)
	}
}

// Отправление сообщения администратору о том, что пользователь отправил картинку
func PicMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Пользователь %s отправил картинку:", userName(update)))
	bot.Send(msg)
	// Определение последней полученной фотографии
	photo := (*update.Message.Photo)[len(*update.Message.Photo)-1]
	fileConfig := tgbotapi.FileConfig{FileID: photo.FileID}
	// Создание объекта tgbotapi.PhotoConfig для отправки фотографии
	photoConfig := tgbotapi.NewPhotoShare(Admin_id, fileConfig.FileID)
	// Отправление фотографии администратору
	_, err := bot.Send(photoConfig)
	if err != nil {
		log.Println("Error sending photo:", err)
	}
}
