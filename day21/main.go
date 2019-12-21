package main

import (
	"flag"
	"fmt"
	"time"

	"aoc2019/computer"
)

var gui bool //nolint:unused

func main() {
	guiPtr := flag.Bool("gui", false, "Add --gui flag to enable graphics")
	flag.Parse()
	gui = *guiPtr

	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %-10v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

func part1() interface{} {
	c := computer.NewComputer(computer.ReadMemory("day21/input.txt"))
	c.IncreaseMemory(4048)

	c.Output = make(chan computer.Msg)
	resCh := make(chan int)

	go func() {
		for {
			msg := <-c.Output
			if msg.Data > 300 {
				resCh <- msg.Data
			}
			fmt.Printf("%c", msg.Data)
		}
	}()
	go c.Run()

	c.SendLine("NOT A J")
	c.SendLine("NOT B T")
	c.SendLine("OR T J")
	c.SendLine("NOT C T")
	c.SendLine("OR T J")
	c.SendLine("AND D J")
	c.SendLine("WALK")

	return <-resCh
}

func part2() interface{} {
	c := computer.NewComputer(computer.ReadMemory("day21/input.txt"))
	c.IncreaseMemory(4048)

	c.Output = make(chan computer.Msg)
	resCh := make(chan int)

	go func() {
		for {
			msg := <-c.Output
			if msg.Data > 300 {
				resCh <- msg.Data
			}
			fmt.Printf("%c", msg.Data)
		}
	}()
	go c.Run()

	c.SendLine("NOT A J")
	c.SendLine("NOT B T")
	c.SendLine("OR T J")
	c.SendLine("NOT C T")
	c.SendLine("OR T J")

	c.SendLine("AND H T")
	c.SendLine("OR E T")
	c.SendLine("AND T J")

	c.SendLine("AND D J")

	c.SendLine("RUN")

	return <-resCh
}
