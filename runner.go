package main

import (
	"io/ioutil"
	"path"
	"sync"

	"go.uber.org/zap"
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
		devLog, _ := zap.Config{
			Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
			Development:      true,
			Encoding:         "console",
			EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
		}.Build()
		s := devLog.Named(outputName).Sugar()
		solution := Solve(s, params)
		s.Infof("score is %d", solution.Scoring())
		assertNoErr(ioutil.WriteFile(path.Join("output", outputName), solution.Output(), 0644))
	}()
}

func (r *Runner) Wait() {
	r.wg.Wait()
}
