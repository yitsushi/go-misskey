package core

import (
	"encoding/json"
	"fmt"
)

const (
	// ResponseReadError occues when error happened while we tried to
	// read the response.
	ResponseReadError = "Error reading body"
	// ResponseReadBodyError occues when error happened while we tried to
	// read the body of the response.
	ResponseReadBodyError = "Error reading body"
	// ErrorResponseParseError occues when the response was an error,
	// but something went wrong with parsing it as an Error.
	ErrorResponseParseError = "Error response parse error"
)

// RequestError happens when something went wrong with the request.
type RequestError struct {
	Message string
	Origin  error
}

func (e RequestError) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Origin.Error())
}

// ErrorResponseWrapper is the wrapper for error responses.
type ErrorResponseWrapper struct {
	Error json.RawMessage `json:"error"`
}

// UnknownError occues when we coudn't determine the source of the error.
type UnknownError struct {
	Response ErrorResponse
}

func (e UnknownError) Error() string {
	return fmt.Sprintf("<%s> %s -> %s", e.Response.Code, e.Response.Info.Param, e.Response.Info.Reason)
}

// InvalidFieldReferenceError occues when we coudn't determine the source of the error.
type InvalidFieldReferenceError struct {
	Name      string
	Type      string
	Reference string
}

func (e InvalidFieldReferenceError) Error() string {
	return fmt.Sprintf(
		"%s refers to %s as %s, but %s is not defined",
		e.Name,
		e.Reference,
		e.Type,
		e.Reference,
	)
}
