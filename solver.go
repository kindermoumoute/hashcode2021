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
	log.Infof("There are %d teams", 4)

	// Compute data structure for this solver here (bitsets, tree, graph)

	// Solve
	s := &Solution{}

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
