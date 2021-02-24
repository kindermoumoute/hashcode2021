package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
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

	r := NewRunner()
	for _, rawInput := range files {
		r.RunSolver(rawInput.FileName+"_latest.txt", SolverParameters{
			Input: DecodeInput(rawInput.Raw),
		})

		// for i := 0; i < 1000; i++ {
		// 	localI := i
		// }
	}

	r.Wait()
}

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}
