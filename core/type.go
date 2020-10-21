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
type RequestHandlerFunc func(Request, interface{}) error

// DataSize is an alias for uint64 with extra functions for easier
// conversion between sizes. The value should contains bytes.
//
// I know I should use github.com/catalint/datasize or something similar,
// but I don't know if I want to support that later or not or I want to add
// extra functionality on it or not.
type DataSize uint64

// Bytes returns with its value as a uint64.
func (d DataSize) Bytes() uint64 {
	return uint64(d)
}

// Kilobytes returns with the calculated kilobytes value as a float.
func (d DataSize) Kilobytes() float64 {
	return float64(d.Bytes()) / 1024
}

// Megabytes returns with the calculated megabytes value as a float.
func (d DataSize) Megabytes() float64 {
	return d.Kilobytes() / 1024
}

// Gigabytes returns with the calculated gigabytes value as a float.
func (d DataSize) Gigabytes() float64 {
	return d.Megabytes() / 1024
}

// Terabytes returns with the calculated terabytes value as a float.
func (d DataSize) Terabytes() float64 {
	return d.Gigabytes() / 1024
}
