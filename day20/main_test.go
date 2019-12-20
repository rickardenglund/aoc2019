package main

import (
	"aoc2019/position"
	"github.com/stretchr/testify/assert"
	"testing"
)

const small = `         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       `

func Test_readMap(t *testing.T) {
	//gui = true
	m := readMap(small)
	portals := findPortals(m)
	assert.Equal(t, portal{name: "BC"}, portals[position.Pos{X: 2, Y: 8}])
	assert.Equal(t, portal{name: "ZZ"}, portals[position.Pos{X: 13, Y: 16}])
}

func Test_availableMoves(t *testing.T) {
	m := readMap(small)
	portals := findPortals(m)
	//start := findPortal(portals, "AA")
	//stop := findPortal(portals, "BC")
	//moves := getAvailableMoves(m, portals, pos{p: start})
	//tree := toTree(m, portals)
	//printTree(tree, portals)

	cost := findCost(m, portals, position.Pos{X: 9, Y: 2}, position.Pos{X: 9, Y: 6})
	assert.Equal(t, 4, cost)
}

func Test_small(t *testing.T) {
	m := readMap(small)
	portals := findPortals(m)
	starts := findPortalsWithName(portals, "AA")
	assert.Equal(t, 1, len(starts))
	stops := findPortalsWithName(portals, "ZZ")
	assert.Equal(t, 1, len(stops))
	cost := findCost(m, portals, starts[0], stops[0])
	assert.Equal(t, 23, cost)
}

const medium = `                   A               
                   A               
  #################.#############  
  #.#...#...................#.#.#  
  #.#.#.###.###.###.#########.#.#  
  #.#.#.......#...#.....#.#.#...#  
  #.#########.###.#####.#.#.###.#  
  #.............#.#.....#.......#  
  ###.###########.###.#####.#.#.#  
  #.....#        A   C    #.#.#.#  
  #######        S   P    #####.#  
  #.#...#                 #......VT
  #.#.#.#                 #.#####  
  #...#.#               YN....#.#  
  #.###.#                 #####.#  
DI....#.#                 #.....#  
  #####.#                 #.###.#  
ZZ......#               QG....#..AS
  ###.###                 #######  
JO..#.#.#                 #.....#  
  #.#.#.#                 ###.#.#  
  #...#..DI             BU....#..LF
  #####.#                 #.#####  
YN......#               VT..#....QG
  #.###.#                 #.###.#  
  #.#...#                 #.....#  
  ###.###    J L     J    #.#.###  
  #.....#    O F     P    #.#...#  
  #.###.#####.#.#####.#####.###.#  
  #...#.#.#...#.....#.....#.#...#  
  #.#####.###.###.#.#.#########.#  
  #...#.#.....#...#.#.#.#.....#.#  
  #.###.#####.###.###.#.#.#######  
  #.#.........#...#.............#  
  #########.###.###.#############  
           B   J   C               
           U   P   P               `

func Test_medium(t *testing.T) {
	m := readMap(medium)
	portals := findPortals(m)
	starts := findPortalsWithName(portals, "AA")
	assert.Equal(t, 1, len(starts))
	stops := findPortalsWithName(portals, "ZZ")
	assert.Equal(t, 1, len(stops))
	cost := findCost(m, portals, starts[0], stops[0])
	assert.Equal(t, 58, cost)
}

//func printTree(tree []node, portals map[position.Pos]portal) {
//	fmt.Printf("#######\n")
//	for i := range tree {
//		fmt.Printf("%v %+v: ", portals[tree[i].pos].name, tree[i].pos)
//		printMoves(tree[i].moves)
//
//	}
//	fmt.Printf("#######\n")
//}
//func printMoves(v []move) {
//	for i := range v {
//		fmt.Printf("(%v_%v:%v) ", v[i].val, v[i].target, v[i].steps)
//	}
//	fmt.Printf("\n")
//}
