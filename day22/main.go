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
	inputs := inputs.GetLine("day22/input.txt")
	res := executeFollow(2019, 10007, inputs)
	return res
}

// too high : 89749636864860
func part2() interface{} {
	times := 101741582076661
	pos := 2020
	inputs := inputs.GetLine("day22/input.txt")
	prev := 0
	for i := 0; i < times; i++ {
		pos = executeFollow(pos, 101741582076661, inputs)

		fmt.Printf("diff: %v\n", pos-prev)
		if i%100_000 == 0 {
			fmt.Printf("progres: %.10f \n", float64(i)/float64(times))
		}
	}
	return pos
}

func executeFollow(pos int, deckLength int, instructions string) int {
	lines := strings.Split(instructions, "\n")
	for _, line := range lines {
		if line == "deal into new stack" {
			pos = dealIntoNewStackFollow(pos, deckLength)
		}
		if strings.HasPrefix(line, "deal with increment") {
			parts := strings.Split(line, " ")
			n, _ := strconv.Atoi(parts[len(parts)-1])
			pos = dealWithIncrementFollow(pos, n, deckLength)
		}
		if strings.HasPrefix(line, "cut") {
			parts := strings.Split(line, " ")
			n, _ := strconv.Atoi(parts[len(parts)-1])
			pos = cutFollow(pos, n, deckLength)
		}
	}

	return pos
}

func dealIntoNewStackFollow(pos, deckLength int) int {
	return (deckLength - 1) - pos
}

func cutFollow(pos, N, decklength int) int {
	if N > 0 {
		if pos < N {
			return decklength - N + pos
		}
		return pos - N
	}

	return (pos + saint.Abs(N)) % decklength
}

func dealWithIncrementFollow(pos, N, deckLength int) int {
	return (pos * N) % deckLength
}
