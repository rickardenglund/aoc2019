package main

import (
	"aoc2019/computer"
	"flag"
	"fmt"
	"time"
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
	resCh := make(chan int)
	go func() {
		sum := 0
		for y := 0; y < 50; y++ {
			for x := 0; x < 50; x++ {
				prog := computer.ReadMemory("day19/input.txt")
				c := computer.NewComputer(prog)
				c.Output = make(chan computer.Msg)
				c.IncreaseMemory(1024)

				go c.Run()

				c.Input <- computer.Msg{Data: x}
				c.Input <- computer.Msg{Data: y}
				d := (<-c.Output).Data
				sum += d
			}
		}
		resCh <- sum
	}()
	return <-resCh
}

type row struct {
	start, stop int
}

var path = "day19/input.txt"

func part2() interface{} {
	if gui {
		for i := 0; i < 100; i++ {
			draw(getRow(i))
		}
	}
	shipSize := 100
	return part2Solve(shipSize)
}

func part2Solve(shipSize int) int {
	for y := 0; ; y++ {
		if isOk(y, shipSize) {
			res := getRow(y + shipSize - 1)
			return res.start*10_000 + y
		} else {
			if gui {
				fmt.Printf("%-5v not OK\n", y)
			}
		}
	}
}

func isOk(y int, shipSize int) bool {
	top := getRow(y)
	r := getRow(y + shipSize - 1)
	rWidthOk := r.stop-r.start >= shipSize
	topStartOk := r.start >= top.start
	rstopOk := top.stop >= r.start+shipSize
	return rWidthOk && topStartOk && rstopOk
}

func getRow(y int) row {
	if y == 1 || y == 3 { // special case for my input
		return row{}
	}
	first := -1
	for x := 0; ; x++ {
		prog := computer.ReadMemory(path)
		c := computer.NewComputer(prog)
		c.Output = make(chan computer.Msg)
		c.IncreaseMemory(1024)

		go c.Run()

		c.Input <- computer.Msg{Data: x}
		c.Input <- computer.Msg{Data: y}
		d := (<-c.Output).Data
		//fmt.Printf("%v", d)
		if first == -1 && d > 0 {
			first = x
		}
		if first != -1 && d == 0 {
			return row{start: first, stop: x}
		}
	}
}

//nolint
func draw(r row) {
	for x := 0; x < r.stop; x++ {
		if x >= r.start && x < r.stop {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("\n")
}
