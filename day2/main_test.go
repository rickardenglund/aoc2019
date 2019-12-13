package main

import "testing"

func Test_part2(t *testing.T) {
	c := make(chan int, 1)

	c <- 2
}
