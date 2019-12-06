package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var input string
var input2 string

func init() {
	input = (`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`)

	input2 = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
}

func Test_createTree(t *testing.T) {
	tree := createTree(strings.Split(input, "\n"))
	assert.Equal(t, 11, len(tree))
}

func Test_getNParents(t *testing.T) {
	tree := createTree(strings.Split(input, "\n"))
	assert.Equal(t, 7, getNParents("L", &tree))
}

func Test_getParents(t *testing.T) {
	tree := createTree(strings.Split(input, "\n"))
	parents := getParents("D", rootNode, &tree)
	assert.Equal(t, "D,C,B,COM", parents)
}

func Test_getParents2(t *testing.T) {
	tree := createTree(strings.Split(input, "\n"))
	parents := getParents("D", "B", &tree)
	assert.Equal(t, "D,C,B", parents)
}

func Test_getAllconnections(t *testing.T) {
	tree := createTree(strings.Split(input, "\n"))
	assert.Equal(t, 42, getAllConnections(&tree))
}

func Test_getDistance(t *testing.T) {
	tree := createTree(strings.Split(input2, "\n"))
	assert.Equal(t, 4, getDistance("YOU", "SAN", &tree))
}

func Test_getCommanAncestor(t *testing.T) {
	tree := createTree(strings.Split(input2, "\n"))
	assert.Equal(t, "B", getCommonAncestor("D", "H", &tree))
}
