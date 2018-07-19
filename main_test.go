package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	err := Handler(Request{})
	assert.IsType(t, nil, err)
}
