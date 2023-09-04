package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/p1xray/port-api/internal/config"
	"github.com/p1xray/port-api/internal/controller"
	"github.com/p1xray/port-api/internal/repository/inmem"
	"github.com/p1xray/port-api/internal/services"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// Читаем конфиг из переменных окружения
	cfg := config.Read()

	// Создаем необходимые репозитории
	portRepository := inmem.NewPortRepository()

	// Создаем необходимые сервисы
	portService := services.NewPortService(portRepository)

	// Создаем необходимые хэндлеры
	portHandler := controller.NewPortHandler(portService)

	// Создаем необходимые роуты
	router := mux.NewRouter()
	router.HandleFunc("/port", portHandler.GetPort).Methods("GET")

	server := &http.Server{
		Addr:    cfg.HttpAddr,
		Handler: router,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HttpAddr)

	// Запуск http сервера
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	return nil
}
