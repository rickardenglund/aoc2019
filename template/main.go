package main

import (
	"aoc2019/computer"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %v in %v\n", p2, time.Since(start2))
}

func part1() int {
	c := computer.Computer{}
	c.ReadMemory("day7/input.txt")
	mem := c.Mem

	max := 0
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						seq := []int{a, b, c, d, e}
						//fmt.Printf("%v\n", seq)
						if uniqueNumbers(seq) {
							out := CalcSignal(seq, mem)
							if out > max {
								max = out
								log.Printf("%v: %v\n", seq, max)
							}
						}
					}
				}
			}
		}
	}
	return max
}

func uniqueNumbers(seq []int) bool {
	ns := make(map[int]bool)
	for _, n := range seq {
		if ns[n] {
			return false
		}
		ns[n] = true
	}
	return true
}

func part2() int {
	return -1
}

func CalcSignal(sequence []int, mem []int) int {
	previousValue := 0
	for i := range sequence {
		c := computer.Computer{Mem: mem}
		c.SetInput([]int{sequence[i], previousValue})
		c.Run()
		previousValue = c.GetLastOutput()
	}

	return previousValue
}

func getSequence(in int) []int {
	var seq []int
	mul := 10
	for in > 0 {
		seq = append(seq, in%mul)
		in = in / mul
	}
	return reverseInts(seq)
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
