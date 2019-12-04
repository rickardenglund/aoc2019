package main

import (
	"fmt"
	"strconv"
)


func main() {
	fmt.Printf("part1: %v\n", part1())
	fmt.Printf("part1: %v\n", part2())
}

const min = 231832
const max = 767346

func part1() int{
	n := 0
	for i := min; i <= max; i++ {
		if isValid(i, loose) {
			n++
		}
	}
	return n
}

func part2() int {
	n := 0
	for i := min; i <= max; i++ {
		if isValid(i, strict) {
			n++
		}
	}
	return n
}

func strict (i int) bool {
	return i == 2
}
func loose (i int) bool {
	return i > 1
}


func isValid(n int, isTuple func(i int) bool) bool{
	str := strconv.Itoa(n)
	numbers := make(map[uint8]int)
	for i := 0; i < len(str) - 1; i++{
		if str[i + 1] < str[i] { return false}
		numbers[str[i]]++
	}
	numbers[str[len(str)-1]]++

	for _, value := range numbers {
		if isTuple(value) {
			return true
		}
	}
	return false
}