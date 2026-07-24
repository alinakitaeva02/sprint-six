package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "APP: ", log.LstdFlags|log.Lshortfile) //Создает логгер

	srv := server.NewServer(logger) //Создает сервер

	log.Printf("Запуск сервера на %s", srv.Server.Addr)

	if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server error: %v", err) //Вывод ошибок при помощи логгера
	}
}
