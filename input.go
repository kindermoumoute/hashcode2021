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
	SimulationTimeSeconds int
	DestinationScore      int
	Intersections         []*Intersection
	Streets               []*Street
	Cars                  []*Car

	// Usefull data
	StreetsByName map[string]*Street
}

func DecodeInput(s string) Input {
	input := Input{
		StreetsByName: make(map[string]*Street),
	}

	lines := strings.Split(s, "\n")
	header := pkg.ParseIntList(lines[0], " ")
	input.SimulationTimeSeconds = header[0]
	input.DestinationScore = header[4]

	totalIntersections := header[1]
	totalStreets := header[2]
	totalCars := header[3]

	for i := 0; i < totalIntersections; i++ {
		input.Intersections = append(input.Intersections, &Intersection{
			ID: i,
		})
	}

	lineIndex := 1
	for i := 0; i < totalStreets; i++ {
		s := &Street{}
		fields := strings.Split(lines[lineIndex], " ")
		s.StartIntersection = input.Intersections[pkg.MustAtoi(fields[0])]
		s.EndIntersection = input.Intersections[pkg.MustAtoi(fields[1])]
		s.Name = fields[2]
		s.TimeToTravelSeconds = pkg.MustAtoi(fields[3])
		input.Streets = append(input.Streets, s)
		input.StreetsByName[s.Name] = s
		lineIndex++
	}

	for i := 0; i < totalCars; i++ {
		c := &Car{}
		fields := strings.Split(lines[lineIndex], " ")
		for _, streetName := range fields[1:] {
			c.Route = append(c.Route, input.StreetsByName[streetName])
		}

		input.Cars = append(input.Cars, c)
		lineIndex++
	}

	return input
}
