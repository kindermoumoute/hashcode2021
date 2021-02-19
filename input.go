package main

import (
	_ "embed"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

//go:embed input/a_example.txt
var inputA string

//go:embed input/b_read_on.txt
var inputB string

//go:embed input/c_incunabula.txt
var inputC string

//go:embed input/d_tough_choices.txt
var inputD string

//go:embed input/e_so_many_books.txt
var inputE string

//go:embed input/f_libraries_of_the_world.txt
var inputF string

func DecodeInput(s string) Input {
	input := Input{}
	lines := strings.Split(s, "\n")
	firstLine := pkg.ParseIntList(lines[0], " ")
	totalLibraries := firstLine[1]
	input.DaysForScanning = firstLine[2]

	for id, score := range pkg.ParseIntList(lines[1], " ") {
		input.Books = append(input.Books, &Book{
			Id:    id,
			Score: score,
		})
	}

	currentLine := 2
	for i := 0; i < totalLibraries; i++ {
		libData := pkg.ParseIntList(lines[currentLine], " ")
		input.Libraries = append(input.Libraries, &Library{
			DaysForSignup:   libData[1],
			DaysForShipping: libData[2],
			BookIds:         pkg.ParseIntList(lines[currentLine+1], " "),
		})

		currentLine += 2
	}

	return input
}

type Input struct {
	DaysForScanning int
	Books           []*Book
	Libraries       []*Library
}

type Book struct {
	Id    int
	Score int
}

type Library struct {
	DaysForSignup   int
	DaysForShipping int
	BookIds         []int
}
