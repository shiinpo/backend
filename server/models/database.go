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
	exerciseSeed(db)
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

// exerciseSeed seeds default exercises
func exerciseSeed(db *sql.DB) {
	for exercise, categoryID := range exercises {
		exer, err := CreateExercise(db, exercise, categoryID)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("CREATED EXERCISE ID:%d, NAME:%s\n", exer.ID, exer.Name)
		}
	}
}
