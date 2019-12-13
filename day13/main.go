package main

import (
	"aoc2019/computer"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	//start := time.Now()
	//p1 := part1()
	//fmt.Printf("part1: %-10v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

type pos struct {
	x int
	y int
}

func part1() int {
	program := computer.ReadMemory("day13/input.txt")
	c := computer.NewComputer(program)
	c.IncreaseMemory(2048)
	c.Output = make(chan computer.Msg)

	//logCh := make(chan string)
	//c.LogChannel = &logCh
	//go func() {
	//	for {
	//		fmt.Printf("%s\n", <-logCh)
	//	}
	//}()

	go c.Run()
	m := map[pos]int{}
	for {
		x, more := <-c.Output
		if !more {
			break
		}
		y := <-c.Output
		t := <-c.Output
		m[pos{x.Data, y.Data}] = t.Data
	}

	res := 0
	for _, v := range m {
		if v == 2 {
			res++
		}
	}

	fmt.Printf("done\n")
	return res
}

func part2() int {
	program := computer.ReadMemory("day13/input.txt")
	program[0] = 2
	c := computer.NewComputer(program)
	c.IncreaseMemory(2048)
	c.Output = make(chan computer.Msg)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go c.RunWithWaithGroup(&wg)
	m := sync.Map{}
	score := 0
	go func() {
		for {
			x, more := <-c.Output
			if !more {
				break
			}
			y := <-c.Output
			t := <-c.Output
			if x.Data == -1 && y.Data == 0 {
				score = t.Data
			} else {
				m.Store(pos{x.Data, y.Data}, t.Data)
			}

		}
	}()

	go func() {
		dir := 1
		for {
			bp := findBall(m)
			pp := findPaddle(m)
			if bp.x+dir > pp.x {
				c.Input <- computer.Msg{Data: 1}
			} else if bp.x+dir < pp.x {
				c.Input <- computer.Msg{Data: -1}
			} else {
				c.Input <- computer.Msg{Data: 0}
			}

			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run() //nolint:errcheck
			draw(m, 0, 40, 0, 22)
			fmt.Printf("%score: v\n", score)
			fmt.Printf("bp: %v\n", bp)

		}
	}()
	wg.Wait()
	return -1
}

func findBall(m sync.Map) pos {
	res := pos{-1, -1}
	m.Range(func(k, v interface{}) bool {
		if v == 4 {
			res = k.(pos)
			return false
		}
		return true
	})

	return res
}

func findPaddle(m sync.Map) pos {
	res := pos{-1, -1}
	m.Range(func(k, v interface{}) bool {
		if v == 3 {
			res = k.(pos)
			return false
		}
		return true
	})

	return res
}

func draw(hull sync.Map, minX, maxX, minY, maxY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			v, _ := hull.Load(pos{x, y})
			switch v {
			case 0:
				fmt.Printf("░░")
			case 1:
				fmt.Printf("██")
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
