package main

import (
	"aoc2019/inputs"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)


func main() {
	fmt.Printf("part1: %v\n", part1())
	fmt.Printf("part1: %v\n", part2())
}

type pos struct {
	x int
	y int
}

func part1() int{
	wireTurns := inputs.GetLines("day3/input.txt")

	path1 := walkpath(wireTurns, 0)
	path2 := walkpath(wireTurns, 1)

	intersections := make(map[pos]bool)
	for p:= range *path1 {
		if (*path2)[p] {
			intersections[p] = true
		}
	}

	//fmt.Printf("%v\n", intersections)
	//fmt.Printf("%v\n", path1)
	//fmt.Printf("%v\n", path2)
	min := math.MaxInt32
	for p := range intersections {
		d := distance(p)
		if d < min {
			min = d
		}
	}
	return min
}

func distance(p pos) int{
	return Abs(p.x) + Abs(p.y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func walkpath(wireTurns []string, i int) *map[pos]bool{
	path := make(map[pos]bool)
	turns := strings.Split(wireTurns[i], ",")
	pos := pos{0, 0}
	for i := range turns {
		length, err := strconv.Atoi(turns[i][1:])
		if err != nil {
			log.Fatal("failed to parse")
		}
		pos = walk(&path, pos, turns[i][0], length)
	}
	return &path
}

func walk(m *map[pos]bool, p pos, direction uint8, length int) pos {
	var newPos pos
	for i := 1; i <= length; i++ {
		switch direction {
		case 'U' : newPos = pos{p.x, p.y+i}
		case 'D' : newPos = pos{p.x, p.y-i}
		case 'L' : newPos = pos{p.x -i, p.y}
		case 'R' : newPos = pos{p.x+i, p.y}
		default:
			log.Fatal("Unknown direction")
		}
		(*m)[newPos] = true
	}
	return newPos
}

func part2() int {
	return algo(6)
}

func algo(n int) int {

	return 2
}
