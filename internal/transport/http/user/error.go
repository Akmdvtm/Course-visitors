package user

import "errors"

var (
	errMethodNotAllowed = errors.New("method not allowed")
)

type errResponse struct {
	Error string `json:"error"`
}

func newErrResponse(err error) errResponse {
	return errResponse{
		Error: err.Error(),
	}
}
