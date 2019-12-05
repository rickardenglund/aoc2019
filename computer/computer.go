package computer

import (
	"aoc2019/inputs"
	"log"
	"strconv"
	"strings"
)

type Computer struct {
	input  int
	mem    []int
	Output []int
}

func (c *Computer) ReadMemory(path string){
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
	c.mem = memory
}

func (c *Computer) Run() {
	pc := 0
	for c.mem[pc] != 99 {
		opcode, _ := c.readInstruction(pc, nil)
		switch opcode {
		case 1: // add
			_, params := c.readInstruction(pc, c.mem[pc+1:pc+4])
			c.mem[c.mem[pc+3]] = params[0] + params[1]
			pc += 4
		case 2: // mul
			_, params := c.readInstruction(pc, c.mem[pc+1:pc+4])
			c.mem[c.mem[pc+3]] = params[0] * params[1]
			pc += 4
		case 3: // read
			c.mem[c.mem[pc+1]] = c.input
			pc += 2
		case 4: // write
			_, params := c.readInstruction(pc, c.mem[pc+1:pc+2])
			c.mem[c.mem[pc+1]] = c.input
			c.Output = append(c.Output, params[0])
			pc+=2
		default:
			log.Fatalf("Unknown opcode: %v pc: %v\n", opcode, pc)
		}
	}
}

func (c *Computer) readParam(pos int) int {
	return c.mem[c.mem[pos]]
}

func (c *Computer) readInstruction(pc int, inParams []int) (opCode int, params []int) {
	var modeList []int
	var err error
	str := strconv.Itoa(c.mem[pc])

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
		if i >=len(modeList) {
			val = c.mem[inParams[i]]
		} else if modeList[i] == 1 {
			val = inParams[i]
		} else if modeList[i] == 0 {
			val = c.mem[inParams[i]]
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
	c.mem = ints

}