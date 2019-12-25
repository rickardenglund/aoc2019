package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
	p := computer.ReadMemory("day25/input.txt")
	c := computer.NewComputer(p)
	c.Output = make(chan computer.Msg)
	c.IncreaseMemory(8000)
	resCh := make(chan string)

	go c.Run()
	go func() {
		for {
			msg := <-c.Output
			fmt.Printf("%c", msg.Data)
		}
	}()

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)
			c.SendLine(line)
		}

	}()
	return <-resCh
}

func part2() interface{} {
	return "-"
}
