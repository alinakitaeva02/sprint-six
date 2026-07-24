package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux() // Создает роутер

	mux.HandleFunc("/", handlers.Handler)
	mux.HandleFunc("/upload", handlers.UploadHandler) //Регистрируем хендлеры из пакета handlers

	srv := &http.Server{ // Структура сервера
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{ // Возврат ссылки на сервер
		Logger: logger,
		Server: srv,
	}
}
