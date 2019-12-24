package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"aoc2019/inputs"
	intmath "aoc2019/mymath"
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
	m := findRepeatingState(strings.TrimSpace(inputs.GetLine("day24/input.txt")))
	return bioRating(m)
}
func bioRating(m [][]bool) int {
	i := 0
	sum := 0
	for y := range m {
		for x := range m[y] {
			if m[y][x] {
				sum += intmath.Exp2(i)
			}
			i++
		}
	}
	return sum
}

func findRepeatingState(str string) [][]bool {
	visited := map[string]bool{}
	m := readState(str)
	visited[toString(m)] = true
	for {
		m = nextState(m)
		s := toString(m)
		if visited[s] {
			break
		}
		visited[s] = true
	}
	return m
}

func nextState(m [][]bool) [][]bool {
	res := make([][]bool, len(m))
	for y := range m {
		res[y] = make([]bool, len(m[y]))
		for x := range m[y] {
			n := countCloseBugs(m, y, x)
			if m[y][x] { // isBug
				res[y][x] = n == 1
			} else { //free space
				res[y][x] = n == 1 || n == 2
			}
		}
	}
	return res

}

func countCloseBugs(m [][]bool, y int, x int) int {
	sum := 0
	if x > 0 && m[y][x-1] {
		sum++
	}
	if x < len(m[0])-1 && m[y][x+1] {
		sum++
	}
	if y > 0 && m[y-1][x] {
		sum++
	}
	if y < len(m)-1 && m[y+1][x] {
		sum++
	}
	return sum

}

func part2() interface{} {
	return "-"
}

func readState(str string) [][]bool {
	lines := strings.Split(str, "\n")
	res := make([][]bool, len(lines))
	for y := range lines {
		res[y] = make([]bool, len(lines[y]))
		line := strings.TrimSpace(lines[y])
		for x := range line {
			c := line[x]
			res[y][x] = '#' == c
		}
	}
	return res
}

func toString(m [][]bool) string {
	sb := strings.Builder{}
	for y := range m {
		for x := range m[y] {
			if m[y][x] {
				sb.WriteRune('#') // nolint: gosec
			} else {
				sb.WriteRune('.') // nolint: gosec
			}
		}
		sb.WriteRune('\n') // nolint: gosec
	}
	return sb.String()
}
