package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http" // Импортируем пакет bd
	"s/model"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)



func GetFights(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		rows, err := db.Query("SELECT id, title, fighter1, fighter2, winner FROM fights")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var fights []model.Fight
		for rows.Next() {
			var fight model.Fight
			if err := rows.Scan(&fight.ID, &fight.Title, &fight.Fighter1, &fight.Fighter2, &fight.Winner); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fights = append(fights, fight)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(fights)
	}
}

func GetFight(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var fight model.Fight
		err = db.QueryRow("SELECT id, title, fighter1, fighter2, winner FROM fights WHERE id = $1", id).
			Scan(&fight.ID, &fight.Title, &fight.Fighter1, &fight.Fighter2, &fight.Winner)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Fight not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(fight)
	}
}

func CreateFight(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var fight model.Fight
		if err := json.NewDecoder(r.Body).Decode(&fight); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := db.QueryRow(
			"INSERT INTO fights (title, fighter1, fighter2, winner) VALUES ($1, $2, $3, $4) RETURNING id",
			fight.Title, fight.Fighter1, fight.Fighter2, fight.Winner,
		).Scan(&fight.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(fight)
	}
}

func UpdateFight(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		var fight model.Fight
		if err := json.NewDecoder(r.Body).Decode(&fight); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(
			"UPDATE fights SET title = $1, fighter1 = $2, fighter2 = $3, winner = $4 WHERE id = $5",
			fight.Title, fight.Fighter1, fight.Fighter2, fight.Winner, id,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fight.ID = id
		json.NewEncoder(w).Encode(map[string]string{"message": "Fight updated successfully"})
	}
}

func DeleteFight(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE FROM fights WHERE id = $1", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
