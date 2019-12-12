package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

const ex1 = `.#..#
.....
#####
....#
...##`

func Test_getAsteroids(t *testing.T) {
	m := getAsteroids(ex1)
	//fmt.Printf("%v\n", m)
	assert.Equal(t, 10, len(m))
}

func Test_NAsteroids(t *testing.T) {
	m := getAsteroids(ex1)
	assert.Equal(t, 7, visibleAsteroids(pos{1, 0}, m))
	assert.Equal(t, 5, visibleAsteroids(pos{4, 2}, m))
	assert.Equal(t, 7, visibleAsteroids(pos{4, 0}, m))
}

func Test_NAsteroids2(t *testing.T) {
	m := getAsteroids(`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`)
	assert.Equal(t, 33, visibleAsteroids(pos{5, 8}, m))
}

const exLarge = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

func Test_NAsteroidsLarge(t *testing.T) {
	m := getAsteroids(exLarge)
	//fmt.Printf("%v\n", m)
	assert.Equal(t, 210, visibleAsteroids(pos{11, 13}, m))
}

func Test_getAngle(t *testing.T) {
	assert.InDelta(t, math.Pi/4, getAngle(pos{0, 0}, pos{1, -1}), float64EqualityThreshold)
	assert.InDelta(t, 0, getAngle(pos{0, 0}, pos{0, -1}), float64EqualityThreshold)
	assert.InDelta(t, math.Pi/2, getAngle(pos{0, 0}, pos{1, 0}), float64EqualityThreshold)
	assert.InDelta(t, math.Pi*1.5, getAngle(pos{0, 0}, pos{-1, 0}), float64EqualityThreshold)
}

func Test_angleOf(t *testing.T) {
	assert.InDelta(t, 0, angleOf(pos{0, -1}), float64EqualityThreshold)
	assert.InDelta(t, math.Pi/4, angleOf(pos{1, -1}), float64EqualityThreshold)
	assert.InDelta(t, math.Pi/2, angleOf(pos{1, 0}), float64EqualityThreshold)
	assert.InDelta(t, math.Pi, angleOf(pos{0, 1}), float64EqualityThreshold)
	assert.InDelta(t, math.Pi*0.5, angleOf(pos{1, 0}), float64EqualityThreshold)
}

func Test_sort(t *testing.T) {
	ints := []int{2, 5, 1, 2}

	sort.Slice(ints,
		func(i, j int) bool {
			return ints[i] < ints[j]
		})
}

func Test_getNExplosion(t *testing.T) {
	m := getAsteroids(exLarge)

	pl := getPolarList(pos{11, 13}, m)

	assert.Equal(t, pos{11, 12}, getNExplosion(1, pl))
	assert.Equal(t, pos{12, 1}, getNExplosion(2, pl))
	assert.Equal(t, pos{8, 2}, getNExplosion(200, pl))
	assert.Equal(t, pos{10, 9}, getNExplosion(201, pl))
}
