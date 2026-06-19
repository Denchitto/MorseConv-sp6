package handlers

import (
	"MorseConv-sp6/internal/service"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {

	data, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	//Читаем файл, заполняем буфер, если буфер заполнен - добавляем в общий буфер и проходимся снова
	buf := make([]byte, 100)
	var bufTotal []byte
	n := 100
	for n == 100 {
		n, err = file.Read(buf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bufTotal = append(bufTotal, buf[:n]...)
	}
	if len(bufTotal) == 0 {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	data := service.MorseConv(string(bufTotal))

	//Присваиваю .txt если у загружаемого файла оно не указано
	fileNameExt := filepath.Ext(handler.Filename)
	if fileNameExt == "" {
		fileNameExt = ".txt"
	}
	//Вместо String() использую Format(), так как двоеточия не воспринимаются
	fileLocal, err := os.OpenFile("../"+time.Now().UTC().Format("2006-01-02 15-04-05")+fileNameExt, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer fileLocal.Close()

	_, err = fileLocal.WriteString(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
