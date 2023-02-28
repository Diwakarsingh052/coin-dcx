package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

var r *chi.Mux

func main() {
	r = chi.NewRouter()
	HandleFunc("/home", Home)
	panic(http.ListenAndServe(":8080", r))

}

func HandleFunc(pattern string, handler HandlerFunc) {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := handler(ctx, w, r)
		if err != nil {
			log.Println(err)
		}
	}
	r.HandleFunc(pattern, f)
}

func Home(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello this is our home page"))
	fmt.Println("Home page invoked")
	return nil
}
