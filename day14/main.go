package main

import (
	"aoc2019/inputs"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func init() {
	fuel = []ingredient{{1, "FUEL"}}
}

var fuel []ingredient

func main() {
	start := time.Now()
	p1 := part1()
	fmt.Printf("part1: %-10v in %v\n", p1, time.Since(start))
	start2 := time.Now()
	p2 := part2()
	fmt.Printf("part2: %-10v in %v\n", p2, time.Since(start2))
}

func part1() int {
	input := inputs.GetLine("day14/input.txt")
	recipes := readRecipes(input)
	res := simplify(recipes, fuel)

	return res[0].amount
}

//const totalOre = 1000000000000

// to low 11758757
func part2() int {
	input := inputs.GetLine("day14/input.txt")
	recipes := readRecipes(input)

	return fuelForOre(recipes)
}

const totalOre = 1000000000000

func fuelForOre(recipes map[string]recipe) int {
	upperBound := totalOre
	lowerBound := 0
	cur := upperBound / 2
	for {
		highOk := canCreate(recipes, cur+1)
		lowOk := canCreate(recipes, cur)

		if !highOk && lowOk {
			break
		} else if highOk {
			lowerBound = cur
		} else if !lowOk {
			upperBound = cur
		}
		cur = lowerBound + (upperBound-lowerBound)/2
		if upperBound == lowerBound+1 {
			return -1
		}
	}
	return cur
}

func canCreate(recipes map[string]recipe, n int) bool {
	res := simplify(recipes, []ingredient{{n, "FUEL"}})
	return res[0].amount < totalOre
}

type recipe struct {
	output      int
	ingredients []ingredient
}

type ingredient struct {
	amount int
	name   string
}

func readRecipes(s string) map[string]recipe {
	lines := strings.Split(s, "\n")
	recipes := map[string]recipe{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " => ")
		ingredientStrings := strings.Split(parts[0], ", ")

		result := getIngredient(parts[1])
		var ingredients []ingredient
		for _, in := range ingredientStrings {
			ingredients = append(ingredients, getIngredient(in))
		}
		recipe := recipe{
			output:      result.amount,
			ingredients: ingredients,
		}
		recipes[result.name] = recipe
	}
	return recipes
}

func getIngredient(in string) ingredient {
	ps := strings.Split(in, " ")
	amount, _ := strconv.Atoi(ps[0])
	return ingredient{
		amount: amount,
		name:   ps[1],
	}
}

func simplify(r map[string]recipe, ingredients []ingredient) []ingredient {
	res, _ := simplifyWithSpares(r, ingredients, map[string]int{})
	return res
}

func simplifyWithSpares(r map[string]recipe, curRequired []ingredient, spares map[string]int) ([]ingredient, map[string]int) {
	curSpares := spares
	for {
		curRequired, spares = useSpares(curRequired, spares)

		nextRequired, nextSpares := simplifyStep(r, curRequired, curSpares)
		//fmt.Printf("%-3v     %v\n", nextSpares, nextRequired)
		if equal(curRequired, nextRequired) && equalDict(curSpares, nextSpares) {
			break
		}
		curRequired = nextRequired
		curSpares = nextSpares
	}
	return curRequired, spares
}

func simplifyStep(r map[string]recipe, requiredIngredients []ingredient, spares map[string]int) ([]ingredient, map[string]int) {
	requiredIngredients, spares = useSpares(requiredIngredients, spares)
	requiredIngredients = merge(requiredIngredients)
	var res []ingredient
	for i := 0; i < len(requiredIngredients); i++ {
		requiredName := requiredIngredients[i].name
		requiredAmount := requiredIngredients[i].amount
		rec := r[requiredName]
		if requiredName == "ORE" {
			res = append(res, requiredIngredients[i])
			continue
		}
		if rec.output > requiredAmount {
			spares[requiredName] += rec.output - requiredAmount
			res = append(res, rec.ingredients...)
		} else if rec.output == requiredAmount {
			res = append(res, rec.ingredients...)
		} else if requiredAmount > rec.output {
			n := requiredAmount / rec.output
			res = append(res, ingredient{requiredAmount % rec.output, requiredName})
			var toAdd []ingredient
			for j := range rec.ingredients {
				toAdd = append(toAdd, ingredient{rec.ingredients[j].amount * n, rec.ingredients[j].name})
			}
			res = append(res, toAdd...)
		}

	}

	res = merge(res)
	return res, spares
}

func useSpares(required []ingredient, spares map[string]int) ([]ingredient, map[string]int) {
	for i := range required {
		nSpare, ok := spares[required[i].name]
		if ok {
			if nSpare > required[i].amount {
				spares[required[i].name] -= required[i].amount
				if spares[required[i].name] == 0 {
					delete(spares, required[i].name)
				}
				required[i].amount = 0
			} else {
				required[i].amount -= nSpare
				delete(spares, required[i].name)
			}
		}
	}
	return required, spares
}

func merge(ingredients []ingredient) []ingredient {
	amounts := map[string]int{}
	res := []ingredient{}
	for _, in := range ingredients {
		amounts[in.name] += in.amount
	}
	for k, v := range amounts {
		if v > 0 {
			res = append(res, ingredient{v, k})
		}
	}

	return res
}

func equal(xs []ingredient, ys []ingredient) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i := range xs {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}
func equalDict(xs, ys map[string]int) bool {
	if len(xs) != len(ys) {
		return false
	}
	for k, v := range xs {
		if ys[k] != v {
			return false
		}
	}
	return true
}
