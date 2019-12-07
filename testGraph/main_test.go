package main

import (
	"fmt"
	"github.com/soniakeys/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_apa(t *testing.T) {
	g := graph.LabeledAdjacencyList{
		1: {{To: 2, Label: 7}, {To: 2, Label: 5}, {To: 3, Label: 9}, {To: 6, Label: 11}},
		2: {{To: 3, Label: 10}, {To: 4, Label: 15}},
		3: {{To: 4, Label: 11}, {To: 6, Label: 2}},
		4: {{To: 5, Label: 7}},
		6: {{To: 5, Label: 9}},
	}
	p, d := g.DijkstraPath(1, 2, weight)
	fmt.Println("Shortest path:", p)
	fmt.Println("Path distance:", d)
}

func weight(w graph.LI) float64 {
	return float64((1))
}

func Test_bepa(t *testing.T) {
	g := graph.LabeledAdjacencyList{
		1: {{To: 2}},
		2: {{To: 3}},
		3: {},
	}

	//g = append(g, 4 : []graph.Half{})
	_, dist := g.DijkstraPath(3, 1, weight)
	assert.Equal(t, float64(2), dist)
}

func Test_cepa(t *testing.T) {
	g := graph.LabeledUndirected{graph.LabeledAdjacencyList{
		1: {{To: 2}},
		2: {{To: 3}},
		3: {},
	}}
	_, dist := g.DijkstraPath(1, 3, weight)

	assert.InDelta(t, 2, dist, 0.1)
}
