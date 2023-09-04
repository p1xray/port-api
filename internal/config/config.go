package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Конфиг
type Config struct {
	HttpAddr string
}

// Читает конфиг из переменных окружения
func Read() Config {
	var config Config

	if err := godotenv.Load("local.env"); err != nil {
		log.Fatalf("Config Read Error: %v", err)
	}

	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HttpAddr = httpAddr
	}

	return config
}
