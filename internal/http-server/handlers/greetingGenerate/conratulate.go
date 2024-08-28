package greetingGenerate

import (
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"strings"
)

func (c *CongratulationsHandler) FetchGreeting() (string, error) {
	//возвращает фиксированное поздравление
	return "С днем рождения, Дима! Мы тебя любим", nil
}

type GreetingProvider interface {
	FetchGreeting() (string, error)
}

func GenerateGreetingHandler(log *slog.Logger, provider GreetingProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const operation = "handlers.greeting.Generate"

		log = log.With(
			slog.String("operation", operation),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		// Получение поздравления
		greetingMessage, err := provider.FetchGreeting()
		if err != nil {
			http.Error(w, "Ошибка получения поздравления", http.StatusInternalServerError)
			return
		}

		//количество повторений сообщения
		repetitionCount := 50

		//строка с повторяющимся сообщением
		repeatedGreetings := strings.Repeat(greetingMessage+"\n", repetitionCount)

		w.Header().Set("Content-Type", "text/plain")

		//отправляем строку с поздравлением
		_, err = w.Write([]byte(repeatedGreetings))
		if err != nil {
			http.Error(w, "Ошибка отправки поздравления", http.StatusInternalServerError)
		}
	}
}
