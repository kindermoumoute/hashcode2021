package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	bitbitset "github.com/willf/bitset"
)

func TestScoreWith23(t *testing.T) {
	onion := Ingredient{"onion", 0}
	pepper := Ingredient{"pepper", 1}
	olive := Ingredient{"olive", 2}
	mushroom := Ingredient{"mushroom", 3}
	tomato := Ingredient{"tomato", 4}
	basil := Ingredient{"basil", 5}

	ingredients1 := bitbitset.New(6).Set(0b000111)
	ingredients2 := bitbitset.New(6).Set(0b110100)

	p1 := Pizza{
		ID: 0,
		Ingredients: []*Ingredient{
			&onion,
			&pepper,
			&olive,
		},
		IngredientsB: ingredients1,
		Used:         false,
	}
	p2 := Pizza{
		ID: 1,
		Ingredients: []*Ingredient{
			&mushroom,
			&tomato,
			&basil,
		},
		IngredientsB: ingredients2,
		Used:         false,
	}
	score := p1.ScoreWith23(p2, 6)

	assert.Equal(t, 1.0, score)
}
