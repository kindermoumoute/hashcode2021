package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"

	"hashcode2022/models"
)

var (
	//go:embed input/a_an_example.in.txt
	a string
	//go:embed input/b_basic.in.txt
	b string
	//go:embed input/c_coarse.in.txt
	c string
	//go:embed input/d_difficult.in.txt
	d string
	//go:embed input/e_elaborate.in.txt
	e string
)

func main() {
	// param1, param2 := 3, 3
	// aOutput := GenerateOutput(models.Output{Pizza: ParseInput(a).FindBestPizzaV2(2, 2)})
	// fmt.Println(ioutil.WriteFile("output/aV2.out", []byte(aOutput), 0755))
	// bOutput := GenerateOutput(models.Output{Pizza: ParseInput(b).FindBestPizzaV2(2, 2)})
	// fmt.Println(ioutil.WriteFile("output/bV2.out", []byte(bOutput), 0755))
	// cOutput := GenerateOutput(models.Output{Pizza: ParseInput(c).FindBestPizzaV2(2, 2)})
	// fmt.Println(ioutil.WriteFile("output/cV2.out", []byte(cOutput), 0755))
	// o := models.Output{Pizza: ParseInput(d).FindBestPizzaV2(param1, param2)}
	// fmt.Println(o.Pizza.Score)
	// dOutput := GenerateOutput(o)
	// fmt.Println(ioutil.WriteFile("output/dV2.out", []byte(dOutput), 0755))
	// fmt.Println("----------------a----------------")
	// output := models.GenerateOutput(models.Output{Pizza: models.MarkIngredients(ParseInput(a))})
	// fmt.Println(ioutil.WriteFile("output/aV9.out", []byte(output), 0755))
	// fmt.Println("----------------n----------------")
	// output = models.GenerateOutput(models.Output{Pizza: models.MarkIngredients(ParseInput(b))})
	// fmt.Println(ioutil.WriteFile("output/bV9.out", []byte(output), 0755))
	// fmt.Println("----------------c----------------")
	// output = models.GenerateOutput(models.Output{Pizza: models.MarkIngredients(ParseInput(c))})
	// fmt.Println(ioutil.WriteFile("output/cV9.out", []byte(output), 0755))
	// fmt.Println("----------------d----------------")
	// output = models.GenerateOutput(models.Output{Pizza: models.MarkIngredients(ParseInput(d))})
	// fmt.Println(ioutil.WriteFile("output/dV9.out", []byte(output), 0755))
	fmt.Println("----------------d----------------")
	output := models.GenerateOutput(models.Output{Pizza: models.MarkIngredients(ParseInput(d))})
	fmt.Println(ioutil.WriteFile("output/dV9.out", []byte(output), 0755))
}
