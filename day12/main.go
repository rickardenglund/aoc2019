package main

import (
	"fmt"
	"strconv"
	"strings"
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

func part2() int {
	moons := []*moon{
		{vec{-8, -10, 0}, vec{}},
		{vec{5, 5, 10}, vec{}},
		{vec{2, -7, 3}, vec{}},
		{vec{9, -8, -3}, vec{}},
	}
	var i int
	visited := map[string]bool{toString(moons): true}
	for {
		grav(moons)
		applyVel(moons)
		i++

		if visited[toString(moons)] {
			break
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
	visited := map[string]bool{}

	for ai := range moons {
		for bi := range moons {
			if ai == bi || visited[toStr(ai, bi)] || visited[toStr(bi, ai)] {
				continue
			}
			gravity(moons[ai], moons[bi])
			visited[toStr(ai, bi)] = true
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

func toStr(a, b int) string {
	return strconv.Itoa(a) + strconv.Itoa(b)
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

func toString(moons []*moon) string {
	var sb strings.Builder
	for i := range moons {
		sb.WriteString(fmt.Sprintf("%v,", moons[i]))
	}
	return sb.String()
}
