package main

type Street struct {
	Name                string
	TimeToTravelSeconds int
	StartIntersection   *Intersection
	EndIntersection     *Intersection

	// Data used for solver
	NbrTimeUsed int
	Score       float64
}
