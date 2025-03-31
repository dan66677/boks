package routes

import (
	"s/handlers"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(db *sqlx.DB) *mux.Router {
	r := mux.NewRouter()

	// Маршруты
	r.HandleFunc("/fights", handlers.GetFights(db)).Methods("GET")
	r.HandleFunc("/fights/{id}", handlers.GetFight(db)).Methods("GET")
	r.HandleFunc("/fights", handlers.CreateFight(db)).Methods("POST")
	r.HandleFunc("/fights/{id}", handlers.UpdateFight(db)).Methods("PUT")
	r.HandleFunc("/fights/{id}", handlers.DeleteFight(db)).Methods("DELETE")

	return r
}
