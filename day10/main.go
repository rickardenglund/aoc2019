package main

import (
	"aoc2019/inputs"
	"fmt"
	"math"
	"strings"
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

func part1() int {
	_, max := getMaxPos()
	return max
}

func getMaxPos() (pos, int) {
	lines := inputs.GetLines("day10/input.txt")
	l := ""
	for i := range lines {
		l = fmt.Sprintf("%s\n%s", l, lines[i])
	}

	//fmt.Printf("%v\n", l)
	m := getAsteroids(l)

	max := 0
	var maxPos pos
	for pos := range m {
		n := visibleAsteroids(pos, m)
		if n > max {
			max = n
			maxPos = pos
		}
	}
	return maxPos, max
}

func part2() int {
	lines := inputs.GetLines("day10/input.txt")
	l := ""
	for i := range lines {
		l = fmt.Sprintf("%s\n%s", l, lines[i])
	}

	m := getAsteroids(l)

	p, _ := getMaxPos()

	fmt.Printf("%v\n", getPolarList(p, m))
	return -1
}

func getPolarList(origo pos, m map[pos]bool) []polar {
	var positions []polar
	for other := range m {
		if eqPos(origo, other) {
			continue
		}
		pp := polar{r: dist(origo, other), angle: getAngle(origo, other)}
		positions = append(positions, pp)
	}

	return positions
}

func getAngle(origo, other pos) (angle float64) {
	dx := other.x - origo.x
	dy := other.y - origo.y
	if dx == 0 {
		if dy > 0 {
			angle = 0
		} else {
			angle = math.Pi / 2
		}
	} else {
		angle = math.Atan(dy / dx)
		if d
	}
	return angle
}

type pos struct {
	x float64
	y float64
}

type polar struct {
	r     float64
	angle float64
}

const float64EqualityThreshold = 0.001

func getAsteroids(str string) map[pos]bool {
	m := make(map[pos]bool)
	lines := strings.Split(str, "\n")
	for y := range lines {
		for x := range lines[y] {
			a := lines[y][x]
			if a == '#' {
				m[pos{float64(x), float64(y)}] = true
			}
		}
	}
	return m
}

func inLine(a, b, c pos) bool {
	dx := a.x - b.x
	dy := a.y - b.y
	k := dy / dx
	m := a.y - k*a.x

	if almostEqual(dx, 0.0) {
		return almostEqual(a.x, c.x)
	}

	cm := c.y - k*c.x
	return almostEqual(cm, m)
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func visibleAsteroids(p pos, m map[pos]bool) int {
	nVisible := 0
	for c := range m {
		if canSee(p, c, m) {
			nVisible++
			//fmt.Printf("%v\n", c)
		}
	}
	return nVisible
}

func canSee(a pos, c pos, m map[pos]bool) bool {
	if eqPos(a, c) {
		return false
	}

	for b := range m {
		if eqPos(b, a) || eqPos(b, c) {
			continue
		}
		if inLine(a, b, c) && between(a, b, c) {
			return false
		}
	}
	return true
}

func between(a, b, c pos) bool {
	return almostEqual(dist(a, b)+dist(b, c), dist(a, c))
}

func eqPos(a, b pos) bool {
	return almostEqual(a.x, b.x) && almostEqual(a.y, b.y)
}

func dist(a, b pos) float64 {
	return math.Sqrt(math.Pow(math.Abs(a.x-b.x), 2) +
		math.Pow(math.Abs(a.y-b.y), 2))
}
