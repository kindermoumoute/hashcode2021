package main

import (
	"fmt"
	"sort"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/willf/bitset"
)

type SolverParameters struct {
	Input
	//Param1 int
}

func Solve(params SolverParameters) *Solution {
	fmt.Println("There are", len(params.Teams), "teams")
	fmt.Println("total ingredients", len(params.Ingredients))
	fmt.Println("total pizzas", len(params.Pizzas))

	cptI := uint(0)
	for _, ingredient := range params.Ingredients {
		ingredient.ID = cptI
		cptI++
	}
	for _, pizza := range params.Input.Pizzas {
		pizza.IngredientsB = bitset.New(uint(len(params.Ingredients)))

		for _, ingredient := range pizza.Ingredients {
			pizza.IngredientsB.Set(ingredient.ID)
		}
	}

	nbrTeam2 := params.Input.GetBinomesCount()
	nbrTeam3 := params.Input.GetTrinomesCount()
	nbrTeam4 := params.Input.GetQuadrinomesCount()
	fmt.Println("max per team : ", nbrTeam2, nbrTeam3, nbrTeam4)

	seuilMax := 0.95
	seuilMin := 0.1

	var pizzas2 []*Pizza2

	// d'abord, on fait des groupes de 2 pizzas
	// à optimiser, comme on commence arbitrairement par assigner la meilleur pizza avec la première du tableau
	// il faudrait modifier ça
	// peut-être d'abord trier les pizzas par nombre d'ingrédients ?

	for _, pizzaA := range params.Pizzas {
		if pizzaA.Used {
			continue
		}
		var bestPizza *Pizza
		bestScore := 0.0
		for _, pizzaB := range params.Pizzas {
			if pizzaA == pizzaB || pizzaB.Used {
				continue
			}
			localScore := pizzaA.ScoreWith23(*pizzaB, float64(len(params.Ingredients)))
			if localScore > bestScore {
				bestScore = localScore
				bestPizza = pizzaB
			}
		}
		if bestPizza == nil {
			continue
		}
		if bestScore < seuilMin {
			continue
		}
		pizzaA.Used = true
		bestPizza.Used = true
		p2 := &Pizza2{
			pizzaA:       pizzaA,
			pizzaB:       bestPizza,
			IngredientsB: bitset.New(uint(len(params.Ingredients))),
			Score:        bestScore,
		}
		p2.IngredientsB = pizzaA.IngredientsB.Union(bestPizza.IngredientsB)
		if bestScore > seuilMax {
			p2.Locked = true
		}
		pizzas2 = append(pizzas2, p2)
		if len(pizzas2) >= nbrTeam2+nbrTeam3+nbrTeam4 {
			break
		}
	}

	fmt.Println("pizzas 2 done", len(pizzas2))

	var pizzas3 []*Pizza3

	for _, pizza := range params.Input.Pizzas {
		if pizza.Used {
			continue
		}

		var bestPizza2 *Pizza2
		var numBestPizza2 int
		bestScore := 0.0
		for numPizza2, pizza2 := range pizzas2 {
			if pizza2.Locked {
				continue
			}
			localScore := pizza2.ScoreWithPizza23(*pizza, float64(len(params.Ingredients)))
			if localScore > bestScore {
				bestScore = localScore
				bestPizza2 = pizza2
				numBestPizza2 = numPizza2
			}
		}
		if bestPizza2 == nil {
			continue
		}
		if bestScore < seuilMin {
			continue
		}
		pizza.Used = true
		p3 := &Pizza3{
			Pizzas2:      bestPizza2,
			pizzaC:       pizza,
			IngredientsB: bitset.New(uint(len(params.Ingredients))),
			Score:        bestScore,
		}
		pizzas2[len(pizzas2)-1], pizzas2[numBestPizza2] = pizzas2[numBestPizza2], pizzas2[len(pizzas2)-1]
		pizzas2 = pizzas2[:len(pizzas2)-1]
		p3.IngredientsB = bestPizza2.IngredientsB.Union(pizza.IngredientsB)
		if bestScore > seuilMax {
			p3.Locked = true
		}
		pizzas3 = append(pizzas3, p3)
		if len(pizzas3) >= nbrTeam3+nbrTeam4 {
			break
		}
	}

	fmt.Println("pizzas 3 done")

	var pizzas4 []*Pizza4

	for _, pizza := range params.Input.Pizzas {
		if pizza.Used {
			continue
		}

		var bestPizza3 *Pizza3
		var numBestPizza3 int
		bestScore := 0.0
		for numPizza3, pizza3 := range pizzas3 {
			if pizza3.Locked {
				continue
			}
			localScore := pizza3.ScoreWithPizza23(*pizza, float64(len(params.Ingredients)))
			if localScore > bestScore {
				bestScore = localScore
				bestPizza3 = pizza3
				numBestPizza3 = numPizza3
			}
		}
		if bestPizza3 == nil {
			continue
		}
		if bestScore < seuilMin {
			continue
		}
		pizza.Used = true
		p4 := &Pizza4{
			Pizzas3:      bestPizza3,
			pizzaD:       pizza,
			IngredientsB: bitset.New(uint(len(params.Ingredients))),
			Score:        bestScore,
		}
		pizzas3[len(pizzas3)-1], pizzas3[numBestPizza3] = pizzas3[numBestPizza3], pizzas3[len(pizzas3)-1]
		pizzas3 = pizzas3[:len(pizzas3)-1]
		p4.IngredientsB = bestPizza3.IngredientsB.Union(pizza.IngredientsB)
		pizzas4 = append(pizzas4, p4)
		if len(pizzas4) >= nbrTeam4 {
			break
		}
	}

	fmt.Println("pizzas 4 done")

	// on devrait avoir une dernière étape pour casserles groupes de pizzas si on a trop
	// de group de pizza et pas assez de team

	for _, pizza := range params.Input.Pizzas {
		if pizza.Used {
			continue
		}

		bestPizza2 := &Pizza2{}
		bestScore2 := 0.0
		bestPizza3 := &Pizza3{}
		bestScore3 := 0.0
		if len(pizzas3) < nbrTeam3 {
			for _, pizza2 := range pizzas2 {
				localScore := pizza2.ScoreWithPizza23(*pizza, float64(len(params.Ingredients)))
				if localScore > bestScore2 {
					bestScore2 = localScore
					bestPizza2 = pizza2
				}
			}
		}
		if len(pizzas4) < nbrTeam4 {
			for _, pizza3 := range pizzas3 {
				localScore := pizza3.ScoreWithPizza23(*pizza, float64(len(params.Ingredients)))
				if localScore > bestScore3 {
					bestScore3 = localScore
					bestPizza3 = pizza3
				}
			}
		}
		bestScore2 = bestScore2 * bestScore2
		bestScore3 = bestScore3 * bestScore3

		pizza.Used = true
		switch {
		case bestScore3 > bestScore2:
			p3 := &Pizza3{Pizzas2: bestPizza2, pizzaC: pizza}
			p3.IngredientsB = bestPizza3.IngredientsB.Union(pizza.IngredientsB)
			pizzas3 = append(pizzas3, p3)
		case bestScore2 > bestScore3:
			p4 := &Pizza4{Pizzas3: bestPizza3, pizzaD: pizza}
			p4.IngredientsB = bestPizza3.IngredientsB.Union(pizza.IngredientsB)
			pizzas4 = append(pizzas4, p4)
		}
	}

	sort.Slice(pizzas2, func(i, j int) bool {
		return pizzas2[i].Score < pizzas2[j].Score
	})
	pizzas2 = pizzas2[:pkg.Min(len(pizzas2), nbrTeam2)]

	sort.Slice(pizzas3, func(i, j int) bool {
		return pizzas3[i].Score < pizzas3[j].Score
	})
	pizzas3 = pizzas3[:pkg.Min(len(pizzas3), nbrTeam3)]

	sort.Slice(pizzas4, func(i, j int) bool {
		return pizzas4[i].Score < pizzas4[j].Score
	})
	pizzas4 = pizzas4[:pkg.Min(len(pizzas4), nbrTeam4)]
	fmt.Println("solver done", len(pizzas2), len(pizzas3), len(pizzas4))

	return pizzasToSolution(pizzas2, pizzas3, pizzas4)
}

func pizzasToSolution(pizza2s []*Pizza2, pizza3s []*Pizza3, pizza4s []*Pizza4) *Solution {
	solution := &Solution{}
	for _, pt2 := range pizza2s {
		solution.PizzaTeams = append(solution.PizzaTeams, &PizzaTeam{
			TeamSize: 2,
			Pizzas:   []*Pizza{pt2.pizzaA, pt2.pizzaB},
		})
	}
	for _, pt3 := range pizza3s {
		solution.PizzaTeams = append(solution.PizzaTeams, &PizzaTeam{
			TeamSize: 3,
			Pizzas:   []*Pizza{pt3.Pizzas2.pizzaA, pt3.Pizzas2.pizzaB, pt3.pizzaC},
		})
	}
	for _, pt4 := range pizza4s {
		solution.PizzaTeams = append(solution.PizzaTeams, &PizzaTeam{
			TeamSize: 4,
			Pizzas: []*Pizza{
				pt4.Pizzas3.Pizzas2.pizzaA,
				pt4.Pizzas3.Pizzas2.pizzaB,
				pt4.Pizzas3.pizzaC,
				pt4.pizzaD,
			},
		})
	}

	return solution
}
