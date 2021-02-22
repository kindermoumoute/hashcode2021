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
	}
	return []byte(output)
}

func (s *Solution) Scoring() {

}
