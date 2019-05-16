package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Exercise struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"category_id"`
}

type Entry struct {
	ID         int `json:"id"`
	Sets       int `json:"sets"`
	Reps       int `json:"reps"`
	RPE        int `json:"rpe"`
	ExerciseID int `json:"exercise_id"`
	UserID     int `json:"user_id"`
}

// type Exercise struct {
// 	ID       int        `json:"id"`
// 	Name     string     `json:"name"`
// 	Category CategoryDB `json:"category"`
// }

// type Entry struct {
// 	ID       int      `json:"id"`
// 	Sets     int      `json:"sets"`
// 	Reps     int      `json:"reps"`
// 	RPE      int      `json:"rpe"`
// 	Exercise Exercise `json:"exercise"`
// 	UserID   int      `json:"user_id"`
// }

// GetAllCategories gets all categories
// GetCategory gets category by ID
// CreateCategory creates new category

// GetAllExercises gets all exercises
// GetExercise gets exercise by ID
// CreateExercises creates new exercise

// GetAllEntry gets all user entry by user ID
// GetEntry gets entry by ID
// CreateEntry creates new entry
// EditEntry edits entry by entry ID
// DelteEntry deletes entry by ID
