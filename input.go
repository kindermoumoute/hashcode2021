package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	bitbitset "github.com/willf/bitset"
)

type RawInput struct {
	Raw      string
	FileName string
}

type Input struct {
	Pizzas []*Pizza
	Teams  []*Team

	// Usefull data
	Ingredients         map[string]*Ingredient
	PizzaPerIngredients map[string][]*Pizza
}

type Pizza struct {
	ID          int
	Ingredients []*Ingredient

	// chacun des tableaux devraient avoir la mmême taille, pour chaque pizza
	IngredientsB *bitbitset.BitSet
	Used         bool
}

// region getters
func (i Input) GetGroupSizeCount(size int) int {
	for _, team := range i.Teams {
		if team.TeamSize == size {
			return team.Count
		}
	}
	return 0
}

func (i Input) GetBinomesCount() int {
	for _, team := range i.Teams {
		if team.TeamSize == 2 {
			return team.Count
		}
	}
	return 0
}

func (i Input) GetTrinomesCount() int {
	for _, team := range i.Teams {
		if team.TeamSize == 3 {
			return team.Count
		}
	}
	return 0
}

func (i Input) GetQuadrinomesCount() int {
	for _, team := range i.Teams {
		if team.TeamSize == 4 {
			return team.Count
		}
	}
	return 0
}

// endregion getters

func (p Pizza) ScoreWith(pizzaB Pizza) int {
	return int(p.IngredientsB.SymmetricDifferenceCardinality(pizzaB.IngredientsB))
}

// Pour les couples I_i, I_j on calcule les distance:
//
// 	d_{i, j} = sum(!I_i & I_j) / sum(!I_i)
// 	d_{j, i} = sum(!I_j & I_i) / sum(!I_j)
//
// 	D_{i, j} = (min(d_{i, j}, d_{j, i})) * ((1 - (sum(I_i & I_j) / NB_DISTINCT_INGREDIENTS))) * ((sum(I_i | I_j) / NB_DISTINCT_INGREDIENTS))
//
func (p Pizza) ScoreWith23(pizzaB Pizza, distinctIngredient float64) float64 {
	A := p.IngredientsB
	B := pizzaB.IngredientsB
	notA := A.Complement()
	notB := B.Complement()
	var AOverB float64
	var BOverA float64
	if int(notA.Count()) == 0 {
		AOverB = 1.0
	} else {
		AOverB = float64(notA.Intersection(B).Count()) / float64(notA.Count()) // d_{a, b}
	}
	if int(notB.Count()) == 0 {
		BOverA = 1.0
	} else {
		BOverA = float64(notB.Intersection(A).Count()) / float64(notB.Count()) // d_{a, b}
	}

	AUnionB := float64(A.Union(B).Count())
	AIntersectB := float64(A.Intersection(B).Count())

	fmt.Println(notA.Intersection(A).Count())
	fmt.Println(notA.Intersection(B).Count())

	return math.Min(AOverB, BOverA) *
		((1 - AIntersectB) / distinctIngredient) *
		(AUnionB / distinctIngredient)
}

type Pizza2 struct {
	pizzaA       *Pizza
	pizzaB       *Pizza
	IngredientsB *bitbitset.BitSet // toutes les pizzas ont ce slice de la meme longueur
	Score        float64
	Locked       bool
}

func (p2 Pizza2) ScoreWithPizza(pizzaB Pizza) int {
	return int(p2.IngredientsB.SymmetricDifferenceCardinality(pizzaB.IngredientsB))
}

func (p2 Pizza2) ScoreWithPizza23(pizzaB Pizza, distinctIngredient float64) float64 {
	A := p2.IngredientsB
	B := pizzaB.IngredientsB
	notA := A.Complement()
	notB := B.Complement()
	var AOverB float64
	var BOverA float64
	if int(notA.Count()) == 0 {
		AOverB = 1.0
	} else {
		AOverB = float64(notA.Intersection(B).Count()) / float64(notA.Count()) // d_{a, b}
	}
	if int(notB.Count()) == 0 {
		BOverA = 1.0
	} else {
		BOverA = float64(notB.Intersection(A).Count()) / float64(notB.Count()) // d_{a, b}
	}

	AUnionB := float64(A.Union(B).Count())
	AIntersectB := float64(A.Intersection(B).Count())

	return math.Min(AOverB, BOverA) *
		((1 - AIntersectB) / distinctIngredient) *
		(AUnionB / distinctIngredient)
}

