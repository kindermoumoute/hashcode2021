package models

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
)

type Ingredient struct {
	Name                     string
	ComplexityScore          float64
	CombinationLikedScore    float64
	CombinationDislikedScore float64
	Score                    float64
	LikedBy                  []*Client
	DislikedBy               []*Client
}

/*
	Computes average number of ingredients on pizza containing that ingredient
*/
func (i *Ingredient) ComputeCombinationLikedScore() {
	if i.CombinationLikedScore != 0.0 {
		return
	}

	var totalIngredientsLiked int

	for _, likedBy := range i.LikedBy {
		totalIngredientsLiked += len(likedBy.Like)
	}

	i.CombinationLikedScore = float64(totalIngredientsLiked) / float64(len(i.LikedBy))
}

func (i *Ingredient) ComputeCombinationDislikedScore() {
	if i.CombinationDislikedScore != 0.0 {
		return
	}

	var totalIngredientsDisliked int

	for _, dislikedBy := range i.DislikedBy {
		totalIngredientsDisliked += len(dislikedBy.Dislike)
	}

	i.CombinationDislikedScore = float64(totalIngredientsDisliked) / float64(len(i.DislikedBy))
}

/*
	Simple complexity (liked / disliked)
*/
func (i *Ingredient) ComputeComplexityScore() {
	if i.ComplexityScore != 0.0 {
		return
	}

	i.ComplexityScore = float64(len(i.LikedBy)) / float64(len(i.DislikedBy))
}

func (i *Ingredient) ComputeScore() {
	if i.ComplexityScore == 0.0 {
		i.ComputeComplexityScore()
	}

	if i.CombinationLikedScore == 0.0 {
		i.ComputeCombinationLikedScore()
	}

	if i.CombinationDislikedScore == 0.0 {
		i.ComputeCombinationDislikedScore()
	}

	i.Score = i.ComplexityScore / ((i.CombinationLikedScore*5 - i.CombinationDislikedScore) + 100000)
}

func (i *Ingredient) String() string {
	return fmt.Sprintf(" (liked: %d, disliked %d)\n", len(i.LikedBy), len(i.LikedBy))
}

func Combination(size, elements int, c chan []int, limit int) {
	var nbr int
	gen := combin.NewCombinationGenerator(size, elements)

	for gen.Next() {
		c <- gen.Combination(nil)

		nbr++
		if nbr >= limit {
			break
		}
	}

	close(c)
}
