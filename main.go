package main

import (
	"log"
	"tgbot3/admin"
	"tgbot3/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("token") // Токен бота t.me/BotFather
	if err != nil {
		log.Println(err)
	}

	bot.Debug = true // Включение отладки чтобы видеть логи в консоли

	log.Printf("Бот t.me/%s запущен", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)

	// Включаем обработчик входящих сообщений
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Отправление стикера администратору от пользователя и отправление стикера пользователю от администратора
		if update.Message.Sticker != nil {
			if update.Message.From.ID != admin.Admin_id {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправлено!")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.StickerMsg(update, bot)
			} else {
				admin.AdminStickerMsg(update, bot)
			}
		}

		// Отправление картинки администратору от пользователя и отправление картинки пользователю от админа
		if update.Message.Photo != nil {
			if update.Message.From.ID != admin.Admin_id {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправлено!")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.PicMsg(update, bot)
			} else {
				admin.AdminPicMsg(update, bot)
			}
		}

		// Отправление видео администратору от пользователя и отправление видео пользователю от администратора
		if update.Message.Video != nil {
			if update.Message.From.ID != admin.Admin_id {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправлено!")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.VideoMsg(update, bot)
			} else {
				admin.AdminVideoMsg(update, bot)
			}
		}

		// Отправление аудио сообщения администратору от пользователя и отправление аудио сообщения пользователю от админа
		if update.Message.Voice != nil {
			if update.Message.From.ID != admin.Admin_id {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправлено!")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.VoiceMsg(update, bot)
			} else {
				admin.AdminVoiceMsg(update, bot)
			}
		}

		// Отправление видео сообщения администратору от пользователя и отправление видео сообщения пользователю от админа
		if update.Message.VideoNote != nil {
			if update.Message.From.ID != admin.Admin_id {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправлено!")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.VideoNoteMsg(update, bot)
			} else {
				admin.AdminVideoNoteMsg(update, bot)
			}
		}

		// Ответы бота на команды
		switch update.Message.Command() {
		case "start":
			if update.Message.From.ID != admin.Admin_id {
				// Запись нового пользователя в базу данных
				database.InsertDb(update, bot)
				// Отправка сообщения с клавиатурой
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Приветствую! Если у вас есть какие либо проблемы, можете написать о них в этот чат и через некоторое время к вам подключится администратор.")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.DefaultMsg(update, bot)
			} else {
				// Запись администратора в базу данных
				database.InsertDb(update, bot)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Бот запущен.")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.DefaultMsg(update, bot)
			}
		case "admin":
			if update.Message.From.ID != admin.Admin_id { // Если команду запрашивает пользователь - отправляем уведомление об этом администратору
				admin.DefaultMsg(update, bot)
				continue
			} else {
				admin.Admin_on(update, bot)
			}
		case "chat_id":
			if update.Message.From.ID != admin.Admin_id { // Если команду запрашивает пользователь - отправляем уведомление об этом администратору
				admin.DefaultMsg(update, bot)
				continue
			} else {
				admin.ChatId(update, bot)
			}
		case "msg":
			if update.Message.From.ID != admin.Admin_id { // Если команду запрашивает пользователь - отправляем уведомление об этом администратору
				admin.DefaultMsg(update, bot)
				continue
			} else {
				admin.AdminMsg(update, bot)
			}
		case "del":
			if update.Message.From.ID != admin.Admin_id { // Если команду запрашивает пользователь - отправляем уведомление об этом администратору
				admin.DefaultMsg(update, bot)
				continue
			} else {
				admin.DelMsg(update, bot)
			}
		case "users":
			if update.Message.From.ID != admin.Admin_id { // Если команду запрашивает пользователь - отправляем уведомление об этом администратору
				admin.DefaultMsg(update, bot)
				continue
			} else {
				database.UsersCount(update, bot)
			}
		case "edit":
			if update.Message.From.ID != admin.Admin_id { // Если команду запрашивает пользователь - отправляем уведомление об этом администратору
				admin.DefaultMsg(update, bot)
				continue
			} else {
				admin.EditMsg(update, bot)
			}
		}
		// Принимаем сообщения от пользователей
		switch update.Message.Text {
		default:
			if update.Message.Sticker != nil || update.Message.Voice != nil || update.Message.Video != nil || update.Message.Photo != nil || update.Message.Command() != "" ||
				update.Message.Text == "/" || update.Message.From.ID == admin.Admin_id || update.Message.VideoNote != nil {
				continue
			} else {
				// Отпралвение сообщения пользователя администратору
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправлено!")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
				admin.DefaultMsg(update, bot)
			}
		}
	}
}
