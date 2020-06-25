package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithArgs(t *testing.T) {
	os.Args = []string{"search", "github"}
	query, err := handleParams()
	assert.Nil(t, err)
	assert.Equal(t, "github", query)
}

func TestWithoutArgs(t *testing.T) {
	os.Args = []string{"search"}
	query, err := handleParams()
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid query", err.Error())
	assert.Equal(t, "", query)
}
