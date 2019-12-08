package main

import (
	"aoc2019/inputs"
	"fmt"
	"log"
	"math"
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

const width = 25
const height = 6

func part1() int {
	lines := inputs.GetLines("day8/input.txt")
	image := lines[0]

	layers := getLayers(image, width, height)

	var minZeroLayer int
	var minZeroCount int = math.MaxInt32
	for l := range layers {
		c := count('0', layers[l])
		if c < minZeroCount {
			minZeroCount = c
			minZeroLayer = l
		}
	}

	return count('1', layers[minZeroLayer]) * count('2', layers[minZeroLayer])
}

func part2() int {
	lines := inputs.GetLines("day8/input.txt")
	image := lines[0]

	layers := getLayers(image, width, height)
	combined := combineLayers(layers, width, height)
	printImage(combined, width)
	return -1
}

func getLayers(image string, w, h int) (layers [][]rune) {
	layerSize := w * h
	nLayers := (len(image) + 1) / layerSize
	layers = make([][]rune, nLayers)
	for currentLayer := 0; currentLayer < nLayers; currentLayer++ {
		layer := make([]rune, layerSize)
		for layerPos := 0; layerPos < layerSize; layerPos++ {
			layer[layerPos] = rune(image[currentLayer*layerSize+layerPos])
		}
		layers[currentLayer] = layer
	}
	return
}

func count(target rune, layer []rune) int {
	m := make(map[rune]int)
	for _, c := range layer {
		m[c]++
	}

	return m[target]
}

func combineLayers(layers [][]rune, w, h int) (combined []rune) {
	combined = make([]rune, w*h)
	for p := 0; p < w*h; p++ {
		comb := '0'
		for l := 0; l < len(layers); l++ {
			switch layers[l][p] {
			case '0':
				comb = '░'
			case '1':
				comb = '█'
			case '2':
				continue
			default:
				log.Fatal("Invalid pixel")
			}
			if comb != '0' {
				break
			}
		}
		combined[p] = comb
	}
	return
}

func printImage(combined []rune, w int) {
	for i := range combined {
		fmt.Printf("%c%c", combined[i], combined[i])
		if i%w == w-1 {
			fmt.Printf("\n")
		}
	}
}
