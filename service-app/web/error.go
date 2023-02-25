package web

// ErrorResponse is the form used for API responses from failures in the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Error struct is used by user of our app to generate error messages
// Later on from this struct we would fetch the error message and send it as response back to the client
type Error struct {
	Err    error
	Status int
}

func (err *Error) Error() string {
	return err.Err.Error()
}

func NewRequestError(err error, status int) error {
	return &Error{
		Err:    err,
		Status: status,
	}

}
