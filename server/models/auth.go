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
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserResponse for login
type UserResponse struct {
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

// GetByUsername gets User by username
func GetByUsername(db *sql.DB, username string) (User, error) {
	var user User
	err := db.QueryRow(
		`SELECT u.id, u.username, u.password, u.email FROM users u WHERE username = $1`,
		username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser returns User by username
func CreateUser(db *sql.DB, username string, hash string, email string) (UserResponse, error) {
	var user UserResponse
	err := db.QueryRow(`INSERT INTO users(username, password, email)
		VALUES
		($1, $2, $3)
		RETURNING id, username, email`, username, hash, email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}
