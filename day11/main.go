package main

import (
	"aoc2019/computer"
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %-10v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

type pos struct {
	x int
	y int
}

type Painter struct {
	Pos       pos
	Computer  computer.Computer
	Direction int
}

func part1() int {
	program := computer.ReadMemory("day11/input.txt")
	c := computer.NewComputer(program)
	c.Output = make(chan computer.Msg)
	c.IncreaseMemory(1024)

	painter := Painter{
		Pos:       pos{0, 0},
		Computer:  c,
		Direction: 0, // up
	}
	hull := make(map[pos]int)
	hull[pos{0, 0}] = 0

	go c.Run()

	for {
		open := tryWrite(c, hull, painter)
		if !open {
			break
		}
		clr := <-c.Output
		turn := <-c.Output
		hull[painter.Pos] = clr.Data
		painter.move(turn.Data)
	}
	return len(hull)
}

func tryWrite(c computer.Computer, hull map[pos]int, painter Painter) (res bool) {
	defer func() {
		err := recover()
		res = err == nil
	}()
	c.Input <- computer.Msg{Data: hull[painter.Pos]}
	return true
}

func (p *Painter) move(turn int) {
	if turn == 0 {
		p.Direction = (p.Direction + 1) % 4
	} else {
		if p.Direction == 0 {
			p.Direction = 3
		} else {
			p.Direction--
		}
	}

	switch p.Direction {
	case 0:
		p.Pos.y++
	case 1:
		p.Pos.x++
	case 2:
		p.Pos.y--
	case 3:
		p.Pos.x--
	default:
		log.Fatal("invalid direction")
	}
}

func part2() int {
	program := computer.ReadMemory("day11/input.txt")
	c := computer.NewComputer(program)
	c.Output = make(chan computer.Msg)
	c.IncreaseMemory(1024)

	painter := Painter{
		Pos:       pos{0, 0},
		Computer:  c,
		Direction: 0, // up
	}
	hull := make(map[pos]int)
	hull[pos{0, 0}] = 1

	go c.Run()

	minX := math.MaxInt64
	maxX := math.MinInt64
	minY := math.MaxInt64
	maxY := math.MinInt64
	for {
		open := tryWrite(c, hull, painter)
		if !open {
			break
		}
		clr := <-c.Output
		turn := <-c.Output
		hull[painter.Pos] = clr.Data
		painter.move(turn.Data)

		if painter.Pos.x < minX {
			minX = painter.Pos.x
		}
		if painter.Pos.x > maxX {
			maxX = painter.Pos.x
		}
		if painter.Pos.y < minY {
			minY = painter.Pos.y
		}
		if painter.Pos.y > maxY {
			maxY = painter.Pos.y
		}
	}

	draw(hull, minX, maxX, minY, maxY)
	return -1
}

func draw(hull map[pos]int, minX, maxX, minY, maxY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if hull[pos{x, y}] == 1 {
				fmt.Printf("██")
			} else {
				fmt.Printf("░░")
			}
		}
		fmt.Printf("\n")
	}
}
