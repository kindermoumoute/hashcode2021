package solution

type Library struct {
	DaysForSignup   int
	DaysForShipping int
	BookIDs         []int
	Score           float64
}

func (l *Library) Scoring() {
	var score float64

	for _, bookID := range l.BookIDs {

	}

	l.Score = score
}
