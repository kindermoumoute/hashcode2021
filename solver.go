package main

import "fmt"

type SolverParameters struct {
	Input
	//Param1 int
}

func Solve(params SolverParameters) *Solution {
	fmt.Println("There are", len(params.Teams), "teams")
	fmt.Println("total ingredients", len(params.Ingredients))
	fmt.Println("total pizzas", len(params.Pizzas))
	//for _, team := range params.Input.Teams {
	//	fmt.Println("There are", team.Count, "teams of size", team.TeamSize)
	//}
	//for ingredient, pizzas := range params.Input.PizzaPerIngredients {
	//	fmt.Println("pizza IDs with ingredient", ingredient, len(pizzas))
	//	//for _, pizza := range pizzas {
	//	//	fmt.Print(pizza.ID, ",")
	//	//}
	//	//fmt.Println()
	//}
	return &Solution{
		PizzaTeams: []*PizzaTeam{
			{
				Pizzas: params.Pizzas,
			},
		},
	}
}
