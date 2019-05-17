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

// GetAllExercises gets all exercises
// GetExercise gets exercise by ID
// CreateExercises creates new exercise
// EditExercise edits categoru by ID
// DeleteExercise deletes exercise by ID

// GetAllEntry gets all user entry by user ID
// GetEntry gets entry by ID
// CreateEntry creates new entry
// EditEntry edits entry by entry ID
// DelteEntry deletes entry by ID
