package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	rawInput, err := ioutil.ReadFile("input/b_example.go")
	assertNoErr(err)

	fmt.Println(rawInput)

	for i := 0; i < 1000; i++ {

		//output := input.Clone().Solve(SolverParameters{
		//	Param1:i,
		//})
	}
}

func assertNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

type SolverParameters struct {
	Param1 int
}
