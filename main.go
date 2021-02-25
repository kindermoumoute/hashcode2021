package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strings"

	"github.com/gammazero/workerpool"
	"go.uber.org/zap"
)

func main() {
	var filePaths []string
	flag.Parse()
	filePaths = append(filePaths, flag.Args()...)
	// filePaths = append(filePaths, "input/a_example.in")

	files := []*RawInput(nil)
	for _, filePath := range filePaths {
		fmt.Println("reading", filePath)
		_, fileName := path.Split(filePath)
		tmp := strings.Split(fileName, ".")
		fileName = strings.Join(tmp[:len(tmp)-1], ".")
		rawInput, err := ioutil.ReadFile(filePath)
		assertNoErr(err)
		files = append(files, &RawInput{
			Raw:      string(rawInput),
			FileName: fileName,
		})
	}
	wp := workerpool.New(runtime.NumCPU())

	//r := NewRunner()
	for _, rawInput := range files {
		rawInput := rawInput
		wp.Submit(func() {
			solverParameters := SolverParameters{
				Input: DecodeInput(rawInput.Raw),
			}
			devLog, _ := zap.Config{
				Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
				Development:      true,
				Encoding:         "console",
				EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
				OutputPaths:      []string{"stdout"},
				ErrorOutputPaths: []string{"stdout"},
			}.Build()
			log := devLog.Named(rawInput.FileName).Sugar()
			solution := Solve(log, solverParameters)
			log.Infof("score is %d", solution.Scoring())
			assertNoErr(ioutil.WriteFile(path.Join("output", rawInput.FileName+"_latest.txt"), solution.Output(), 0644))
		})
	}

	wp.StopWait()
}

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}
