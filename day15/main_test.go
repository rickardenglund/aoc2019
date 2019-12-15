package main

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func Test_flow(t *testing.T) {
	m := map[pos]int{
		{0, 0}:   open,
		{1, 0}:   open,
		{-1, 0}:  open,
		{1, -1}:  open,
		{2, -1}:  open,
		{-1, -1}: open,
		{-1, -2}: open,
		{0, -2}:  open,
	}
	assert2.Equal(t, 4, flow(m, []pos{{0, 0}}))
	//drawSquare(m, 5)
}
