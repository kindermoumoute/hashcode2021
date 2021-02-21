package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var fileNames []string
	flag.Parse()
	fileNames = append(fileNames, flag.Args()...)
	fileNames = append(fileNames, "input/a_example.txt")
	fmt.Println("start")

	files := []*RawInput(nil)
	for _, fileName := range fileNames {
		fmt.Println("reading", fileName)
		rawInput, err := ioutil.ReadFile(fileName)
		assertNoErr(err)
		files = append(files, &RawInput{
			Raw:      string(rawInput),
			FileName: fileName,
		})
	}

	r := NewRunner()
	for _, rawInput := range files {
		r.RunSolver(rawInput.FileName+"_output.txt", SolverParameters{
			Input:  DecodeInput(rawInput.Raw),
			Param1: 2,
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
