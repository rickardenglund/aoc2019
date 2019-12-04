package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	assert.True(t, isValid(111111, loose))
	assert.False(t, isValid(223450, loose))
	assert.False(t, isValid(123789, loose))
}

func Test_isValidStrict(t *testing.T) {
	assert.True(t, isValid(112233, strict))
	assert.False(t, isValid(123444, strict))
	assert.True(t, isValid(111122, strict))
	assert.False(t, isValid(699922, strict))
}