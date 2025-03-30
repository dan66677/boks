package routes

import (
	"s/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Маршруты
	r.HandleFunc("/fights", handlers.GetFights).Methods("GET")
	r.HandleFunc("/fights/{id}", handlers.GetFight).Methods("GET")
	r.HandleFunc("/fights", handlers.CreateFight).Methods("POST")
	r.HandleFunc("/fights/{id}", handlers.UpdateFight).Methods("PUT")
	r.HandleFunc("/fights/{id}", handlers.DeleteFight).Methods("DELETE")

	return r
}
