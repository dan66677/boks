package handlers

import (
	"encoding/json"
	"net/http"
	"s/bd" // Импортируем пакет bd
	"s/model"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var fights []model.Fight
var idCounter int

func GetFights(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Используем bd.DB для доступа к базе данных
	rows, err := bd.DB.Query("SELECT * FROM fights")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Считываем данные из rows в fights
	for rows.Next() {
		var fight model.Fight
		if err := rows.Scan(&fight.ID, &fight.Name, &fight.Date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fights = append(fights, fight)
	}

	json.NewEncoder(w).Encode(fights)
}

func GetFight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, fight := range fights {
		if fight.ID == id {
			json.NewEncoder(w).Encode(fight)
			return
		}
	}
	http.Error(w, "Fight not found", http.StatusNotFound)
}

func CreateFight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var fight model.Fight
	_ = json.NewDecoder(r.Body).Decode(&fight)
	fight.ID = idCounter
	idCounter++
	fights = append(fights, fight)
	json.NewEncoder(w).Encode(fight)
}

func UpdateFight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, fight := range fights {
		if fight.ID == id {
			fights = append(fights[:index], fights[index+1:]...)
			var updatedFight model.Fight
			_ = json.NewDecoder(r.Body).Decode(&updatedFight)
			updatedFight.ID = id
			fights = append(fights, updatedFight)
			json.NewEncoder(w).Encode(updatedFight)
			return
		}
	}
	http.Error(w, "Fight not found", http.StatusNotFound)
}

func DeleteFight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, fight := range fights {
		if fight.ID == id {
			fights = append(fights[:index], fights[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Fight not found", http.StatusNotFound)
}
