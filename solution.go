package main

import "fmt"

type Solution struct {
}

type SolverParameters struct {
	Input Input
	//Param1 int
}

func Solve(params SolverParameters) *Solution {
	for _, team := range params.Input.Teams {
		fmt.Println("There are", team.Count, "teams of size", team.TeamSize)
	}
	fmt.Println("total ingredients", len(params.Input.Ingredients))
	for ingredient, pizzas := range params.Input.PizzaPerIngredients {
		fmt.Println("pizza IDs with ingredient", ingredient)
		for _, pizza := range pizzas {
			fmt.Print(pizza.ID, ",")
		}
		fmt.Println()
	}
	return &Solution{}
}

func (s *Solution) Output() []byte {
	return []byte("plop")
}

func (s *Solution) Scoring() {

}
