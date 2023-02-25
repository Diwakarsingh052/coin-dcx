package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {

	v, ok := ctx.Value(KeyValue).(*Values) // type assertion
	if !ok {
		return fmt.Errorf("web.Values missing from the context")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	v.StatusCode = statusCode
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		return err
	}
	return nil
}

func RespondError(ctx context.Context, w http.ResponseWriter, err error) error {

	var webErr *Error
	//if user used it to create error message or not
	ok := errors.As(err, &webErr)

	// if webErr is being used to construct an Error then I will send the content of it to the end user
	if ok {
		er := ErrorResponse{Error: webErr.Err.Error()}
		err := Respond(ctx, w, er, webErr.Status)
		if err != nil {
			return err
		}
		return nil // it means RespondError got success in sending an error response back to the client
	}

	er := ErrorResponse{Error: http.StatusText(http.StatusInternalServerError)}
	err = Respond(ctx, w, er, http.StatusInternalServerError)

	if err != nil {
		return err
	}
	return nil

}

func Decode(r *http.Request, val any) error {

	err := json.NewDecoder(r.Body).Decode(&val)

	if err != nil {
		return err
	}

	return nil
}
