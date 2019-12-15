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
	case north:
		return south
	case south:
		return north
	case west:
		return east
	case east:
		return west
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
	oxygen        = 3
)

func part1() int {
	program := computer.ReadMemory("day15/input.txt")
	c := computer.NewComputer(program)
	resultCh := make(chan int)
	exploreCompletedCh := make(chan bool)
	c.Output = make(chan computer.Msg)

	go c.Run()

	go func() {
		m := map[pos]int{}
		path := []pos{{0, 0}}

		visit(&c, resultCh, m, path)
		exploreCompletedCh <- true
	}()
	return <-resultCh
}

func part2() int {
	program := computer.ReadMemory("day15/input.txt")
	c := computer.NewComputer(program)
	resultCh := make(chan int)
	result2Ch := make(chan int)
	c.Output = make(chan computer.Msg)

	go c.Run()

	go func() {
		m := map[pos]int{}
		path := []pos{{0, 0}}

		visit(&c, resultCh, m, path)
		if os.Getenv("GUI") == "true" {
			drawSquare(m, drawSize)
		}

		// incorrect solution but correct answer
		time := timeToFill(m)
		result2Ch <- time
	}()
	<-resultCh

	return <-result2Ch
}

func timeToFill(m map[pos]int) int {
	gp := getGeneratorPos(m)

	res := flow(m, []pos{gp})
	return res
}

func flow(m map[pos]int, toVisit []pos) int {

	i := 0
	for ; len(toVisit) > 0; i++ {
		//time.Sleep(200 * time.Millisecond)
		var newList []pos
		for _, p := range toVisit {
			m[p] = oxygen
			newList = append(newList, p)
		}

		toVisit = []pos{}
		for i := range newList {
			toVisit = append(toVisit, getOpenNeighbours(m, newList[i])...)
		}
	}
	return i - 1
}

func getOpenNeighbours(m map[pos]int, p pos) (res []pos) {
	for i := 1; i <= 4; i++ {
		if m[p.step(i)] == open {
			res = append(res, p.step(i))
		}
	}
	return
}

func getGeneratorPos(m map[pos]int) pos {
	for p, v := range m {
		if v == oxygenStation {
			return p
		}
	}
	log.Fatalln("Station not found")
	return pos{}
}

const drawSize = 30

func visit(c *computer.Computer, resultCh chan int, m map[pos]int, path []pos) {
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
		c.Input <- computer.Msg{Data: opposite(dir)}
		assert(open, (<-c.Output).Data)
		resultCh <- len(path)
	}
	return resp.Data
}

func assert(expected interface{}, value interface{}) {
	if expected != value {
		log.Fatalln(fmt.Sprintf("%v ist not equal to %v", value, expected))
	}

}
func drawSquare(m map[pos]int, size int) {
	draw(m, -size, size, -size, size)
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
			case wall:
				fmt.Printf("██")
			case open:
				fmt.Printf("░░")
			case oxygenStation:
				fmt.Printf("##")
			case oxygen:
				fmt.Printf("..")
			case 4:
				fmt.Printf("<>")
			}
		}
		fmt.Printf("\n")
	}
}
