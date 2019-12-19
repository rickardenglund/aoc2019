package main

import (
	"aoc2019/position"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getSteps(t *testing.T) {
	m := map[position.Pos]rune{
		{X: 1}:  '.',
		{X: -1}: '.',
		{Y: 1}:  'A',
		{Y: -1}: 'a',
	}
	v := map[position.Pos]bool{}
	assert.Equal(t, []position.Pos{{X: 1}, {X: -1}, {Y: -1}, {Y: 1}}, getSteps(m, v, pos{}))
}

func Test_readMap(t *testing.T) {
	m, player := readMap("###\n#@.\n###")
	assert.Equal(t, position.Pos{X: 1, Y: 1}, player)
	assert.Equal(t, '.', m[position.Pos{X: 2, Y: 1}])
}

const mini2 = `######
#a..@.b#
######A#
######.#
########`

const mediumX = `########################
#f.D.E.e.C.@.........c.#
######################.#
#d.....................#
########################`

const medium = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`

func Test_getAvailableMoves(t *testing.T) {
	m, playerPos := readMap(small)
	res := []move{
		{val: 'a', steps: 2, target: position.Pos{X: 7, Y: 1}},
		{val: 'A', steps: 2, target: position.Pos{X: 3, Y: 1}}}

	assert.Contains(t, getAvailableMoves(m, pos{playerPos, 0}), res[0])
	assert.Contains(t, getAvailableMoves(m, pos{playerPos, 0}), res[1])

	m, playerPos = readMap(mediumX)
	res = []move{
		{val: 'C', steps: 2, target: position.Pos{X: 9, Y: 1}},
		{val: 'c', steps: 10, target: position.Pos{X: 21, Y: 1}}}
	assert.Contains(t, getAvailableMoves(m, pos{playerPos, 0}), res[0])
	assert.Contains(t, getAvailableMoves(m, pos{playerPos, 0}), res[1])
}

func Test_getTree(t *testing.T) {
	m, start := readMap(small)
	tree := toTree(m, start)
	assert.Equal(t, position.Pos{X: 5, Y: 1}, start)
	assert.Contains(t, tree[start], move{val: 'a', target: position.Pos{X: 7, Y: 1}, steps: 2})
	assert.Contains(t, tree[start], move{val: 'A', target: position.Pos{X: 3, Y: 1}, steps: 2})

	assert.Contains(t, tree[position.Pos{X: 3, Y: 1}], move{val: 'b', target: position.Pos{X: 1, Y: 1}, steps: 2})
	assert.Contains(t, tree[position.Pos{X: 3, Y: 1}], move{val: 'a', target: position.Pos{X: 7, Y: 1}, steps: 4})
}

const small = `#########
#b.A.@.a#
#########`

func Test_findCostB(t *testing.T) {
	gui = true
	m, playerPos := readMap(small)
	assert.Equal(t, 8, findCostMap(m, playerPos))
}

func Test_findCost22B(t *testing.T) {
	m, playerPos := readMap(mini2)
	assert.Equal(t, 7, findCostMap(m, playerPos))
}

const mini = `######
#a..@.b#
######A#
######c#
########`

func Test_find(t *testing.T) {
	s := state{
		pos:           position.Pos{X: 1, Y: 1},
		collectedKeys: map[rune]bool{},
		totalKeys:     1,
		cost:          0,
		//visited:       nil,
		//path:          nil,
	}
	tree := map[position.Pos][]move{
		{X: 1, Y: 1}: {{'a', 2, position.Pos{X: 2, Y: 2}}},
	}
	moves := filter(tree, &s)
	assert.Contains(t, moves, move{'a', 2, position.Pos{X: 2, Y: 2}})
}

func Test_findCost2B(t *testing.T) {
	m, playerPos := readMap(mini)
	assert.Equal(t, 10, findCostMap(m, playerPos))
}

func Test_findCost3B(t *testing.T) {
	m, playerPos := readMap(medium)
	assert.Equal(t, 86, findCostMap(m, playerPos))
}

const medium2 = `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`

func Test_findCost4B(t *testing.T) {
	m, playerPos := readMap(medium2)
	assert.Equal(t, 132, findCostMap(m, playerPos))
}

const medium3 = `#################
#i.G..c...e..H.gp#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`

func Test_findCost5B(t *testing.T) {
	gui = true
	m, playerPos := readMap(medium3)
	assert.Equal(t, 136, findCostMap(m, playerPos))
}

const medium4 = `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`

func Test_findCost6B(t *testing.T) {
	m, playerPos := readMap(medium4)
	assert.Equal(t, 81, findCostMap(m, playerPos))
}

func Test_copyState(t *testing.T) {
	a := state{}
	b := a
	a.pos = position.Pos{X: 1, Y: 2}
	assert.NotEqual(t, a, b)
}
func Test_append(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, len(a)+1)
	copy(b, a)
	b = append(b[:2], 4)
	a[0] = 0
	assert.NotEqual(t, a[2], b[2])
	assert.NotEqual(t, a[0], b[0])
}

const apa = `#######
#b.a.c#
#.###.#
#..d..#
#e#####
#######`

func Test_removeNode(t *testing.T) {
	m, p := readMap(apa)
	tree := toTree(m, p)

	printTree(tree)

	newTree := removeNode(tree, position.Pos{X: 1, Y: 1})
	fmt.Printf("########\n")
	printTree(newTree)

}

func Test_appendMin(t *testing.T) {
	moves := []move{{val: 'a', steps: 2}}
	moves = appendMin(moves, move{val: 'a', steps: 1})
	moves = appendMin(moves, move{val: 'a', steps: 4})
	printMoves(moves)
	assert.Equal(t, 1, len(moves))
}
