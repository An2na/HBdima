package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"time"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	// Подключение к базе данных
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

type Greetings struct {
	Id      int
	Message string
}

func (s *Storage) GetRandomGreeting() (Greetings, error) { //метод возвращает Greetings и ошибку
	// Определите количество записей в таблице
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM greetings").Scan(&count)
	if err != nil {
		return Greetings{}, err
	}

	//генерация рандомного числа для айди поздравлений
	rand.Seed(time.Now().UnixNano())
	randomID := rand.Intn(count) + 1 // ID начинается с 1

	//случайное поздравление
	var greeting Greetings
	err = s.db.QueryRow("SELECT id, message FROM greetings WHERE id = ?", randomID).Scan(&greeting.Id, &greeting.Message)
	if err != nil {
		return Greetings{}, err
	}

	return greeting, nil
}
