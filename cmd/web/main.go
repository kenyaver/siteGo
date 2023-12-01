package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Инициализируем FileServer, он будет обрабатывать
	// HTTP-запросы к статическим файлам из папки "./ui/static".
	// Обратите внимание, что переданный в функцию http.Dir путь
	// является относительным корневой папке проекта
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Используем функцию mux.Handle() для регистрации обработчика для 
	// всех запросов, которые начинаются с "/static/". Мы убираем
	// префикс "/static" перед тем как запрос достигнет http.FileServer
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Println("Запуск сервера на http://127.0.0.1:8080")
    log.Println("Запуск сервера на http://127.0.0.1:8080/snippet?id=17")
    log.Println("Запуск сервера на http://127.0.0.1:8080/snippet/create")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}