package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	pwd, err := os.Getwd()
	if err != nil {
		http.Error(w, "internal error", 1)
		return
	}
	path := filepath.Join(filepath.Dir(pwd), "index.html")

	f, err := os.Open(path)
	if err != nil {
		http.Error(w, "internal error", 2)
		return
	}
	defer f.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, f)
	if err != nil {
		http.Error(w, "internal error", 3)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "parse form error", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "get file error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "read file error", http.StatusInternalServerError)
		return
	}

	result, err := service.HandleStr(string(data))
	if err != nil {
		http.Error(w, "incorrect file data", http.StatusBadRequest)
	}

	fileName := time.Now().UTC().String()
	localFile, err := os.Create(fileName + filepath.Ext(header.Filename))
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	localFile.WriteString(result)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
