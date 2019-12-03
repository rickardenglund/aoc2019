package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_walkpath(t *testing.T) {
	fmt.Printf("%v\n", walkpath([]string{"U2,L2"}, 0))
}

func Test_getWireDistance(t *testing.T) {
	assert.Equal(t, 4, getWireDistance([]string{"U2,L2"}, 0, pos{-2, 2}))
}