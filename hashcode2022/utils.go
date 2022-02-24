package main

import "gonum.org/v1/gonum/stat/combin"

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
