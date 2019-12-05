package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFail(t *testing.T) {
	assert.Fail(t, "expected Failure")
}
