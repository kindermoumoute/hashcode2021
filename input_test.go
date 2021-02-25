package main

import (
	"io/ioutil"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {
	rawInput, err := ioutil.ReadFile("input/a.txt")
	require.NoError(t, err)
	textInput := &RawInput{
		Raw: string(rawInput),
	}
	input := DecodeInput(textInput.Raw)

	sort.SliceStable(input.Cars, func(i, j int) bool {
		return input.Cars[i].GetPathDuration() < input.Cars[j].GetPathDuration()
	})

	nbCars := float64(len(input.Cars))
	for i, car := range input.Cars {
		car.GlobalScore := (input.AlphaSort + ((nbCars - i)/nbCars)) * ((float64(input.SimulationTimeSeconds) - car.GetPathDuration()) / float64(input.DestinationScore))
	}

	score := 1.0
	assert.Equal(t, 1.0, score)
}
