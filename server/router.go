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
