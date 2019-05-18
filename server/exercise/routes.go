package exercise

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reynld/shinpo/server/models"
)

// GetUserRecords the user Records handler
func GetUserRecords(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("ID").(int)
	records, err := models.GetAllEntries(db, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(records)
}

// AddUserRecord the add new user record handler
func AddUserRecord(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ID").(int)
	var payload models.Entry
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	record, err := models.CreateEntry(
		db,
		models.Entry{
			ID:            0,
			Weight:        payload.Weight,
			Reps:          payload.Reps,
			RPE:           payload.RPE,
			DatePerformed: payload.DatePerformed,
			ExerciseID:    payload.ExerciseID,
			UserID:        userID,
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(record)
}

// EditUserRecord the edit record handler
func EditUserRecord(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ID").(int)
	var payload models.Entry
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
		case "weight":
			if value == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case "reps":
			if value == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case "rpe":
			if value == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}

	record, err := models.EditEntry(db, userID, payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(record)
}

// DeleteUserRecord the delete record handler
func DeleteUserRecord(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("ID").(int)

	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	count, err := models.DeleteEntry(db, userID, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"count": count})

}
