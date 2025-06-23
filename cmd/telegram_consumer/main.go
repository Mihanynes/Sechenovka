package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const welcomeMessage = `Приветствуем вас в боте для рассылки уведомлений о состоянии пациентов!👋

Введите команду /register, чтобы подписаться на уведомления о состоянии пациентов.🫀
`

// UserState хранит состояние пользователя
type UserState struct {
	WaitingForLogin    bool
	WaitingForPassword bool
	Login              string
}

var (
	bot        *tgbotapi.BotAPI
	userStates = make(map[int64]*UserState) // chat_id -> состояние
)

func main() {
	token := "6658100298:AAEJIXBUpevzDcbQSBu47TRikbQ44XS0nPo"

	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		text := update.Message.Text

		// Если у пользователя есть состояние - обрабатываем его
		if state, exists := userStates[chatID]; exists {
			handleUserState(chatID, text, state)
			continue
		}

		// Обработка команд
		if update.Message.IsCommand() {
			switch text {
			case "/register":
				// Начинаем регистрацию
				userStates[chatID] = &UserState{
					WaitingForLogin: true,
				}
				msg := tgbotapi.NewMessage(chatID, "Введите логин, который вы использовали при регистрации на сайте 💻:")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(chatID, welcomeMessage)
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(chatID, welcomeMessage)
			bot.Send(msg)
		}

	}
}

// Обработка состояния пользователя
func handleUserState(chatID int64, text string, state *UserState) {
	if state.WaitingForLogin {
		// Сохраняем логин и ждём пароль
		state.Login = text
		state.WaitingForLogin = false
		state.WaitingForPassword = true

		msg := tgbotapi.NewMessage(chatID, "Введите пароль, который вы использовали при регистрации на сайте 💻:")
		bot.Send(msg)
		return
	}

	if state.WaitingForPassword {
		// Пароль получен - завершаем регистрацию
		password := text

		baseUrl := "http://go_app:8080/api/auth/telegram/login"
		params := url.Values{}
		params.Set("password", password)
		params.Set("username", state.Login)
		params.Set("chatId", strconv.FormatInt(chatID, 10))
		fullURL := baseUrl + "?" + params.Encode()

		resp, err := http.Post(fullURL, "application/x-www-form-urlencoded", strings.NewReader(""))
		if err != nil {
			msg := tgbotapi.NewMessage(chatID, "Ошибка при авторизации: "+err.Error())
			bot.Send(msg)
			delete(userStates, chatID)
			return

		}
		if resp.StatusCode != 200 {
			msg := tgbotapi.NewMessage(chatID, "Ошибка при авторизации")
			bot.Send(msg)
			delete(userStates, chatID)
			return
		}

		msg := tgbotapi.NewMessage(chatID, "✅ Регистрация завершена!\nЛогин: "+state.Login+"\nПароль: "+password)
		bot.Send(msg)

		// Удаляем состояние
		delete(userStates, chatID)
	}
}
