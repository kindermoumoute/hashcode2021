package model

type Book struct {
	ID        int
	Score     int
	ScannedBy *Library
}
