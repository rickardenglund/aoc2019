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
		if isValid(i) {
			n++
		}
	}
	return n
}

func isValid(n int) bool{
	str := strconv.Itoa(n)
	hasDouble := false
	for i := 0; i < len(str) - 1; {
		if str[i + 1] < str[i] { return false}
		gs := groupSize(str, i)
		if gs > 1 {
			hasDouble = true
		}
		i++; gs++
	}
	return hasDouble
}

func part2() int {
	n := 0
	for i := min; i <= max; i++ {
		if isValidStrict(i) {
			println(i)
			n++
		}
	}
	return n
}

func isValidStrict(n int) bool{
	str := strconv.Itoa(n)
	hasDouble := false
	for i := 0; i < len(str) - 1; {
		if str[i + 1] < str[i] { return false}
		gs := groupSize(str, i)
		if gs == 2 {
			hasDouble = true
		}

		if gs == 1 {
			i++
		} else {
			i += gs -1
		}
	}
	return hasDouble
}

func groupSize(str string, i int) int {
	n := str[i]
	gs := 1
	for {
		if i == len(str) -1 || str[i + 1] != n{
			return gs
		}
		gs++
		i++
	}


}

func isStrictPair(str string, i int) bool {
	if str[i] == str[i+1] {
		if i == 0 || i == 4 {
			if i == 0 && str[i+3] != str[i] {
				return true
			}
			if i == 4 && str[i-1] != str[i] {
				return true
			}
		} else
		if str[i-1] != str[i] && str[i+1] != str[i+2] {
			return true
		}
	}
	return false
}
