package computer

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_run(t *testing.T) {
	program := []int{1, 0, 0, 0, 99}
	c := NewComputer(program)
	var wg sync.WaitGroup

	wg.Add(1)
	go c.RunWithWaithGroup(&wg)
	wg.Wait()

	assert.Equal(t, []int{2, 0, 0, 0, 99}, c.Mem)
}

func Test_run2(t *testing.T) {
	program := []int{2, 3, 0, 3, 99}
	c := NewComputer(program)

	var wg sync.WaitGroup
	wg.Add(1)
	go c.RunWithWaithGroup(&wg)
	wg.Wait()

	assert.Equal(t, []int{2, 3, 0, 6, 99}, c.Mem)
}

func Test_run3(t *testing.T) {
	program := []int{2, 4, 4, 5, 99, 0}
	c := NewComputer(program)

	var wg sync.WaitGroup
	wg.Add(1)
	go c.RunWithWaithGroup(&wg)
	wg.Wait()
	assert.Equal(t, []int{2, 4, 4, 5, 99, 9801}, c.Mem)
}
func Test_run4(t *testing.T) {
	program := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	c := NewComputer(program)

	var wg sync.WaitGroup
	wg.Add(1)
	go c.RunWithWaithGroup(&wg)
	wg.Wait()
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, c.Mem)
}
