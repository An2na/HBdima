package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local" env-required:"true"` //env-default дефолтное значение
	StoragePath string `yaml:"storage_path" env-required:"true"`                      //env-required — делает параметры обязательными. Если такой параметр не указан, мы будем получать ошибку.
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	//Phone       string        `yaml:"phone" env-required:"true"`
	//Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() *Config {
	// Получаем путь до конфиг-файла из env-переменной CONFIG_PATH
	//configPath := os.Getenv("CONFIG_PATH")
	//configPath := "./path/to/local.yaml"
	configPath := "/home/user/GoProjects/hbDima/config/local.yaml"

	if configPath == " " {
		log.Fatal("CONFIG_PATH is not set")
	}

	// Проверяем существование конфиг-файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}
	var cfg Config
	// Читаем конфиг-файл и заполняем нашу структуру
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
