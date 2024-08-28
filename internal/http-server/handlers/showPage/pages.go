package showPage

import (
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func Home(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "homePage.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func Bio(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "bio.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func Nastya(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "nastyaPage.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func Mom(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "momPage.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func Anya(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "anyaPage.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func Cats(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "catsPage.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func Birth(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "happyB.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
func Foto(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.signIN.SignIN.IndexPage"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		workDir, _ := os.Getwd()
		fileLocation := filepath.Join(workDir, "html", "foto.html")
		html, err := os.ReadFile(fileLocation)
		_, err = w.Write(html)
		fmt.Println(err)

	}
}
