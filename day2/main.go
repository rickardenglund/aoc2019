package main

import (
	"aoc2019/inputs"
	"fmt"
	"log"
	"strconv"
	"strings"
)


func main() {
	fmt.Printf("part1: %v\n", part1())
	fmt.Printf("part1: %v\n", part2())
}
// guess: 394702
func readMemory()[]int {
	lines := inputs.GetLines("day2/input.txt")
	ints := strings.Split(lines[0], ",")
	var program []int = make([]int, len(ints))
	for i, code := range ints {
		n, err := strconv.Atoi(code)
		if err != nil {
			log.Fatal(err)
		}
		program[i] = n
	}
	return program
}

func part1() int{

	program := readMemory()

	program[1] = 12
	program[2] = 2
	result := run(program)
	fmt.Printf("%v\n", result)
	return result[0]
}

func part2() int {

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb< 100; verb++ {
			program := readMemory()
			program[1] = noun
			program[2] = verb
			res := run(program)
			if res[0] == 19690720 {
				return 100 * noun + verb
			}

		}
	}
	return -1
}

func run(program []int) []int {
	pc := 0
	for program[pc] != 99 {
		switch program[pc] {
		case 1:
			program[program[pc + 3]] = read(program, program[pc+1]) + read(program, program[pc+2])
		case 2:
			program[program[pc + 3]] = read(program, program[pc+1]) * read(program,program[pc+2])
		}
		pc += 4
	}
	return program
}

func read(program []int, pos int)int {
	return program[pos]
}

func algo(n int) int {

	return -1
}
