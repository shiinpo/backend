package exercise

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reynld/shinpo/server/models"
)

// GetAllExercises the user Exercises handler
func GetAllExercises(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	exercises, err := models.GetAllExercises(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(exercises)
}

// AddExercise the add new user record handler
func AddExercise(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var payload models.Exercise
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	exercise, err := models.CreateExercise(db, payload.Name, payload.CategoryID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(exercise)
}

// EditExercise the edit record handler
func EditExercise(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var payload models.Exercise
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
		case "exercise":
			if value == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		}
	}

	exercise, err := models.EditExercise(db, payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(exercise)
}

// DeleteExercise the delete record handler
func DeleteExercise(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	count, err := models.DeleteExercise(db, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"count": count})

}
