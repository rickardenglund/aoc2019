package main

import (
	"aoc2019/computer"
	"fmt"
	"log"
	"os"
	"os/exec"
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

func (p *pos) step(direction int) pos {
	switch direction {
	case 1:
		return pos{p.x, p.y - 1}
	case 2:
		return pos{p.x, p.y + 1}
	case 3:
		return pos{p.x - 1, p.y}
	case 4:
		return pos{p.x + 1, p.y}
	default:
		log.Fatalln("Unknown direction")
		return pos{0, 0}
	}
}
func opposite(dir int) int {
	switch dir {
	case 1:
		return 2
	case 2:
		return 1
	case 3:
		return 4
	case 4:
		return 3
	default:
		log.Fatalln("Invalid direction")
		return -1
	}
}

const (
	north = 1
	south = 2
	west  = 3
	east  = 4

	wall          = 0
	open          = 1
	oxygenStation = 2
)

func part1() int {
	program := computer.ReadMemory("day15/input.txt")
	c := computer.NewComputer(program)
	resultCh := make(chan int)
	c.Output = make(chan computer.Msg)

	go c.Run()

	go func() {
		m := map[pos]int{}
		path := []pos{{0, 0}}

		visit(&c, resultCh, m, path)
	}()
	return <-resultCh
}

const drawSize = 30

func visit(c *computer.Computer, resultCh chan int, m map[pos]int, path []pos) {
	draw(m, -drawSize, drawSize, -drawSize, drawSize)
	cur := path[len(path)-1]
	dirToTry := []int{}
	for dir := 1; dir <= 4; dir++ {
		_, ok := m[cur.step(dir)]
		if ok {
			continue
		}
		res := testDirection(cur, dir, c, resultCh, path)
		m[cur.step(dir)] = res
		if res == open {
			dirToTry = append(dirToTry, dir)
		}
	}

	for _, dir := range dirToTry {
		newPos := path[len(path)-1].step(dir)
		c.Input <- computer.Msg{Data: dir}
		assert(open, (<-c.Output).Data)
		visit(c, resultCh, m, append(path, newPos))
		c.Input <- computer.Msg{Data: opposite(dir)}
		assert(open, (<-c.Output).Data)
	}

}

func testDirection(cur pos, dir int, c *computer.Computer, resultCh chan int, path []pos) int {
	c.Input <- computer.Msg{Data: dir}
	resp := <-c.Output
	switch resp.Data {
	case wall:
	case open:
		c.Input <- computer.Msg{Data: opposite(dir)}
		assert(open, (<-c.Output).Data)
	case oxygenStation:
		resultCh <- len(path)
	}
	return resp.Data
}

func assert(expected interface{}, value interface{}) {
	if expected != value {
		log.Fatalln(fmt.Sprintf("%v ist not equal to %v", value, expected))
	}

}

func part2() int {
	return -1
}
func draw(m map[pos]int, minX, maxX, minY, maxY int) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run() //nolint:errcheck
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			v, ok := m[pos{x, y}]
			if !ok {
				fmt.Printf("  ")
				continue
			}
			switch v {
			case 0:
				fmt.Printf("██")
			case 1:
				fmt.Printf("░░")
			case 2:
				fmt.Printf("##")
			case 3:
				fmt.Printf("XX")
			case 4:
				fmt.Printf("<>")
			}
		}
		fmt.Printf("\n")
	}
}
