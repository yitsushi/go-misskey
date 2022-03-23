package permissions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey/core/permissions"
)

func TestWrite(t *testing.T) {
	testCases := map[permissions.Resource]permissions.Permission{
		permissions.Account:   "write:account",
		permissions.PageLikes: "write:page-likes",
	}

	for input, expected := range testCases {
		assert.Equal(t, expected, permissions.Write(input))
	}
}

func TestRead(t *testing.T) {
	testCases := map[permissions.Resource]permissions.Permission{
		permissions.Account:   "read:account",
		permissions.PageLikes: "read:page-likes",
	}

	for input, expected := range testCases {
		assert.Equal(t, expected, permissions.Read(input))
	}
}
