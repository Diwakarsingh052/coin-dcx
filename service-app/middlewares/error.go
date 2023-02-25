package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"service-app/web"
)

func (m *Mid) Error(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		v, ok := ctx.Value(web.KeyValue).(*web.Values) // type assertion and making sure values struct with the trace id is present
		if !ok {
			return fmt.Errorf("web.Values missing from the context")
		}
		// we call the next thing in the chain first to capture any errors at next stage
		err := next(ctx, w, r)

		if err != nil {
			// Log the error.
			m.log.Printf("%s : ERROR : %v", v.TraceId, err)
			err := web.RespondError(ctx, w, err)

			if err != nil {
				return err
			}
		}
		return nil
	}
}
