package main

import (
	"aoc2019/inputs"
	"flag"
	"fmt"
	"github.com/bxcodec/saint"
	"strconv"
	"strings"
	"time"
)

var gui bool //nolint:unused

func main() {
	guiPtr := flag.Bool("gui", false, "Add --gui flag to enable graphics")
	flag.Parse()
	gui = *guiPtr

	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %-10v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

func part1() interface{} {
	input := inputs.GetLine("day22/input.txt")
	deck := getDeck(10007)
	execute(deck, input)

	for i := range deck {
		if deck[i] == 2019 {
			return i
		}
	}
	return -1
}

func part2() interface{} {
	input := inputs.GetLine("day22/input.txt")
	deck := getDeck(101741582076661)
	execute(deck, input)

	for i := range deck {
		if deck[i] == 2020 {
			return i
		}
	}
	return -1
}

//7350 to high
func execute(deck []int, instructions string) {
	lines := strings.Split(instructions, "\n")
	for _, line := range lines {
		if line == "deal into new stack" {
			dealIntoNewStack(deck)
		}
		if strings.HasPrefix(line, "deal with increment") {
			parts := strings.Split(line, " ")
			n, _ := strconv.Atoi(parts[len(parts)-1])
			dealWithIncrement(deck, n)
		}
		if strings.HasPrefix(line, "cut") {
			parts := strings.Split(line, " ")
			n, _ := strconv.Atoi(parts[len(parts)-1])
			cut(deck, n)
		}
	}
}

func dealWithIncrement(deck []int, inc int) {
	tmp := make([]int, len(deck))
	copy(tmp, deck)

	for i := 0; i < len(deck); i++ {
		deck[i*inc%len(deck)] = tmp[i]
	}
}

func cut(deck []int, n int) {
	if n == 0 {
		return
	}

	if n > 0 {
		tmp := make([]int, n)
		copy(tmp, deck[:n])
		for i := 0; i < len(deck)-n; i++ {
			deck[i] = deck[i+n]
		}
		for i := range tmp {
			deck[len(deck)-n+i] = tmp[i]
		}
		return
	}

	//negative
	tmp := make([]int, saint.Abs(n))
	copy(tmp, deck[len(deck)+n:])
	copy(deck[-n:], deck[:len(deck)+n])
	copy(deck, tmp)
}

func dealIntoNewStack(deck []int) {
	var tmp int
	for i := 0; i < len(deck)/2; i++ {
		tmp = deck[i]
		backIndex := len(deck) - 1 - i
		deck[i] = deck[backIndex]
		deck[backIndex] = tmp
	}
}

func getDeck(size int) []int {
	deck := make([]int, size)
	for i := range deck {
		deck[i] = i
	}
	return deck
}
