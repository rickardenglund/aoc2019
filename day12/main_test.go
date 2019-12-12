package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_gravity(t *testing.T) {
	moons := []*moon{
		{vec{-1, 0, 2}, vec{}},
		{vec{2, -10, -7}, vec{}},
		{vec{4, -8, 8}, vec{}},
		{vec{3, 5, -1}, vec{}},
	}

	grav(moons)
	applyVel(moons)
	assert.Equal(t, 3, moons[0].vel.x)
	assert.Equal(t, -1, moons[0].vel.y)
	assert.Equal(t, -1, moons[0].vel.z)

	assert.Equal(t, 1, moons[1].vel.x)
	assert.Equal(t, 3, moons[1].vel.y)
	assert.Equal(t, 3, moons[1].vel.z)

	assert.Equal(t, -3, moons[2].vel.x)
	assert.Equal(t, 1, moons[2].vel.y)
	assert.Equal(t, -3, moons[2].vel.z)

	assert.Equal(t, 2, moons[0].pos.x)
	assert.Equal(t, -1, moons[0].pos.y)
	assert.Equal(t, 1, moons[0].pos.z)

	assert.Equal(t, 3, moons[1].pos.x)
	assert.Equal(t, -7, moons[1].pos.y)
	assert.Equal(t, -4, moons[1].pos.z)
}

func Test_pot1(t *testing.T) {
	moons := []*moon{
		{vec{-1, 0, 2}, vec{}},
		{vec{2, -10, -7}, vec{}},
		{vec{4, -8, 8}, vec{}},
		{vec{3, 5, -1}, vec{}},
	}
	for i := 0; i < 10; i++ {
		grav(moons)
		applyVel(moons)
	}

	assert.Equal(t, 179, energy(moons))
}

func Test_repeat1(t *testing.T) {
	moons := []*moon{
		{vec{-1, 0, 2}, vec{}},
		{vec{2, -10, -7}, vec{}},
		{vec{4, -8, 8}, vec{}},
		{vec{3, 5, -1}, vec{}},
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
	assert.Equal(t, 2772, i)
}

func Test_repeat2(t *testing.T) {
	moons := []*moon{
		{vec{-8, -10, 0}, vec{}},
		{vec{5, 5, 10}, vec{}},
		{vec{2, -7, 3}, vec{}},
		{vec{9, -8, -3}, vec{}},
	}
	var i int
	visited := map[string]bool{toString(moons): true}
	start := time.Now()
	for {
		grav(moons)
		applyVel(moons)
		i++

		if visited[toString(moons)] {
			break
		}
		const Batch = 100000
		if i%Batch == 0 {
			fmt.Printf("%v: %v\n", i, time.Since(start)/Batch)
			start = time.Now()
		}

	}
	assert.Equal(t, 4686774924, i)
}
