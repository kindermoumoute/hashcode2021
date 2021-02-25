package main

import "strconv"

type Solution struct {
	Intersections []*IntersectionSolution
}

type IntersectionSolution struct {
	ID              int
	StreetSolutions []*StreetSolution
}

type StreetSolution struct {
	Name               string
	GreenLightDuration int
}

func (s *Solution) Output() []byte {
	output := ""
	output += strconv.Itoa(len(s.Intersections)) + "\n"
	for _, intersection := range s.Intersections {
		output += strconv.Itoa(intersection.ID) + "\n"
		output += strconv.Itoa(len(intersection.StreetSolutions)) + "\n"
		for _, street := range intersection.StreetSolutions {
			output += street.Name + " " + strconv.Itoa(street.GreenLightDuration) + "\n"
		}
	}

	return []byte(output)
}

func (s *Solution) Scoring() int {
	sum := 0

	return sum
}
