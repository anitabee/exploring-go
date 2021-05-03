package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestgetResponseBody(t *testing.T) {
	// todo: build me mock server
	response := getResponseBody("https://api.github.com/emojis")

	assert.EqualValues(t, response, "https://github.githubassets.com/images/icons/emoji/unicode/1f44b.png?v8")
}
