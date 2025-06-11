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

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Ошибка при чтении index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlContent)
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Ошибка парсинга", http.StatusInternalServerError)
		return

	}
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка получения файла из формы", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения содержимого файла", http.StatusInternalServerError)
		return
	}

	text := string(fileContent)
	result := service.MorzeDetect(text)

	filename := time.Now().UTC().String() + filepath.Ext(header.Filename)
	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Ошибка создания выходного файла", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	if _, err := outputFile.WriteString(result); err != nil {
		http.Error(w, "Ошибка записи результата в выходной файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, result)
}
