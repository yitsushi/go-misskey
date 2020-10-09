package core

// String is simple string that can be null.
type String *string

// NewString transforms a string into a nullable string.
func NewString(value string) String {
	return &value
}

// StringValue returns the value of a core.String as string.
// If th value is nil, it returns with an empty string.
func StringValue(s String) string {
	if s == nil {
		return ""
	}

	return *s
}

// RequestHandlerFunc is the function signature for request
// handlers called by services to make requests.
type RequestHandlerFunc func(*BaseRequest, interface{}) error
