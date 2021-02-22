package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
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
}

type Ingredient struct {
	Name string
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
