package main

import (
	"go.uber.org/zap"
)

type SolverParameters struct {
	Input
	//Param1 int
}

func Solve(log *zap.SugaredLogger, params SolverParameters) *Solution {
	// Log basic info about this input
	log.Infof("There are %d teams", len(params.Intersections))
	log.Infof("There are %d teams", len(params.Streets))
	log.Infof("There are %d teams", len(params.Cars))

	// Compute data structure for this solver here (bitsets, tree, graph)

	// Solve
	s := &Solution{
		Intersections: []*IntersectionSolution{
			{
				ID: 1,
				StreetSolutions: []*StreetSolution{
					{
						Name:               "rue-d-athenes",
						GreenLightDuration: 3,
					},
				},
			},
		},
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
