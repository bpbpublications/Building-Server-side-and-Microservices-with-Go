package util

import (
	"errors"
	"log"
	"net/http"
)

var (
	// MapErrorTypeToHTTPStatus maps errors to corresponding HTTP status code
	MapErrorTypeToHTTPStatus = mapErrorTypeToHTTPStatus

	// IsError returns underlying error type
	IsError = isError

	// NewError creates new Error object
	NewError = newError
)

var (
	// ErrBadRequest is returned when there is something wrong with the request itself
	// (i.e. invalid request body)
	// Should be returned as 400 HTTP Status
	ErrBadRequest = errors.New("Bad request")

	// ErrInternal is returned when something bad happened during processing API request
	// Should be returned as 500 HTTP Status
	ErrInternal = errors.New("Internal error")

	// ErrInvalidAPICall is returned when API call issued to server is invalid
	// (API is not supported or URI is invalid)
	// Should be returned as 404 HTTP Status
	ErrInvalidAPICall = errors.New("Invalid API call")

	// ErrNotAuthenticated is returned when user is not authenticated
	// (authentication token is missing or is invalid)
	// Should be returned as 401 HTTP Status
	ErrNotAuthenticated = errors.New("Not authenticated")

	// ErrResourceNotFound is returned when API called with a request
	// for nonexisting system resource
	// Should be returned as 401
	ErrResourceNotFound = errors.New("Resource not found")
)

// Error codes used when returning errors
const (
	ErrorCodeInternal           = 0   // 500; Internal server error
	ErrorCodeInvalidJSONBody    = 30  // 400; Request body contains invalid JSON
	ErrorCodeInvalidCredentials = 201 // 401 or 403; Username or password are invalid
	ErrorCodeEntityNotFound     = 404 // 404; Entity not found
	ErrorCodeValidation         = 500 // 400; Value provided for a field is not allowed (Validation error)
)

// ErrorResponse is sent back to client as JSON
type ErrorResponse struct {
	ErrorCode int
	Cause     string
}

type serverError struct {
	code      int
	cause     string
	errorType error
}

func (e serverError) Error() string {
	return e.cause
}

func mapErrorTypeToHTTPStatus(err error) int {
	switch err {
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrInternal:
		return http.StatusInternalServerError
	case ErrInvalidAPICall, ErrResourceNotFound:
		return http.StatusNotFound
	case ErrNotAuthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

// returns: IsServerError, ErrorCode, Cause, ErrorType
func isError(errorType error) (bool, int, string, error) {
	err, isError := errorType.(serverError)
	if !isError {
		return false, 0, "", errorType
	}

	return true, err.code, err.cause, err.errorType
}

func newError(cause string, code int, errorType, err error) error {
	// Error logging
	if err != nil {
		log.Printf("error: %v: %v", cause, err)
	} else {
		log.Printf("error: %v", cause)
	}

	return serverError{code, cause, errorType}
}
