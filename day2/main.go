package main

import (
	"aoc2019/computer"
	"fmt"
)

func main() {
	fmt.Printf("part1: %v\n", part1())
	fmt.Printf("part1: %v\n", part2())
}

func part1() int {
	c := computer.Computer{}
	c.ReadMemory("day2/input.txt")

	c.Mem[1] = 12
	c.Mem[2] = 2
	c.Run()
	return c.Mem[0]
}

func part2() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			c := computer.Computer{}
			c.ReadMemory("day2/input.txt")
			c.Mem[1] = noun
			c.Mem[2] = verb
			c.Run()
			if c.Mem[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return -1
}
