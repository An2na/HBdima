package greetingGenerate

import (
	"github.com/go-chi/chi/v5/middleware"
	"hbDima/internal/lib/api/response"
	"hbDima/internal/storage/sqlite"
	"log/slog"
	"net/http"
)

type Response struct {
	response.Response
	Congrats string //
}
type CongratulationsHandler struct {
	Storage *sqlite.Storage
}

func (c *CongratulationsHandler) GetRandomGreeting() (string, error) {
	greeting, err := c.Storage.GetRandomGreeting()
	if err != nil {
		return "", err
	}
	return greeting.Message, nil
}

type Congratulations interface {
	GetRandomGreeting() (string, error)
}

func Congratulate(log *slog.Logger, useCongrats Congratulations) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.categories.Categ"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		//получение случайного поздравления
		randomGreet, err := useCongrats.GetRandomGreeting()
		if err != nil {
			http.Error(w, "Ошибка получения поздравления", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")

		//отправляем строку с поздравлением
		_, err = w.Write([]byte(randomGreet))
		if err != nil {
			http.Error(w, "Ошибка отправки поздравления", http.StatusInternalServerError)
		}
	}
}
