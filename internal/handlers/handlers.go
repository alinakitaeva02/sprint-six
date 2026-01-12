package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only the POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20) // Парсит html-форму
	if err != nil {
		http.Error(w, "couldn't make out the form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile") // Получает данные из файла
	if err != nil {
		http.Error(w, "Couldn't get the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Couldn't read the file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	converted, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "Conversion error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename) // Берет расширение от исходного файла
	if ext == "" {
		ext = ".txt"
	}

	t := time.Now()
	filename := t.Format("2006-01-02_15-04-05.000000000") + ext

	f, err := os.Create(filename) // Создает файл
	if err != nil {
		http.Error(w, "failed to create a file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = f.WriteString(converted) // Записывает результат конвертации
	if err != nil {
		http.Error(w, "couldn't write to a file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, converted)
}
