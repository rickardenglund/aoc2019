package main

import (
	"aoc2019/inputs"
	"fmt"
	"log"
	"strconv"
)


func main() {
	fmt.Printf("part1: %v\n", part1())
	fmt.Printf("part1: %v\n", part2())
}

func part1() int{
	lines := inputs.GetLines("day1/input.txt")

	var totalFuel int
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		totalFuel += fuelNeeded(mass)
	}
	return totalFuel
}

func fuelNeeded(mass int) int {
	return mass/3 - 2
}

func fuelNeededAdvanced(mass int) int {
	val := fuelNeeded(mass)
	if val < 0 {
		return 0
	}

	return val + fuelNeededAdvanced(val)
}

func part2() int {
	lines := inputs.GetLines("day1/input.txt")

	var totalFuel int
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		totalFuel += fuelNeededAdvanced(mass)
	}
	return totalFuel
}
