package models

import (
	"database/sql"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Credentials a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// User response from database
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Claims a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// JWTResponse struct returned from generate token
type JWTResponse struct {
	Token string
	Time  time.Time
}

// InitialUserResponse struct
type InitialUserResponse struct {
	User       User       `json:"user"`
	Records    []Record   `json:"records"`
	Categories []Category `json:"categories"`
	Exercises  []Exercise `json:"exercises"`
	Token      string     `json:"token"`
}

// GetByUsername gets User by username
func GetByUsername(db *sql.DB, username string) (User, string, error) {
	var user User
	var password string
	err := db.QueryRow(
		`SELECT u.id, u.username, u.password, u.email FROM users u WHERE username = $1`,
		username).Scan(&user.ID, &user.Username, &password, &user.Email)
	if err != nil {
		return user, password, err
	}
	return user, password, nil
}

// GetUserByID gets User by ID
func GetUserByID(db *sql.DB, id int) (User, error) {
	var user User
	err := db.QueryRow(
		`SELECT u.id, u.username, u.email FROM users u WHERE id = $1`,
		id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser returns User by username
func CreateUser(db *sql.DB, username string, hash string, email string) (User, error) {
	var user User
	err := db.QueryRow(`INSERT INTO users(username, password, email)
		VALUES
		($1, $2, $3)
		RETURNING id, username, email`, username, hash, email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetInitialUserResponse initial user Response
func GetInitialUserResponse(db *sql.DB, user User) (InitialUserResponse, error) {
	var userRes InitialUserResponse
	var err error

	userRes.User = user

	userRes.Categories, err = GetAllCategories(db)
	if err != nil {
		return userRes, err
	}

	userRes.Exercises, err = GetAllExercises(db)
	if err != nil {
		return userRes, err
	}

	userRes.Records, err = GetAllRecords(db, userRes.User.ID)
	if err != nil {
		return userRes, err
	}

	return userRes, err
}
