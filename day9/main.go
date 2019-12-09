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

	out := make(chan computer.Msg)
	c.Output = out
	logCh := make(chan string)
	c.LogChannel = &logCh

	go c.Run()
	c.Input <- computer.Msg{
		Sender: "P1",
		Data:   1,
	}

	go func() {
		for {
			log := <-logCh
			fmt.Printf("Log: %v\n", log)
		}
	}()

	for {
		out, more := <-out
		if !more {
			break
		}
		fmt.Printf("%v\n", out)
	}

	return c.GetLastOutput()

}

func part2() int {
	return -1
}
