package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"service-app/web"
)

func (m *Mid) Panic(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
		v, ok := ctx.Value(web.KeyValue).(*web.Values) // type assertion and making sure values struct with the trace id is present
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}
		defer func() {
			r := recover()

			if r != nil { // panic happened
				s := fmt.Sprintf("PANIC :%v", r)

				//err is a named return value // this would be returned automatically when this function return
				err = errors.New(s)
				// Log the Go stack trace for this panic's goroutine.
				log.Printf("%s :\n%s", v.TraceId, debug.Stack())
			}
		}()

		return next(ctx, w, r)
	}
}
