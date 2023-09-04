package main

import (
	"log"
	"os"

	"github.com/p1xray/port-api/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// Читаем конфиг из переменных окружения
	_ = config.Read()

	return nil
}
