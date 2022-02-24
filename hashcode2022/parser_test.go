package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"hashcode2022/models"

	"github.com/maxatome/go-testdeep/td"
)

func TestParse(t *testing.T) {
	in := ParseInput(`3
2 cheese peppers
0
1 basil
1 pineapple
2 mushrooms tomatoes
1 basil`)

	td.Cmp(t, in.Clients, td.Len(3))
	td.Cmp(t, in.AllIngredients, td.Len(6))
}

func TestGenerator(t *testing.T) {
	in := ParseInput(d)
	max := 1707.0
	for {
		m := map[string]*models.Ingredient{}
		for ingName := range in.AllIngredients {
			if rand.Int()%3 == 0 {
				continue
			}
			m[ingName] = in.AllIngredients[ingName]
		}
		p := models.NewPizzaFromMap(m)
		p.ComputeScore(in)
		if p.Score > max {
			s := []string{strconv.Itoa(len(m))}
			for n := range m {
				s = append(s, n)
			}
			fmt.Println(strings.Join(s, " "))
			fmt.Println(p.Score)
			max = p.Score
		}
	}
}

func rndIng() int {
	return (rand.Int() % 9999) + 1
}
