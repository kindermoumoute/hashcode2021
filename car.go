package main

type Car struct {
	ID    int
	Route []*Street

	// Scoring
	duration    float64
	GlobalScore float64
}

func (c *Car) GetPathDuration() float64 {
	if c.duration != 0.0 {
		return c.duration
	}
	var duration float64
	for _, street := range c.Route {
		duration += float64(street.TimeToTravelSeconds)
	}
	c.duration = duration
	return duration
}
