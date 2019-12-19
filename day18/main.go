package main

import (
	"aoc2019/inputs"
	"aoc2019/position"
	"flag"
	"fmt"
	"math"
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

// 6719 to high
// 6640, 6632 wrong
func part1() interface{} {
	m, pos := readMap(inputs.GetLine("day18/input.txt"))
	//tree := toTree(m, pos)

	return findKeysCost(m, pos)
}

func toTree(m map[position.Pos]rune, start position.Pos) map[position.Pos][]move {
	tree := map[position.Pos][]move{}
	for k, v := range m {
		if v >= 'A' && v <= 'z' || v == '@' {
			neighbours := getAvailableMoves(m, pos{p: k, dist: 0})
			for n := range neighbours {
				tree[k] = append(tree[k], neighbours[n])
			}
		}
	}

	neighbours := getAvailableMoves(m, pos{p: start, dist: 0})
	for n := range neighbours {
		tree[start] = append(tree[start], neighbours[n])
	}
	return tree

}

type state struct {
	pos           position.Pos
	collectedKeys int
	totalKeys     int
}

var best = math.MaxInt64

func findKeysCost(m map[position.Pos]rune, startingPos position.Pos) int {
	s := state{}
	s.totalKeys = countKeys(m)
	s.pos = startingPos
	best = math.MaxInt32
	cost, path := findCost(m, s, 0, []move{})
	if gui {
		printPath(path)
	}
	return cost
}

func countKeys(m map[position.Pos]rune) int {
	sum := 0
	for _, v := range m {
		if v >= 'a' && v <= 'z' {
			sum++
		}
	}
	return sum
}

func printPath(path []move) {
	for i := range path {
		fmt.Printf("%c, ", path[i].val)
	}
	fmt.Println()
}

func findCost(m map[position.Pos]rune, s state, acc int, path []move) (int, []move) {
	if acc > best {
		return math.MaxInt32, path
	}
	moves := getAvailableMoves(m, pos{s.pos, 0})

	moves = filterMoves(path, moves)
	if len(moves) == 0 || s.totalKeys == s.collectedKeys {
		if acc < best {
			best = acc
		}
		if gui {
			fmt.Printf("No more:%-6v -  %-7v ", best, acc)
			printPath(path)
		}
		return acc, path
	}

	var paths [][]move
	for i := range moves {
		newMap := copyMap(m)
		newState := s

		moveTo(&newState, moves[i], newMap)

		newPath := copyAppend(path, moves[i])
		_, resPath := findCost(newMap, newState, acc+moves[i].steps, newPath)

		paths = append(paths, resPath)
	}

	min := math.MaxInt64
	//maxlength := 0
	var minPath []move
	for i := range paths {
		cost := pathCost(paths[i])
		if cost < min {
			min = cost
			minPath = paths[i]
		}
	}

	return min, minPath
}

func pathCost(moves []move) int {
	sum := 0
	for i := range moves {
		sum += moves[i].steps
	}
	return sum
}

func copyAppend(path []move, m move) []move {
	newPath := make([]move, len(path))
	copy(newPath, path)
	newPath = append(newPath, m)
	return newPath
}

func moveTo(s *state, move move, m map[position.Pos]rune) {
	s.pos = move.p
	if m[s.pos] >= 'a' && m[s.pos] <= 'z' {
		s.collectedKeys++
		m[s.pos] = '.'
	}
	if m[s.pos] >= 'A' && m[s.pos] <= 'Z' {
		m[s.pos] = '.'
	}
}

func copyMap(a map[position.Pos]rune) map[position.Pos]rune {
	res := map[position.Pos]rune{}
	for k, v := range a {
		res[k] = v
	}
	return res
}

func filterMoves(path []move, moves []move) []move {
	var filtered []move
	for i := range moves {
		if moves[i].val <= 'Z' && moves[i].val >= 'A' {
			if hasKey(path, moves[i].val) {
				filtered = append(filtered, moves[i])
			}
		} else {
			filtered = append(filtered, moves[i])
		}
	}
	return filtered
}

func hasKey(path []move, door rune) bool {
	for i := range path {
		if path[i].val-32 == door {
			return true
		}
	}
	return false
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
