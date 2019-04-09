package main

import (
	"fmt"
	"log"
	"net/http"

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

func main() {

	http.HandleFunc("/signin", auth.Signin)
	http.HandleFunc("/welcome", auth.Protected(Welcome))
	// http.HandleFunc("/refresh", Refresh)

	// start the server on port 8000
	fmt.Printf("server live on port 9001\n")
	log.Fatal(http.ListenAndServe(":9001", nil))
}
