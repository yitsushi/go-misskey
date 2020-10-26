package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey/core"
)

// ValidateRequests validates requests with their Validate function and check if
// error occurred or not.
func ValidateRequests(t *testing.T, invalidRequests []core.BaseRequest, validRequests []core.BaseRequest) {
	for _, testCase := range invalidRequests {
		assert.Error(t, testCase.Validate())
	}

	for _, testCase := range validRequests {
		assert.NoError(t, testCase.Validate())
	}
}
