package signIN

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"hbDima/internal/lib/logger/sl"
	"hbDima/internal/storage/sqlite"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

type Request struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type Response struct {
	resp.Response
	Token    string `json:"token"`
	UserID   int    `json:"user_id"`
	ExpireAt int64  `json:"expire_at"`
}
type SignINHandler interface {
	PostUserIN(usersInPost sqlite.Users1) (int, int64, error)
}

func Autoriz(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "signIN.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func SignIN(log *slog.Logger, pUsers1 SignINHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.SignIN"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			// Такую ошибку встретим, если получили запрос с пустым телом.
			// Обработаем её отдельно
			log.Error("request body is empty")

			render.JSON(w, r, resp.Error("empty request"))

			return
		}
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}
		Token := sqlite.GenerateToken(32)
		Phone := req.Phone
		Password := req.Password
		fmt.Println(Phone, Password)

		user_id, exp, err := pUsers1.PostUserIN(sqlite.Users1{
			Phone:    Phone,
			Password: Password,
			Token:    Token,
		})
		if err != nil {
			log.Error("failed sign in", sl.Err(err))
			render.JSON(w, r, resp.Error("failed sign in"))
			return
		}
		responseOK(w, r, Token, user_id, exp)
	}
}
func responseOK(w http.ResponseWriter, r *http.Request, Token string, user_id int, expire_at int64) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Token:    Token,
		UserID:   user_id,
		ExpireAt: expire_at,
	})
}
