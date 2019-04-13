package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/go-auth/pkg/auth"
)

// LoggingMiddleware logs HTTP request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// Welcome route handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username, ok := ctx.Value(string("username")).(string)
	id, ok := ctx.Value(string("id")).(int)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("user:%s, id:%d", username, id)))
}

// GetServerIsUp '/' endpoint cheks if server is up
func GetServerIsUp(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("server is live"))
}

// initializeDB connects to DB
func (s *Server) initializeRoutes() {
	s.Router = mux.NewRouter()
	s.Router.Use(LoggingMiddleware)

	s.Router.HandleFunc("/", GetServerIsUp).Methods("GET")
	s.Router.HandleFunc("/signin", s.Signin).Methods("POST")
	s.Router.HandleFunc("/register", s.Register).Methods("POST")
	s.Router.HandleFunc("/welcome", auth.Protected(Welcome)).Methods("GET")
}
