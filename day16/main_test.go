package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_goRoutine(t *testing.T) {
	list := splitInput("123")
	assert.Equal(t, []int{1, 2, 3}, list)
}
func Test_getPattern(t *testing.T) {
	assert.Equal(t, []int{0, 0, 1, 1, 0, 0, -1, -1}, genPattern(2))
}

func Test_sum(t *testing.T) {
	assert.Equal(t, 4, do([]int{0, 1, 0, -1}, []int{1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 8, do([]int{0, 0, 1, 1, 0, 0, -1, -1}, []int{1, 2, 3, 4, 5, 6, 7, 8}))
}

func Test_fft(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8}
	out := []int{4, 8, 2, 2, 6, 1, 5, 8}
	assert.Equal(t, out, fft(in))

	out2 := []int{3, 4, 0, 4, 0, 4, 3, 8}
	assert.Equal(t, out2, fft(fft(in)))
}

func Test_100fft(t *testing.T) {
	in := "80871224585914546619083218645595"
	assert.Equal(t, 24176176, fft100(in))
}
