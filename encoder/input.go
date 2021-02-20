package encoder

import (
	"github.com/kindermoumoute/hashcode2021/model"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type Input struct {
	NbrBooks     int
	NbrLibraries int
	NbrDays      int
	Books        []*model.Book
	Libraries    []*model.Library
}

func DecodeInput(s string) Input {
	input := Input{}

	lines := strings.Split(s, "\n")
	firstLine := pkg.ParseIntList(lines[0], " ")

	input.NbrBooks = firstLine[0]
	input.NbrLibraries = firstLine[1]
	input.NbrDays = firstLine[2]

	input.Books = make([]*model.Book, input.NbrBooks)
	input.Libraries = make([]*model.Library, input.NbrLibraries)


	for bookID, bookScore := range pkg.ParseIntList(lines[1], " ") {
		input.Books = append(input.Books, &model.Book{

			ID:    bookID,
			Score: bookScore,
		})
	}

	for libID, numLine := 0, 2; numLine < len(lines); libID, numLine = libID+1, numLine+2 {
		if len(lines[numLine]) == 0 {
			break
		}
		libraryDetails := pkg.ParseIntList(lines[numLine], " ")
		library := &model.Library{
			ID:            libID,
			Books:         make([]*model.Book, libraryDetails[0]),
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
