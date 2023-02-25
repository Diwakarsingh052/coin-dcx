package handlers

import (
	"context"
	"net/http"
	"service-app/web"
)

func check(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{Status: "ok"}

	_ = status
	//panic("some kind of problem")
	//err := web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
	//return fmt.Errorf("checking error %w", err)
	//err := errors.New("this is an internal error")

	return web.Respond(ctx, w, status, http.StatusOK)

	//return json.NewEncoder(w).Encode(status)
}
