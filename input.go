package main

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type RawInput struct {
	Raw      string
	FileName string
}

type Input struct {
	NbrDays   int
	Books     []*Book
	Libraries []*Library
}

func DecodeInput(s string) Input {
	input := Input{}

	lines := strings.Split(s, "\n")
	firstLine := pkg.ParseIntList(lines[0], " ")

	input.NbrDays = firstLine[2]

	input.Books = []*Book(nil)
	input.Libraries = []*Library(nil)

	for bookID, bookScore := range pkg.ParseIntList(lines[1], " ") {
		input.Books = append(input.Books, &Book{
			ID:    bookID,
			Score: bookScore,
		})
	}

	for libID, numLine := 0, 2; numLine < len(lines); libID, numLine = libID+1, numLine+2 {
		if len(lines[numLine]) == 0 {
			break
		}
		libraryDetails := pkg.ParseIntList(lines[numLine], " ")
		library := &Library{
			ID:            libID,
			Books:         make([]*Book, libraryDetails[0]),
			DaysForSignup: libraryDetails[1],
			BooksPerDay:   libraryDetails[2],
		}

		booksID := pkg.ParseIntList(lines[numLine+1], " ")
		for _, bookID := range booksID {
			library.Books = append(library.Books, input.Books[bookID])
		}

		input.Libraries = append(input.Libraries, library)
	}

	return input
}
