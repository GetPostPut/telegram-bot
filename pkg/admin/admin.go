package admin

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Admin struct {
	adminId int
	p_id int64
	BlockUserId int64
	block_list []int
}

func NewAdmin(adminId int) *Admin {
	return &Admin {
		adminId: adminId,
	}
}

// AdminHelp - функция для вывода списка команд
func (a Admin) AdminHelp(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	//...
	return nil
}

// userName - функция для получения имени пользователя
func (a Admin) UserName() error {
	//...
	return nil
}

// ChatId - функция для получения ID чата
func (a Admin) ChatId(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	//...
	return nil
}

// .
// .
// .

// // Отправление сообщения пользователю через бота
// func AdminMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	if update.Message.Chat.ID == Admin_id {
// 		messagep := strings.TrimSpace(update.Message.CommandArguments()) // Сообщение, которое будет отправлено пользователю /message [текст который увидит пользователь], с помощью strings.TrimSpace убираем лишние пробелы по бокам
// 		if messagep == "" {
// 			msg := tgbotapi.NewMessage(Admin_id, "Для отправки сообщения другому пользователю напиши /msg [сообщение пользователю]")
// 			bot.Send(msg)
// 			msg = tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Для изменения id получателя напиши /chat_id [ID получателя]\nТекущий ID получателя %d.", P_id))
// 			bot.Send(msg)
// 		} else {
// 			msg := tgbotapi.NewMessage(P_id, strings.TrimSpace(update.Message.CommandArguments()))
// 			sentMessage, err := bot.Send(msg)
// 			if err != nil {
// 				msg = tgbotapi.NewMessage(Admin_id, "Ошибка. Сообщение не было отправлено. Попробуй изменить ID получателя командой /chat_id [ID получателя]")
// 				bot.Send(msg)
// 			} else {
// 				messageID := sentMessage.MessageID
// 				msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Сообщение [MessageID: %d] пользователю с ID %d было успешно отправлено!", messageID, P_id))
// 				bot.Send(msg)
// 			}
// 		}
// 	}
// }

// // Отправление картинки пользователю через бота
// func AdminPicMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	// Определение последней полученной фотографии
// 	photo := (*update.Message.Photo)[len(*update.Message.Photo)-1]
// 	fileConfig := tgbotapi.FileConfig{FileID: photo.FileID}
// 	// Создание объекта tgbotapi.PhotoConfig для отправки фотографии
// 	photoConfig := tgbotapi.NewPhotoShare(P_id, fileConfig.FileID)
// 	// Отправка фотографии пользователю
// 	sentMessage, err := bot.Send(photoConfig)
// 	if err != nil {
// 		log.Println("Error sending photo:", err)
// 	}
// 	messageID := sentMessage.MessageID
// 	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Картинка [MessageID: %d] пользователю с ID %d была успешно отправлена!", messageID, P_id))
// 	bot.Send(msg)
// }

// // Отправление видео пользователю через бота
// func AdminVideoMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	// Определение последней полученной фотографии
// 	video := (*update.Message.Video)
// 	fileConfig := tgbotapi.FileConfig{FileID: video.FileID}
// 	// Создание объекта tgbotapi.VideoShare для отправки фотографии
// 	videoConfig := tgbotapi.NewVideoShare(P_id, fileConfig.FileID)
// 	// Отправка фотографии пользователю
// 	sentMessage, err := bot.Send(videoConfig)
// 	if err != nil {
// 		log.Println("Error sending video:", err)
// 	}
// 	messageID := sentMessage.MessageID
// 	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Видео [MessageID: %d] пользователю с ID %d было успешно отправлено!", messageID, P_id))
// 	bot.Send(msg)
// }

// // Отправление аудио сообщения пользователю через бота
// func AdminVoiceMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	// Определение последней полученной аудио сообщения
// 	voice := (*update.Message.Voice)
// 	fileConfig := tgbotapi.FileConfig{FileID: voice.FileID}
// 	// Создание объекта tgbotapi.VideoShare для отправки аудио сообщения
// 	voiceConfig := tgbotapi.NewVoiceShare(P_id, fileConfig.FileID)
// 	// Отправка аудио сообщения пользователю
// 	sentMessage, err := bot.Send(voiceConfig)
// 	if err != nil {
// 		log.Println("Error sending voice:", err)
// 	}
// 	messageID := sentMessage.MessageID
// 	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Аудио сообщение [MessageID: %d] пользователю с ID %d было успешно отправлено!", messageID, P_id))
// 	bot.Send(msg)

// }

// // Отправление видео сообщения пользователю через бота
// func AdminVideoNoteMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	// Определение последнего полученного видео сообщения
// 	videoNote := (*update.Message.VideoNote)
// 	fileConfig := tgbotapi.FileConfig{FileID: videoNote.FileID}
// 	// Создание объекта tgbotapi.VideoNoteShare для отправки видео сообщения
// 	videoNoteConfig := tgbotapi.NewVideoNoteShare(P_id, 0, fileConfig.FileID)
// 	// Отправка видео сообщения пользователю
// 	sentMessage, err := bot.Send(videoNoteConfig)
// 	if err != nil {
// 		log.Println("Error sending video note:", err)
// 	}
// 	messageID := sentMessage.MessageID
// 	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Видео сообщение [MessageID: %d] пользователю с ID %d было успешно отправлено!", messageID, P_id))
// 	bot.Send(msg)
// }

