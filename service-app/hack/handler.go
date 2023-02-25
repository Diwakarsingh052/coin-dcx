package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var mux = chi.NewRouter()

// HandlerFunc is a custom type like http.HandlerFunc func in standard library
type HandlerFunc func(context.Context, http.ResponseWriter, *http.Request) error

func HandleFunc(method string, path string, handler HandlerFunc) {

	f := func(w http.ResponseWriter, r *http.Request) {
		//taking the context out from the request
		ctx := r.Context()

		//wrapping the actual request
		err := handler(ctx, w, r)
		if err != nil {
			log.Println("error captured", err)
			return
		}
	}

	//exec the request
	mux.HandleFunc(path, f)
	//add := func(int, int) {}
	//sum(add)
}

func sum(func(int, int)) {

}

func main() {

	log.Println("main invoked")
	//mux.HandleFunc()
	HandleFunc(http.MethodGet, "/home", test)
	http.ListenAndServe(":8080", mux)

}

func test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "this is a test going on")
	return errors.New("testing custom handler")
}
