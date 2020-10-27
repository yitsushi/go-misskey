package core

import (
	"encoding/json"
	"fmt"
)

const (
	// ResponseReadError occurs when error happened while we tried to
	// read the response.
	ResponseReadError = "Error reading body"
	// ResponseReadBodyError occurs when error happened while we tried to
	// read the body of the response.
	ResponseReadBodyError = "Error reading body"
	// ErrorResponseParseError occurs when the response was an error,
	// but something went wrong with parsing it as an Error.
	ErrorResponseParseError = "Error response parse error"
	// UndefinedRequiredField occurs when a mandatory field is not defined
	// in a request.
	UndefinedRequiredField = "Undefined required field"
	// OutOfRangeError is a template error where a given value
	// has to be in a given range.
	OutOfRangeError = "Out of range [%d..%d]"
	// ExceedMaximumLengthError occurs when a parameter is longer
	// than expected on an endpoint.
	ExceedMaximumLengthError = "Value is too long"
)

// NewRangeError generates an error message for an OutOfRangeError.
func NewRangeError(from, to int64) string {
	return fmt.Sprintf(OutOfRangeError, from, to)
}

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

// UnknownError occurs when we couldn't determine the source of the error.
type UnknownError struct {
	Response ErrorResponse
}

func (e UnknownError) Error() string {
	return fmt.Sprintf("<%s> %s -> %s", e.Response.Code, e.Response.Info.Param, e.Response.Info.Reason)
}

// InvalidFieldReferenceError occurs when we couldn't determine the source of the error.
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

// RequestValidationError occurs when one or more
// mandatory fields are missing.
type RequestValidationError struct {
	Request BaseRequest
	Message string
	Field   string
}

// FieldError is the detailed error on a given field in a request.
type FieldError struct {
	Name  string
	Issue string
}

func (e RequestValidationError) Error() string {
	return fmt.Sprintf(
		"%T request validation failed: [%s] %s",
		e.Request,
		e.Field,
		e.Message,
	)
}
