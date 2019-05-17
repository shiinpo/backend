package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"golang.org/x/crypto/bcrypt"
)

var categories = []string{
	"upper body",
	"lower body",
	"arms",
	"legs",
	"cardio",
}

var exercises = map[string]int{
	"squat":    4,
	"bench":    1,
	"deadlift": 2,
}

var userEntries = []Entry{
	{ID: 0, Weight: 225, Reps: 10, RPE: 10, DatePerformed: "2019/05/28", ExerciseID: 2, UserID: 1},
	{ID: 0, Weight: 235, Reps: 10, RPE: 10, DatePerformed: "2019/05/29", ExerciseID: 2, UserID: 1},
	{ID: 0, Weight: 215, Reps: 10, RPE: 10, DatePerformed: "2019/05/30", ExerciseID: 2, UserID: 1},
	{ID: 0, Weight: 255, Reps: 10, RPE: 10, DatePerformed: "2019/05/31", ExerciseID: 2, UserID: 1},
	{ID: 0, Weight: 265, Reps: 10, RPE: 10, DatePerformed: "2019/06/01", ExerciseID: 2, UserID: 1},
	{ID: 0, Weight: 325, Reps: 10, RPE: 10, DatePerformed: "2019/05/28", ExerciseID: 1, UserID: 1},
	{ID: 0, Weight: 335, Reps: 10, RPE: 10, DatePerformed: "2019/05/29", ExerciseID: 1, UserID: 1},
	{ID: 0, Weight: 315, Reps: 10, RPE: 10, DatePerformed: "2019/05/30", ExerciseID: 1, UserID: 1},
	{ID: 0, Weight: 355, Reps: 10, RPE: 10, DatePerformed: "2019/05/31", ExerciseID: 1, UserID: 1},
	{ID: 0, Weight: 365, Reps: 10, RPE: 10, DatePerformed: "2019/06/01", ExerciseID: 1, UserID: 1},
	{ID: 0, Weight: 425, Reps: 10, RPE: 10, DatePerformed: "2019/05/28", ExerciseID: 3, UserID: 1},
	{ID: 0, Weight: 435, Reps: 10, RPE: 10, DatePerformed: "2019/05/29", ExerciseID: 3, UserID: 1},
	{ID: 0, Weight: 415, Reps: 10, RPE: 10, DatePerformed: "2019/05/30", ExerciseID: 3, UserID: 1},
	{ID: 0, Weight: 455, Reps: 10, RPE: 10, DatePerformed: "2019/05/31", ExerciseID: 3, UserID: 1},
	{ID: 0, Weight: 465, Reps: 10, RPE: 10, DatePerformed: "2019/06/01", ExerciseID: 3, UserID: 1},
}

// RunMigrations runs migrations on database
func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

// RunDBSeeds migrates and seeds databse with JSON file from scraper
func RunDBSeeds(db *sql.DB) {
	userSeeds(db)
	categorySeeds(db)
	exerciseSeeds(db)
	userEntrySeeds(db)
}

// userSeeds seeds default user
func userSeeds(db *sql.DB) {
	hash, err := bcrypt.GenerateFromPassword([]byte("pass"), 10)
	if err != nil {
		log.Fatal("error hasing seed password")
	}

	var id int
	err = CreateUser(db, &id, "rey", string(hash))
	if err != nil {
		log.Fatal("error seeding user")
	}
}

// categorySeeds seeds default categories
func categorySeeds(db *sql.DB) {
	for _, category := range categories {
		cat, err := CreateCategory(db, category)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("CREATED CATEGORY ID:%d, NAME:%s\n", cat.ID, cat.Name)
		}
	}
}

// exerciseSeeds seeds default exercises
func exerciseSeeds(db *sql.DB) {
	for exercise, categoryID := range exercises {
		exer, err := CreateExercise(db, exercise, categoryID)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("CREATED EXERCISE ID:%d, NAME:%s\n", exer.ID, exer.Name)
		}
	}
}

// userEntrySeeds seeds default user entries
func userEntrySeeds(db *sql.DB) {
	for _, entry := range userEntries {
		ent, err := CreateEntry(db, entry)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("CREATED ENTRY ID:%d\n", ent.ID)
		}
	}
}
