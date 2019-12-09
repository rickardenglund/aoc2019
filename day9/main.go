package main

import (
	"aoc2019/computer"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %v in %v\n", p2, time.Since(start2))
}

func part1() int {
	c := computer.NewComputerWithName("Day9", computer.ReadMemory("day9/input.txt"))
	c.Output = make(chan computer.Msg)

	go c.Run()

	c.Input <- computer.Msg{
		Sender: "P2",
		Data:   1,
	}

	return (<-c.Output).Data
}

func part2() int {
	c := computer.NewComputerWithName("Day9", computer.ReadMemory("day9/input.txt"))
	c.Output = make(chan computer.Msg)

	go c.Run()

	c.Input <- computer.Msg{
		Sender: "P2",
		Data:   2,
	}

	return (<-c.Output).Data
}
