package main

import (
	"log"
	"net/http"
	"s/bd"
	"s/routes"
)

func main() {

	db, err := bd.NewPosgresDB(bd.Config{
		Host:     "localhost",
		Port:     "8008",
		Username: "postgres",
		Password: "1234",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	// Инициализация маршрутов
	r := routes.SetupRoutes()

	// Запуск сервера
	log.Println("Server started on :8000")
	http.ListenAndServe(":8000", r)
}
