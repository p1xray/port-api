package config

import "os"

// Конфиг
type Config struct {
	HttpAddr string
}

// Читает конфиг из переменных окружения
func Read() Config {
	var config Config

	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HttpAddr = httpAddr
	}

	return config
}
