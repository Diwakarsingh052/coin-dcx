package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// req -> handlerFunc() -> CreateUser // normal flow
// req -> mid -> handlerFunc() -> CreateUser

type key int

const RequestIDKey key = 123

func main() {
	http.HandleFunc("/home", RequestIdMid(LoggingMid(home)))
	panic(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("In home Page handler")
	fmt.Fprintln(w, "hello this is our first home page")

}
func RequestIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		traceId := uuid.NewString()
		ctx = context.WithValue(ctx, RequestIDKey, traceId)
		next(w, r.WithContext(ctx))
	}
}

func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//do middleware specific stuff first and when it is over then go to the actual handler func to exec it

		reqID, ok := r.Context().Value(RequestIDKey).(string)
		if !ok {
			reqID = "unknown"
		}
		log.Printf("%s : started   : %s %s ",
			reqID,
			r.Method, r.URL.Path)
		if r.Method != http.MethodGet {
			http.Error(w, "method must be get", http.StatusInternalServerError)
			return
		}

		next(w, r) // executing the next handlerFunc or the middleware in the chain

		log.Printf("%s : completed : %s %s ",
			reqID,
			r.Method, r.URL.Path,
		)

	}
}
