package encoder

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

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
