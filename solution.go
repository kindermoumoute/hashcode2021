package main

import "fmt"

type Solution struct {
}

type SolverParameters struct {
	Input Input
	//Param1 int
}

func Solve(params SolverParameters) *Solution {
	fmt.Println("There are", len(params.Input.Teams), "teams")
	fmt.Println("total ingredients", len(params.Input.Ingredients))
	fmt.Println("total pizzas", len(params.Input.Pizzas))
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
	return &Solution{}
}

func (s *Solution) Output() []byte {
	return []byte("plop")
}

func (s *Solution) Scoring() {

}
