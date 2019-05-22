package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/reynld/shinpo/server/auth"
	"github.com/rs/cors"
)

// Server has db, router and cache instances
type Server struct {
	DB     *sql.DB
	Router *mux.Router
	Cache  *redis.Client
}

// Initialize creates DB, Router and Cache instances
func (s *Server) Initialize() {
	s.setRouter()
	s.connectDB()
}

// setRouter creates and connects mux router to server struct
func (s *Server) setRouter() {
	s.Router = mux.NewRouter()
	s.Router.Use(s.loggingMiddleware)

	// Auth + Default Endpoints
	s.Router.HandleFunc("/", s.getServerIsUp).Methods("GET")
	s.Router.HandleFunc("/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/register", s.Register).Methods("POST")
	s.Router.HandleFunc("/user", auth.Protected(s.GetUserInfo)).Methods("GET")

	// User Record Endpoints
	s.Router.HandleFunc("/record/all", auth.Protected(s.GetUserRecords)).Methods("GET")
	s.Router.HandleFunc("/record/add", auth.Protected(s.AddUserRecord)).Methods("POST")
	s.Router.HandleFunc("/record/edit", auth.Protected(s.EditUserRecord)).Methods("PUT")
	s.Router.HandleFunc("/record/delete/{id}", auth.Protected(s.DeleteUserRecord)).Methods("DELETE")

	// Exercise Endpoints
	s.Router.HandleFunc("/exercise/all", auth.Protected(s.GetAllExercises)).Methods("GET")
	s.Router.HandleFunc("/exercise/add", auth.Protected(s.AddExercise)).Methods("POST")
	s.Router.HandleFunc("/exercise/edit", auth.Protected(s.EditExercise)).Methods("PUT")
	// s.Router.HandleFunc("/exercise/delete/{id}", auth.Protected(s.DeleteExercise)).Methods("DELETE")

	// Category Endpoints
	s.Router.HandleFunc("/category/all", auth.Protected(s.GetAllCategories)).Methods("GET")
	s.Router.HandleFunc("/category/add", auth.Protected(s.AddCategory)).Methods("POST")
	s.Router.HandleFunc("/category/edit", auth.Protected(s.EditCategory)).Methods("PUT")
	// s.Router.HandleFunc("/category/delete/{id}", auth.Protected(s.DeleteCategory)).Methods("DELETE")

	s.Router.NotFoundHandler = http.HandlerFunc(s.routeNotFound)
}

// connectDB connects to DB
func (s *Server) connectDB() {
	dburi, err := s.GetDBUri()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", dburi)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	s.DB = db
}

// Run runs the server
func (s *Server) Run() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                                               // All origins
		AllowedHeaders: []string{"Authorization", "Content-Type"},                   // All headers
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowing only get, just an example
	})

	fmt.Printf("server live on port%s\n", port)
	log.Fatal(http.ListenAndServe(port, c.Handler(s.Router)))
}
