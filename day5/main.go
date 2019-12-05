package main

import (
	"aoc2019/computer"
	"fmt"
)

func main() {
	fmt.Printf("part1: %v\n", part1())
	fmt.Printf("part1: %v\n", part2())
}

// guess: 7441400
func part1() int{
	c := computer.Computer{}
	c.ReadMemory("day5/input.txt")
	c.SetInput(1)
	c.Run()
	return c.Output[len(c.Output)-1]
}

func part2() int {
	c := computer.Computer{}
	c.ReadMemory("day5/input.txt")
	c.SetInput(5)
	c.Run()
	return c.Output[len(c.Output)-1]
}
