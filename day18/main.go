package main

import (
	"aoc2019/inputs"
	"aoc2019/position"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math"
	"sort"
	"time"
	"unicode"
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
// 6640, 6632, 6364, 6362 wrong
func part1() interface{} {
	m, pos := readMap(inputs.GetLine("day18/input.txt"))
	return findCostMap(m, pos)
}

func part2() int {

	return -1
}

func findCostMap(m map[position.Pos]rune, start position.Pos) int {
	tree := toTree(m, start)

	startState := vState{
		pos:  start,
		keys: nil,
		cost: 0,
		name: m[start],
		path: []rune{},
	}
	totalKeys := countKeys(m)
	res := findCost(m, tree, startState, totalKeys)
	longestPath = nil
	return res
}

var longestPath []rune

func findCost(m map[position.Pos]rune, tree []node, startState vState, totalKeys int) int {
	pq := make(PriorityQueue, 0)
	var visited = make(map[position.Pos][]*vState)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		value:    &startState,
		priority: 0,
	})

	for len(pq) > 0 {
		s := heap.Pop(&pq).(*Item).value

		if len(s.keys) == totalKeys {
			return s.cost
		}
		if gui {
			if len(s.keys) > len(longestPath) {
				longestPath = CopyRunes(s.keys)
				fmt.Printf("pq: %v, visited: %v __ ", len(pq), len(visited))
				printRunePath(longestPath)
			}
		}
		visited = addToVisited(visited, s)

		possibleMoves := getMoves(tree, s)
		for i := range possibleMoves {
			newState := createState(s, possibleMoves[i])
			if shouldVisit(newState, visited) {
				heap.Push(&pq, &Item{
					value:    newState,
					priority: calcPriority(tree, newState, totalKeys),
				})
			}
		}

	}

	return -1
}

func printRunePath(path []rune) {
	fmt.Printf("path: ")
	for i := range path {
		fmt.Printf("%c, ", path[i])
	}
	fmt.Printf("\n")

}

func shouldVisit(state *vState, visited map[position.Pos][]*vState) bool {
	list, ok := visited[state.pos]
	if ok {
		for i := range list {
			if isBetterVersion(list[i], state) {
				return false
			}
		}
	}
	return true
}

func addToVisited(visited map[position.Pos][]*vState, s *vState) map[position.Pos][]*vState {
	listOfPos, ok := visited[s.pos]
	if ok {
		for i := range listOfPos {
			if isBetterVersion(s, listOfPos[i]) {
				listOfPos[i] = s
				return visited
			}
		}
	}
	visited[s.pos] = append(visited[s.pos], s)
	return visited
}

func isBetterVersion(a *vState, b *vState) bool {
	if a.pos == b.pos && allOrMore(a.keys, b.keys) {
		if a.cost <= b.cost {
			return true
		}
	}
	return false

}

