package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %-10v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

//9127
func part1() int {
	moons := []*moon{
		{vec{-8, -9, -7}, vec{}},
		{vec{-5, 2, -1}, vec{}},
		{vec{11, 8, -14}, vec{}},
		{vec{1, -4, -11}, vec{}},
	}
	for i := 0; i < 1000; i++ {
		grav(moons)
		applyVel(moons)
	}
	return energy(moons)
}

func part2() int {
	moons := []*moon{
		{vec{-8, -9, -7}, vec{}},
		{vec{-5, 2, -1}, vec{}},
		{vec{11, 8, -14}, vec{}},
		{vec{1, -4, -11}, vec{}},
	}

	return findLoop(moons)
}

type state struct {
	pos   int //nolint:structcheck
	speed int //nolint:structcheck
}

func findLoop(moons []*moon) int {
	startX := getXState(moons)
	startY := getYState(moons)
	startZ := getZState(moons)

	xVisited := map[[4]state]int{startX: 0}
	yVisited := map[[4]state]int{startY: 0}
	zVisited := map[[4]state]int{startZ: 0}

	var i int
	for {
		grav(moons)
		applyVel(moons)
		i++

		xPrev, xok := xVisited[getXState(moons)]
		yPrev, yok := yVisited[getYState(moons)]
		zPrev, zok := zVisited[getZState(moons)]

		if xok && yok && zok {
			return LCMFold([]int{i - xPrev, i - yPrev, i - zPrev})
		}

		xVisited[getXState(moons)] = i
		yVisited[getYState(moons)] = i
		zVisited[getZState(moons)] = i
	}
}

func getZState(moons []*moon) [4]state {
	return [4]state{
		{moons[0].pos.z, moons[0].vel.z},
		{moons[1].pos.z, moons[1].vel.z},
		{moons[2].pos.z, moons[2].vel.z},
		{moons[3].pos.z, moons[3].vel.z},
	}
}

func getYState(moons []*moon) [4]state {
	return [4]state{
		{moons[0].pos.y, moons[0].vel.y},
		{moons[1].pos.y, moons[1].vel.y},
		{moons[2].pos.y, moons[2].vel.y},
		{moons[3].pos.y, moons[3].vel.y},
	}
}

func getXState(moons []*moon) [4]state {
	return [4]state{
		{moons[0].pos.x, moons[0].vel.x},
		{moons[1].pos.x, moons[1].vel.x},
		{moons[2].pos.x, moons[2].vel.x},
		{moons[3].pos.x, moons[3].vel.x},
	}
}

type vec struct {
	x int
	y int
	z int
}

type moon struct {
	pos vec
	vel vec
}

func applyVel(moons []*moon) {
	for i := range moons {
		moons[i].pos.x += moons[i].vel.x
		moons[i].pos.y += moons[i].vel.y
		moons[i].pos.z += moons[i].vel.z
	}
}

func grav(moons []*moon) {
	nMoons := len(moons)
	for ai := 0; ai < nMoons; ai++ {
		for bi := ai + 1; bi < nMoons; bi++ {
			gravity(moons[ai], moons[bi])
		}
	}
}

func gravity(a, b *moon) {
	if a.pos.x > b.pos.x {
		a.vel.x--
		b.vel.x++
	} else if a.pos.x < b.pos.x {
		a.vel.x++
		b.vel.x--
	}

	if a.pos.y > b.pos.y {
		a.vel.y--
		b.vel.y++
	} else if a.pos.y < b.pos.y {
		a.vel.y++
		b.vel.y--
	}

	if a.pos.z > b.pos.z {
		a.vel.z--
		b.vel.z++
	} else if a.pos.z < b.pos.z {
		a.vel.z++
		b.vel.z--
	}

}

func energy(moons []*moon) int {
	total := 0
	for i := range moons {
		total += pot(moons[i]) * kin(moons[i])
	}
	return total
}

func kin(m *moon) int {
	return Abs(m.vel.x) + Abs(m.vel.y) + Abs(m.vel.z)

}

func pot(m *moon) int {
	return Abs(m.pos.x) + Abs(m.pos.y) + Abs(m.pos.z)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func LCMFold(xs []int) int {
	agg := LCM(xs[0], xs[1])
	for i := 2; i < len(xs); i++ {
		agg = LCM(agg, xs[i])
	}
	return agg
}

func LCM(x, y int) int {
	var i int
	for {
		i++
		if i*x%y == 0 {
			return i * x
		}
	}

}