// // Отправление видео сообщения пользователю через бота
// func AdminStickerMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	// Определение последнего полученного стикера
// 	sticker := (*update.Message.Sticker)
// 	fileConfig := tgbotapi.FileConfig{FileID: sticker.FileID}
// 	// Создание объекта tgbotapi.stickerConfig для отправки стикера
// 	stickerConfig := tgbotapi.NewStickerShare(P_id, fileConfig.FileID)
// 	// Отправка стикера пользователю
// 	sentMessage, err := bot.Send(stickerConfig)
// 	if err != nil {
// 		log.Println("Error sending sticker:", err)
// 	}
// 	messageID := sentMessage.MessageID
// 	msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Стикер [MessageID: %d] пользователю с ID %d был успешно отправлен!", messageID, P_id))
// 	bot.Send(msg)
// }

// // Удаление отправленного пользователю сообщения
// func DelMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	delargs := strings.TrimSpace(update.Message.CommandArguments()) // Аргументы для удаления сообщения (ChatId и MessageID), с помощью strings.TrimSpace убираем лишние пробелы по бокам
// 	// Разбиваем аргументы на две переменные
// 	argDelList := strings.Fields(delargs)
// 	if len(argDelList) < 2 || delargs == "" { // Если аргументы не указаны или их не хватает выводим подсказку
// 		msg := tgbotapi.NewMessage(Admin_id, "Для удаления сообщения необходимо ввести /del [ChatId] [MessageID].")
// 		bot.Send(msg)
// 	}
// 	if len(argDelList) >= 2 {
// 		DChatId := argDelList[0]               // ChatId получателя в виде строки
// 		DMessageId := argDelList[1]            // MessageIdполучателя в виде строки
// 		DChatIdInt, _ := strconv.Atoi(DChatId) // ChatId получателя в виде числа
// 		DChatIdInt64 := int64(DChatIdInt)
// 		DMessageIdInt, _ := strconv.Atoi(DMessageId) // MessageID получателя в виде числа
// 		_, err := bot.DeleteMessage(tgbotapi.NewDeleteMessage(DChatIdInt64, DMessageIdInt))
// 		if err != nil {
// 			log.Println(err)
// 			msg := tgbotapi.NewMessage(Admin_id, "Ты что-то сделал не так, попробуй ещё раз.")
// 			bot.Send(msg)
// 		} else {
// 			msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Сообщение [MessageID: %d] пользователю с ID %d было успешно удалено!", DMessageIdInt, DChatIdInt64))
// 			bot.Send(msg)
// 		}
// 	}
// }

// // Функция для редактирования отправленного сообщения
// func EditMsg(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	editargs := strings.TrimSpace(update.Message.CommandArguments()) // Аргументы для удаления сообщения (ChatId и MessageID), с помощью strings.TrimSpace убираем лишние пробелы по бокам
// 	argEditList := strings.Fields(editargs)
// 	if len(argEditList) < 3 || editargs == "" { // Если аргументы не указаны или их не хватает выводим подсказку
// 		msg := tgbotapi.NewMessage(Admin_id, "Для редактирования сообщения необходимо ввести /edit [ChatId] [MessageID] [отредактированный текст].")
// 		bot.Send(msg)
// 	}
// 	if len(argEditList) >= 3 {
// 		EChatId := argEditList[0]              // ChatId получателя в виде строки
// 		EMessageId := argEditList[1]           // MessageIdполучателя в виде строки
// 		EChatIdInt, _ := strconv.Atoi(EChatId) // ChatId получателя в виде числа
// 		EChatIdInt64 := int64(EChatIdInt)
// 		EMessageIdInt, _ := strconv.Atoi(EMessageId) // MessageID получателя в виде числа
// 		EText := argEditList[2]
// 		edit := tgbotapi.NewEditMessageText(EChatIdInt64, EMessageIdInt, EText)
// 		_, err := bot.Send(edit)
// 		if err != nil {
// 			log.Println(err)
// 			msg := tgbotapi.NewMessage(Admin_id, "Ты что-то сделал не так, попробуй ещё раз.")
// 			bot.Send(msg)
// 		} else {
// 			msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Сообщение [MessageID: %d] отправленное пользователю с ID %d было успешно отредактированно!", EMessageIdInt, EChatIdInt64))
// 			bot.Send(msg)
// 		}
// 	}
// }

// // Функция для блокировки пользователя
// func BlockUser(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
// 	if update.Message.CommandArguments() == "" {
// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Необходимо ввести ID пользователя. Например: /block 5451284197")
// 		bot.Send(msg)
// 	} else {
// 		blockUserIdStr := strings.TrimSpace(update.Message.CommandArguments()) // id получателя в виде строки
// 		blockUserIdInt, _ := strconv.Atoi(blockUserIdStr)                      // id получателя в виде числа
// 		BlockUserId := int64(blockUserIdInt)
// 		if BlockUserId == 0 || len(blockUserIdStr) > 12 {
// 			msg := tgbotapi.NewMessage(Admin_id, "Ты что-то сделал не так. Попробуй ещё раз.")
// 			bot.Send(msg)
// 		} else {
// 			Block_list = append(Block_list, int(BlockUserId))
// 			msg := tgbotapi.NewMessage(Admin_id, fmt.Sprintf("Готово! Пользователь с ID %d заблокирован.", BlockUserId))
// 			bot.Send(msg)
// 		}
// 	}
// }

// // Функция, которая проверяет, заблокирован ли пользователь с определенным ID
// func UserBlocked(update tgbotapi.Update, bot *tgbotapi.BotAPI) bool {
// 	for _, blockedID := range Block_list {
// 		if blockedID == int(update.Message.Chat.ID) {
// 			return true
// 		}
// 	}
// 	return false
// }
