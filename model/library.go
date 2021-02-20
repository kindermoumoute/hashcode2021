package model

import (
	"fmt"
)

type Library struct {
	ID            int
	DaysForSignup int
	BooksPerDay   int

	// Should be sorted by score
	Books []*Book

	// For algo of resolution (2021) only
	Score float64

	// For solution scoring only
	NumDaySignedUp *int
	BooksToScan    []*Book
}

func (l *Library) GetAssignedBooks() []*Book {
	var books []*Book
	for _, book := range l.Books {
		if book.ScannedBy == l {
			books = append(books, book)
		}
	}

	return books
}

func (l *Library) GetScannedBooksID() string {
	var IDs string
	for _, book := range l.Books {
		if book.ScannedBy == l {
			IDs += fmt.Sprintf("%d ", book.ID)
		}
	}

	return IDs
}

func (l *Library) Scoring() {
	var score float64

	for _, book := range l.Books {
		score += float64(book.Score)
	}

	l.Score = score
}

func (l *Library) ScoreOfTheDay(numDay int) float64 {
	if l.NumDaySignedUp == nil {
		return 0
	}

	var (
		nbrBookScanned int
		score          float64
	)
	for _, book := range l.BooksToScan {
		score += float64(book.Score)
		nbrBookScanned++
		if nbrBookScanned >= l.BooksPerDay {
			break
		}
	}
	l.BooksToScan = l.BooksToScan[nbrBookScanned:]
	return score
}
