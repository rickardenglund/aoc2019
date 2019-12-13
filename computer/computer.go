package computer

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"aoc2019/inputs"
)

type Computer struct {
	Mem          []int
	relativeBase int
	Input        chan Msg
	Output       chan Msg
	Outputs      []int
	name         string
	LogChannel   *chan string
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

func (c *Computer) RunWithWaithGroup(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Run()
	c.log("Halting")
}

func (c *Computer) Run() {
	_ = map[int]string{
		1: "add",
		2: "mul",
		3: "read",
		4: "write",
		5: "jumpIfTrue",
		6: "JumpIfFalse",
		7: "Less than",
		8: "equal",
		9: "AdjBase",
	}

	pc := 0
loop:
	for {
		opcode := c.getOpcode(pc)
		//c.log(fmt.Sprintf("pc: %v, %v: %v", pc, names[opcode], c.Mem[pc:pc+5]))
		switch opcode {
		case 1: // add
			params := c.getParamValues(pc, 2)
			outputAddress := c.getOutputAddress(pc, 3)
			c.Mem[outputAddress] = params[0] + params[1]
			pc += 4
		case 2: // mul
			params := c.getParamValues(pc, 2)
			outputAddress := c.getOutputAddress(pc, 3)
			c.Mem[outputAddress] = params[0] * params[1]
			pc += 4
		case 3: // Input
			c.log(fmt.Sprintf("reading input"))
			msg := <-c.Input
			outputAddress := c.getOutputAddress(pc, 1)
			c.Mem[outputAddress] = msg.Data
			c.log(fmt.Sprintf("got: %v from %v", msg.Data, msg.Sender))
			pc += 2
		case 4: // Output
			params := c.getParamValues(pc, 1)
			c.Outputs = append(c.Outputs, params[0])
			c.log(fmt.Sprintf("outputs: %v", params[0]))
			c.trySend(params[0])
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
			outputAddress := c.getOutputAddress(pc, 3)
			if params[0] < params[1] {
				c.Mem[outputAddress] = 1
			} else {
				c.Mem[outputAddress] = 0
			}
			pc += 4
		case 8: // equal
			params := c.getParamValues(pc, 2)
			outputAddress := c.getOutputAddress(pc, 3)
			if params[0] == params[1] {
				c.Mem[outputAddress] = 1
			} else {
				c.Mem[outputAddress] = 0
			}
			pc += 4
		case 9: // adjust relative base
			params := c.getParamValues(pc, 1)
			c.relativeBase += params[0]
			pc += 2
		case 99:
			if c.Input != nil {
				close(c.Input)
			}
			if c.Output != nil {
				close(c.Output)
			}
			c.log("Stop")
			break loop
		default:
			log.Fatalf("Unknown opcode: %v pc: %v\n %v\n", opcode, pc, c.Mem)
		}
	}
}

func (c *Computer) trySend(data int) {
	defer func() {
		recover() // nolint: errcheck
	}()

	msg := Msg{c.name, data}
	if c.Output != nil {
		c.Output <- msg
	}
}

func (c *Computer) log(msg string) {
	if c.LogChannel != nil {
		*c.LogChannel <- fmt.Sprintf("%s : %s", c.name, msg)
	}
}

func (c *Computer) getOpcode(pc int) int {
	return c.Mem[pc] % 100
}

func (c *Computer) getOutputAddress(pc int, outputPos int) (address int) {
	modeList := getModes(c.Mem[pc], outputPos)
	if modeList[outputPos-1] == 2 {
		return c.relativeBase + c.Mem[pc+outputPos] // relative mode
	}
	return c.Mem[pc+outputPos] //position mode
}

func (c *Computer) getParamValues(pc int, nParams int) (values []int) {
	inParams := c.Mem[pc+1 : pc+1+nParams]
	modeList := getModes(c.Mem[pc], nParams)

	for i := 0; i < nParams; i++ {
		var val int
		switch modeList[i] {
		case 1:
			val = inParams[i]
		case 2:
			val = c.Mem[inParams[i]+c.relativeBase]
		case 0:
			val = c.Mem[inParams[i]]
		default:
			log.Fatal("Invalid param mode")
		}

		values = append(values, val)
	}

	return values
}

func getModes(op, nParams int) []int {
	number := op / 100

	modes := []int{number % 10}
	var divisor = 10
	for i := 1; i < nParams; i++ {
		modes = append(modes, (number/divisor)%divisor)
		divisor *= 10
	}
	return modes
}

type Msg struct {
	Sender string
	Data   int
}

func NewComputerWithName(name string, mem []int) Computer {
	memCopy := make([]int, len(mem))

	copy(memCopy, mem)

	return Computer{
		Mem:   memCopy,
		Input: make(chan Msg),
		//Output: make(chan Msg),
		name: name,
	}
}

func NewComputer(mem []int) Computer {
	return NewComputerWithName("Name", mem)
}

func (c *Computer) setMem(ints []int) {
	c.Mem = ints
}

func (c *Computer) IncreaseMemory(size int) {
	c.Mem = append(c.Mem, make([]int, size)...)
}

func (c *Computer) GetLastOutput() int {
	return c.Outputs[len(c.Outputs)-1]
}

func ReadMemory(path string) []int {
	c := Computer{}
	c.ReadMemory(path)
	return c.Mem
}
