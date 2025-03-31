package main

import (
	"log"
	"net/http"
	"os"
	"s/bd"
	"s/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env varialbles: %s", err.Error())
	}

	db, err := bd.InitDB(bd.Config{
		Host:     "localhost",
		Port:     os.Getenv("port"),
		Username: os.Getenv("post"),
		Password: os.Getenv("DB"),
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	r := routes.SetupRoutes(db) // Передаем правильное соединение

	log.Println("Сервер запущен на :8008")
	http.ListenAndServe(":8008", r)
}
