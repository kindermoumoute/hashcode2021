package main

import (
	"math"
	"sort"

	"go.uber.org/zap"
)

type SolverParameters struct {
	Input

	AlphaSort float64
	Beta      float64
	Gamma     float64
}

func Solve(log *zap.SugaredLogger, params SolverParameters) *Solution {
	// Log basic info about this input
	log.Infof("There are %d intersections", len(params.Intersections))
	log.Infof("There are %d streets", len(params.Streets))
	log.Infof("There are %d cars", len(params.Cars))

	// Compute data structure for this solver here (bitsets, tree, graph)

	for _, street := range params.Streets {
		street.EndIntersection.StreetEnds = append(street.EndIntersection.StreetEnds, street)
	}
	log.Infof("streets added to StreetEnds")
	sort.SliceStable(params.Cars, func(i, j int) bool {
		return params.Cars[i].GetPathDuration() < params.Cars[j].GetPathDuration()
	})

	nbCars := float64(len(params.Cars))
	for i, car := range params.Cars {
		ordonancementScorePerCar := math.Exp(params.Beta*(nbCars-float64(i))/nbCars + 1.0)
		ponderationSurLesBonusDeTemps := math.Exp(params.Gamma*(params.AlphaSort+(float64(params.SimulationTimeSeconds)-car.GetPathDuration())/float64(params.DestinationScore)) + 1.0)
		y := math.Log(math.Max(1, ordonancementScorePerCar*ponderationSurLesBonusDeTemps))
		// log.Infof("%d GetPathDuration %f", i, car.GetPathDuration())
		// log.Infof("%d ordonancementScorePerCar %f", i, ordonancementScorePerCar)
		// log.Infof("%d ponderationSurLesBonusDeTemps %f", i, ponderationSurLesBonusDeTemps)
		// log.Infof("%d y %f", i, y)
		car.GlobalScore = math.Max(params.AlphaSort, y)
	}

	for _, car := range params.Cars {
		for _, street := range car.Route {
			street.Score += car.GlobalScore
		}
	}
	log.Infof("Score computed")
	s := &Solution{}
	for _, noeud := range params.Intersections {
		var utilsiationFeus []float64
		minimum := math.MaxFloat64
		maximum := 0.0
		for _, feu := range noeud.StreetEnds {
			utilsiationFeus = append(utilsiationFeus, feu.Score)
			if minimum > feu.Score {
				minimum = feu.Score
			}
			if maximum < feu.Score {
				maximum = feu.Score
			}
		}
		if minimum == 0 {
			minimum = 1
		}
		if maximum == 0 {
			maximum = 1
		}

		interSolu := &IntersectionSolution{
			ID: noeud.ID,
		}
		for _, feu := range noeud.StreetEnds {
			var time int
			// time = int(math.Ceil(feu.Score / minimum))
			// if time > params.SimulationTimeSeconds/len(noeud.StreetEnds)/10 {
			// 	time = params.SimulationTimeSeconds / len(noeud.StreetEnds) / 10
			// }
			// if time <= 0 {
			// 	time = 1
			// 	// continue
			// }
			time = 1
			if maximum != 0.0 {
				time = int(feu.Score*5/maximum) + 1
			}

			interSolu.StreetSolutions = append(interSolu.StreetSolutions, &StreetSolution{
				Name:               feu.Name,
				GreenLightDuration: time,
			})
		}
		if len(interSolu.StreetSolutions) == 0 {
			for _, feu := range noeud.StreetEnds {
				interSolu.StreetSolutions = append(interSolu.StreetSolutions, &StreetSolution{
					Name:               feu.Name,
					GreenLightDuration: 1,
				})
			}
		}
		s.Intersections = append(s.Intersections, interSolu)
	}

	return s
}

// Note: we can use workerpool for parallelizing intensive calculus:
//	wp := workerpool.New(2)
//	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
//
//	for _, r := range requests {
//		r := r
//		wp.Submit(func() {
//			fmt.Println("Handling request:", r)
//		})
//	}
//
//	wp.StopWait()
//
