package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fuelNeeded(t *testing.T) {
	assert.Equal(t, fuelNeeded(12),2 )
}

func Test_fuelNeededRec(t *testing.T) {
	assert.Equal(t, 2, fuelNeededAdvanced(14))
	assert.Equal(t, 966, fuelNeededAdvanced(1969))
	assert.Equal(t, 50346, fuelNeededAdvanced(100756))
}