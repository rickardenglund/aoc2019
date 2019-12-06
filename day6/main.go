package main

import (
	"aoc2019/inputs"
	"fmt"
	"log"
	"strings"
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

const rootNode = "COM"

func part1() int {
	lines := inputs.GetLines("day6/input.txt")
	tree := createTree(lines)

	return getAllConnections(&tree)
}

func part2() int {
	lines := inputs.GetLines("day6/input.txt")
	tree := createTree(lines)

	return getDistance("YOU", "SAN", &tree)
}

func getDistance(s string, s2 string, tree *map[string]string) int {
	ancestor := getCommonAncestor(s, s2, tree)
	return getParentDistance((*tree)[s], ancestor, tree) + getParentDistance((*tree)[s2], ancestor, tree)
}

func getCommonAncestor(s string, s2 string, tree *map[string]string) string {
	parents1 := strings.Split(getParents(s, rootNode, tree), ",")
	parents2 := strings.Split(getParents(s2, rootNode, tree), ",")
	for _, parent := range parents1 {
		if contains(parents2, parent) {
			return parent
		}
	}
	log.Fatal("No Common Ancestor")
	return "COM"
}

func contains(parents []string, parent string) bool {
	for i := range parents {
		if parents[i] == parent {
			return true
		}
	}
	return false
}

func getAllConnections(tree *map[string]string) int {
	n := 0
	for key := range *tree {
		n += getNParents(key, tree)
	}
	return n
}

func createTree(lines []string) map[string]string {
	tree := make(map[string]string)
	for _, line := range lines {
		planets := strings.Split(line, ")")
		tree[planets[1]] = planets[0]
	}
	return tree
}

func getNParents(name string, tree *map[string]string) int {
	return getParentDistance(name, rootNode, tree)
}

func getParentDistance(parent, root string, tree *map[string]string) int {
	parents := getParents(parent, root, tree)
	return len(strings.Split(parents, ",")) - 1
}

func getParents(parent, root string, tree *map[string]string) string {
	if parent == root {
		return root
	}

	return parent + "," + getParents((*tree)[parent], root, tree)
}
