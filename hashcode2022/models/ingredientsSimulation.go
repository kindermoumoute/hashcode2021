package models

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type IngredientsStats struct {
	Ingredient *Ingredient
	ScoreTotal float64
	NbrUsed    int64

	Score float64
}

type PizzasStats struct {
	Size       int
	ScoreTotal float64
	NbrUsed    int64

	Score float64
}

func MarkIngredients(in Input) Pizza {
	rand.Seed(time.Now().Unix())

	// lancer des simulations avec ingrédients randoms
	// noter la pizza et donc noter les ingrédients dedans
	// avoir une moyenne score/nombre de fois dans une pizza

	pizzasStats := make(map[int]PizzasStats)
	ingredientsStats := make(map[string]*IngredientsStats)
	for _, ingredient := range in.AllIngredients {
		ingredientsStats[ingredient.Name] = &IngredientsStats{Ingredient: ingredient}
	}

	var bestPizza Pizza
	var wgGenerator sync.WaitGroup
	var mutex sync.Mutex

	fmt.Println("nbr Ingredients:", len(in.AllIngredients))

	guard := make(chan struct{}, 4)
	for i := 0; i < 10000; i++ {
		wgGenerator.Add(1)
		guard <- struct{}{}
		go func() {
			defer wgGenerator.Done()
			var ingredients []*Ingredient
			// nbrIngredients := rand.Intn(len(in.AllIngredients))
			nbrIngredients := 350 + rand.Intn(50)

			var cpt int
			for _, ingredient := range in.AllIngredients {
				ingredients = append(ingredients, ingredient)

				cpt++
				if cpt >= nbrIngredients {
					break
				}
			}

			pizza := NewPizza(ingredients)
			pizza.ComputeScore(in)

			if pizza.Score > bestPizza.Score {
				bestPizza = pizza
				fmt.Println("new gigh score !", nbrIngredients, pizza.Score)
			}

			mutex.Lock()
			for _, ingredient := range ingredients {
				ingredientsStats[ingredient.Name].ScoreTotal += pizza.Score
				ingredientsStats[ingredient.Name].NbrUsed++
			}
			pizzaStats := pizzasStats[nbrIngredients]
			pizzaStats.ScoreTotal += pizza.Score
			pizzaStats.NbrUsed++
			pizzasStats[nbrIngredients] = pizzaStats
			mutex.Unlock()
			<-guard
		}()
	}
	wgGenerator.Wait()

	var pizzasStatsSorted []PizzasStats
	for size, pizzaStatsSorted := range pizzasStats {
		pizzaStatsSorted.Size = size
		pizzaStatsSorted.Score = pizzaStatsSorted.Score / float64(pizzaStatsSorted.NbrUsed)
		pizzasStatsSorted = append(pizzasStatsSorted, pizzaStatsSorted)
	}
	sort.Slice(pizzasStatsSorted, func(i, j int) bool {
		return pizzasStatsSorted[i].Score > pizzasStatsSorted[j].Score
	})
	fmt.Println("best size is", pizzasStatsSorted[0].Size, pizzasStatsSorted[0].NbrUsed)

	for _, ingredientStats := range ingredientsStats {
		in.AllIngredients[ingredientStats.Ingredient.Name].Score = ingredientStats.ScoreTotal / float64(ingredientStats.NbrUsed)
	}

	var ingredients []*Ingredient
	for _, ingr := range in.AllIngredients {
		ingredients = append(ingredients, ingr)
		// ingr.ComputeScore()
	}

	sort.Slice(ingredients, func(i, j int) bool {
		return ingredients[i].Score > ingredients[j].Score
	})

	fmt.Println("let's generate them")

	for _, p := range pizzasStatsSorted {
		wgGenerator.Add(1)
		guard <- struct{}{}

		go func(size int) {
			defer wgGenerator.Done()
			c := make(chan []int)
			go Combination(len(ingredients), size, c, 1000)
			for set := range c {
				pizzaTemp := NewPizzaFromCombination(ingredients, set)
				pizzaTemp.ComputeScore(in)

				mutex.Lock()
				if pizzaTemp.Score > bestPizza.Score {
					bestPizza = pizzaTemp
					fmt.Println("new gigh score !", size, pizzaTemp.Score)
				}
				mutex.Unlock()
			}
			<-guard
		}(p.Size)
	}
	wgGenerator.Wait()

	return bestPizza
}
