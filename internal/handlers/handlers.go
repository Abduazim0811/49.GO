package handlers

import (
	"cat-fact-service/internal/db"
	"cat-fact-service/internal/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetCatFactHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catFact, err := fetchCatFact()
		if err != nil {
			http.Error(w, "Failed to fetch cat fact", http.StatusInternalServerError)
			return
		}

		id, err := db.SaveCatFact(database, catFact.Fact)
		if err != nil {
			http.Error(w, "Failed to save cat fact", http.StatusInternalServerError)
			return
		}

		response := models.CatFactResponse{ID: id}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
func CreateTable(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS cat_facts (
        id SERIAL PRIMARY KEY,
        fact TEXT NOT NULL
    );`
    _, err := db.Exec(query)
    return err
}

func fetchCatFact() (models.CatFact, error) {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return models.CatFact{}, err
	}
	defer resp.Body.Close()

	var catFact models.CatFact
	err = json.NewDecoder(resp.Body).Decode(&catFact)
	if err != nil {
		return models.CatFact{}, err
	}

	return catFact, nil
}