func allOrMore(as []rune, bs []rune) bool {
	if len(as) >= len(bs) {
		for i := range bs {
			if bs[i] != as[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func printState(s *vState) {
	fmt.Printf("%c: c=%v - ", s.name, s.cost)
	printKeys(s.keys)
	fmt.Printf("\n")
	fmt.Printf("path: [")
	for i := range s.path {
		fmt.Printf("%c, ", s.path[i])
	}
	fmt.Printf("]\n\n")
}

func createState(s *vState, m move) *vState {
	newState := vState{
		pos:  m.target,
		keys: CopyRunes(s.keys), // TODO: maybe optimize
		cost: s.cost + m.steps,
		name: m.val,
		path: CopyRunesAppend(s.path, m.val),
	}
	if isLower(m.val) {
		newState.keys = addKey(newState.keys, m.val)
	}
	return &newState
}

func addKey(keys []rune, val rune) []rune {
	for i := range keys {
		if keys[i] == val {
			return keys
		}
	}

	keys = append(keys, val)
	sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
	return keys
}

func CopyRunes(path []rune) []rune {
	res := make([]rune, len(path))
	copy(res, path)
	return res
}
func CopyRunesAppend(path []rune, val rune) []rune {
	res := make([]rune, len(path)+1)
	copy(res, path)
	res[len(path)] = val
	return res
}

func getMoves(tree []node, s *vState) []move {
	moves := get(tree, s.pos)
	var res []move
	for i := range moves {
		if isLower(moves[i].val) || canUnlock(s.keys, moves[i].val) {
			res = append(res, moves[i])
		}
	}

	return res
}

func canUnlock(keys []rune, val rune) bool {
	for i := range keys {
		if keys[i] == val+32 {
			return true
		}
	}
	return false
}

type vState struct {
	pos  position.Pos
	keys []rune
	cost int
	name rune
	path []rune
}

//nolint
func calcPriority(tree []node, s *vState, totalKeys int) int {
	gn := s.cost

	moves := get(tree, s.pos)
	closestKey := math.MaxInt32
	found := false
	for i := range moves {
		if isLower(moves[i].val) && !containsRune(s.keys, moves[i].val) {
			if moves[i].steps < closestKey {
				closestKey = moves[i].steps
				found = true
			}
		}
	}
	if !found {
		closestKey = 0
	}
	hg := closestKey
	return hg + gn
}

func containsRune(keys []rune, val rune) bool {
	for i := range keys {
		if keys[i] == val {
			return true
		}
	}
	return false
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func getI(tree []node, p position.Pos) int {
	for i := range tree {
		if tree[i].pos == p {
			return i
		}
	}
	log.Fatalln("Not found")
	return -1
}

func get(tree []node, p position.Pos) []move {
	for i := range tree {
		if tree[i].pos == p {
			return tree[i].moves
		}
	}
	log.Fatalln("Not found")
	return nil
}

type node struct {
	pos   position.Pos
	moves []move
}

func toTree(m map[position.Pos]rune, start position.Pos) []node {
	tree := []node{}
	for k, v := range m {
		if v >= 'A' && v <= 'z' || v == '@' {
			neighbours := getAvailableMoves(m, pos{p: k, dist: 0})
			tree = append(tree, node{k, []move{}})
			i := getI(tree, k)
			for n := range neighbours {
				tree[i].moves = append(tree[i].moves, neighbours[n])
			}
		}
	}

	neighbours := getAvailableMoves(m, pos{p: start, dist: 0})
	tree = append(tree, node{start, []move{}})
	i := getI(tree, start)
	for n := range neighbours {
		tree[i].moves = append(tree[i].moves, neighbours[n])
	}

	return tree
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

//nolint
func printPath(m map[position.Pos]rune, path []position.Pos) {
	for i := range path {
		fmt.Printf("%c, ", m[path[i]])
	}
	fmt.Printf("\n")
}

//nolint
func CopyArray(path []move) []move {
	res := make([]move, len(path))
	copy(res, path)
	return res
}

//nolint
func CopyMap(a map[position.Pos]bool) map[position.Pos]bool {
	res := map[position.Pos]bool{}
	for k, v := range a {
		res[k] = v
	}
	return res
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
	if gui {
		fmt.Printf("\n")
	}
	m[playerPos] = '.'
	return m, playerPos
}

type move struct {
	val    rune
	steps  int
	target position.Pos
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
				moves = append(moves, move{val: v, steps: cur.dist + 1, target: neighbours[i]})
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

//nolint
func printTree(tree []node, m map[position.Pos]rune) {
	fmt.Printf("#######\n")
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			c, ok := m[position.Pos{X: x, Y: y}]
			if ok && unicode.IsLetter(c) {
				fmt.Printf("%c", c)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("#######\n")
	for i := range tree {
		fmt.Printf("%c %+v: ", m[tree[i].pos], tree[i].pos)
		printMoves(tree[i].moves)

	}
	fmt.Printf("#######\n")
}

func printMoves(v []move) {
	for i := range v {
		fmt.Printf("(%c_%v:%v) ", v[i].val, v[i].target, v[i].steps)
	}
	fmt.Printf("\n")
}

func printKeys(keys []rune) {
	for k := range keys {
		fmt.Printf("%c, ", keys[k])
	}
}
