package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"service-app/web"
	"time"
)

func (m *Mid) Logger(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		v, ok := ctx.Value(web.KeyValue).(*web.Values) // type assertion and making sure values struct with the trace id is present
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}

		m.log.Printf("%s : started   : %s %s ",
			v.TraceId,
			r.Method, r.URL.Path)

		err := next(ctx, w, r) // executing the next handlerFunc or the middleware in the chain

		m.log.Printf("%s : completed : %s %s ->  (%d)  (%s)",
			v.TraceId,
			r.Method, r.URL.Path,
			v.StatusCode, time.Since(v.Now),
		)

		if err != nil {
			return err
		}

		return nil
	}
}
