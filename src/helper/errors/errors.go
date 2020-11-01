package errors

import (
	"errors"
	"net/http"
)

// ErrInternalServer error
var ErrInternalServer = errors.New("Internal Server Error")

// ErrNotFound error
var ErrNotFound = errors.New("Not Found")

// ErrUnauthorized error
var ErrUnauthorized = errors.New("Unauthorized")

// ErrForbidden error
var ErrForbidden = errors.New("Forbidden")

// ErrUnprocessableEntitiy error
var ErrUnprocessableEntitiy = errors.New("Unprocessable Entity")

// ErrBadRequest error
var ErrBadRequest = errors.New("Bad Request")

// ErrMethodNotAllowed error
var ErrMethodNotAllowed = errors.New("Method Not Allowed")

// ErrLocked error
var ErrLocked = errors.New("Locked")

// GetHTTPStatusCodeByError returns http status code by given error.
//
// Example of return:
//
// When the given error is "errors.ErrNotFound" then the return will be constant value of "http.StatusNotFound".
// Return value will be "http.StatusInternalServerError" when the given error is not included in helper/errors
func GetHTTPStatusCodeByError(err error) int {
	switch err {
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrNotFound:
		return http.StatusNotFound
	case ErrForbidden:
		return http.StatusForbidden
	case ErrInternalServer:
		return http.StatusInternalServerError
	case ErrUnprocessableEntitiy:
		return http.StatusUnprocessableEntity
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrMethodNotAllowed:
		return http.StatusMethodNotAllowed
	case ErrLocked:
		return http.StatusLocked

	default:
		return http.StatusInternalServerError
	}
}
