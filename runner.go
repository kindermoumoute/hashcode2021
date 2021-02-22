package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"sync"
)

type Runner struct {
	wg sync.WaitGroup
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) RunSolver(outputName string, params SolverParameters) {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		solution := Solve(params)
		fmt.Println("solving", outputName, "score is", solution.Scoring())
		assertNoErr(ioutil.WriteFile(path.Join("output", outputName), solution.Output(), 0644))
	}()
}

func (r *Runner) Wait() {
	r.wg.Wait()
}
