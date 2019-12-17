package main

import (
	"aoc2019/computer"
	"aoc2019/position"
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

const (
	newline      = 10
	scaffold     = 35
	intersection = 65
	nothing      = 46
)

func part1() int {
	program := computer.ReadMemory("day17/input.txt")
	c := computer.NewComputer(program)
	c.Output = make(chan computer.Msg)
	c.IncreaseMemory(4048)

	go c.Run()

	m := map[position.Pos]int{}

	go func() {
		y := 0
		for x := 0; true; x++ {
			data := <-c.Output
			m[position.Pos{x, y}] = (data).Data
			if gui {
				fmt.Printf("%c", data.Data)
			}
			if data.Data == newline {
				y++
				x = -1
			}
		}
	}()

	time.Sleep(1 * time.Second)
	sum := getAlignment(m)

	return sum
}

func getAlignment(m map[position.Pos]int) int {
	intersections := map[position.Pos]bool{}
	for k := range m {
		if m[k] == scaffold &&
			m[position.Pos{k.X + 1, k.Y}] == scaffold &&
			m[position.Pos{k.X - 1, k.Y}] == scaffold &&
			m[position.Pos{k.X, k.Y + 1}] == scaffold &&
			m[position.Pos{k.X, k.Y - 1}] == scaffold {
			intersections[k] = true
		}
	}
	sum := 0
	for k := range intersections {
		m[k] = intersection
		sum += k.X * k.Y
	}

	fmt.Printf("Intersections: %v\n", intersections)
	if gui {
		draw(m, 0, 40, 0, 75)
	}
	return sum
}

func part2() int {
	return -1
}
func draw(m map[position.Pos]int, minX, maxX, minY, maxY int) {
	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run() //nolint:errcheck
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			v, ok := m[position.Pos{x, y}]
			if !ok {
				fmt.Printf("  ")
				continue
			}
			switch v {
			case scaffold:
				fmt.Printf("██")
			case nothing:
				fmt.Printf("░░")
			case newline:
				fmt.Printf("n")
			default:
				fmt.Printf("%c%c", v, v)
			}
		}
		fmt.Printf("\n")
	}
}
