package cron

import (
	"Sechenovka/db"
	"Sechenovka/internal/storage/user"
	"Sechenovka/internal/storage/user_result"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	for {
		db := db.ConnectDB()
		userStorage := user.New(db)
		userResultStorage := user_result.New(db)
		
		// Отправляем уведомление врачу
		go func() {
			payload := map[string]string{
				"title": "Тревога",
				"body":  "Пациенту стало плохо. Срочно проверьте его состояние!",
			}
			body, _ := json.Marshal(payload)

			resp, err := http.Post("http://push_sender:8081/api/notify", "application/json", bytes.NewReader(body))
			if err != nil {
				fmt.Println("Ошибка при отправке уведомления:", err)
				return
			}
			defer resp.Body.Close()
		}()
	}
	return
}
