package main

import (
	"aoc2019/inputs"
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

// 8089 to high
func part1() interface{} {
	m, pos := readMap(inputs.GetLine("day18/input.txt"))

	return findKeysCost(m, pos)
}

type state struct {
	keys map[rune]bool
	pos  position.Pos
}

func findKeysCost(m map[position.Pos]rune, startingPos position.Pos) int {
	cost := 0
	s := state{keys: map[rune]bool{}}
	s.pos = startingPos

	for {
		moves := getAvailableMoves(m, pos{s.pos, 0})

		moves = filterMoves(s, moves)
		if len(moves) == 0 {
			return cost
		}

		cost += moves[0].steps
		s.pos = moves[0].p
		if m[s.pos] >= 'a' && m[s.pos] <= 'z' {
			s.keys[m[s.pos]] = true
			m[s.pos] = '.'
		}
		if m[s.pos] >= 'A' && m[s.pos] <= 'Z' {
			m[s.pos] = '.'
		}
	}
}

func filterMoves(s state, moves []move) []move {
	filtered := []move{}
	for i := range moves {
		if moves[i].val <= 'Z' {

			if s.keys[moves[i].val+32] {
				filtered = append(filtered, moves[i])
			}
		} else {
			filtered = append(filtered, moves[i])
		}
	}
	return filtered
}

func readMap(str string) (map[position.Pos]rune, position.Pos) {
	m := map[position.Pos]rune{}
	y := 0
	x := 0
	var playerPos position.Pos
	for i := 0; i < len(str); i++ {
		data := rune(str[i])
		m[position.Pos{X: x, Y: y}] = data
		if gui {
			fmt.Printf("%c", data)
		}
		if data == '@' {
			playerPos = position.Pos{X: x, Y: y}
		}
		if data == '\n' {
			y++
			x = 0
		} else {
			x++
		}
	}
	m[playerPos] = '.'
	return m, playerPos
}

type move struct {
	val   rune
	steps int
	p     position.Pos
}

type pos struct {
	p    position.Pos
	dist int
}

func getAvailableMoves(m map[position.Pos]rune, cur pos) []move {
	visited := map[position.Pos]bool{}
	toVisit := map[pos]bool{cur: true}
	var moves []move
	steps := 0
	for ; ; steps++ {
		delete(toVisit, cur)
		visited[cur.p] = true
		neighbours := getSteps(m, visited, cur)
		for i := range neighbours {
			v := m[neighbours[i]]

			if v >= 'A' && v <= 'z' {
				moves = append(moves, move{val: v, steps: cur.dist + 1, p: neighbours[i]})
			} else {
				toVisit[pos{neighbours[i], cur.dist + 1}] = true
			}
		}

		if len(toVisit) == 0 {
			break
		}
		for k := range toVisit {
			cur = k
			break
		}

	}

	return moves
}

func getSteps(m map[position.Pos]rune, visited map[position.Pos]bool, p pos) []position.Pos {
	var res []position.Pos
	if isOk(m, visited, position.Pos{X: p.p.X + 1, Y: p.p.Y}) {
		res = append(res, position.Pos{X: p.p.X + 1, Y: p.p.Y})
	}
	if isOk(m, visited, position.Pos{X: p.p.X - 1, Y: p.p.Y}) {
		res = append(res, position.Pos{X: p.p.X - 1, Y: p.p.Y})
	}
	if isOk(m, visited, position.Pos{X: p.p.X, Y: p.p.Y - 1}) {
		res = append(res, position.Pos{X: p.p.X, Y: p.p.Y - 1})
	}
	if isOk(m, visited, position.Pos{X: p.p.X, Y: p.p.Y + 1}) {
		res = append(res, position.Pos{X: p.p.X, Y: p.p.Y + 1})
	}
	return res
}

func isOk(m map[position.Pos]rune, visited map[position.Pos]bool, p position.Pos) bool {
	v, ok := m[p]
	return ok && v != '#' && !visited[p]
}

func part2() int {
	return -1
}
