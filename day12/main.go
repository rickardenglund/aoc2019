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

// 4686774924 to low
func part2() int {
	moons := []*moon{
		{vec{-8, -9, -7}, vec{}},
		{vec{-5, 2, -1}, vec{}},
		{vec{11, 8, -14}, vec{}},
		{vec{1, -4, -11}, vec{}},
	}

	i := findLoop(moons)
	return i
}

func findLoop(moons []*moon) int {
	var i int
	startPos := [4]moon{*moons[0], *moons[1], *moons[2], *moons[3]}
	visited := map[[4]moon]bool{startPos: true}
	start := time.Now()
	for {
		grav(moons)
		applyVel(moons)
		i++

		if visited[[4]moon{*moons[0], *moons[1], *moons[2], *moons[3]}] {
			break
		}

		const Batch = 10000000
		if i%Batch == 0 {
			fmt.Printf("%v: %v\n", i, time.Since(start)/Batch)
			start = time.Now()
		}
	}
	return i
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
