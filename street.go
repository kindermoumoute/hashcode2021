package main

type Street struct {
	Name                string
	TimeToTravelSeconds int
	StartIntersection   *Intersection
	EndIntersection     *Intersection
}
