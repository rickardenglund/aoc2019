package computer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputer_Run(t *testing.T) {
	c := NewComputer([]int{3, 0, 4, 0, 99})
	go c.Run()
	c.Input <- Msg{"", 2}
	assert.Equal(t, 2, (<-c.Output).Data)
}

func TestComputer_Run5(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1003, 0, 4, 0, 99})
	values := c.getParamValues(0, 3)
	assert.Equal(t, []int{1003, 4, 1003}, values)
}

func TestComputer_Run4(t *testing.T) {
	c := NewComputer([]int{101, 1, 6, 7, 4, 7, 99, 0})
	go c.Run()
	res := <-c.Output
	<-c.Input
	assert.Equal(t, 100, res.Data)
}

func TestComputer_Run3(t *testing.T) {
	c := NewComputer([]int{1002, 4, 2, 5, 99, 0})
	go c.Run()
	<-c.Input // wait for termination
	assert.Equal(t, []int{1002, 4, 2, 5, 99, 198}, c.Mem)
}

func TestEqual(t *testing.T) {
	c := NewComputer([]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8})

	go c.Run()
	c.Input <- Msg{"test", 2}
	assert.Equal(t, []int{0}, c.outputs)
}

func TestGetParams(t *testing.T) {
	n := 12302
	assert.Equal(t, []int{3, 2, 1, 0}, getModes(n, 3))
}
