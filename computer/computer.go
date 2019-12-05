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
		opcode, _ := c.readInstruction(pc, nil)
		switch opcode {
		case 1: // add
			_, params := c.readInstruction(pc, c.Mem[pc+1:pc+4])
			c.Mem[c.Mem[pc+3]] = params[0] + params[1]
			pc += 4
		case 2: // mul
			_, params := c.readInstruction(pc, c.Mem[pc+1:pc+4])
			c.Mem[c.Mem[pc+3]] = params[0] * params[1]
			pc += 4
		case 3: // input
			c.Mem[c.Mem[pc+1]] = c.input
			pc += 2
		case 4: // output
			_, params := c.readInstruction(pc, c.Mem[pc+1:pc+2])
			c.Output = append(c.Output, params[0])
			pc += 2
		case 5: // jump if true
			_, params := c.readInstruction(pc, c.Mem[pc+1:pc+3])
			if params[0] != 0 {
				pc = params[1]
			} else {
				pc += 3
			}
		case 6: // jump if false
			_, params := c.readInstruction(pc, c.Mem[pc+1:pc+3])
			if params[0] == 0 {
				pc = params[1]
			} else {
				pc += 3
			}
		case 7: // less than
			_, params := c.readInstruction(pc, c.Mem[pc+1:pc+4])
			if params[0] < params[1] {
				c.Mem[c.Mem[pc+3]] = 1
			} else {
				c.Mem[c.Mem[pc+3]] = 0
			}
			pc += 4
		case 8: // equal
			_, params := c.readInstruction(pc, c.Mem[pc+1:pc+4])
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
			//log.Fatalf("Unknown opcode: %v pc: %v\n %v\n", opcode, pc, c.Mem[240:245])
		}
	}
}

func (c *Computer) readParam(pos int) int {
	return c.Mem[c.Mem[pos]]
}

func (c *Computer) readInstruction(pc int, inParams []int) (opCode int, params []int) {
	var modeList []int
	var err error
	str := strconv.Itoa(c.Mem[pc])

	if len(str) < 2 {
		opCode, err = strconv.Atoi(str)
		checkError(err)
	} else {

		opCode, err = strconv.Atoi(str[len(str)-2:])
		checkError(err)

		for i := len(str) - 3; i >= 0; i-- {
			mode := 0
			if str[i] == '1' {
				mode = 1
			}
			modeList = append(modeList, mode)
		}
	}

	for i := 0; i < len(inParams); i++ {
		var val int
		if i >= len(modeList) {
			val = c.Mem[inParams[i]]
		} else if modeList[i] == 1 {
			val = inParams[i]
		} else if modeList[i] == 0 {
			val = c.Mem[inParams[i]]
		}

		params = append(params, val)
	}

	return opCode, params
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Computer) SetInput(input int) {
	c.input = input
}

func (c *Computer) setMem(ints []int) {
	c.Mem = ints

}
