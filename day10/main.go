package main

import (
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
	return -1
}

func part2() int {
	return -1
}

type pos struct {
	x float64
	y float64
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
			fmt.Printf("%v\n", c)
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
		if inLine(a, b, c) {
			return false
		}
	}
	return true
}

func eqPos(a, b pos) bool {
	return almostEqual(a.x, b.x) && almostEqual(a.y, b.y)
}

func between(a, b, x float64) bool {
	min := math.Min(a, b)
	max := math.Max(a, b)
	return min > a && x < max
}
