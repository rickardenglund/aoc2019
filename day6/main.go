package main

import (
	"fmt"
	"time"
)


func main() {
	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %v in %v\n", p1, time.Now().Sub(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %v in %v\n", p2, time.Now().Sub(start2))
}

func part1() int{
	return -1
}

func part2() int {
	return -1
}
