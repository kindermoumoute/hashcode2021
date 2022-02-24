package main

import (
	"strings"

	"hashcode2022/models"
)

func ParseInput(s string) models.Input {
	in := models.Input{
		AllIngredients: make(map[string]*models.Ingredient),
	}
	lines := strings.Split(s, "\n")
	for i := 0; i < (len(lines)-1)/2; i++ {
		currentClient := &models.Client{
			ID:      i,
			Like:    map[string]*models.Ingredient{},
			Dislike: map[string]*models.Ingredient{},
		}
		in.Clients = append(in.Clients, currentClient)

		likedIngredientStrs := strings.Split(lines[i*2+1], " ")[1:]
		dislikedIngredientStrs := strings.Split(lines[i*2+2], " ")[1:]
		for liked, ingredientStrs := range [][]string{likedIngredientStrs, dislikedIngredientStrs} {
			for _, ingredientStr := range ingredientStrs {
				ingredient, exist := in.AllIngredients[ingredientStr]
				if !exist {
					ingredient = &models.Ingredient{
						Name: ingredientStr,
					}
					in.AllIngredients[ingredientStr] = ingredient
				}

				if liked == 0 {
					ingredient.LikedBy = append(ingredient.LikedBy, currentClient)
					currentClient.Like[ingredient.Name] = ingredient
				} else {
					ingredient.DislikedBy = append(ingredient.DislikedBy, currentClient)
					currentClient.Dislike[ingredient.Name] = ingredient
				}
			}
		}
	}
	return in
}
