package main

import (
	"aoc2019/computer"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

//const testName = "test"

func Test_computerPrintItself(t *testing.T) {
	program := []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	c := computer.NewComputer(program)

	c.Run()
	assert.Equal(t, program, c.Outputs)
}

func Test_computerCalcLargeNum(t *testing.T) {
	program := []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	c := computer.NewComputer(program)

	c.Run()

	length := len(strconv.Itoa(c.GetLastOutput()))
	assert.Equal(t, 16, length)
}

func Test_computerOutputLargeNum2(t *testing.T) {
	program := []int{104, 1125899906842624, 99}
	c := computer.NewComputer(program)

	c.Run()

	assert.Equal(t, 1125899906842624, c.GetLastOutput())
}

func Test_computerRead(t *testing.T) {
	program := []int{9, 50, 203, 0, 99}
	c := computer.NewComputer(program)

	go c.Run()
	c.Input <- computer.Msg{Data: 55}

	assert.Equal(t, 55, c.Mem[50])
}

func Test_computerAdj(t *testing.T) {
	program := []int{109, 5, 9, 1, 209, -7, 2101, 1, 0, 50, 99, 8}
	c := computer.NewComputer(program)

	c.Run()

	assert.Equal(t, 9, c.Mem[50])
}
