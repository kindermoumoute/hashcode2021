package models

type Pizza struct {
	Ingredients map[string]*Ingredient
	Score       float64
}

func NewPizzaFromMap(m map[string]*Ingredient) Pizza {
	return Pizza{
		Ingredients: m,
	}
}

func NewPizzaFromCombination(ingredients []*Ingredient, combinations []int) Pizza {
	m := make(map[string]*Ingredient, len(combinations))
	for _, combination := range combinations {
		m[ingredients[combination].Name] = ingredients[combination]
	}

	return Pizza{
		Ingredients: m,
	}
}

func NewPizza(ingredients []*Ingredient) Pizza {
	m := make(map[string]*Ingredient, len(ingredients))
	for _, ingredient := range ingredients {
		m[ingredient.Name] = ingredient
	}

	return Pizza{
		Ingredients: m,
	}
}

func (p *Pizza) ComputeScore(input Input) {
	var score int

	for _, client := range input.Clients {
		if client.DoesLikePizza(*p) {
			score++
		}
	}

	p.Score = float64(score)
}

func (p Pizza) RemoveIngredients(toDelete map[string]*Ingredient) Pizza {
	ingredientsToAdd := []*Ingredient(nil)
	for _, ingredient := range p.Ingredients {
		if _, exist := toDelete[ingredient.Name]; !exist {
			ingredientsToAdd = append(ingredientsToAdd, ingredient)
		}
	}
	return NewPizza(ingredientsToAdd)
}

func (p Pizza) AddIngredients(toAdd map[string]*Ingredient) Pizza {
	for s, ingredient := range p.Ingredients {
		toAdd[s] = ingredient
	}
	return NewPizzaFromMap(toAdd)
}
