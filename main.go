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
	// Parse files
	var filePaths []string
	flag.Parse()
	filePaths = append(filePaths, flag.Args()...)
	// filePaths = append(filePaths, "input/a_example.in")

	// Create logger
	log := createLogger()

	// Parse files
	files := []*RawInput(nil)
	for _, filePath := range filePaths {
		log.Infof("reading %s", filePath)
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

	// alphas := []float64{}
	// for i := 0.05; i < 1.0; i += 0.05 {
	// 	alphas = append(alphas, i)
	// }

	alphasPerInput := map[string]float64{
		"a": 0.1545,
		"b": 0.1545,
		"c": 0.1545,
		"d": 0.1545,
		"e": 0.1545,
		"f": 0.1545,
	}

	// Solve each input in a different worker
	wp := workerpool.New(runtime.NumCPU())
	for _, rawInput := range files {
		rawInput := rawInput
		taskLogger := log.Named(rawInput.FileName)

		bestAlpha := alphasPerInput[rawInput.FileName]
		tryingAlphas := []float64{bestAlpha}
		for _, alpha := range tryingAlphas {
			alpha := alpha
			wp.Submit(func() {
				solverParameters := SolverParameters{
					Input:     DecodeInput(rawInput.Raw),
					AlphaSort: alpha,
					Beta:      1.2,
					Gamma:     0.4,
				}
				solution := Solve(taskLogger, solverParameters)
				output := solution.Output()
				taskLogger.Infof("score is %d", solution.Scoring())
				assertNoErr(ioutil.WriteFile(path.Join("output", rawInput.FileName+fmt.Sprintf("_%2f_latest.txt", alpha)), output, 0644))
			})
		}

	}
	wp.StopWait()
}

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func createLogger() *zap.SugaredLogger {
	devLog, _ := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}.Build()
	return devLog.Sugar()
}
