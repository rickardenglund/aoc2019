package main

import (
	"aoc2019/computer"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	p1 := part1()
	timeP1 := time.Since(start)
	fmt.Printf("part1: %v in %v\n", p1, timeP1)
	start2 := time.Now()
	p2 := part2()
	timeP2 := time.Since(start2)
	fmt.Printf("part2: %v in %v\n", p2, timeP2)
}

func part1() int {
	c := computer.Computer{}
	c.ReadMemory("day5/input.txt")
	c.SetInput([]int{1})
	c.Run()
	return c.Output[len(c.Output)-1]
}

func part2() int {
	c := computer.Computer{}
	c.ReadMemory("day5/input.txt")
	c.SetInput([]int{5})
	c.Run()
	return c.Output[len(c.Output)-1]
}