type Pizza3 struct {
	Pizzas2      *Pizza2
	pizzaC       *Pizza
	IngredientsB *bitbitset.BitSet // toutes les pizzas ont ce slice de la meme longueur
	Score        float64
	Locked       bool
}

func (p3 Pizza3) ScoreWithPizza(pizzaC Pizza) int {
	return int(p3.IngredientsB.SymmetricDifferenceCardinality(pizzaC.IngredientsB))
}

func (p3 Pizza3) ScoreWithPizza23(pizzaB Pizza, distinctIngredient float64) float64 {
	A := p3.IngredientsB
	B := pizzaB.IngredientsB
	notA := A.Complement()
	notB := B.Complement()
	var AOverB float64
	var BOverA float64
	if int(notA.Count()) == 0 {
		AOverB = 1.0
	} else {
		AOverB = float64(notA.Intersection(B).Count()) / float64(notA.Count()) // d_{a, b}
	}
	if int(notB.Count()) == 0 {
		BOverA = 1.0
	} else {
		BOverA = float64(notB.Intersection(A).Count()) / float64(notB.Count()) // d_{a, b}
	}

	AUnionB := float64(A.Union(B).Count())
	AIntersectB := float64(A.Intersection(B).Count())

	return math.Min(AOverB, BOverA) *
		((1 - AIntersectB) / distinctIngredient) *
		(AUnionB / distinctIngredient)
}

type Pizza4 struct {
	Pizzas3      *Pizza3
	pizzaD       *Pizza
	IngredientsB *bitbitset.BitSet // toutes les pizzas ont ce slice de la meme longueur
	Score        float64
}

func (p4 Pizza4) ScoreWithPizza(pizzaD Pizza) int {
	return int(p4.IngredientsB.SymmetricDifferenceCardinality(pizzaD.IngredientsB))
}

func (p4 Pizza4) ScoreWithPizza23(pizzaB Pizza, distinctIngredient float64) float64 {
	A := p4.IngredientsB
	B := pizzaB.IngredientsB
	notA := A.Complement()
	notB := B.Complement()
	var AOverB float64
	var BOverA float64
	if int(notA.Count()) == 0 {
		AOverB = 1.0
	} else {
		AOverB = float64(notA.Intersection(B).Count()) / float64(notA.Count()) // d_{a, b}
	}
	if int(notB.Count()) == 0 {
		BOverA = 1.0
	} else {
		BOverA = float64(notB.Intersection(A).Count()) / float64(notB.Count()) // d_{a, b}
	}

	AUnionB := float64(A.Union(B).Count())
	AIntersectB := float64(A.Intersection(B).Count())

	return math.Min(AOverB, BOverA) *
		((1 - AIntersectB) / distinctIngredient) *
		(AUnionB / distinctIngredient)
}

type Ingredient struct {
	Name string
	ID   uint
}

type Team struct {
	TeamSize int
	Count    int
}

func DecodeInput(s string) Input {
	input := Input{
		Ingredients:         make(map[string]*Ingredient),
		PizzaPerIngredients: make(map[string][]*Pizza),
	}

	lines := strings.Split(s, "\n")
	header := pkg.ParseIntList(lines[0], " ")
	for teamType, count := range header[1:] {
		input.Teams = append(input.Teams, &Team{
			TeamSize: teamType + 2,
			Count:    count,
		})
	}

	for id, line := range lines[1:] {
		if len(line) == 0 { // trailing line
			break
		}
		pizza := &Pizza{
			ID: id,
		}
		for _, newIngredient := range strings.Split(line, " ")[1:] {
			ingredient, exist := input.Ingredients[newIngredient]
			if !exist {
				ingredient = &Ingredient{
					Name: newIngredient,
				}
				input.Ingredients[newIngredient] = ingredient
			}
			pizza.Ingredients = append(pizza.Ingredients, ingredient)
			input.PizzaPerIngredients[newIngredient] = append(input.PizzaPerIngredients[newIngredient], pizza)
		}
		input.Pizzas = append(input.Pizzas, pizza)
	}

	return input
}
