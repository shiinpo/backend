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
	ID         int `json:"id"`
	Sets       int `json:"sets"`
	Reps       int `json:"reps"`
	RPE        int `json:"rpe"`
	ExerciseID int `json:"exercise_id"`
	UserID     int `json:"user_id"`
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
// GetAllEntry gets all user entry by user ID
// GetEntry gets entry by ID
// CreateEntry creates new entry
// EditEntry edits entry by entry ID
// DelteEntry deletes entry by ID
