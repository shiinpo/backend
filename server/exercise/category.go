package exercise

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reynld/shinpo/server/models"
)

// GetAllCategories the user Categories handler
func GetAllCategories(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategories(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(categories)
}

// AddCategory the add new user record handler
func AddCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var payload models.Category
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	category, err := models.CreateCategory(db, payload.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(category)
}

// EditCategory the edit record handler
func EditCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var payload models.Category
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var checkStruct map[string]interface{}
	jsonStruct, _ := json.Marshal(payload)
	json.Unmarshal(jsonStruct, &checkStruct)

	for key, value := range checkStruct {
		switch key {
		case "id":
			if value == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case "name":
			if value == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}

	category, err := models.EditCategory(db, payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(category)
}

// DeleteCategory the delete record handler
func DeleteCategory(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	count, err := models.DeleteCategory(db, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"count": count})
}
