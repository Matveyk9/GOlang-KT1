package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
	
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		resp := Response{
			Message:   "Привет! Это JSON-ответ от Go сервиса",
			Timestamp: time.Now(),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {

			http.Error(w, fmt.Sprintf("Ошибка кодирования JSON: %v", err), http.StatusInternalServerError)
			return
		}
	})

	port := 8080
	fmt.Printf("Сервер запущен на http://localhost:%d/api\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {

		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
