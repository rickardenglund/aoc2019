package main

import (
	"aoc2019/position"
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

const small = `#########
#b.A.@.a#
#########`

const mediumx = `########################
#f.D.E.e.C.@.........c.#
######################.#
#d.....................#
########################`

const medium = `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`

const medium2 = `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`

func Test_getAvailableMoves(t *testing.T) {
	m, playerPos := readMap(small)
	res := []move{
		{val: 'a', steps: 2, p: position.Pos{X: 7, Y: 1}},
		{val: 'A', steps: 2, p: position.Pos{3, 1}}}

	assert.Equal(t, res, getAvailableMoves(m, pos{playerPos, 0}))

	m, playerPos = readMap(mediumx)
	res = []move{
		{val: 'C', steps: 2, p: position.Pos{9, 1}},
		{val: 'c', steps: 10, p: position.Pos{21, 1}}}
	assert.Equal(t, res, getAvailableMoves(m, pos{playerPos, 0}))
}

func Test_filterMoves(t *testing.T) {
	s := state{
		keys: nil,
		pos:  position.Pos{},
	}
	avMoves := []move{{val: 'a', steps: 2}, {val: 'A', steps: 2}}
	assert.Equal(t, []move{{
		val:   'a',
		steps: 2,
	}}, filterMoves(s, avMoves))

	s = state{
		keys: map[rune]bool{'a': true},
		pos:  position.Pos{},
	}
	avMoves = []move{{val: 'b', steps: 2}, {val: 'A', steps: 2}}
	assert.Equal(t, []move{{val: 'b', steps: 2}, {val: 'A', steps: 2}}, filterMoves(s, avMoves))
}

func Test_findCost(t *testing.T) {
	m, playerPos := readMap(small)
	assert.Equal(t, 8, findKeysCost(m, playerPos))

	m, playerPos = readMap(medium)
	assert.Equal(t, 86, findKeysCost(m, playerPos))

	m, playerPos = readMap(medium2)
	assert.Equal(t, 132, findKeysCost(m, playerPos))
}
