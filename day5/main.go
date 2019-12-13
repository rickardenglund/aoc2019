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
	return do(1)
}

func part2() int {
	return do(5)
}

func do(in int) int {
	c := computer.NewComputer(computer.ReadMemory("day5/input.txt"))

	go func() {
		defer func() {
			recover() //nolint:errcheck
		}()
		for {
			c.Input <- computer.Msg{Sender: "p1", Data: in}
		}
	}()

	go c.Run()

	<-c.HaltChannel
	return c.GetLastOutput()
}
