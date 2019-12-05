package computer

import (
	"log"
	"strconv"
	"strings"

	"aoc2019/inputs"
)

type Computer struct {
	input  int
	Mem    []int
	Output []int
}

func (c *Computer) ReadMemory(path string) {
	lines := inputs.GetLines(path)
	ints := strings.Split(lines[0], ",")
	var memory []int = make([]int, len(ints))
	for i, code := range ints {
		n, err := strconv.Atoi(code)
		if err != nil {
			log.Fatal(err)
		}
		memory[i] = n
	}
	c.Mem = memory
}

func (c *Computer) Run() {
	pc := 0
loop:
	for {
		opcode := c.getOpcode(pc)
		switch opcode {
		case 1: // add
			params := c.getParamValues(pc, 2)
			c.Mem[c.Mem[pc+3]] = params[0] + params[1]
			pc += 4
		case 2: // mul
			params := c.getParamValues(pc, 2)
			c.Mem[c.Mem[pc+3]] = params[0] * params[1]
			pc += 4
		case 3: // input
			c.Mem[c.Mem[pc+1]] = c.input
			pc += 2
		case 4: // output
			params := c.getParamValues(pc, 1)
			c.Output = append(c.Output, params[0])
			pc += 2
		case 5: // jump if true
			params := c.getParamValues(pc, 2)
			if params[0] != 0 {
				pc = params[1]
			} else {
				pc += 3
			}
		case 6: // jump if false
			params := c.getParamValues(pc, 2)
			if params[0] == 0 {
				pc = params[1]
			} else {
				pc += 3
			}
		case 7: // less than
			params := c.getParamValues(pc, 2)
			if params[0] < params[1] {
				c.Mem[c.Mem[pc+3]] = 1
			} else {
				c.Mem[c.Mem[pc+3]] = 0
			}
			pc += 4
		case 8: // equal
			params := c.getParamValues(pc, 2)
			if params[0] == params[1] {
				c.Mem[c.Mem[pc+3]] = 1
			} else {
				c.Mem[c.Mem[pc+3]] = 0
			}
			pc += 4
		case 99:
			break loop
		default:
			log.Fatalf("Unknown opcode: %v pc: %v\n %v\n", opcode, pc, c.Mem)
		}
	}
}

func (c *Computer) readParam(pos int) int {
	return c.Mem[c.Mem[pos]]
}

func (c *Computer) getOpcode(pc int) int {
	return c.Mem[pc] % 100
}

func (c *Computer) getParamValues(pc int, nParams int) (values []int) {
	inParams := c.Mem[pc + 1: pc + 1 + nParams]
	modeList := getModes(c.Mem[pc], nParams)

	for i := 0; i < nParams; i++ {
		var val int
		if modeList[i] == 1 {
			val = inParams[i]
		} else if modeList[i] == 0 {
			val = c.Mem[inParams[i]]
		}

		values = append(values, val)
	}

	return values
}

func getModes(op, nParams int) []int {
	number := op / 100

	modes := []int{number % 10}
	var divisor = 10
	for i := 0; i < nParams; i++ {
		modes = append(modes, (number/divisor)%divisor)
		divisor *= 10
	}
	return modes
}

func (c *Computer) SetInput(input int) {
	c.input = input
}

func (c *Computer) setMem(ints []int) {
	c.Mem = ints
}
