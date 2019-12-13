package main

import (
	computer "aoc2019/computer"
	"fmt"
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
						if uniqueNumbers(seq) {
							out := CalcSignal(seq, mem)
							if out > max {
								max = out
							}
						}
					}
				}
			}
		}
	}
	return max
}

func part2() int {
	mem := computer.ReadMemory("day7/input.txt")
	max := 0
	for a := 5; a <= 9; a++ {
		for b := 5; b <= 9; b++ {
			for c := 5; c <= 9; c++ {
				for d := 5; d <= 9; d++ {
					for e := 5; e <= 9; e++ {
						seq := []int{a, b, c, d, e}
						if uniqueNumbers(seq) {
							out := CalcSignalFeedback(seq, mem)
							if out > max {
								max = out
							}
						}
					}
				}
			}
		}
	}
	return max
}

func CalcSignalFeedback(seq, mem []int) int {
	const N = 5
	var computers = make([]computer.Computer, N)
	logCh := make(chan string)
	for i := 0; i < N; i++ {
		computers[i] = computer.NewComputerWithName(string('A'+i), mem)
		computers[i].LogChannel = &logCh
		computers[i].Output = make(chan computer.Msg, 1)
	}

	go computer.LogTask(logCh, false)

	for i := 0; i < N; i++ {
		j := (i + 1) % N
		computers[j].Input = computers[i].Output
	}

	for i := 0; i < N; i++ {
		go computers[i].Run()
		computers[i].Input <- computer.Msg{Sender: "Init", Data: seq[i]}
	}

	computers[0].Input <- computer.Msg{
		Sender: "Init",
		Data:   0,
	}

	<-computers[4].HaltChannel
	return computers[N-1].GetLastOutput()

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

func CalcSignal(sequence []int, mem []int) int {
	previousValue := 0
	for i := range sequence {
		c := computer.NewComputer(mem)
		c.Output = make(chan computer.Msg)
		go c.Run()
		c.Input <- computer.Msg{Sender: "Main", Data: sequence[i]}
		c.Input <- computer.Msg{Sender: "Main", Data: previousValue}
		response := <-c.Output
		previousValue = response.Data
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
