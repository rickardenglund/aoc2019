package main

import (
	"aoc2019/inputs"
	"aoc2019/position"
	"container/heap"
	"flag"
	"fmt"
	"log"
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
// 6640, 6632, 6364, 6362 wrong
func part1() interface{} {
	m, pos := readMap(inputs.GetLine("day18/input.txt"))
	return findCostMap(m, pos)
}

func findCostMap(m map[position.Pos]rune, start position.Pos) int {
	tree := toTree(m, start)

	s := state{
		pos:           start,
		collectedKeys: make(map[rune]bool),
		totalKeys:     countKeys(m),
		path:          make([]position.Pos, 0),
		//path:          []move{},
		tree: tree,
	}
	pq = make(PriorityQueue, 0)
	heap.Init(&pq)

	res := findCost(m, &s)
	return res
}

var pq PriorityQueue

type vState struct {
	pos  position.Pos
	keys map[rune]bool
}

func findCost(m map[position.Pos]rune, startingState *state) int {
	var visited []vState
	var longestPath []position.Pos
	heap.Push(&pq, &Item{
		value:    startingState,
		priority: calcPrio(m, startingState),
	})
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		workingState := item.value
		possibleMoves := filter(workingState)
		for i := range possibleMoves {
			ns := copyState(workingState)
			doMove(ns, possibleMoves[i])

			if gui && len(ns.path) > len(longestPath) {
				longestPath = ns.path
				printPath(m, longestPath)
				fmt.Printf("pq: %v - %v\n", len(pq), len(ns.path))
			}
			if len(ns.collectedKeys) == ns.totalKeys {
				if gui {
					printPath(m, ns.path)
				}
				return ns.cost
			}
			filter(ns)
			if !isVisited(visited, ns) {
				heap.Push(&pq, &Item{
					value:    ns,
					priority: calcPrio(m, ns),
				})

				visited = append(visited, vState{
					pos:  ns.pos,
					keys: CopyMapRune(ns.collectedKeys),
				})
			}
		}
	}
	return -1
}

func isVisited(visited []vState, ns *state) bool {
	for i := range visited {
		if visited[i].pos == ns.pos {
			for k := range ns.collectedKeys {
				if !visited[i].keys[k] {
					return false
				}
			}
			return true
		}
	}
	return false
}

func calcPrio(m map[position.Pos]rune, s *state) int {
	gn := s.cost

	var hg int
	max := 0

	moves := get(s.tree, s.pos)
	for i := range moves {
		if moves[i].steps > max && !contains(s.path, moves[i].target) {
			max = moves[i].steps
		}
	}
	hg = max * 2

	return hg + gn
}

func appendMin(newMoves []move, newMove move) []move {
	for i := range newMoves {
		if newMoves[i].val == newMove.val {
			if newMoves[i].steps > newMove.steps {
				newMoves[i] = newMove
				return newMoves
			} else {
				return newMoves
			}
		}
	}
	newMoves = append(newMoves, newMove)
	return newMoves
}

func doMove(s *state, m move) {
	s.pos = m.target
	s.cost += m.steps
	s.path = add(s.path, s.pos)
	if isLower(m.val) {
		s.collectedKeys[m.val] = true
	}
	//s.path = append(s.path, m)
}

func add(visited []position.Pos, p position.Pos) []position.Pos {
	for i := range visited {
		if visited[i] == p {
			return visited
		}
	}

	return append(visited, p)
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}
func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func filter(s *state) []move {
	tree := s.tree
	checked := make([]position.Pos, 0)
	cur := s.pos

	return find(tree, s, checked, cur, 0)
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
func find(tree []node, s *state, checked []position.Pos, cur position.Pos, dist int) []move {
	var res []move
	checked = add(checked, cur)
	moves := get(tree, cur)
	var alreadyVisited []move
	for m := range moves {
		if contains(s.path, moves[m].target) {
			if !contains(checked, moves[m].target) {
				alreadyVisited = append(alreadyVisited, moves[m])
			}
		} else if isUpper(moves[m].val) {
			if s.collectedKeys[moves[m].val+32] {
				res = append(res, addDist(moves[m], dist))
			}
		} else {
			res = append(res, addDist(moves[m], dist))
		}

		checked = add(checked, moves[m].target)
	}

	for i := range alreadyVisited {
		res = append(res, find(tree, s, checked, alreadyVisited[i].target, dist+alreadyVisited[i].steps)...)
	}
	return res
}

func contains(visited []position.Pos, target position.Pos) bool {
	for i := range visited {
		if visited[i] == target {
			return true
		}
	}
	return false
}

func addDist(m move, dist int) move {
	return move{
		val:    m.val,
		steps:  m.steps + dist,
		target: m.target,
	}
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

type state struct {
	pos           position.Pos
	collectedKeys map[rune]bool
	totalKeys     int
	cost          int
	path          []position.Pos
	tree          []node
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

func printPath(m map[position.Pos]rune, path []position.Pos) {
	for i := range path {
		fmt.Printf("%c, ", m[path[i]])
	}
	fmt.Printf("\n")
}

func copyState(s *state) *state {
	res := state{
		pos:           s.pos,
		collectedKeys: CopyMapRune(s.collectedKeys),
		totalKeys:     s.totalKeys,
		cost:          s.cost,
		path:          CopyArray2(s.path),
		tree:          s.tree,
	}
	return &res
}

func CopyArray2(visited []position.Pos) []position.Pos {
	res := make([]position.Pos, len(visited))
	copy(res, visited)
	return res
}

func CopyArray(path []move) []move {
	res := make([]move, len(path))
	copy(res, path)
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

func part2() int {
	return -1
}
func printTree(tree []node) {
	fmt.Printf("#######\n")
	for i := range tree {
		fmt.Printf("%+v: ", tree[i].pos)
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

// An Item is something we manage in a priority queue.
type Item struct {
	value    *state // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
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

//// update modifies the priority and value of an Item in the queue.
//func (pq *PriorityQueue) update(item *Item, value *state, priority int) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}
