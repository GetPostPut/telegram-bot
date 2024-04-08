# tg-feedback-bot

Telegram bot для обратной связи. ...

## Setup
### Создание базы данных

```sql
CREATE DATABASE tgbot_users;

CREATE TABLE users (
username VARCHAR(35),
registration_date DATE NOT NULL,
tg_id BIGINT NOT NULL,
UNIQUE (tg_id)
);
```

### Конфигурация

Установить значения в файле config.yml:
1. `bot-token` - api токен tg бота
2. `admin-id` - ...
3. `bd-password` - пароль от db
4. `db-host` - ...
5. `db-port` - порт от db
6. `db-user` - ...
7. `db-name` -  название db

## Run

```bush
$ go run cmd/bot/main.go
```

## Functions

### tg client
1. `/admin` - получение всех доступных администратору команд с подробным описанием
