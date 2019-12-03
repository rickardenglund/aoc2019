package main

import (
	"fmt"
	"testing"
)

func Test_walkpath(t *testing.T) {
	fmt.Printf("%v\n", walkpath([]string{"U2,L2"}, 0))
}