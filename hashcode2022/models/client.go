package models

import "fmt"

type Client struct {
	ID                       int
	ComplexityScore          float64
	CombinationLikedScore    float64
	CombinationDislikedScore float64
	Score                    float64
	Like                     map[string]*Ingredient
	Dislike                  map[string]*Ingredient
}

/*
	Computes average number of ingredients on pizza containing that ingredient
*/
func (c *Client) ComputeCombinationLikedScore() {
	if c.CombinationLikedScore != 0.0 {
		return
	}

	var totalLikingClients int

	for _, likedIngredient := range c.Like {
		totalLikingClients += len(likedIngredient.LikedBy)
	}

	c.CombinationLikedScore = float64(totalLikingClients) / float64(len(c.Like))
}

func (c *Client) ComputeCombinationDislikedScore() {
	if c.CombinationDislikedScore != 0.0 {
		return
	}

	var totalDislikingClients int

	for _, dislikedIngredient := range c.Dislike {
		totalDislikingClients += len(dislikedIngredient.LikedBy)
	}

	c.CombinationDislikedScore = float64(totalDislikingClients) / float64(len(c.Dislike))
}

/*
	Simple complexity (like / dislike)
*/
func (c *Client) ComputeComplexityScore() {
	if c.CombinationLikedScore != 0.0 {
		return
	}

	c.ComplexityScore = float64(len(c.Like)) / float64(len(c.Dislike))
}

/*
  - Ratio ingrédients aimés/détestés
  - Pondéré par le nombre d'ingrédients aimés
*/
func (c *Client) ComputeScore() {
	c.Score = c.ComplexityScore / c.CombinationLikedScore
}

func (c *Client) String() string {
	return fmt.Sprintf("\n client %d (like: %d, dislike %d)", c.ID, len(c.Like), len(c.Dislike))
}

/*
	- No disliked ingredient is in Pizza
	- All liked ingredients are in Pizza
*/
func (c Client) DoesLikePizza(pizza Pizza) bool {
	if len(c.Like) > len(pizza.Ingredients) {
		return false
	}

	for _, ingredient := range pizza.Ingredients {
		if _, exists := c.Dislike[ingredient.Name]; exists {
			return false
		}
	}

	for _, ingredientLiked := range c.Like {
		if _, exists := pizza.Ingredients[ingredientLiked.Name]; !exists {
			return false
		}
	}

	return true
}

func (c Client) DislikedIngredients(pizza Pizza) map[string]*Ingredient {
	ingredientsDisliked := make(map[string]*Ingredient)

	for _, ingredient := range pizza.Ingredients {
		if ingredient, exists := c.Dislike[ingredient.Name]; exists {
			ingredientsDisliked[ingredient.Name] = ingredient
		}
	}

	return ingredientsDisliked
}

func (c Client) MissingLikedIngredients(pizza Pizza) map[string]*Ingredient {
	ingredientsLiked := make(map[string]*Ingredient)

	for _, ingredient := range pizza.Ingredients {
		if _, exists := c.Like[ingredient.Name]; !exists {
			ingredientsLiked[ingredient.Name] = ingredient
		}
	}

	return ingredientsLiked
}

func (c Client) PositiveRatio() bool {
	return c.ComplexityScore > 10
}
