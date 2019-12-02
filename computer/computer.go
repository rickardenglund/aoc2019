package computer

import (
	"aoc2019/inputs"
	"log"
	"strconv"
	"strings"
)

func ReadMemory(path string)[]int {
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
	return memory
}

func Run(mem []int) []int {
	pc := 0
	for mem[pc] != 99 {
		switch mem[pc] {
		case 1:
			mem[mem[pc + 3]] = readRel(mem, pc+1) + readRel(mem, pc+2)
			pc += 4
		case 2:
			mem[mem[pc + 3]] = readRel(mem, pc+1) * readRel(mem, pc+2)
			pc += 4
		}
	}
	return mem
}

func readRel(mem []int, pos int) int {
	return mem[mem[pos]]
}