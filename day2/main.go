package main

import (
	"aoc2019/computer"
	"fmt"
)


func main() {
	fmt.Printf("part1: %v\n", part1())
	fmt.Printf("part1: %v\n", part2())
}


func part1() int{
	memory := computer.ReadMemory("day2/input.txt")

	memory[1] = 12
	memory[2] = 2
	result := computer.Run(memory)
	return result[0]
}

func part2() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb< 100; verb++ {
			program := computer.ReadMemory("day2/input.txt")
			program[1] = noun
			program[2] = verb
			res := computer.Run(program)
			if res[0] == 19690720 {
				return 100 * noun + verb
			}
		}
	}
	return -1
}
