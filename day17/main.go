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
			m[position.Pos{X: x, Y: y}] = (data).Data
			if gui {
				fmt.Printf("%c", data.Data)
			}
			if data.Data == newline {
				y++
				x = -1
			}
		}
	}()

	time.Sleep(200 * time.Millisecond)
	sum := getAlignment(m)

	return sum
}

func part2() int {
	program := computer.ReadMemory("day17/input.txt")
	program[0] = 2
	c := computer.NewComputer(program)
	c.Output = make(chan computer.Msg)
	c.IncreaseMemory(4048)
	resCh := make(chan int)
	go c.Run()

	go func() {
		// Do some manual processing to find the path
		c.SendLine("A,B,A,B,C,B,C,A,C,C")

		c.SendLine("R,12,L,10,L,10")
		c.SendLine("L,6,L,12,R,12,L,4")
		c.SendLine("L,12,R,12,L,6")
		c.SendLine("n")
	}()

	go func() {
		for {
			msg := <-c.Output
			if msg.Data < 255 {
				if gui {
					fmt.Printf("%c", msg.Data)
				}
			} else {
				resCh <- msg.Data
			}
		}
	}()

	return <-resCh
}

func getAlignment(m map[position.Pos]int) int {
	intersections := map[position.Pos]bool{}
	for k := range m {
		if m[k] == scaffold &&
			m[position.Pos{X: k.X + 1, Y: k.Y}] == scaffold &&
			m[position.Pos{X: k.X - 1, Y: k.Y}] == scaffold &&
			m[position.Pos{X: k.X, Y: k.Y + 1}] == scaffold &&
			m[position.Pos{X: k.X, Y: k.Y - 1}] == scaffold {
			intersections[k] = true
		}
	}
	sum := 0
	for k := range intersections {
		m[k] = intersection
		sum += k.X * k.Y
	}

	if gui {
		fmt.Printf("Intersections: %v\n", intersections)
		draw(m, 0, 40, 0, 75)
	}
	return sum
}
func draw(m map[position.Pos]int, minX, maxX, minY, maxY int) {
	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run() //nolint:errcheck
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			v, ok := m[position.Pos{X: x, Y: y}]
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
