package errors

import "net/http"

type Errors struct {
	Message string
	Status  int
	Error   string
}

func NewInternalServerError(message string) *Errors {
	return &Errors{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func NewBadRequestError(message string) *Errors {
	return &Errors{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request_error",
	}
}
