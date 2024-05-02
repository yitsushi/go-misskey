package test

import (
	"fmt"
	"os"
)

// LoadFixture loads the content of a fixture file.
func LoadFixture(name string) ([]byte, error) {
	content, err := os.ReadFile(fmt.Sprintf("fixtures/%s", name))
	if err != nil {
		path, _ := os.Getwd()

		return []byte{}, FixtureFileNotFoundError{path, name}
	}

	return content, nil
}

// Must be content. Basically go and panic if error happened.
func Must(content []byte, err error) []byte {
	if err != nil {
		panic(err)
	}

	return content
}
