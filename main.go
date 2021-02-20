package main

import (
	_ "embed"
	"fmt"
	"sort"

	"github.com/kindermoumoute/hashcode2021/encoder"
)

var (
	//go:embed input/a_example.txt
	inputA string

	//go:embed input/b_read_on.txt
	inputB string

	//go:embed input/c_incunabula.txt
	inputC string

	//go:embed input/d_tough_choices.txt
	inputD string

	//go:embed input/e_so_many_books.txt
	inputE string

	//go:embed input/f_libraries_of_the_world.txt
	inputF string
)

func main() {
	file := inputA
	input := encoder.DecodeInput(file)
	fmt.Println(input)
}

func stats(i encoder.Input) {
	fmt.Println("Days for scanning", i.NbrDays)

	scoreCount := make(map[int]int)
	for _, book := range i.Books {
		scoreCount[book.Score]++
	}
	scoreCountSlice := [][2]int(nil)
	for score, count := range scoreCount {
		scoreCountSlice = append(scoreCountSlice, [2]int{score, count})
	}
	sort.Slice(scoreCountSlice, func(i, j int) bool {
		return scoreCountSlice[i][1] > scoreCountSlice[j][1]
	})

	fmt.Println("Score distribution")
	fmt.Println(scoreCountSlice)

}
