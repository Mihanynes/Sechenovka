package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"strconv"
)

var bot *tgbotapi.BotAPI

func main() {
	token := "6658100298:AAEJIXBUpevzDcbQSBu47TRikbQ44XS0nPo"

	// Инициализируем бота
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	log.Printf("Бот %s запущен", bot.Self.UserName)

	http.HandleFunc("/send", sendNotificationHandler)
	log.Println("Сервер запущен на :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Получен запрос на отправку сообщения")
	if r.Method != "POST" {
		log.Println("Метод не POST")
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	chatId, err := strconv.ParseInt(r.URL.Query().Get("chatId"), 10, 64)
	if err != nil {
		log.Println("Ошибка парсинга chatId: " + err.Error())
		http.Error(w, "Ошибка парсинга chatId: "+err.Error(), http.StatusBadRequest)
		return
	}

	message := r.URL.Query().Get("message")

	if message == "" {
		log.Println("Не указан message")
		http.Error(w, "Не указан message", http.StatusBadRequest)
		return
	}

	// Отправляем сообщение в Telegram
	msg := tgbotapi.NewMessage(chatId, message)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("Ошибка отправки сообщения: " + err.Error())
		http.Error(w, "Ошибка отправки сообщения: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Сообщение отправлено: %s", message)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Сообщение отправлено"))
}
