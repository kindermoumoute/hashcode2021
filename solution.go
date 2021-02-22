package main

import (
	"strconv"
	"strings"
)

type Solution struct {
	PizzaTeams []*PizzaTeam
}

type PizzaTeam struct {
	TeamSize int
	Pizzas   []*Pizza // Pizzas' length should match TeamSize
}

func (s *Solution) Output() []byte {
	output := strconv.Itoa(len(s.PizzaTeams)) + "\n"
	for _, pizzaTeam := range s.PizzaTeams {
		pizzas := []string(nil)
		for _, p := range pizzaTeam.Pizzas {
			pizzas = append(pizzas, strconv.Itoa(p.ID))
		}
		output += strconv.Itoa(pizzaTeam.TeamSize) + " " + strings.Join(pizzas, " ")
	} // Prepare pizza boolean arrays

	return []byte(output)

	/*
		Pour les couples I_i, I_j on calcule les distance:

		d_{i, j} = sum(!I_i & I_j) / sum(!I_i)
		d_{j, i} = sum(!I_j & I_i) / sum(!I_j)

		D_{i, j} = (min(d_{i, j}, d_{j, i})) * ((1 - (sum(I_i & I_j) / NB_DISTINCT_INGREDIENTS))) * ((sum(I_i | I_j) / NB_DISTINCT_INGREDIENTS))

		On procède par un scoring en binômes (pour répondre au nombre d'équipes de 2)
		Puis, pour les pires pizzas restantes, on les associe en trinômes (pour répondre au nombre d'équipes de 3)
		Et pareil pour les 4


	*/
}

func (s *Solution) Scoring() int {
	sum := 0
	for _, pt := range s.PizzaTeams {
		ingredientBitset := pt.Pizzas[0].IngredientsB.Clone()
		for _, p := range pt.Pizzas[1:] {
			ingredientBitset.InPlaceUnion(p.IngredientsB)
		}
		ingredientCount := int(ingredientBitset.Count())
		sum += ingredientCount * ingredientCount
	}

	return sum
}
