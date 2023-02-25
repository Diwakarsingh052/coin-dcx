package web

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type Values struct {
	TraceId    string
	Now        time.Time
	StatusCode int
}

// creating custom type for storing the value in the context
type ctxKey int

const KeyValue ctxKey = 1

type App struct {
	*chi.Mux // Mux impls http.Handler interface
	// and embedding this to app struct would also make our app struct to impl the http.Handler interface automatically
}

// HandlerFunc is a custom type like http.HandlerFunc func in standard library
type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func (a *App) HandleFunc(method string, path string, handler HandlerFunc) {
	// wrapping the actual call in a variable which have a same signature as a handlerFunc ,
	//so we can pass it to chi router to process it
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // taking context out of the Request
		v := &Values{
			TraceId: uuid.NewString(), // gen a unique id
			Now:     time.Now(),
		}

		//putting values struct in the context //
		// we can access these values during the request lifetime
		ctx = context.WithValue(ctx, KeyValue, v)

		err := handler(ctx, w, r)
		if err != nil {
			log.Println("error escaped from the middleware ", err)
			return
		}
	}
	// chi router can accept the f var because the signature of f var matches
	//to func(w http.ResponseWriter, r *http.Request)
	a.Mux.MethodFunc(method, path, f) // executing the handler func

}
