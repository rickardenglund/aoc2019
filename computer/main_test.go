package computer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run(t *testing.T) {
	program := []int{1, 0, 0, 0, 99}
	c := NewComputer(program)

	go c.Run()
	<-c.HaltChannel

	assert.Equal(t, []int{2, 0, 0, 0, 99}, c.Mem)
}

func Test_run2(t *testing.T) {
	program := []int{2, 3, 0, 3, 99}
	c := NewComputer(program)

	go c.Run()
	<-c.HaltChannel

	assert.Equal(t, []int{2, 3, 0, 6, 99}, c.Mem)
}

func Test_run3(t *testing.T) {
	program := []int{2, 4, 4, 5, 99, 0}
	c := NewComputer(program)

	go c.Run()
	<-c.HaltChannel
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, c.Mem)
}
func Test_run4(t *testing.T) {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	c := NewComputer(program)

	go c.Run()
	<-c.HaltChannel
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, c.Mem)
}

func Test_NoRoutines(t *testing.T) {
	c := NewComputer([]int{1101, 1, 1, 5, 99, 0})
	go c.Run()
	<-c.HaltChannel
	assert.Equal(t, 2, c.Mem[5])
}

func Test_NoRoutineOutput(t *testing.T) {
	c := NewComputer([]int{1101, 1, 1, 7, 4, 7, 99, 0})
	go c.Run()
	<-c.HaltChannel
	assert.Equal(t, 2, c.Mem[7])
}
