package main

import (
	"aoc2019/inputs"
	"aoc2019/position"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"math"
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

func part1() interface{} {
	str := inputs.GetLine("day20/input.txt")
	m := readMap(str)
	portals := findPortals(m)
	starts := findPortalsWithName(portals, "AA")
	stops := findPortalsWithName(portals, "ZZ")
	if len(starts) != 1 || len(stops) != 1 {
		log.Fatalln("Invalid input")
	}
	cost := findCost(m, portals, starts[0], stops[0])
	return cost
}

func part2() interface{} {
	return "-"
}

type portal struct {
	name string
}

type state struct {
	cost    int
	pos     position.Pos
	target  position.Pos
	visited []position.Pos
}
type stateMove struct {
	*state
	move
}

func findCost(m map[position.Pos]rune, portals map[position.Pos]portal, start position.Pos, stop position.Pos) int {
	tree := toTree(m, portals)

	s := state{
		cost:   0,
		pos:    start,
		target: stop,
	}

	pq = make(PriorityQueue, 0)
	heap.Init(&pq)

	res := search(tree, portals, s)
	return res
}

var pq PriorityQueue

func search(tree []node, portals map[position.Pos]portal, startingState state) int {
	s := &startingState
	for {
		possibleMoves := getMoves(tree, portals, s)
		for i := range possibleMoves {
			heap.Push(&pq, &Item{
				value:    &stateMove{s, possibleMoves[i]},
				priority: calcPrio(tree, portals, s, possibleMoves[i]),
			})
		}

		item := heap.Pop(&pq).(*Item)
		sm := item.value
		s = doMove(sm.state, sm.move)
		if s.pos == s.target {
			return s.cost
		}

		if len(pq) == 0 {
			return -1
		}
	}
}

func calcPrio(tree []node, portals map[position.Pos]portal, s *state, mv move) int {
	gn := s.cost + mv.steps
	hn := 0

	moves := getMoves(tree, portals, s)
	targetDist := math.MaxInt32
	for i := range moves {
		if moves[i].steps < targetDist {
			targetDist = moves[i].steps
		}
	}
	hn = targetDist
	if mv.target == s.target {
		hn = 0
	}
	return gn + hn
}

func doMove(s *state, m move) *state {
	newState := copyState(s)
	newState.pos = m.target
	newState.cost += m.steps
	newState.visited = append(newState.visited, newState.pos)
	return newState
}

func copyState(s *state) *state {
	ns := state{
		cost:    s.cost,
		pos:     s.pos,
		target:  s.target,
		visited: make([]position.Pos, len(s.visited)),
	}
	copy(ns.visited, s.visited)
	return &ns

}

func getMoves(tree []node, portals map[position.Pos]portal, s *state) []move {
	allMoves := get(tree, s.pos)
	filtered := []move{}
	for m := range allMoves {
		if !contains(s.visited, allMoves[m].target) {
			filtered = append(filtered, allMoves[m])
		}
	}

	p, ok := getCompanionPortal(portals, s.pos)
	if ok {
		filtered = append(filtered, move{
			val:    portals[s.pos],
			steps:  1,
			target: p,
		})
	}

	return filtered
}

func getCompanionPortal(portals map[position.Pos]portal, p position.Pos) (position.Pos, bool) {
	list := findPortalsWithName(portals, portals[p].name)
	for i := range list {
		if list[i] != p {
			return list[i], true
		}
	}
	return p, false

}
func findPortalsWithName(portals map[position.Pos]portal, name string) []position.Pos {
	res := []position.Pos{}
	var target = portal{name}
	for k, v := range portals {
		if v == target {
			res = append(res, k)
		}
	}
	return res
}

type pos struct {
	p    position.Pos
	dist int
}

type move struct {
	val    portal
	steps  int
	target position.Pos
}

type node struct {
	pos   position.Pos
	moves []move
}

func toTree(m map[position.Pos]rune, portals map[position.Pos]portal) []node {
	tree := []node{}
	for k := range portals {
		neighbours := getAvailableMoves(m, portals, pos{k, 0})
		tree = append(tree, node{k, []move{}})
		i := getI(tree, k)
		for ni := range neighbours {
			tree[i].moves = append(tree[i].moves, neighbours[ni])
		}
	}
	return tree
}

func getAvailableMoves(m map[position.Pos]rune, portals map[position.Pos]portal, cur pos) []move {
	visited := map[position.Pos]bool{}
	toVisit := map[pos]bool{cur: true}
	var moves []move
	steps := 0
	for ; ; steps++ {
		delete(toVisit, cur)
		visited[cur.p] = true
		neighbours := getSteps(m, visited, cur)
		for i := range neighbours {
			portal, ok := portals[neighbours[i]]
			if ok {
				moves = append(moves, move{val: portal, steps: cur.dist + 1, target: neighbours[i]})
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
	return ok && v == '.' && !visited[p]
}

func findPortals(m map[position.Pos]rune) map[position.Pos]portal {
	res := make(map[position.Pos]portal)
	for k, v := range m {
		var start position.Pos
		var name string
		if unicode.IsLetter(v) {
			isLetter := func(p position.Pos) bool {
				return unicode.IsLetter(m[p])
			}
			isOpen := func(p position.Pos) bool {
				return m[p] == '.'
			}
			k2 := getNearby(isLetter, []position.Pos{k})
			start = getNearby(isOpen, []position.Pos{k, k2})
			name = getName(m, k, k2)
			res[start] = portal{name: name}
		}

	}
	return res
}

func getName(m map[position.Pos]rune, k position.Pos, k2 position.Pos) string {
	if k.X < k2.X || k.Y < k2.Y {
		return string([]rune{m[k], m[k2]})
	}
	return string([]rune{m[k2], m[k]})
}

func getNearby(isOk func(a position.Pos) bool, ps []position.Pos) position.Pos {
	for _, p := range ps {
		if isOk(position.Pos{X: p.X + 1, Y: p.Y}) {
			return position.Pos{X: p.X + 1, Y: p.Y}
		}
		if isOk(position.Pos{X: p.X - 1, Y: p.Y}) {
			return position.Pos{X: p.X - 1, Y: p.Y}
		}
		if isOk(position.Pos{X: p.X, Y: p.Y - 1}) {
			return position.Pos{X: p.X, Y: p.Y - 1}
		}
		if isOk(position.Pos{X: p.X, Y: p.Y + 1}) {
			return position.Pos{X: p.X, Y: p.Y + 1}
		}
	}
	log.Fatalln("Letter not found")
	return position.Pos{}
}
func readMap(str string) map[position.Pos]rune {
	m := map[position.Pos]rune{}
	y := 0
	x := 0
	for i := 0; i < len(str); i++ {
		data := rune(str[i])
		m[position.Pos{X: x, Y: y}] = data
		if gui {
			fmt.Printf("%c", data)
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
	return m
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

func contains(visited []position.Pos, target position.Pos) bool {
	for i := range visited {
		if visited[i] == target {
			return true
		}
	}
	return false
}

type Item struct {
	value    *stateMove // The value of the item; arbitrary.
	priority int        // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
