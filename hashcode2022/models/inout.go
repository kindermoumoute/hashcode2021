package models

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type Input struct {
	AllIngredients map[string]*Ingredient
	Clients        []*Client
}

type Output struct {
	Pizza Pizza
}

func (in Input) FindBestPizza() Pizza {
	var ingredients []*Ingredient
	for _, ingr := range in.AllIngredients {
		ingredients = append(ingredients, ingr)
		ingr.ComputeScore()
	}

	sort.Slice(ingredients, func(i, j int) bool {
		return ingredients[i].Score > ingredients[j].Score
	})

	var bestPizza Pizza

	// nbrIngr := 6120

	for n := 6126; n < len(ingredients); n += 500 {
		pizzaTemp := NewPizza(ingredients[:n])
		pizzaTemp.ComputeScore(in)
		if pizzaTemp.Score > bestPizza.Score {
			bestPizza = pizzaTemp
			eOutput := GenerateOutput(Output{Pizza: bestPizza})
			fmt.Println(n, ioutil.WriteFile("output/eV2.out", []byte(eOutput), 0755), bestPizza.Score)
		}
	}

	return bestPizza
}

func (in Input) FindBestPizzaV2(paramLoop6, paramLoop8 int) Pizza {
	var ingredients []*Ingredient
	for _, ingr := range in.AllIngredients {
		ingredients = append(ingredients, ingr)
		ingr.ComputeScore()
	}

	// 1. Ajouter tous les ingrédients non détestés
	ingredientsNotDisliked := make(map[string]*Ingredient)

	for _, ingredient := range ingredients {
		if len(ingredient.DislikedBy) == 0 {
			ingredientsNotDisliked[ingredient.Name] = ingredient
		}
	}

	basePizza := Pizza{Ingredients: ingredientsNotDisliked}
	basePizza.ComputeScore(in)

	// 2. Supprimer tous les ingrédients qui ne sont pas aimés
	// var ingredientsLiked []*Ingredient
	// for _, ingredient := range ingredients {
	// 	if len(ingredient.LikedBy) == 0 {
	// 		continue
	// 	}
	// 	ingredientsLiked = append(ingredientsLiked, ingredient)
	// }

	// 3. Mettre de côté les clients.DoesLikePizza(bestPizza) && (client.ComplexityScore > 1)
	var simpleClients []*Client
	var complexClients []*Client

	for _, client := range in.Clients {
		if client.DoesLikePizza(basePizza) && client.PositiveRatio() {
			simpleClients = append(simpleClients, client)
		} else {
			complexClients = append(complexClients, client)
		}
	}

	// Loop start
	clientsToProcess := make([]*Client, len(simpleClients))
	copy(clientsToProcess, simpleClients)

	for nbrLoop8 := 0; nbrLoop8 < paramLoop8; nbrLoop8++ {
		for nbrLoop6 := 0; nbrLoop6 < paramLoop6; nbrLoop6++ {
			// 4. simpleClients.sort(client.score)
			sort.Slice(clientsToProcess, func(i, j int) bool {
				return clientsToProcess[i].Score > clientsToProcess[j].Score
			})

			// 5. Pour chaque clientToProcess :
			//		if doesNotLike(bestPizza) (becauseDislikedIngredients)
			// 			delete(ingredients);
			//			if score decreases => cancel
			//		else if doesNotLike(bestPizza) (becauseMissingLikedIngredient)
			//	 		add(missingIngredients)
			// 			if bestPizza.score decreases => cancel
			for _, clientToProcess := range clientsToProcess {
				if dislikedIngredients := clientToProcess.DislikedIngredients(basePizza); len(dislikedIngredients) > 0 {
					pizzaUpdated := basePizza.RemoveIngredients(dislikedIngredients)
					pizzaUpdated.ComputeScore(in)

					if pizzaUpdated.Score > basePizza.Score {
						basePizza = pizzaUpdated
					}
				} else if missingLikedIngredients := clientToProcess.MissingLikedIngredients(basePizza); len(missingLikedIngredients) > 0 {
					pizzaUpdated := basePizza.RemoveIngredients(missingLikedIngredients)
					pizzaUpdated.ComputeScore(in)

					if pizzaUpdated.Score > basePizza.Score {
						basePizza = pizzaUpdated
					}
				}
			}

			// 6. Loop while bestPizza is optimizable
		}

		// 7. Add complexClients to clientsToProcess
		clientsToProcess = append(clientsToProcess, complexClients...)

		// 8. Loop to 4 (twice because 6. loops to 3.)
	}

	return basePizza
}
