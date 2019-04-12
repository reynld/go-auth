package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reynld/go-auth/pkg/auth"
)

// Welcome route handler
func Welcome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username, ok := ctx.Value(string("username")).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("Welcome %s!", username)))
}

// initializeDB connects to DB
func (s *Server) initializeRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/sigin", auth.Signin).Methods("GET")
	s.Router.HandleFunc("/welcome", auth.Protected(Welcome)).Methods("GET")
	// s.router.HandleFunc("/referesh", auth.Protected(auth.Refresh)).Methods("GET")
}
