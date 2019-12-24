package intmath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExp2(t *testing.T) {
	assert.Equal(t, 4, Exp2(2))
	assert.Equal(t, 1, Exp2(0))
	assert.Equal(t, 2, Exp2(1))
	assert.Equal(t, 8, Exp2(3))
	assert.Equal(t, 16, Exp2(4))
}
