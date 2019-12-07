package computer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputer_Run(t *testing.T) {
	c := Computer{Input: []int{0}}
	c.setMem([]int{3, 0, 4, 0, 99})
	c.Run()
	assert.Equal(t, []int{0}, c.Output)
}

func TestComputer_Run2(t *testing.T) {
	c := Computer{}
	c.setMem([]int{101003, 0, 4, 0, 99})
	//instruction, params := c.getParamValues(0, c.Mem[1:2])
	//fmt.Printf("op: %v params: %v\n", instruction, params)
}

func TestComputer_Run5(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1003, 0, 4, 0, 99})
	//instruction, params := c.getParamValues(0, c.Mem[1:4])
	//fmt.Printf("op: %v params: %v\n", instruction, params)
}

func TestComputer_Run4(t *testing.T) {
	c := Computer{}
	c.setMem([]int{101, 1, 6, 7, 4, 7, 99, 0})
	c.Run()
	fmt.Printf("Output %v\n", c.Output)
}

func TestComputer_Run3(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1002, 4, 2, 5, 99, 0})
	c.Run()
	assert.Equal(t, []int{1002, 4, 2, 5, 99, 198}, c.Mem)
}

func TestEqual(t *testing.T) {
	c := Computer{
		Mem:   []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		Input: []int{2},
	}
	c.Run()
	assert.Equal(t, []int{0}, c.Output)
}

func TestGetParams(t *testing.T) {
	n := 12302
	assert.Equal(t, []int{3, 2, 1, 0}, getModes(n, 3))
}
