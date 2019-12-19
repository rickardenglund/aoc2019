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
	return findCostMap(m, pos)
}

func findCostMap(m map[position.Pos]rune, start position.Pos) int {
	tree := toTree(m, start)
	best = math.MaxInt32

	s := state{
		pos:           start,
		collectedKeys: make(map[rune]bool),
		totalKeys:     countKeys(m),
		visited:       make(map[position.Pos]bool),
		path:          []move{},
	}

	res := findCost2(tree, &s)
	printPath(s.path)
	return res
}
func findCost2(tree map[position.Pos][]move, s *state) int {
	possibleMoves := filter(tree, s)

	if s.cost > best {
		//fmt.Printf("too bad: ")
		//printPath(s.path)
		return math.MaxInt32
	}

	if len(possibleMoves) == 0 || len(s.collectedKeys) == s.totalKeys {
		if gui {
			fmt.Printf("cost: %v - ", s.cost)
			printPath(s.path)
		}
		if best > s.cost {
			best = s.cost
		}
		return s.cost
	}

	minCost := math.MaxInt32
	for i := range possibleMoves {
		newState := copyState(s)
		doMove(newState, possibleMoves[i])
		cost := findCost2(tree, newState)
		if cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func doMove(s *state, m move) {
	s.pos = m.p
	s.cost += m.steps
	s.visited[s.pos] = true
	if isLower(m.val) {
		s.collectedKeys[m.val] = true
	}
	s.path = append(s.path, m)
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func filter(tree map[position.Pos][]move, s *state) []move {
	checked := make(map[position.Pos]bool)
	cur := s.pos

	return find(tree, s, checked, cur, 0)
}

func find(tree map[position.Pos][]move, s *state, checked map[position.Pos]bool, cur position.Pos, dist int) []move {
	var res []move
	checked[cur] = true
	moves := tree[cur]
	var alreadyVisited []move
	for m := range moves {
		if s.visited[moves[m].p] {
			if !checked[moves[m].p] {
				alreadyVisited = append(alreadyVisited, moves[m])
			}
		} else if isUpper(moves[m].val) {
			if s.collectedKeys[moves[m].val+32] {
				res = append(res, addDist(moves[m], dist))
			}
		} else {
			res = append(res, addDist(moves[m], dist))
		}

		checked[moves[m].p] = true
	}

	for i := range alreadyVisited {
		res = append(res, find(tree, s, checked, alreadyVisited[i].p, dist+alreadyVisited[i].steps)...)
	}
	return res
}

func addDist(m move, dist int) move {
	return move{
		val:   m.val,
		steps: m.steps + dist,
		p:     m.p,
	}
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
	collectedKeys map[rune]bool
	totalKeys     int
	cost          int
	visited       map[position.Pos]bool
	path          []move
}

var best = math.MaxInt64

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

func CopyAppend(path []move, m move) []move {
	newPath := make([]move, len(path))
	copy(newPath, path)
	newPath = append(newPath, m)
	return newPath
}

func copyState(s *state) *state {
	res := state{
		pos:           s.pos,
		collectedKeys: CopyMapRune(s.collectedKeys),
		totalKeys:     s.totalKeys,
		cost:          s.cost,
		visited:       CopyMap(s.visited),
		path:          CopyArray(s.path),
	}
	return &res
}

func CopyArray(path []move) []move {
	res := make([]move, len(path))
	for i := range path {
		res[i] = path[i]
	}
	return res
}

func CopyMapRune(a map[rune]bool) map[rune]bool {
	res := map[rune]bool{}
	for k, v := range a {
		res[k] = v
	}
	return res
}
func CopyMap(a map[position.Pos]bool) map[position.Pos]bool {
	res := map[position.Pos]bool{}
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
