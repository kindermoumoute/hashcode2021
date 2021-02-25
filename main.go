package main

import (
	"flag"
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

	// Solve each input in a different worker
	wp := workerpool.New(runtime.NumCPU())
	for _, rawInput := range files {
		rawInput := rawInput
		taskLogger := log.Named(rawInput.FileName)
		wp.Submit(func() {
			solverParameters := SolverParameters{
				Input:     DecodeInput(rawInput.Raw),
				AlphaSort: 0.5,
			}
			solution := Solve(taskLogger, solverParameters)
			taskLogger.Infof("score is %d", solution.Scoring())
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
