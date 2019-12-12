package main

import (
	"aoc2019/inputs"
	"fmt"
	"log"
	"math"
	"sort"
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
	input := inputs.GetLine("day10/input.txt")
	m := getAsteroids(input)
	_, max := getMaxPos(m)
	return max
}

func getMaxPos(m map[pos]bool) (pos, int) {
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
	input := inputs.GetLine("day10/input.txt")
	m := getAsteroids(input)

	p, _ := getMaxPos(m)
	fmt.Printf("maxpos: %v\n", p)

	list := getPolarList(p, m)
	explosion200 := getNExplosion(200, list)

	fmt.Printf("200 pos: %v\n", explosion200)
	return explosion200.x*100 + explosion200.y
}

func getNExplosion(n int, pl []polar) pos {
	if n > len(pl) || n <= 0 {
		log.Fatal("not implemented")
	}
	cur := pl[0]
	nExploded := 1
	for i := 1; i < len(pl); i++ {
		if almostEqual(cur.angle, pl[i].angle) {
			continue
		}

		if nExploded == n {
			break
		}
		cur = pl[i]
		nExploded++

		if i+1 == len(pl) {
			log.Fatal("Too many")
		}
	}

	return pos{cur.x, cur.y}
}

func getPolarList(origo pos, m map[pos]bool) []polar {
	var positions []polar
	for other := range m {
		if origo == other {
			continue
		}
		pp := polar{
			r:     dist(origo, other),
			angle: getAngle(origo, other),
			x:     other.x,
			y:     other.y,
		}
		positions = append(positions, pp)
	}

	sort.Slice(positions, func(i, j int) bool {
		if !almostEqual(positions[i].angle, positions[j].angle) {
			return positions[i].angle < positions[j].angle
		}
		return positions[i].r < positions[j].r
	})

	return positions
}

func getAngle(origo, other pos) (angle float64) {
	dx := other.x - origo.x
	dy := other.y - origo.y

	return angleOf(pos{dx, dy})
}

func angleOf(p pos) float64 {
	//angle := math.Mod(math.Atan2(float64(p.y), float64(-p.x))-math.Pi/+2*math.Pi, 2*math.Pi)
	angle := math.Mod(4.5*math.Pi+math.Atan2(float64(p.y), float64(p.x)), 2*math.Pi)
	return angle
}

type pos struct {
	x int
	y int
}

type polar struct {
	r     float64
	angle float64
	x     int
	y     int
}

const float64EqualityThreshold = 0.00000001

func getAsteroids(str string) map[pos]bool {
	m := make(map[pos]bool)
	lines := strings.Split(str, "\n")
	for y := range lines {
		for x := range lines[y] {
			a := lines[y][x]
			if a == '#' {
				m[pos{x, y}] = true
			}
		}
	}
	return m
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func visibleAsteroids(p pos, m map[pos]bool) int {
	coords := getPolarList(p, m)

	nVisible := 1
	for i := 1; i < len(coords); i++ {
		if !almostEqual(coords[i].angle, coords[i-1].angle) {
			nVisible++
		}
	}
	return nVisible
}

func dist(origo, other pos) float64 {
	a := float64(origo.x - other.x)
	b := float64(origo.y - other.y)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func draw(hull map[pos]bool, minX, maxX, minY, maxY int) {
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if x == 25 && y == 32 {
				fmt.Printf("X")
			} else if x == 11 && y == 12 {
				fmt.Printf("P")
			} else if !hull[pos{x, y}] {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
}
