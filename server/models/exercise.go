package models

import "database/sql"

// Category is the DB response struct from category table
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Exercise is the DB response struct from exercise table
type Exercise struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}

// Entry is the DB response struct from user_entry table
type Entry struct {
	ID            int    `json:"id"`
	Weight        int    `json:"weight"`
	Reps          int    `json:"reps"`
	RPE           int    `json:"rpe"`
	DatePerformed string `json:"date_performed"`
	ExerciseID    int    `json:"exercise_id"`
	UserID        int    `json:"user_id"`
}

//////////////////
//// CATEGORY ////
//////////////////

// GetAllCategories gets all categories
func GetAllCategories(db *sql.DB) ([]Category, error) {
	rows, err := db.Query(`SELECT c.id, c.name FROM category c`)
	if err != nil {
		return nil, err
	}

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil

}

// GetCategory gets category by ID
func GetCategory(db *sql.DB, id int) (Category, error) {
	var category Category
	err := db.QueryRow(`SELECT c.id, c.name FROM category c WHERE id = ($1)`,
		id).Scan(&category.ID, &category.Name)

	if err != nil {
		return category, err
	}

	return category, nil

}

// CreateCategory creates new category
func CreateCategory(db *sql.DB, name string) (Category, error) {
	var category Category
	err := db.QueryRow(`INSERT INTO category(name)
		VALUES
		(UPPER($1))
		RETURNING id, name`, name).Scan(&category.ID, &category.Name)

	if err != nil {
		return category, err
	}

	return category, nil
}

// EditCategory edits category by ID
func EditCategory(db *sql.DB, name string, id int) (Category, error) {
	var category Category
	err := db.QueryRow(`UPDATE category
		SET name = $1
		WHERE id = $2
		RETURNING id, name`, name, id).Scan(&category.ID, &category.Name)

	if err != nil {
		return category, err
	}

	return category, nil
}

// DeleteCategory deletes category
func DeleteCategory(db *sql.DB, id int) (int, error) {
	var count = 0
	rows, err := db.Query(`DELETE FROM category WHERE id = $1 RETURNING *`, id)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		count++
	}
	return count, nil
}

//////////////////
//// EXERCISE ////
//////////////////

// GetAllExercises gets all exercises
func GetAllExercises(db *sql.DB) ([]Exercise, error) {
	rows, err := db.Query(`SELECT e.id, e.name, e.category_id FROM exercise e`)
	if err != nil {
		return nil, err
	}

	var exercises []Exercise
	for rows.Next() {
		var exercise Exercise
		err := rows.Scan(
			&exercise.ID,
			&exercise.Name,
			&exercise.CategoryID,
		)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil
}

// GetExercise gets exercise by ID
func GetExercise(db *sql.DB, id int) (Exercise, error) {
	var exercise Exercise
	err := db.QueryRow(`SELECT e.id, e.name, e.category_id FROM exercise e WHERE id = ($1)`,
		id).Scan(&exercise.ID, &exercise.Name, &exercise.CategoryID)

	if err != nil {
		return exercise, err
	}

	return exercise, nil
}

// CreateExercise creates a new exercise
func CreateExercise(db *sql.DB, name string, categoryID int) (Exercise, error) {
	var exercise Exercise
	err := db.QueryRow(`INSERT INTO exercise(name, category_id)
		VALUES
		(UPPER($1), $2)
		RETURNING id, name, category_id`, name, categoryID).Scan(&exercise.ID, &exercise.Name, &exercise.CategoryID)

	if err != nil {
		return exercise, err
	}

	return exercise, nil
}

// EditExercise edits categoru by ID
func EditExercise(db *sql.DB, e Exercise) (Exercise, error) {
	var exercise Exercise
	err := db.QueryRow(`UPDATE exercise
		SET name = $1, category_id = $2
		WHERE id = $3
		RETURNING id, name, category_id`, e.Name, e.CategoryID, e.ID).Scan(&exercise.ID, &exercise.Name, &exercise.CategoryID)

	if err != nil {
		return exercise, err
	}

	return exercise, nil
}

// DeleteExercise deletes exercise by ID
func DeleteExercise(db *sql.DB, id int) (int, error) {
	var count = 0
	rows, err := db.Query(`DELETE FROM exercise WHERE id = $1 RETURNING *`, id)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		count++
	}
	return count, nil
}

//////////////////
////   ENTRY  ////
//////////////////

// GetAllEntries gets all user entry by user ID
func GetAllEntries(db *sql.DB, id int) ([]Entry, error) {
	rows, err := db.Query(`SELECT * FROM user_entry WHERE user_id = $1`, id)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(
			&entry.ID,
			&entry.Weight,
			&entry.Reps,
			&entry.RPE,
			&entry.DatePerformed,
			&entry.ExerciseID,
			&entry.UserID,
		)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// GetEntry gets entry by ID
func GetEntry(db *sql.DB, id int) (Entry, error) {
	var entry Entry
	err := db.QueryRow(`SELECT * FROM entry WHERE id = ($1)`, id).Scan(
		&entry.ID,
		&entry.Weight,
		&entry.Reps,
		&entry.RPE,
		&entry.DatePerformed,
		&entry.ExerciseID,
		&entry.UserID,
	)

	if err != nil {
		return entry, err
	}

	return entry, nil
}

// CreateEntry creates new entry
func CreateEntry(db *sql.DB, e Entry) (Entry, error) {
	var entry Entry
	err := db.QueryRow(`
		INSERT INTO user_entry(weight, reps, rpe, date_performed, exercise_id, user_id)
		VALUES
		($1, $2, $3, $4, $5, $6)
		RETURNING *`,
		e.Weight, e.Reps, e.RPE, e.DatePerformed, e.ExerciseID, e.UserID,
	).Scan(
		&entry.ID,
		&entry.Weight,
		&entry.Reps,
		&entry.RPE,
		&entry.DatePerformed,
		&entry.ExerciseID,
		&entry.UserID,
	)

	if err != nil {
		return entry, err
	}

	return entry, nil
}

// EditEntry edits entry by entry ID
func EditEntry(db *sql.DB, id int, e Entry) (Entry, error) {
	var entry Entry
	err := db.QueryRow(`UPDATE user_entry
		SET weight = $1, reps = $2, rpe = $3
		WHERE id = $4
		RETURNING *`,
		e.Weight, e.Reps, e.RPE, id,
	).Scan(
		&entry.ID,
		&entry.Weight,
		&entry.Reps,
		&entry.RPE,
		&entry.DatePerformed,
		&entry.ExerciseID,
		&entry.UserID,
	)

	if err != nil {
		return entry, err
	}

	return entry, nil
}

// DeleteEntry deletes entry by ID
func DeleteEntry(db *sql.DB, userID int, id int) (int, error) {
	var count = 0
	rows, err := db.Query(`DELETE FROM user_entry WHERE id = $1 AND user_id = $2 RETURNING *`, id, userID)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		count++
	}
	return count, nil
}
