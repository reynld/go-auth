package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	s.Router.HandleFunc("/", s.getServerIsUp).Methods("GET")
	s.Router.HandleFunc("/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/register", s.Register).Methods("POST")
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

	fmt.Printf("server live on port%s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(s.Router)))
}
