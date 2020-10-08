package core

// String is simple string that can be null.
type String *string

// NewString transforms a string into a nullable string.
func NewString(value string) String {
	return &value
}

func StringValue(s String) string {
	if s == nil {
		return ""
	}

	return *s
}
