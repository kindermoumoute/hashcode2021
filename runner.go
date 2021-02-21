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
		fmt.Println("solving", outputName)

		assertNoErr(ioutil.WriteFile(path.Join("output", outputName), Solve(params).Output(), 0644))
	}()
}

func (r *Runner) Wait() {
	r.wg.Wait()
}
