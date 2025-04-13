package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/SherClockHolmes/webpush-go"
)

//
//// Подписки (можно заменить на базу)
//var subscriptions = make([]*webpush.Subscription, 0)
//var mu sync.Mutex
//
//// Твоя почта + VAPID ключи (можно вынести в env)
//var (
//	vapidPublicKey  = "ТВОЙ_PUBLIC_KEY"
//	vapidPrivateKey = "ТВОЙ_PRIVATE_KEY"
//)
//
//func main() {
//	http.HandleFunc("/api/subscribe", handleSubscribe)
//	http.HandleFunc("/api/notify", handleNotify)
//
//	fmt.Println("Сервер запущен на http://localhost:8080")
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}
//
//func handleSubscribe(w http.ResponseWriter, r *http.Request) {
//	var sub webpush.Subscription
//	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
//		http.Error(w, "Ошибка разбора подписки", http.StatusBadRequest)
//		return
//	}
//
//	mu.Lock()
//	subscriptions = append(subscriptions, &sub)
//	mu.Unlock()
//
//	fmt.Println("Новая подписка:", sub.Endpoint)
//	w.WriteHeader(http.StatusCreated)
//}
//
//func handleNotify(w http.ResponseWriter, r *http.Request) {
//	mu.Lock()
//	defer mu.Unlock()
//
//	payload := map[string]string{
//		"title": "ЧекАп напоминает",
//		"body":  "Самое время проверить своё здоровье! 💊",
//	}
//	body, _ := json.Marshal(payload)
//
//	for _, sub := range subscriptions {
//		resp, err := webpush.SendNotification(body, sub, &webpush.Options{
//			TTL:             60,
//			VAPIDPublicKey:  vapidPublicKey,
//			VAPIDPrivateKey: vapidPrivateKey,
//			Subscriber:      "mailto:you@example.com",
//		})
//		if err != nil {
//			fmt.Println("Ошибка отправки:", err)
//			continue
//		}
//		defer resp.Body.Close()
//	}
//
//	fmt.Println("Уведомления отправлены!")
//	w.Write([]byte("ОК"))
//}
