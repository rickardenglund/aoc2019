package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCalc(t *testing.T) {
	mem := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}

	assert.Equal(t, 43210, CalcSignal([]int{4, 3, 2, 1, 0}, mem))
}

func TestCalc2(t *testing.T) {
	mem := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}

	assert.Equal(t, 54321, CalcSignal([]int{0, 1, 2, 3, 4}, mem))
}

func TestCalcFeedback(t *testing.T) {
	mem := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	assert.Equal(t, 139629729, CalcSignalFeedback([]int{9, 8, 7, 6, 5}, mem))
}

func TestCalcFeedback2(t *testing.T) {
	mem := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

	assert.Equal(t, 18216, CalcSignalFeedback([]int{9, 7, 8, 5, 6}, mem))
}

func TestGetSequence(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, getSequence(123))
}

func TestChannels(t *testing.T) {
	s := []int{1, 2, 3}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c

	assert.Equal(t, 6, x+y)
}

func sum(ints []int, c chan int) {
	sum := 0
	for _, n := range ints {
		sum += n
	}
	c <- sum
}

func TestWG(t *testing.T) {
	var wg = sync.WaitGroup{}
	wg.Add(2)
}

func TestPanic(t *testing.T) {
	defer func() {
		val := recover()
		assert.Equal(t, 42, val)
	}()

	panic(42)
}
