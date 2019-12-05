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
	assert.Equal(t, []int{3}, c.Output)
}

func TestComputer_Run2(t *testing.T) {
	c := Computer{}
	c.setMem([]int{101003,0,4,0,99})
	instruction, params := c.readInstruction(0, c.mem[1:2])
	fmt.Printf("op: %v params: %v\n", instruction, params)
}

func TestComputer_Run5(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1003,0,4,0,99})
	instruction, params := c.readInstruction(0, c.mem[1:4])
	fmt.Printf("op: %v params: %v\n", instruction, params)
}

func TestComputer_Run4(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1002,4,10,4,33,99})
	c.Run()
	fmt.Printf("output %v\n", c.Output)

}

func TestComputer_Run3(t *testing.T) {
	c := Computer{}
	c.setMem([]int{1002,4,10,4,33,99})
	c.Run()
}