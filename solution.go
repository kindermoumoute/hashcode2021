package main

type Solution struct {
	NbrDays int
	// must be sorted by score
	Libraries []*Library
	Score     float64
}

type SolverParameters struct {
	Input  Input
	Param1 int
}

func Solve(params SolverParameters) *Solution {
	return &Solution{}
}

func (s *Solution) Output() []byte {
	return []byte("plop")
}

func (s *Solution) Scoring() {
	var (
		score       float64
		signUpUntil int
	)

	for numDay := 0; numDay < s.NbrDays; numDay++ {
		for _, library := range s.Libraries {
			score += library.ScoreOfTheDay(numDay)
		}

		if numDay >= signUpUntil {
			for _, library := range s.Libraries {
				if library.NumDaySignedUp == nil {
					*library.NumDaySignedUp = numDay
					signUpUntil = numDay + library.DaysForSignup
					break
				}
			}
		}
	}

	s.Score = score
}
