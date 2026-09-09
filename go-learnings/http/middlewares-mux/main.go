package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Auth struct {
}

func (auth *Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("GET /r1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("main router:  open")
	})

	auth := new(Auth)
	sub1 := r.NewRoute().Subrouter()
	sub1.Use(auth.Authenticate)
	sub1.HandleFunc("GET /r2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("sub router 1: auth.authenticate")
	})

	sub2 := r.NewRoute().Subrouter()
	sub2.Use(Authenticate)
	sub2.HandleFunc("GET /r3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("sub router 2: authenticate")
	})
}

// this is function signature for middleware
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware")
		next.ServeHTTP(w, r)
	})
}
