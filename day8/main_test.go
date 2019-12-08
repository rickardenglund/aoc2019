package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getLayers(t *testing.T) {
	code := "123456789012"
	layers := getLayers(code, 3, 2)
	assert.Equal(t, [][]rune{
		0: {'1', '2', '3', '4', '5', '6'},
		1: {'7', '8', '9', '0', '1', '2'},
	}, layers)
}

func Test_count(t *testing.T) {
	l := []rune{'7', '8', '9', '8', '1', '2'}
	c := count('8', l)
	assert.Equal(t, 2, c)
}

func Test_combine(t *testing.T) {
	data := "0222112222120000"
	layers := getLayers(data, 2, 2)
	combined := combineLayers(layers, 2, 2)
	printImage(combined, 2)
}

func Test_print(t *testing.T) {
	printImage([]rune{'a', 'b', 'a', 'b'}, 2)
}
