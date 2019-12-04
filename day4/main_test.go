package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	assert.True(t, isValid(111111))
	assert.False(t, isValid(223450))
	assert.False(t, isValid(123789))
}

func Test_isValidStrict(t *testing.T) {
	assert.True(t, isValidStrict(112233))
	assert.False(t, isValidStrict(123444))
	assert.True(t, isValidStrict(111122))
	assert.False(t, isValidStrict(699922))
}

func Test_groupSize(t *testing.T) {
	assert.Equal(t, 3, groupSize("111", 0))
	assert.Equal(t, 1, groupSize("0111", 0))
	assert.Equal(t, 2, groupSize("0111", 2))
}