package computer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComputer_Run(t *testing.T) {
	c := Computer{}
	c.setMem([]int{3,0,4,0,99})
	c.Run()
	assert.Equal(t, []int{0}, c.Output)
}

func TestComputer_Run2(t *testing.T) {
	c := Computer{}
	c.setMem([]int{101003,0,4,0,99})
	instruction, params := c.readInstruction(0, c.Mem[1:2])
	fmt.Printf("op: %v params: %v\n", instruction, params)
}

func TestComputer_Run5(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1003,0,4,0,99})
	instruction, params := c.readInstruction(0, c.Mem[1:4])
	fmt.Printf("op: %v params: %v\n", instruction, params)
}

func TestComputer_Run4(t *testing.T) {
	c := Computer{}
	c.setMem([]int{101, 1, 6, 7, 4, 7, 99,  0})
	c.Run()
	fmt.Printf("output %v\n", c.Output)
}

func TestComputer_Run3(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1002,4,2,5,99, 0})
	c.Run()
	assert.Equal(t, []int{1002, 4, 2, 5, 99, 198}, c.Mem)
}
