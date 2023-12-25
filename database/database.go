package database

import (
	"database/sql"
	"fmt"
	"log"
	"tgbot3/admin"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq"
)

// Настройки базы данных
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "tgbot_users"
)

// Функция для занесения нового пользователя в базу данных
func InsertDb(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	// Создаем подключение к базе данных
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	// Проверяем, что подключение к базе данных успешно
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	// Получаем информацию о новом пользователе
	user := update.Message.From

	// Записываем пользователя в базу данных
	err = InsertUser(db, user.ID, user.UserName)
	if err != nil {
		log.Println("Ошибка при записи пользователя в базу данных:", err)
	} else {
		log.Println("Пользователь успешно записан в базу данных.")
	}
}

// Заносим пользователя в базу данных
func InsertUser(db *sql.DB, userID int, username string) error {
	if username == "" {
		username = "unknown"
	}
	_, err := db.Exec("INSERT INTO users (registration_date, username, tg_id) VALUES ($1, $2, $3)", time.Now(), username, userID)
	return err
}

// Функция для отправки сообщения с количеством пользователей
func UsersCount(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	usersCount, err := SelectUsersCount()
	if err != nil {
		log.Println("Ошибка получения количества пользователей: ", err)
	} else {
		msg := tgbotapi.NewMessage(admin.Admin_id, fmt.Sprintf("Количество пользователей в базе данных: %d.", usersCount))
		bot.Send(msg)
	}
}

// Функция для получения количества пользователей в базе данных
func SelectUsersCount() (int, error) {
	// Подключение к базе данных PostgreSQL
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	// Запрос к базе данных для получения количества пользователей
	var usersCount int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&usersCount)
	if err != nil {
		log.Println(err)
	}

	return usersCount, nil
}
