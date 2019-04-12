package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server struct
type Server struct {
	DB     *sql.DB
	Router *mux.Router
}

// Initialize Server struct
func (s *Server) Initialize() {
	s.initializeDB()
	s.initializeRoutes()
}

// Run runs the router
func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(s.Router)))
}
