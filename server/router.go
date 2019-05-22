package server

import (
	"log"
	"net/http"

	"github.com/reynld/shinpo/server/auth"
	"github.com/reynld/shinpo/server/exercise"
)

// loggingMiddleware logs HTTP request
func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// getServerIsUp '/' endpoint cheks if server is up
func (s *Server) getServerIsUp(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("server is live"))
}

// routeNotFound '/*' endpoint for undefined routes
func (s *Server) routeNotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("route not found"))
}

// Login route wrapper
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	auth.Login(s.DB, w, r)
}

// Register route wrapper
func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	auth.Register(s.DB, w, r)
}

// GetUserInfo route wrapper
func (s *Server) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	auth.UserInfo(s.DB, w, r)
}

//////////////////
////  RECORD  ////
//////////////////

// GetUserRecords route wrapper
func (s *Server) GetUserRecords(w http.ResponseWriter, r *http.Request) {
	exercise.GetUserRecords(s.DB, w, r)
}

// AddUserRecord route wrapper
func (s *Server) AddUserRecord(w http.ResponseWriter, r *http.Request) {
	exercise.AddUserRecord(s.DB, w, r)
}

// EditUserRecord route wrapper
func (s *Server) EditUserRecord(w http.ResponseWriter, r *http.Request) {
	exercise.EditUserRecord(s.DB, w, r)
}

// DeleteUserRecord route wrapper
func (s *Server) DeleteUserRecord(w http.ResponseWriter, r *http.Request) {
	exercise.DeleteUserRecord(s.DB, w, r)
}

//////////////////
//// Exercise ////
//////////////////

// GetAllExercises route wrapper
func (s *Server) GetAllExercises(w http.ResponseWriter, r *http.Request) {
	exercise.GetAllExercises(s.DB, w, r)
}

// AddExercise route wrapper
func (s *Server) AddExercise(w http.ResponseWriter, r *http.Request) {
	exercise.AddExercise(s.DB, w, r)
}

// EditExercise route wrapper
func (s *Server) EditExercise(w http.ResponseWriter, r *http.Request) {
	exercise.EditExercise(s.DB, w, r)
}

// DeleteExercise route wrapper
func (s *Server) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	exercise.DeleteExercise(s.DB, w, r)
}

//////////////////
//// Category ////
//////////////////

// GetAllCategories route wrapper
func (s *Server) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	exercise.GetAllCategories(s.DB, w, r)
}

// AddCategory route wrapper
func (s *Server) AddCategory(w http.ResponseWriter, r *http.Request) {
	exercise.AddCategory(s.DB, w, r)
}

// EditCategory route wrapper
func (s *Server) EditCategory(w http.ResponseWriter, r *http.Request) {
	exercise.EditCategory(s.DB, w, r)
}

// DeleteCategory route wrapper
func (s *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	exercise.DeleteCategory(s.DB, w, r)
}
